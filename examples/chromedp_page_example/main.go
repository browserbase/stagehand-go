package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/browserbase/stagehand-go"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load() // Load .env from current directory (run from repo root)

	// Environment variables required (same as other examples):
	// - BROWSERBASE_API_KEY
	// - BROWSERBASE_PROJECT_ID
	// - MODEL_API_KEY
	client := stagehand.NewClient()

	// 1) Start a Stagehand session and get the Browserbase CDP URL.
	startResponse, err := client.Sessions.Start(context.TODO(), stagehand.SessionStartParams{
		ModelName: "openai/gpt-5-nano",
		Browser: stagehand.SessionStartParamsBrowser{
			Type: "browserbase",
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	sessionID := startResponse.Data.SessionID
	cdpURL := startResponse.Data.CdpURL
	if cdpURL == "" {
		log.Fatalf("start response missing cdpUrl (sessionID=%s); cannot attach chromedp", sessionID)
	}
	fmt.Printf("Session started: %s\n", sessionID)
	fmt.Printf("CDP URL: %s\n", cdpURL)

	// 2) Connect chromedp to the same browser over CDP.
	allocatorCtx, cancelAllocator := chromedp.NewRemoteAllocator(context.Background(), cdpURL)
	defer cancelAllocator()

	tabCtx, cancelTab := chromedp.NewContext(allocatorCtx)
	defer cancelTab()

	// 3) Use chromedp to navigate, then directly call CDP Page.getFrameTree to get the frame ID.
	var frameID string
	err = chromedp.Run(
		tabCtx,
		chromedp.Navigate("https://example.com"),
		chromedp.WaitReady("body", chromedp.ByQuery),
		chromedp.ActionFunc(func(ctx context.Context) error {
			tree, err := page.GetFrameTree().Do(ctx)
			if err != nil {
				return err
			}
			if tree.Frame == nil {
				return fmt.Errorf("Page.getFrameTree returned nil root frame")
			}
			frameID = string(tree.Frame.ID)
			if frameID == "" {
				return fmt.Errorf("Page.getFrameTree returned empty frame id")
			}
			return nil
		}),
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Resolved frameId via chromedp Page.getFrameTree: %s\n", frameID)

	// 4) Pass that frameId into Stagehand methods.
	observeResponse, err := client.Sessions.Observe(
		context.TODO(),
		sessionID,
		stagehand.SessionObserveParams{
			Instruction: stagehand.String("Find the most relevant click target on this page"),
			FrameID:     stagehand.String(frameID),
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Observed %d possible actions\n", len(observeResponse.Data.Result))

	// 5) Demonstrate using Stagehand to click something in the same tab/frame.
	_, err = client.Sessions.Act(
		context.TODO(),
		sessionID,
		stagehand.SessionActParams{
			Input: stagehand.SessionActParamsInputUnion{
				OfString: stagehand.String("Click on the 'More information...' link"),
			},
			FrameID: stagehand.String(frameID),
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	// Give navigation a moment to settle.
	time.Sleep(2 * time.Second)

	// 6) Extract from the same frame.
	schema := map[string]any{
		"type": "object",
		"properties": map[string]any{
			"title": map[string]any{"type": "string"},
			"url":   map[string]any{"type": "string"},
		},
	}
	extractResponse, err := client.Sessions.Extract(
		context.TODO(),
		sessionID,
		stagehand.SessionExtractParams{
			Instruction: stagehand.String("Extract the page title and current URL"),
			Schema:      schema,
			FrameID:     stagehand.String(frameID),
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Extracted: %+v\n", extractResponse.Data.Result)

	// 7) End session.
	_, err = client.Sessions.End(context.TODO(), sessionID, stagehand.SessionEndParams{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Session ended")
}

