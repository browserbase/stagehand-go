package main

import (
	"context"
	"fmt"
	"os"

	"github.com/browserbase/stagehand-go/v3"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load() // Load .env from current directory (run from repo root)

	client := stagehand.NewClient() // Uses env vars: BROWSERBASE_API_KEY, BROWSERBASE_PROJECT_ID, MODEL_API_KEY

	startResponse, err := client.Sessions.Start(context.TODO(), stagehand.SessionStartParams{
		ModelName: "openai/gpt-5-nano",
	})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("Session started: %s\n", startResponse.Data.SessionID)

	sessionID := startResponse.Data.SessionID

	// Navigate to Hacker News
	_, err = client.Sessions.Navigate(
		context.TODO(),
		sessionID,
		stagehand.SessionNavigateParams{
			URL: "https://news.ycombinator.com",
		},
	)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Navigated to Hacker News")

	// Observe to find possible actions
	observeResponse, err := client.Sessions.Observe(
		context.TODO(),
		sessionID,
		stagehand.SessionObserveParams{
			Instruction: stagehand.String("find the link to view comments for the top post"),
		},
	)
	if err != nil {
		panic(err.Error())
	}

	actions := observeResponse.Data.Result
	fmt.Printf("Found %d possible actions\n", len(actions))

	if len(actions) == 0 {
		fmt.Println("No actions found")
		return
	}

	// Use the first action
	action := actions[0]
	fmt.Printf("Acting on: %s\n", action.Description)

	// Pass the action to Act
	actResponse, err := client.Sessions.Act(
		context.TODO(),
		sessionID,
		stagehand.SessionActParams{
			Input: stagehand.SessionActParamsInputUnion{
				OfAction: &stagehand.ActionParam{
					Description: action.Description,
					Selector:    action.Selector,
					Method:      stagehand.String(action.Method),
					Arguments:   action.Arguments,
				},
			},
		},
	)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("Act completed: %s\n", actResponse.Data.Result.Message)

	// Extract data from the page
	// We're now on the comments page, so extract the top comment text
	extractResponse, err := client.Sessions.Extract(
		context.TODO(),
		sessionID,
		stagehand.SessionExtractParams{
			Instruction: stagehand.String("extract the text of the top comment on this page"),
			Schema: map[string]any{
				"type": "object",
				"properties": map[string]any{
					"commentText": map[string]any{
						"type":        "string",
						"description": "The text content of the top comment",
					},
					"author": map[string]any{
						"type":        "string",
						"description": "The username of the comment author",
					},
				},
				"required": []string{"commentText"},
			},
		},
	)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("Extracted data: %+v\n", extractResponse.Data.Result)

	// Get the author from the extracted data
	extractedData := extractResponse.Data.Result.(map[string]any)
	author := extractedData["author"].(string)
	fmt.Printf("Looking up profile for author: %s\n", author)

	// Use the Agent to find the author's profile
	// Execute runs an autonomous agent that can navigate and interact with pages
	executeResponse, err := client.Sessions.Execute(
		context.TODO(),
		sessionID,
		stagehand.SessionExecuteParams{
			ExecuteOptions: stagehand.SessionExecuteParamsExecuteOptions{
				Instruction: fmt.Sprintf(
					"Find any personal website, GitHub, LinkedIn, or other best profile URL for the Hacker News user '%s'. "+
						"Click on their username to go to their profile page and look for any links they have shared. "+
						"Use Google Search with their username or other details from their profile if you dont find any direct links.",
					author,
				),
				MaxSteps: stagehand.Float(15),
			},
			AgentConfig: stagehand.SessionExecuteParamsAgentConfig{
				Model: stagehand.ModelConfigUnionParam{
					OfModelConfigModelConfigObject: &stagehand.ModelConfigModelConfigObjectParam{
						ModelName: "openai/gpt-5-nano",
						APIKey:    stagehand.String(os.Getenv("MODEL_API_KEY")),
					},
				},
				Cua: stagehand.Bool(false),
			},
		},
	)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("Agent completed: %s\n", executeResponse.Data.Result.Message)
	fmt.Printf("Agent success: %t\n", executeResponse.Data.Result.Success)
	fmt.Printf("Agent actions taken: %d\n", len(executeResponse.Data.Result.Actions))

	// End the session to clean up resources
	_, err = client.Sessions.End(
		context.TODO(),
		sessionID,
		stagehand.SessionEndParams{},
	)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Session ended")
}
