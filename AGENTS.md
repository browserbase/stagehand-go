# Overview

This is a stateless SDK client for the Stagehand API provided by Browserbase.com, built using Stainless.

The Stagehand API allows users to control browsers using a natural language interface with these high-level primitives:

- `Act("do xyz on this page")` - Perform actions on the page
- `Observe("look for xyz elements on this page")` - Find interactive elements
- `Extract("find xyz information on this page")` - Extract structured data from pages

The other calls provided are `Start()` and `End()` to begin and end a browser session, and `Navigate()` which is a helper to visit a specific URL.

Stagehand supports two modes:
- **Cloud mode**: Uses Browserbase cloud browsers (requires BROWSERBASE_API_KEY)
- **Local mode**: Uses a local browser on your machine (only requires MODEL_API_KEY)

These primitives are intended to be combined with your browser driver library of choice, e.g. chromedp, go-rod, etc.

**Links:**
- GitHub: https://github.com/browserbase/stagehand-go
- Documentation: https://docs.stagehand.dev/v3/sdk/go
- Go Package: https://pkg.go.dev/github.com/browserbase/stagehand-go/v3

## Usage

Refer to the README.md "## Usage" section and `./examples` directory for detailed usage examples.

For installation instructions, see the "## Installation" section of the README.

For running examples, see the "### Running the Examples" section of the README.

## Common Tasks

```bash
# Install the SDK
go get -u 'github.com/browserbase/stagehand-go/v3'

# Set environment variables for cloud mode
export BROWSERBASE_API_KEY="your-bb-api-key"
export BROWSERBASE_PROJECT_ID="your-bb-project-uuid"
export MODEL_API_KEY="sk-proj-your-llm-api-key"

# Run cloud example
go run examples/basic/main.go

# Run local example (only needs MODEL_API_KEY)
go run examples/local/main.go
```

```go
// Cloud mode - uses Browserbase
client := stagehand.NewClient()
startResponse, _ := client.Sessions.Start(context.TODO(), stagehand.SessionStartParams{
    ModelName: "openai/gpt-5-nano",
})
sessionID := startResponse.Data.SessionID
client.Sessions.Navigate(context.TODO(), sessionID, stagehand.SessionNavigateParams{
    URL: "https://example.com",
})
client.Sessions.Act(context.TODO(), sessionID, stagehand.SessionActParams{
    Input: stagehand.SessionActParamsInputUnion{OfString: stagehand.String("click login")},
})
client.Sessions.End(context.TODO(), sessionID, stagehand.SessionEndParams{})
```

```go
// Local mode - uses local browser
client := stagehand.NewClient(option.WithServer("local"))
defer client.Close()
startResponse, _ := client.Sessions.Start(context.TODO(), stagehand.SessionStartParams{
    ModelName: "openai/gpt-5-nano",
    Browser: stagehand.SessionStartParamsBrowser{
        Type: "local",
        LaunchOptions: stagehand.SessionStartParamsBrowserLaunchOptions{
            Headless: stagehand.Bool(true),
        },
    },
})
// ... same API as cloud mode
```
