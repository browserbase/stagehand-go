// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stagehand

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/browserbase/stagehand-go/internal/apijson"
	"github.com/browserbase/stagehand-go/internal/requestconfig"
	"github.com/browserbase/stagehand-go/option"
	"github.com/browserbase/stagehand-go/packages/param"
	"github.com/browserbase/stagehand-go/packages/respjson"
)

// SessionService contains methods and other services that help with interacting
// with the stagehand API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSessionService] method instead.
type SessionService struct {
	Options []option.RequestOption
}

// NewSessionService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSessionService(opts ...option.RequestOption) (r SessionService) {
	r = SessionService{}
	r.Options = opts
	return
}

// Performs a browser action based on natural language instruction or a specific
// action object returned by observe().
func (r *SessionService) Act(ctx context.Context, sessionID string, params SessionActParams, opts ...option.RequestOption) (res *SessionActResponse, err error) {
	if !param.IsOmitted(params.XStreamResponse) {
		opts = append(opts, option.WithHeader("x-stream-response", fmt.Sprintf("%s", params.XStreamResponse)))
	}
	opts = slices.Concat(r.Options, opts)
	if sessionID == "" {
		err = errors.New("missing required sessionId parameter")
		return
	}
	path := fmt.Sprintf("sessions/%s/act", sessionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Closes the browser and cleans up all resources associated with the session.
func (r *SessionService) End(ctx context.Context, sessionID string, opts ...option.RequestOption) (res *SessionEndResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if sessionID == "" {
		err = errors.New("missing required sessionId parameter")
		return
	}
	path := fmt.Sprintf("sessions/%s/end", sessionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Runs an autonomous agent that can perform multiple actions to complete a complex
// task.
func (r *SessionService) ExecuteAgent(ctx context.Context, sessionID string, params SessionExecuteAgentParams, opts ...option.RequestOption) (res *SessionExecuteAgentResponse, err error) {
	if !param.IsOmitted(params.XStreamResponse) {
		opts = append(opts, option.WithHeader("x-stream-response", fmt.Sprintf("%s", params.XStreamResponse)))
	}
	opts = slices.Concat(r.Options, opts)
	if sessionID == "" {
		err = errors.New("missing required sessionId parameter")
		return
	}
	path := fmt.Sprintf("sessions/%s/agentExecute", sessionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Extracts data from the current page using natural language instructions and
// optional JSON schema for structured output.
func (r *SessionService) Extract(ctx context.Context, sessionID string, params SessionExtractParams, opts ...option.RequestOption) (res *SessionExtractResponseUnion, err error) {
	if !param.IsOmitted(params.XStreamResponse) {
		opts = append(opts, option.WithHeader("x-stream-response", fmt.Sprintf("%s", params.XStreamResponse)))
	}
	opts = slices.Concat(r.Options, opts)
	if sessionID == "" {
		err = errors.New("missing required sessionId parameter")
		return
	}
	path := fmt.Sprintf("sessions/%s/extract", sessionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Navigates the browser to the specified URL and waits for page load.
func (r *SessionService) Navigate(ctx context.Context, sessionID string, params SessionNavigateParams, opts ...option.RequestOption) (res *SessionNavigateResponse, err error) {
	if !param.IsOmitted(params.XStreamResponse) {
		opts = append(opts, option.WithHeader("x-stream-response", fmt.Sprintf("%s", params.XStreamResponse)))
	}
	opts = slices.Concat(r.Options, opts)
	if sessionID == "" {
		err = errors.New("missing required sessionId parameter")
		return
	}
	path := fmt.Sprintf("sessions/%s/navigate", sessionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Returns a list of candidate actions that can be performed on the page,
// optionally filtered by natural language instruction.
func (r *SessionService) Observe(ctx context.Context, sessionID string, params SessionObserveParams, opts ...option.RequestOption) (res *[]Action, err error) {
	if !param.IsOmitted(params.XStreamResponse) {
		opts = append(opts, option.WithHeader("x-stream-response", fmt.Sprintf("%s", params.XStreamResponse)))
	}
	opts = slices.Concat(r.Options, opts)
	if sessionID == "" {
		err = errors.New("missing required sessionId parameter")
		return
	}
	path := fmt.Sprintf("sessions/%s/observe", sessionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Initializes a new Stagehand session with a browser instance. Returns a session
// ID that must be used for all subsequent requests.
func (r *SessionService) Start(ctx context.Context, body SessionStartParams, opts ...option.RequestOption) (res *SessionStartResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "sessions/start"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type Action struct {
	// Arguments for the method
	Arguments []string `json:"arguments,required"`
	// Human-readable description of the action
	Description string `json:"description,required"`
	// Method to execute (e.g., "click", "fill")
	Method string `json:"method,required"`
	// CSS or XPath selector for the element
	Selector string `json:"selector,required"`
	// CDP backend node ID
	BackendNodeID int64 `json:"backendNodeId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Arguments     respjson.Field
		Description   respjson.Field
		Method        respjson.Field
		Selector      respjson.Field
		BackendNodeID respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Action) RawJSON() string { return r.JSON.raw }
func (r *Action) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this Action to a ActionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ActionParam.Overrides()
func (r Action) ToParam() ActionParam {
	return param.Override[ActionParam](json.RawMessage(r.RawJSON()))
}

// The properties Arguments, Description, Method, Selector are required.
type ActionParam struct {
	// Arguments for the method
	Arguments []string `json:"arguments,omitzero,required"`
	// Human-readable description of the action
	Description string `json:"description,required"`
	// Method to execute (e.g., "click", "fill")
	Method string `json:"method,required"`
	// CSS or XPath selector for the element
	Selector string `json:"selector,required"`
	// CDP backend node ID
	BackendNodeID param.Opt[int64] `json:"backendNodeId,omitzero"`
	paramObj
}

func (r ActionParam) MarshalJSON() (data []byte, err error) {
	type shadow ActionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ModelConfigParam struct {
	// API key for the model provider
	APIKey param.Opt[string] `json:"apiKey,omitzero"`
	// Custom base URL for API
	BaseURL param.Opt[string] `json:"baseURL,omitzero" format:"uri"`
	// Model name
	Model param.Opt[string] `json:"model,omitzero"`
	// Any of "openai", "anthropic", "google".
	Provider ModelConfigProvider `json:"provider,omitzero"`
	paramObj
}

func (r ModelConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow ModelConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ModelConfigParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ModelConfigProvider string

const (
	ModelConfigProviderOpenAI    ModelConfigProvider = "openai"
	ModelConfigProviderAnthropic ModelConfigProvider = "anthropic"
	ModelConfigProviderGoogle    ModelConfigProvider = "google"
)

type SessionActResponse struct {
	// Actions that were executed
	Actions []Action `json:"actions,required"`
	// Result message
	Message string `json:"message,required"`
	// Whether the action succeeded
	Success bool `json:"success,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Actions     respjson.Field
		Message     respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SessionActResponse) RawJSON() string { return r.JSON.raw }
func (r *SessionActResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SessionEndResponse struct {
	Success bool `json:"success"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SessionEndResponse) RawJSON() string { return r.JSON.raw }
func (r *SessionEndResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SessionExecuteAgentResponse struct {
	// Final message from the agent
	Message string `json:"message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SessionExecuteAgentResponse) RawJSON() string { return r.JSON.raw }
func (r *SessionExecuteAgentResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SessionExtractResponseUnion contains all possible properties and values from
// [SessionExtractResponseExtraction], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfSessionExtractResponseCustomItem]
type SessionExtractResponseUnion struct {
	// This field will be present if the value is a [any] instead of an object.
	OfSessionExtractResponseCustomItem any `json:",inline"`
	// This field is from variant [SessionExtractResponseExtraction].
	Extraction string `json:"extraction"`
	JSON       struct {
		OfSessionExtractResponseCustomItem respjson.Field
		Extraction                         respjson.Field
		raw                                string
	} `json:"-"`
}

func (u SessionExtractResponseUnion) AsSessionExtractResponseExtraction() (v SessionExtractResponseExtraction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SessionExtractResponseUnion) AsCustom() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SessionExtractResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *SessionExtractResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Default extraction result
type SessionExtractResponseExtraction struct {
	Extraction string `json:"extraction"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Extraction  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SessionExtractResponseExtraction) RawJSON() string { return r.JSON.raw }
func (r *SessionExtractResponseExtraction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Navigation response (may be null)
type SessionNavigateResponse struct {
	Ok     bool   `json:"ok"`
	Status int64  `json:"status"`
	URL    string `json:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Ok          respjson.Field
		Status      respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SessionNavigateResponse) RawJSON() string { return r.JSON.raw }
func (r *SessionNavigateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SessionStartResponse struct {
	// Whether the session is ready to use
	Available bool `json:"available,required"`
	// Unique identifier for the session
	SessionID string `json:"sessionId,required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Available   respjson.Field
		SessionID   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SessionStartResponse) RawJSON() string { return r.JSON.raw }
func (r *SessionStartResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SessionActParams struct {
	// Natural language instruction
	Input SessionActParamsInputUnion `json:"input,omitzero,required"`
	// Frame ID to act on (optional)
	FrameID param.Opt[string]       `json:"frameId,omitzero"`
	Options SessionActParamsOptions `json:"options,omitzero"`
	// Any of "true", "false".
	XStreamResponse SessionActParamsXStreamResponse `header:"x-stream-response,omitzero" json:"-"`
	paramObj
}

func (r SessionActParams) MarshalJSON() (data []byte, err error) {
	type shadow SessionActParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SessionActParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type SessionActParamsInputUnion struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfAction *ActionParam      `json:",omitzero,inline"`
	paramUnion
}

func (u SessionActParamsInputUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfAction)
}
func (u *SessionActParamsInputUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *SessionActParamsInputUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfAction) {
		return u.OfAction
	}
	return nil
}

type SessionActParamsOptions struct {
	// Timeout in milliseconds
	Timeout param.Opt[int64] `json:"timeout,omitzero"`
	Model   ModelConfigParam `json:"model,omitzero"`
	// Template variables for instruction
	Variables map[string]string `json:"variables,omitzero"`
	paramObj
}

func (r SessionActParamsOptions) MarshalJSON() (data []byte, err error) {
	type shadow SessionActParamsOptions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SessionActParamsOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SessionActParamsXStreamResponse string

const (
	SessionActParamsXStreamResponseTrue  SessionActParamsXStreamResponse = "true"
	SessionActParamsXStreamResponseFalse SessionActParamsXStreamResponse = "false"
)

type SessionExecuteAgentParams struct {
	AgentConfig    SessionExecuteAgentParamsAgentConfig    `json:"agentConfig,omitzero,required"`
	ExecuteOptions SessionExecuteAgentParamsExecuteOptions `json:"executeOptions,omitzero,required"`
	FrameID        param.Opt[string]                       `json:"frameId,omitzero"`
	// Any of "true", "false".
	XStreamResponse SessionExecuteAgentParamsXStreamResponse `header:"x-stream-response,omitzero" json:"-"`
	paramObj
}

func (r SessionExecuteAgentParams) MarshalJSON() (data []byte, err error) {
	type shadow SessionExecuteAgentParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SessionExecuteAgentParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SessionExecuteAgentParamsAgentConfig struct {
	// Enable Computer Use Agent mode
	Cua          param.Opt[bool]                                `json:"cua,omitzero"`
	SystemPrompt param.Opt[string]                              `json:"systemPrompt,omitzero"`
	Model        SessionExecuteAgentParamsAgentConfigModelUnion `json:"model,omitzero"`
	// Any of "openai", "anthropic", "google".
	Provider string `json:"provider,omitzero"`
	paramObj
}

func (r SessionExecuteAgentParamsAgentConfig) MarshalJSON() (data []byte, err error) {
	type shadow SessionExecuteAgentParamsAgentConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SessionExecuteAgentParamsAgentConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SessionExecuteAgentParamsAgentConfig](
		"provider", "openai", "anthropic", "google",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type SessionExecuteAgentParamsAgentConfigModelUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfModelConfig *ModelConfigParam `json:",omitzero,inline"`
	paramUnion
}

func (u SessionExecuteAgentParamsAgentConfigModelUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfModelConfig)
}
func (u *SessionExecuteAgentParamsAgentConfigModelUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *SessionExecuteAgentParamsAgentConfigModelUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfModelConfig) {
		return u.OfModelConfig
	}
	return nil
}

// The property Instruction is required.
type SessionExecuteAgentParamsExecuteOptions struct {
	// Task for the agent to complete
	Instruction string `json:"instruction,required"`
	// Visually highlight the cursor during actions
	HighlightCursor param.Opt[bool] `json:"highlightCursor,omitzero"`
	// Maximum number of steps the agent can take
	MaxSteps param.Opt[int64] `json:"maxSteps,omitzero"`
	paramObj
}

func (r SessionExecuteAgentParamsExecuteOptions) MarshalJSON() (data []byte, err error) {
	type shadow SessionExecuteAgentParamsExecuteOptions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SessionExecuteAgentParamsExecuteOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SessionExecuteAgentParamsXStreamResponse string

const (
	SessionExecuteAgentParamsXStreamResponseTrue  SessionExecuteAgentParamsXStreamResponse = "true"
	SessionExecuteAgentParamsXStreamResponseFalse SessionExecuteAgentParamsXStreamResponse = "false"
)

type SessionExtractParams struct {
	// Frame ID to extract from
	FrameID param.Opt[string] `json:"frameId,omitzero"`
	// Natural language instruction for extraction
	Instruction param.Opt[string]           `json:"instruction,omitzero"`
	Options     SessionExtractParamsOptions `json:"options,omitzero"`
	// JSON Schema for structured output
	Schema map[string]any `json:"schema,omitzero"`
	// Any of "true", "false".
	XStreamResponse SessionExtractParamsXStreamResponse `header:"x-stream-response,omitzero" json:"-"`
	paramObj
}

func (r SessionExtractParams) MarshalJSON() (data []byte, err error) {
	type shadow SessionExtractParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SessionExtractParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SessionExtractParamsOptions struct {
	// Extract only from elements matching this selector
	Selector param.Opt[string] `json:"selector,omitzero"`
	Timeout  param.Opt[int64]  `json:"timeout,omitzero"`
	Model    ModelConfigParam  `json:"model,omitzero"`
	paramObj
}

func (r SessionExtractParamsOptions) MarshalJSON() (data []byte, err error) {
	type shadow SessionExtractParamsOptions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SessionExtractParamsOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SessionExtractParamsXStreamResponse string

const (
	SessionExtractParamsXStreamResponseTrue  SessionExtractParamsXStreamResponse = "true"
	SessionExtractParamsXStreamResponseFalse SessionExtractParamsXStreamResponse = "false"
)

type SessionNavigateParams struct {
	// URL to navigate to
	URL     string                       `json:"url,required" format:"uri"`
	FrameID param.Opt[string]            `json:"frameId,omitzero"`
	Options SessionNavigateParamsOptions `json:"options,omitzero"`
	// Any of "true", "false".
	XStreamResponse SessionNavigateParamsXStreamResponse `header:"x-stream-response,omitzero" json:"-"`
	paramObj
}

func (r SessionNavigateParams) MarshalJSON() (data []byte, err error) {
	type shadow SessionNavigateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SessionNavigateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SessionNavigateParamsOptions struct {
	// When to consider navigation complete
	//
	// Any of "load", "domcontentloaded", "networkidle".
	WaitUntil string `json:"waitUntil,omitzero"`
	paramObj
}

func (r SessionNavigateParamsOptions) MarshalJSON() (data []byte, err error) {
	type shadow SessionNavigateParamsOptions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SessionNavigateParamsOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SessionNavigateParamsOptions](
		"waitUntil", "load", "domcontentloaded", "networkidle",
	)
}

type SessionNavigateParamsXStreamResponse string

const (
	SessionNavigateParamsXStreamResponseTrue  SessionNavigateParamsXStreamResponse = "true"
	SessionNavigateParamsXStreamResponseFalse SessionNavigateParamsXStreamResponse = "false"
)

type SessionObserveParams struct {
	// Frame ID to observe
	FrameID param.Opt[string] `json:"frameId,omitzero"`
	// Natural language instruction to filter actions
	Instruction param.Opt[string]           `json:"instruction,omitzero"`
	Options     SessionObserveParamsOptions `json:"options,omitzero"`
	// Any of "true", "false".
	XStreamResponse SessionObserveParamsXStreamResponse `header:"x-stream-response,omitzero" json:"-"`
	paramObj
}

func (r SessionObserveParams) MarshalJSON() (data []byte, err error) {
	type shadow SessionObserveParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SessionObserveParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SessionObserveParamsOptions struct {
	// Observe only elements matching this selector
	Selector param.Opt[string] `json:"selector,omitzero"`
	Timeout  param.Opt[int64]  `json:"timeout,omitzero"`
	Model    ModelConfigParam  `json:"model,omitzero"`
	paramObj
}

func (r SessionObserveParamsOptions) MarshalJSON() (data []byte, err error) {
	type shadow SessionObserveParamsOptions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SessionObserveParamsOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SessionObserveParamsXStreamResponse string

const (
	SessionObserveParamsXStreamResponseTrue  SessionObserveParamsXStreamResponse = "true"
	SessionObserveParamsXStreamResponseFalse SessionObserveParamsXStreamResponse = "false"
)

type SessionStartParams struct {
	// API key for Browserbase Cloud
	BrowserbaseAPIKey string `json:"BROWSERBASE_API_KEY,required"`
	// Project ID for Browserbase
	BrowserbaseProjectID string `json:"BROWSERBASE_PROJECT_ID,required"`
	// Timeout in ms to wait for DOM to settle
	DomSettleTimeout param.Opt[int64] `json:"domSettleTimeout,omitzero"`
	// AI model to use for actions (must be prefixed with provider/)
	Model param.Opt[string] `json:"model,omitzero"`
	// Enable self-healing for failed actions
	SelfHeal param.Opt[bool] `json:"selfHeal,omitzero"`
	// Custom system prompt for AI actions
	SystemPrompt param.Opt[string] `json:"systemPrompt,omitzero"`
	// Logging verbosity level
	Verbose param.Opt[int64] `json:"verbose,omitzero"`
	paramObj
}

func (r SessionStartParams) MarshalJSON() (data []byte, err error) {
	type shadow SessionStartParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SessionStartParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
