// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stagehand

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/browserbase/stagehand-go/internal/apijson"
	"github.com/browserbase/stagehand-go/internal/requestconfig"
	"github.com/browserbase/stagehand-go/option"
	"github.com/browserbase/stagehand-go/packages/param"
	"github.com/browserbase/stagehand-go/packages/respjson"
	"github.com/browserbase/stagehand-go/shared/constant"
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

// Executes a browser action using natural language instructions or a predefined
// Action object.
func (r *SessionService) Act(ctx context.Context, id string, params SessionActParams, opts ...option.RequestOption) (res *SessionActResponse, err error) {
	if !param.IsOmitted(params.XLanguage) {
		opts = append(opts, option.WithHeader("x-language", fmt.Sprintf("%s", params.XLanguage)))
	}
	if !param.IsOmitted(params.XSDKVersion) {
		opts = append(opts, option.WithHeader("x-sdk-version", fmt.Sprintf("%s", params.XSDKVersion.Value)))
	}
	if !param.IsOmitted(params.XSentAt) {
		opts = append(opts, option.WithHeader("x-sent-at", fmt.Sprintf("%s", params.XSentAt.Value)))
	}
	if !param.IsOmitted(params.XStreamResponse) {
		opts = append(opts, option.WithHeader("x-stream-response", fmt.Sprintf("%s", params.XStreamResponse)))
	}
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/sessions/%s/act", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Terminates the browser session and releases all associated resources.
func (r *SessionService) End(ctx context.Context, id string, body SessionEndParams, opts ...option.RequestOption) (res *SessionEndResponse, err error) {
	if !param.IsOmitted(body.XLanguage) {
		opts = append(opts, option.WithHeader("x-language", fmt.Sprintf("%s", body.XLanguage)))
	}
	if !param.IsOmitted(body.XSDKVersion) {
		opts = append(opts, option.WithHeader("x-sdk-version", fmt.Sprintf("%s", body.XSDKVersion.Value)))
	}
	if !param.IsOmitted(body.XSentAt) {
		opts = append(opts, option.WithHeader("x-sent-at", fmt.Sprintf("%s", body.XSentAt.Value)))
	}
	if !param.IsOmitted(body.XStreamResponse) {
		opts = append(opts, option.WithHeader("x-stream-response", fmt.Sprintf("%s", body.XStreamResponse)))
	}
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/sessions/%s/end", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Runs an autonomous AI agent that can perform complex multi-step browser tasks.
func (r *SessionService) Execute(ctx context.Context, id string, params SessionExecuteParams, opts ...option.RequestOption) (res *SessionExecuteResponse, err error) {
	if !param.IsOmitted(params.XLanguage) {
		opts = append(opts, option.WithHeader("x-language", fmt.Sprintf("%s", params.XLanguage)))
	}
	if !param.IsOmitted(params.XSDKVersion) {
		opts = append(opts, option.WithHeader("x-sdk-version", fmt.Sprintf("%s", params.XSDKVersion.Value)))
	}
	if !param.IsOmitted(params.XSentAt) {
		opts = append(opts, option.WithHeader("x-sent-at", fmt.Sprintf("%s", params.XSentAt.Value)))
	}
	if !param.IsOmitted(params.XStreamResponse) {
		opts = append(opts, option.WithHeader("x-stream-response", fmt.Sprintf("%s", params.XStreamResponse)))
	}
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/sessions/%s/agentExecute", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Extracts structured data from the current page using AI-powered analysis.
func (r *SessionService) Extract(ctx context.Context, id string, params SessionExtractParams, opts ...option.RequestOption) (res *SessionExtractResponse, err error) {
	if !param.IsOmitted(params.XLanguage) {
		opts = append(opts, option.WithHeader("x-language", fmt.Sprintf("%s", params.XLanguage)))
	}
	if !param.IsOmitted(params.XSDKVersion) {
		opts = append(opts, option.WithHeader("x-sdk-version", fmt.Sprintf("%s", params.XSDKVersion.Value)))
	}
	if !param.IsOmitted(params.XSentAt) {
		opts = append(opts, option.WithHeader("x-sent-at", fmt.Sprintf("%s", params.XSentAt.Value)))
	}
	if !param.IsOmitted(params.XStreamResponse) {
		opts = append(opts, option.WithHeader("x-stream-response", fmt.Sprintf("%s", params.XStreamResponse)))
	}
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/sessions/%s/extract", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Navigates the browser to the specified URL.
func (r *SessionService) Navigate(ctx context.Context, id string, params SessionNavigateParams, opts ...option.RequestOption) (res *SessionNavigateResponse, err error) {
	if !param.IsOmitted(params.XLanguage) {
		opts = append(opts, option.WithHeader("x-language", fmt.Sprintf("%s", params.XLanguage)))
	}
	if !param.IsOmitted(params.XSDKVersion) {
		opts = append(opts, option.WithHeader("x-sdk-version", fmt.Sprintf("%s", params.XSDKVersion.Value)))
	}
	if !param.IsOmitted(params.XSentAt) {
		opts = append(opts, option.WithHeader("x-sent-at", fmt.Sprintf("%s", params.XSentAt.Value)))
	}
	if !param.IsOmitted(params.XStreamResponse) {
		opts = append(opts, option.WithHeader("x-stream-response", fmt.Sprintf("%s", params.XStreamResponse)))
	}
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/sessions/%s/navigate", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Identifies and returns available actions on the current page that match the
// given instruction.
func (r *SessionService) Observe(ctx context.Context, id string, params SessionObserveParams, opts ...option.RequestOption) (res *SessionObserveResponse, err error) {
	if !param.IsOmitted(params.XLanguage) {
		opts = append(opts, option.WithHeader("x-language", fmt.Sprintf("%s", params.XLanguage)))
	}
	if !param.IsOmitted(params.XSDKVersion) {
		opts = append(opts, option.WithHeader("x-sdk-version", fmt.Sprintf("%s", params.XSDKVersion.Value)))
	}
	if !param.IsOmitted(params.XSentAt) {
		opts = append(opts, option.WithHeader("x-sent-at", fmt.Sprintf("%s", params.XSentAt.Value)))
	}
	if !param.IsOmitted(params.XStreamResponse) {
		opts = append(opts, option.WithHeader("x-stream-response", fmt.Sprintf("%s", params.XStreamResponse)))
	}
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/sessions/%s/observe", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Creates a new browser session with the specified configuration. Returns a
// session ID used for all subsequent operations.
func (r *SessionService) Start(ctx context.Context, params SessionStartParams, opts ...option.RequestOption) (res *SessionStartResponse, err error) {
	if !param.IsOmitted(params.XLanguage) {
		opts = append(opts, option.WithHeader("x-language", fmt.Sprintf("%s", params.XLanguage)))
	}
	if !param.IsOmitted(params.XSDKVersion) {
		opts = append(opts, option.WithHeader("x-sdk-version", fmt.Sprintf("%s", params.XSDKVersion.Value)))
	}
	if !param.IsOmitted(params.XSentAt) {
		opts = append(opts, option.WithHeader("x-sent-at", fmt.Sprintf("%s", params.XSentAt.Value)))
	}
	if !param.IsOmitted(params.XStreamResponse) {
		opts = append(opts, option.WithHeader("x-stream-response", fmt.Sprintf("%s", params.XStreamResponse)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v1/sessions/start"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Action object returned by observe and used by act
type Action struct {
	// Human-readable description of the action
	Description string `json:"description,required"`
	// CSS selector or XPath for the element
	Selector string `json:"selector,required"`
	// Arguments to pass to the method
	Arguments []string `json:"arguments"`
	// The method to execute (click, fill, etc.)
	Method string `json:"method"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description respjson.Field
		Selector    respjson.Field
		Arguments   respjson.Field
		Method      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Action) RawJSON() string { return r.JSON.raw }
func (r *Action) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func ModelConfigParamOfModelConfigObject(modelName string) ModelConfigUnionParam {
	var variant ModelConfigObjectParam
	variant.ModelName = modelName
	return ModelConfigUnionParam{OfModelConfigObject: &variant}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ModelConfigUnionParam struct {
	OfString            param.Opt[string]       `json:",omitzero,inline"`
	OfModelConfigObject *ModelConfigObjectParam `json:",omitzero,inline"`
	paramUnion
}

func (u ModelConfigUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfModelConfigObject)
}
func (u *ModelConfigUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ModelConfigUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfModelConfigObject) {
		return u.OfModelConfigObject
	}
	return nil
}

// The property ModelName is required.
type ModelConfigObjectParam struct {
	ModelName string            `json:"modelName,required"`
	APIKey    param.Opt[string] `json:"apiKey,omitzero"`
	BaseURL   param.Opt[string] `json:"baseURL,omitzero" format:"uri"`
	paramObj
}

func (r ModelConfigObjectParam) MarshalJSON() (data []byte, err error) {
	type shadow ModelConfigObjectParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ModelConfigObjectParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SessionActResponse struct {
	Data SessionActResponseData `json:"data,required"`
	// Indicates whether the request was successful
	Success bool `json:"success,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
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

type SessionActResponseData struct {
	Result SessionActResponseDataResult `json:"result,required"`
	// Action ID for tracking
	ActionID string `json:"actionId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Result      respjson.Field
		ActionID    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SessionActResponseData) RawJSON() string { return r.JSON.raw }
func (r *SessionActResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SessionActResponseDataResult struct {
	// Description of the action that was performed
	ActionDescription string `json:"actionDescription,required"`
	// List of actions that were executed
	Actions []Action `json:"actions,required"`
	// Human-readable result message
	Message string `json:"message,required"`
	// Whether the action completed successfully
	Success bool `json:"success,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActionDescription respjson.Field
		Actions           respjson.Field
		Message           respjson.Field
		Success           respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SessionActResponseDataResult) RawJSON() string { return r.JSON.raw }
func (r *SessionActResponseDataResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SessionEndResponse struct {
	Success bool `json:"success,required"`
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

type SessionExecuteResponse struct {
	Data SessionExecuteResponseData `json:"data,required"`
	// Indicates whether the request was successful
	Success bool `json:"success,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SessionExecuteResponse) RawJSON() string { return r.JSON.raw }
func (r *SessionExecuteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SessionExecuteResponseData struct {
	Result SessionExecuteResponseDataResult `json:"result,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Result      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SessionExecuteResponseData) RawJSON() string { return r.JSON.raw }
func (r *SessionExecuteResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SessionExecuteResponseDataResult struct {
	Actions []SessionExecuteResponseDataResultAction `json:"actions,required"`
	// Whether the agent finished its task
	Completed bool `json:"completed,required"`
	// Summary of what the agent accomplished
	Message string `json:"message,required"`
	// Whether the agent completed successfully
	Success  bool                                  `json:"success,required"`
	Metadata map[string]any                        `json:"metadata"`
	Usage    SessionExecuteResponseDataResultUsage `json:"usage"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Actions     respjson.Field
		Completed   respjson.Field
		Message     respjson.Field
		Success     respjson.Field
		Metadata    respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SessionExecuteResponseDataResult) RawJSON() string { return r.JSON.raw }
func (r *SessionExecuteResponseDataResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SessionExecuteResponseDataResultAction struct {
	// Type of action taken
	Type        string `json:"type,required"`
	Action      string `json:"action"`
	Instruction string `json:"instruction"`
	PageText    string `json:"pageText"`
	PageURL     string `json:"pageUrl"`
	// Agent's reasoning for taking this action
	Reasoning     string `json:"reasoning"`
	TaskCompleted bool   `json:"taskCompleted"`
	// Time taken for this action in ms
	TimeMs      float64        `json:"timeMs"`
	ExtraFields map[string]any `json:",extras"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type          respjson.Field
		Action        respjson.Field
		Instruction   respjson.Field
		PageText      respjson.Field
		PageURL       respjson.Field
		Reasoning     respjson.Field
		TaskCompleted respjson.Field
		TimeMs        respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SessionExecuteResponseDataResultAction) RawJSON() string { return r.JSON.raw }
func (r *SessionExecuteResponseDataResultAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SessionExecuteResponseDataResultUsage struct {
	InferenceTimeMs   float64 `json:"inference_time_ms,required"`
	InputTokens       float64 `json:"input_tokens,required"`
	OutputTokens      float64 `json:"output_tokens,required"`
	CachedInputTokens float64 `json:"cached_input_tokens"`
	ReasoningTokens   float64 `json:"reasoning_tokens"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InferenceTimeMs   respjson.Field
		InputTokens       respjson.Field
		OutputTokens      respjson.Field
		CachedInputTokens respjson.Field
		ReasoningTokens   respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SessionExecuteResponseDataResultUsage) RawJSON() string { return r.JSON.raw }
func (r *SessionExecuteResponseDataResultUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SessionExtractResponse struct {
	Data SessionExtractResponseData `json:"data,required"`
	// Indicates whether the request was successful
	Success bool `json:"success,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SessionExtractResponse) RawJSON() string { return r.JSON.raw }
func (r *SessionExtractResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SessionExtractResponseData struct {
	// Extracted data matching the requested schema
	Result any `json:"result,required"`
	// Action ID for tracking
	ActionID string `json:"actionId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Result      respjson.Field
		ActionID    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SessionExtractResponseData) RawJSON() string { return r.JSON.raw }
func (r *SessionExtractResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SessionNavigateResponse struct {
	Data SessionNavigateResponseData `json:"data,required"`
	// Indicates whether the request was successful
	Success bool `json:"success,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SessionNavigateResponse) RawJSON() string { return r.JSON.raw }
func (r *SessionNavigateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SessionNavigateResponseData struct {
	// Navigation response (Playwright Response object or null)
	Result any `json:"result,required"`
	// Action ID for tracking
	ActionID string `json:"actionId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Result      respjson.Field
		ActionID    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SessionNavigateResponseData) RawJSON() string { return r.JSON.raw }
func (r *SessionNavigateResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SessionObserveResponse struct {
	Data SessionObserveResponseData `json:"data,required"`
	// Indicates whether the request was successful
	Success bool `json:"success,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SessionObserveResponse) RawJSON() string { return r.JSON.raw }
func (r *SessionObserveResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SessionObserveResponseData struct {
	Result []Action `json:"result,required"`
	// Action ID for tracking
	ActionID string `json:"actionId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Result      respjson.Field
		ActionID    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SessionObserveResponseData) RawJSON() string { return r.JSON.raw }
func (r *SessionObserveResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SessionStartResponse struct {
	Data SessionStartResponseData `json:"data,required"`
	// Indicates whether the request was successful
	Success bool `json:"success,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SessionStartResponse) RawJSON() string { return r.JSON.raw }
func (r *SessionStartResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SessionStartResponseData struct {
	Available bool `json:"available,required"`
	// Unique session identifier
	SessionID string `json:"sessionId,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Available   respjson.Field
		SessionID   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SessionStartResponseData) RawJSON() string { return r.JSON.raw }
func (r *SessionStartResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SessionActParams struct {
	// Natural language instruction or Action object
	Input SessionActParamsInputUnion `json:"input,omitzero,required"`
	// Target frame ID for the action
	FrameID param.Opt[string] `json:"frameId,omitzero"`
	// Version of the Stagehand SDK
	XSDKVersion param.Opt[string] `header:"x-sdk-version,omitzero" json:"-"`
	// ISO timestamp when request was sent
	XSentAt param.Opt[time.Time]    `header:"x-sent-at,omitzero" format:"date-time" json:"-"`
	Options SessionActParamsOptions `json:"options,omitzero"`
	// Client SDK language
	//
	// Any of "typescript", "python", "playground".
	XLanguage SessionActParamsXLanguage `header:"x-language,omitzero" json:"-"`
	// Whether to stream the response via SSE
	//
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
	OfString                      param.Opt[string]                 `json:",omitzero,inline"`
	OfSessionActsInputActionInput *SessionActParamsInputActionInput `json:",omitzero,inline"`
	paramUnion
}

func (u SessionActParamsInputUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfSessionActsInputActionInput)
}
func (u *SessionActParamsInputUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *SessionActParamsInputUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfSessionActsInputActionInput) {
		return u.OfSessionActsInputActionInput
	}
	return nil
}

// Action object returned by observe and used by act
//
// The properties Description, Selector are required.
type SessionActParamsInputActionInput struct {
	// Human-readable description of the action
	Description string `json:"description,required"`
	// CSS selector or XPath for the element
	Selector string `json:"selector,required"`
	// The method to execute (click, fill, etc.)
	Method param.Opt[string] `json:"method,omitzero"`
	// Arguments to pass to the method
	Arguments []string `json:"arguments,omitzero"`
	paramObj
}

func (r SessionActParamsInputActionInput) MarshalJSON() (data []byte, err error) {
	type shadow SessionActParamsInputActionInput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SessionActParamsInputActionInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SessionActParamsOptions struct {
	// Timeout in ms for the action
	Timeout param.Opt[float64]    `json:"timeout,omitzero"`
	Model   ModelConfigUnionParam `json:"model,omitzero"`
	// Variables to substitute in the action instruction
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

// Client SDK language
type SessionActParamsXLanguage string

const (
	SessionActParamsXLanguageTypescript SessionActParamsXLanguage = "typescript"
	SessionActParamsXLanguagePython     SessionActParamsXLanguage = "python"
	SessionActParamsXLanguagePlayground SessionActParamsXLanguage = "playground"
)

// Whether to stream the response via SSE
type SessionActParamsXStreamResponse string

const (
	SessionActParamsXStreamResponseTrue  SessionActParamsXStreamResponse = "true"
	SessionActParamsXStreamResponseFalse SessionActParamsXStreamResponse = "false"
)

type SessionEndParams struct {
	// Version of the Stagehand SDK
	XSDKVersion param.Opt[string] `header:"x-sdk-version,omitzero" json:"-"`
	// ISO timestamp when request was sent
	XSentAt param.Opt[time.Time] `header:"x-sent-at,omitzero" format:"date-time" json:"-"`
	// Client SDK language
	//
	// Any of "typescript", "python", "playground".
	XLanguage SessionEndParamsXLanguage `header:"x-language,omitzero" json:"-"`
	// Whether to stream the response via SSE
	//
	// Any of "true", "false".
	XStreamResponse SessionEndParamsXStreamResponse `header:"x-stream-response,omitzero" json:"-"`
	paramObj
}

// Client SDK language
type SessionEndParamsXLanguage string

const (
	SessionEndParamsXLanguageTypescript SessionEndParamsXLanguage = "typescript"
	SessionEndParamsXLanguagePython     SessionEndParamsXLanguage = "python"
	SessionEndParamsXLanguagePlayground SessionEndParamsXLanguage = "playground"
)

// Whether to stream the response via SSE
type SessionEndParamsXStreamResponse string

const (
	SessionEndParamsXStreamResponseTrue  SessionEndParamsXStreamResponse = "true"
	SessionEndParamsXStreamResponseFalse SessionEndParamsXStreamResponse = "false"
)

type SessionExecuteParams struct {
	AgentConfig    SessionExecuteParamsAgentConfig    `json:"agentConfig,omitzero,required"`
	ExecuteOptions SessionExecuteParamsExecuteOptions `json:"executeOptions,omitzero,required"`
	// Target frame ID for the agent
	FrameID param.Opt[string] `json:"frameId,omitzero"`
	// Version of the Stagehand SDK
	XSDKVersion param.Opt[string] `header:"x-sdk-version,omitzero" json:"-"`
	// ISO timestamp when request was sent
	XSentAt param.Opt[time.Time] `header:"x-sent-at,omitzero" format:"date-time" json:"-"`
	// Client SDK language
	//
	// Any of "typescript", "python", "playground".
	XLanguage SessionExecuteParamsXLanguage `header:"x-language,omitzero" json:"-"`
	// Whether to stream the response via SSE
	//
	// Any of "true", "false".
	XStreamResponse SessionExecuteParamsXStreamResponse `header:"x-stream-response,omitzero" json:"-"`
	paramObj
}

func (r SessionExecuteParams) MarshalJSON() (data []byte, err error) {
	type shadow SessionExecuteParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SessionExecuteParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SessionExecuteParamsAgentConfig struct {
	// Enable Computer Use Agent mode
	Cua param.Opt[bool] `json:"cua,omitzero"`
	// Custom system prompt for the agent
	SystemPrompt param.Opt[string]     `json:"systemPrompt,omitzero"`
	Model        ModelConfigUnionParam `json:"model,omitzero"`
	paramObj
}

func (r SessionExecuteParamsAgentConfig) MarshalJSON() (data []byte, err error) {
	type shadow SessionExecuteParamsAgentConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SessionExecuteParamsAgentConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Instruction is required.
type SessionExecuteParamsExecuteOptions struct {
	// Natural language instruction for the agent
	Instruction string `json:"instruction,required"`
	// Whether to visually highlight the cursor during execution
	HighlightCursor param.Opt[bool] `json:"highlightCursor,omitzero"`
	// Maximum number of steps the agent can take
	MaxSteps param.Opt[float64] `json:"maxSteps,omitzero"`
	paramObj
}

func (r SessionExecuteParamsExecuteOptions) MarshalJSON() (data []byte, err error) {
	type shadow SessionExecuteParamsExecuteOptions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SessionExecuteParamsExecuteOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Client SDK language
type SessionExecuteParamsXLanguage string

const (
	SessionExecuteParamsXLanguageTypescript SessionExecuteParamsXLanguage = "typescript"
	SessionExecuteParamsXLanguagePython     SessionExecuteParamsXLanguage = "python"
	SessionExecuteParamsXLanguagePlayground SessionExecuteParamsXLanguage = "playground"
)

// Whether to stream the response via SSE
type SessionExecuteParamsXStreamResponse string

const (
	SessionExecuteParamsXStreamResponseTrue  SessionExecuteParamsXStreamResponse = "true"
	SessionExecuteParamsXStreamResponseFalse SessionExecuteParamsXStreamResponse = "false"
)

type SessionExtractParams struct {
	// Target frame ID for the extraction
	FrameID param.Opt[string] `json:"frameId,omitzero"`
	// Natural language instruction for what to extract
	Instruction param.Opt[string] `json:"instruction,omitzero"`
	// Version of the Stagehand SDK
	XSDKVersion param.Opt[string] `header:"x-sdk-version,omitzero" json:"-"`
	// ISO timestamp when request was sent
	XSentAt param.Opt[time.Time]        `header:"x-sent-at,omitzero" format:"date-time" json:"-"`
	Options SessionExtractParamsOptions `json:"options,omitzero"`
	// JSON Schema defining the structure of data to extract
	Schema map[string]any `json:"schema,omitzero"`
	// Client SDK language
	//
	// Any of "typescript", "python", "playground".
	XLanguage SessionExtractParamsXLanguage `header:"x-language,omitzero" json:"-"`
	// Whether to stream the response via SSE
	//
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
	// CSS selector to scope extraction to a specific element
	Selector param.Opt[string] `json:"selector,omitzero"`
	// Timeout in ms for the extraction
	Timeout param.Opt[float64]    `json:"timeout,omitzero"`
	Model   ModelConfigUnionParam `json:"model,omitzero"`
	paramObj
}

func (r SessionExtractParamsOptions) MarshalJSON() (data []byte, err error) {
	type shadow SessionExtractParamsOptions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SessionExtractParamsOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Client SDK language
type SessionExtractParamsXLanguage string

const (
	SessionExtractParamsXLanguageTypescript SessionExtractParamsXLanguage = "typescript"
	SessionExtractParamsXLanguagePython     SessionExtractParamsXLanguage = "python"
	SessionExtractParamsXLanguagePlayground SessionExtractParamsXLanguage = "playground"
)

// Whether to stream the response via SSE
type SessionExtractParamsXStreamResponse string

const (
	SessionExtractParamsXStreamResponseTrue  SessionExtractParamsXStreamResponse = "true"
	SessionExtractParamsXStreamResponseFalse SessionExtractParamsXStreamResponse = "false"
)

type SessionNavigateParams struct {
	// URL to navigate to
	URL string `json:"url,required"`
	// Target frame ID for the navigation
	FrameID param.Opt[string] `json:"frameId,omitzero"`
	// Version of the Stagehand SDK
	XSDKVersion param.Opt[string] `header:"x-sdk-version,omitzero" json:"-"`
	// ISO timestamp when request was sent
	XSentAt param.Opt[time.Time]         `header:"x-sent-at,omitzero" format:"date-time" json:"-"`
	Options SessionNavigateParamsOptions `json:"options,omitzero"`
	// Client SDK language
	//
	// Any of "typescript", "python", "playground".
	XLanguage SessionNavigateParamsXLanguage `header:"x-language,omitzero" json:"-"`
	// Whether to stream the response via SSE
	//
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
	// Referer header to send with the request
	Referer param.Opt[string] `json:"referer,omitzero"`
	// Timeout in ms for the navigation
	Timeout param.Opt[float64] `json:"timeout,omitzero"`
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

// Client SDK language
type SessionNavigateParamsXLanguage string

const (
	SessionNavigateParamsXLanguageTypescript SessionNavigateParamsXLanguage = "typescript"
	SessionNavigateParamsXLanguagePython     SessionNavigateParamsXLanguage = "python"
	SessionNavigateParamsXLanguagePlayground SessionNavigateParamsXLanguage = "playground"
)

// Whether to stream the response via SSE
type SessionNavigateParamsXStreamResponse string

const (
	SessionNavigateParamsXStreamResponseTrue  SessionNavigateParamsXStreamResponse = "true"
	SessionNavigateParamsXStreamResponseFalse SessionNavigateParamsXStreamResponse = "false"
)

type SessionObserveParams struct {
	// Target frame ID for the observation
	FrameID param.Opt[string] `json:"frameId,omitzero"`
	// Natural language instruction for what actions to find
	Instruction param.Opt[string] `json:"instruction,omitzero"`
	// Version of the Stagehand SDK
	XSDKVersion param.Opt[string] `header:"x-sdk-version,omitzero" json:"-"`
	// ISO timestamp when request was sent
	XSentAt param.Opt[time.Time]        `header:"x-sent-at,omitzero" format:"date-time" json:"-"`
	Options SessionObserveParamsOptions `json:"options,omitzero"`
	// Client SDK language
	//
	// Any of "typescript", "python", "playground".
	XLanguage SessionObserveParamsXLanguage `header:"x-language,omitzero" json:"-"`
	// Whether to stream the response via SSE
	//
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
	// CSS selector to scope observation to a specific element
	Selector param.Opt[string] `json:"selector,omitzero"`
	// Timeout in ms for the observation
	Timeout param.Opt[float64]    `json:"timeout,omitzero"`
	Model   ModelConfigUnionParam `json:"model,omitzero"`
	paramObj
}

func (r SessionObserveParamsOptions) MarshalJSON() (data []byte, err error) {
	type shadow SessionObserveParamsOptions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SessionObserveParamsOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Client SDK language
type SessionObserveParamsXLanguage string

const (
	SessionObserveParamsXLanguageTypescript SessionObserveParamsXLanguage = "typescript"
	SessionObserveParamsXLanguagePython     SessionObserveParamsXLanguage = "python"
	SessionObserveParamsXLanguagePlayground SessionObserveParamsXLanguage = "playground"
)

// Whether to stream the response via SSE
type SessionObserveParamsXStreamResponse string

const (
	SessionObserveParamsXStreamResponseTrue  SessionObserveParamsXStreamResponse = "true"
	SessionObserveParamsXStreamResponseFalse SessionObserveParamsXStreamResponse = "false"
)

type SessionStartParams struct {
	// Model name to use for AI operations
	ModelName string `json:"modelName,required"`
	// Timeout in ms for act operations
	ActTimeoutMs param.Opt[float64] `json:"actTimeoutMs,omitzero"`
	// Existing Browserbase session ID to resume
	BrowserbaseSessionID param.Opt[string] `json:"browserbaseSessionID,omitzero"`
	DebugDom             param.Opt[bool]   `json:"debugDom,omitzero"`
	// Timeout in ms to wait for DOM to settle
	DomSettleTimeoutMs param.Opt[float64] `json:"domSettleTimeoutMs,omitzero"`
	Experimental       param.Opt[bool]    `json:"experimental,omitzero"`
	// Enable self-healing for failed actions
	SelfHeal param.Opt[bool] `json:"selfHeal,omitzero"`
	// Custom system prompt for AI operations
	SystemPrompt param.Opt[string] `json:"systemPrompt,omitzero"`
	// Logging verbosity level (0=quiet, 1=normal, 2=debug)
	Verbose              param.Opt[int64] `json:"verbose,omitzero"`
	WaitForCaptchaSolves param.Opt[bool]  `json:"waitForCaptchaSolves,omitzero"`
	// Version of the Stagehand SDK
	XSDKVersion param.Opt[string] `header:"x-sdk-version,omitzero" json:"-"`
	// ISO timestamp when request was sent
	XSentAt                        param.Opt[time.Time]                             `header:"x-sent-at,omitzero" format:"date-time" json:"-"`
	Browser                        SessionStartParamsBrowser                        `json:"browser,omitzero"`
	BrowserbaseSessionCreateParams SessionStartParamsBrowserbaseSessionCreateParams `json:"browserbaseSessionCreateParams,omitzero"`
	// Client SDK language
	//
	// Any of "typescript", "python", "playground".
	XLanguage SessionStartParamsXLanguage `header:"x-language,omitzero" json:"-"`
	// Whether to stream the response via SSE
	//
	// Any of "true", "false".
	XStreamResponse SessionStartParamsXStreamResponse `header:"x-stream-response,omitzero" json:"-"`
	paramObj
}

func (r SessionStartParams) MarshalJSON() (data []byte, err error) {
	type shadow SessionStartParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SessionStartParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SessionStartParamsBrowser struct {
	// Chrome DevTools Protocol URL for connecting to existing browser
	CdpURL        param.Opt[string]                      `json:"cdpUrl,omitzero"`
	LaunchOptions SessionStartParamsBrowserLaunchOptions `json:"launchOptions,omitzero"`
	// Browser type to use
	//
	// Any of "local", "browserbase".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r SessionStartParamsBrowser) MarshalJSON() (data []byte, err error) {
	type shadow SessionStartParamsBrowser
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SessionStartParamsBrowser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SessionStartParamsBrowser](
		"type", "local", "browserbase",
	)
}

type SessionStartParamsBrowserLaunchOptions struct {
	AcceptDownloads     param.Opt[bool]                                              `json:"acceptDownloads,omitzero"`
	CdpURL              param.Opt[string]                                            `json:"cdpUrl,omitzero"`
	ChromiumSandbox     param.Opt[bool]                                              `json:"chromiumSandbox,omitzero"`
	ConnectTimeoutMs    param.Opt[float64]                                           `json:"connectTimeoutMs,omitzero"`
	DeviceScaleFactor   param.Opt[float64]                                           `json:"deviceScaleFactor,omitzero"`
	Devtools            param.Opt[bool]                                              `json:"devtools,omitzero"`
	DownloadsPath       param.Opt[string]                                            `json:"downloadsPath,omitzero"`
	ExecutablePath      param.Opt[string]                                            `json:"executablePath,omitzero"`
	HasTouch            param.Opt[bool]                                              `json:"hasTouch,omitzero"`
	Headless            param.Opt[bool]                                              `json:"headless,omitzero"`
	IgnoreHTTPSErrors   param.Opt[bool]                                              `json:"ignoreHTTPSErrors,omitzero"`
	Locale              param.Opt[string]                                            `json:"locale,omitzero"`
	PreserveUserDataDir param.Opt[bool]                                              `json:"preserveUserDataDir,omitzero"`
	UserDataDir         param.Opt[string]                                            `json:"userDataDir,omitzero"`
	Args                []string                                                     `json:"args,omitzero"`
	IgnoreDefaultArgs   SessionStartParamsBrowserLaunchOptionsIgnoreDefaultArgsUnion `json:"ignoreDefaultArgs,omitzero"`
	Proxy               SessionStartParamsBrowserLaunchOptionsProxy                  `json:"proxy,omitzero"`
	Viewport            SessionStartParamsBrowserLaunchOptionsViewport               `json:"viewport,omitzero"`
	paramObj
}

func (r SessionStartParamsBrowserLaunchOptions) MarshalJSON() (data []byte, err error) {
	type shadow SessionStartParamsBrowserLaunchOptions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SessionStartParamsBrowserLaunchOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type SessionStartParamsBrowserLaunchOptionsIgnoreDefaultArgsUnion struct {
	OfBool        param.Opt[bool] `json:",omitzero,inline"`
	OfStringArray []string        `json:",omitzero,inline"`
	paramUnion
}

func (u SessionStartParamsBrowserLaunchOptionsIgnoreDefaultArgsUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfStringArray)
}
func (u *SessionStartParamsBrowserLaunchOptionsIgnoreDefaultArgsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *SessionStartParamsBrowserLaunchOptionsIgnoreDefaultArgsUnion) asAny() any {
	if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

// The property Server is required.
type SessionStartParamsBrowserLaunchOptionsProxy struct {
	Server   string            `json:"server,required"`
	Bypass   param.Opt[string] `json:"bypass,omitzero"`
	Password param.Opt[string] `json:"password,omitzero"`
	Username param.Opt[string] `json:"username,omitzero"`
	paramObj
}

func (r SessionStartParamsBrowserLaunchOptionsProxy) MarshalJSON() (data []byte, err error) {
	type shadow SessionStartParamsBrowserLaunchOptionsProxy
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SessionStartParamsBrowserLaunchOptionsProxy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Height, Width are required.
type SessionStartParamsBrowserLaunchOptionsViewport struct {
	Height float64 `json:"height,required"`
	Width  float64 `json:"width,required"`
	paramObj
}

func (r SessionStartParamsBrowserLaunchOptionsViewport) MarshalJSON() (data []byte, err error) {
	type shadow SessionStartParamsBrowserLaunchOptionsViewport
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SessionStartParamsBrowserLaunchOptionsViewport) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SessionStartParamsBrowserbaseSessionCreateParams struct {
	ExtensionID     param.Opt[string]                                               `json:"extensionId,omitzero"`
	KeepAlive       param.Opt[bool]                                                 `json:"keepAlive,omitzero"`
	ProjectID       param.Opt[string]                                               `json:"projectId,omitzero"`
	Timeout         param.Opt[float64]                                              `json:"timeout,omitzero"`
	BrowserSettings SessionStartParamsBrowserbaseSessionCreateParamsBrowserSettings `json:"browserSettings,omitzero"`
	Proxies         SessionStartParamsBrowserbaseSessionCreateParamsProxiesUnion    `json:"proxies,omitzero"`
	// Any of "us-west-2", "us-east-1", "eu-central-1", "ap-southeast-1".
	Region       string         `json:"region,omitzero"`
	UserMetadata map[string]any `json:"userMetadata,omitzero"`
	paramObj
}

func (r SessionStartParamsBrowserbaseSessionCreateParams) MarshalJSON() (data []byte, err error) {
	type shadow SessionStartParamsBrowserbaseSessionCreateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SessionStartParamsBrowserbaseSessionCreateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SessionStartParamsBrowserbaseSessionCreateParams](
		"region", "us-west-2", "us-east-1", "eu-central-1", "ap-southeast-1",
	)
}

type SessionStartParamsBrowserbaseSessionCreateParamsBrowserSettings struct {
	AdvancedStealth param.Opt[bool]                                                            `json:"advancedStealth,omitzero"`
	BlockAds        param.Opt[bool]                                                            `json:"blockAds,omitzero"`
	ExtensionID     param.Opt[string]                                                          `json:"extensionId,omitzero"`
	LogSession      param.Opt[bool]                                                            `json:"logSession,omitzero"`
	RecordSession   param.Opt[bool]                                                            `json:"recordSession,omitzero"`
	SolveCaptchas   param.Opt[bool]                                                            `json:"solveCaptchas,omitzero"`
	Context         SessionStartParamsBrowserbaseSessionCreateParamsBrowserSettingsContext     `json:"context,omitzero"`
	Fingerprint     SessionStartParamsBrowserbaseSessionCreateParamsBrowserSettingsFingerprint `json:"fingerprint,omitzero"`
	Viewport        SessionStartParamsBrowserbaseSessionCreateParamsBrowserSettingsViewport    `json:"viewport,omitzero"`
	paramObj
}

func (r SessionStartParamsBrowserbaseSessionCreateParamsBrowserSettings) MarshalJSON() (data []byte, err error) {
	type shadow SessionStartParamsBrowserbaseSessionCreateParamsBrowserSettings
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SessionStartParamsBrowserbaseSessionCreateParamsBrowserSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property ID is required.
type SessionStartParamsBrowserbaseSessionCreateParamsBrowserSettingsContext struct {
	ID      string          `json:"id,required"`
	Persist param.Opt[bool] `json:"persist,omitzero"`
	paramObj
}

func (r SessionStartParamsBrowserbaseSessionCreateParamsBrowserSettingsContext) MarshalJSON() (data []byte, err error) {
	type shadow SessionStartParamsBrowserbaseSessionCreateParamsBrowserSettingsContext
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SessionStartParamsBrowserbaseSessionCreateParamsBrowserSettingsContext) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SessionStartParamsBrowserbaseSessionCreateParamsBrowserSettingsFingerprint struct {
	// Any of "chrome", "edge", "firefox", "safari".
	Browsers []string `json:"browsers,omitzero"`
	// Any of "desktop", "mobile".
	Devices []string `json:"devices,omitzero"`
	// Any of "1", "2".
	HTTPVersion string   `json:"httpVersion,omitzero"`
	Locales     []string `json:"locales,omitzero"`
	// Any of "android", "ios", "linux", "macos", "windows".
	OperatingSystems []string                                                                         `json:"operatingSystems,omitzero"`
	Screen           SessionStartParamsBrowserbaseSessionCreateParamsBrowserSettingsFingerprintScreen `json:"screen,omitzero"`
	paramObj
}

func (r SessionStartParamsBrowserbaseSessionCreateParamsBrowserSettingsFingerprint) MarshalJSON() (data []byte, err error) {
	type shadow SessionStartParamsBrowserbaseSessionCreateParamsBrowserSettingsFingerprint
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SessionStartParamsBrowserbaseSessionCreateParamsBrowserSettingsFingerprint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SessionStartParamsBrowserbaseSessionCreateParamsBrowserSettingsFingerprint](
		"httpVersion", "1", "2",
	)
}

type SessionStartParamsBrowserbaseSessionCreateParamsBrowserSettingsFingerprintScreen struct {
	MaxHeight param.Opt[float64] `json:"maxHeight,omitzero"`
	MaxWidth  param.Opt[float64] `json:"maxWidth,omitzero"`
	MinHeight param.Opt[float64] `json:"minHeight,omitzero"`
	MinWidth  param.Opt[float64] `json:"minWidth,omitzero"`
	paramObj
}

func (r SessionStartParamsBrowserbaseSessionCreateParamsBrowserSettingsFingerprintScreen) MarshalJSON() (data []byte, err error) {
	type shadow SessionStartParamsBrowserbaseSessionCreateParamsBrowserSettingsFingerprintScreen
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SessionStartParamsBrowserbaseSessionCreateParamsBrowserSettingsFingerprintScreen) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SessionStartParamsBrowserbaseSessionCreateParamsBrowserSettingsViewport struct {
	Height param.Opt[float64] `json:"height,omitzero"`
	Width  param.Opt[float64] `json:"width,omitzero"`
	paramObj
}

func (r SessionStartParamsBrowserbaseSessionCreateParamsBrowserSettingsViewport) MarshalJSON() (data []byte, err error) {
	type shadow SessionStartParamsBrowserbaseSessionCreateParamsBrowserSettingsViewport
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SessionStartParamsBrowserbaseSessionCreateParamsBrowserSettingsViewport) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type SessionStartParamsBrowserbaseSessionCreateParamsProxiesUnion struct {
	OfBool                                                    param.Opt[bool]                                                         `json:",omitzero,inline"`
	OfSessionStartsBrowserbaseSessionCreateParamsProxiesArray []SessionStartParamsBrowserbaseSessionCreateParamsProxiesArrayItemUnion `json:",omitzero,inline"`
	paramUnion
}

func (u SessionStartParamsBrowserbaseSessionCreateParamsProxiesUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfSessionStartsBrowserbaseSessionCreateParamsProxiesArray)
}
func (u *SessionStartParamsBrowserbaseSessionCreateParamsProxiesUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *SessionStartParamsBrowserbaseSessionCreateParamsProxiesUnion) asAny() any {
	if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfSessionStartsBrowserbaseSessionCreateParamsProxiesArray) {
		return &u.OfSessionStartsBrowserbaseSessionCreateParamsProxiesArray
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type SessionStartParamsBrowserbaseSessionCreateParamsProxiesArrayItemUnion struct {
	OfSessionStartsBrowserbaseSessionCreateParamsProxiesArrayItemBrowserbaseProxyConfig *SessionStartParamsBrowserbaseSessionCreateParamsProxiesArrayItemBrowserbaseProxyConfig `json:",omitzero,inline"`
	OfSessionStartsBrowserbaseSessionCreateParamsProxiesArrayItemExternalProxyConfig    *SessionStartParamsBrowserbaseSessionCreateParamsProxiesArrayItemExternalProxyConfig    `json:",omitzero,inline"`
	paramUnion
}

func (u SessionStartParamsBrowserbaseSessionCreateParamsProxiesArrayItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfSessionStartsBrowserbaseSessionCreateParamsProxiesArrayItemBrowserbaseProxyConfig, u.OfSessionStartsBrowserbaseSessionCreateParamsProxiesArrayItemExternalProxyConfig)
}
func (u *SessionStartParamsBrowserbaseSessionCreateParamsProxiesArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *SessionStartParamsBrowserbaseSessionCreateParamsProxiesArrayItemUnion) asAny() any {
	if !param.IsOmitted(u.OfSessionStartsBrowserbaseSessionCreateParamsProxiesArrayItemBrowserbaseProxyConfig) {
		return u.OfSessionStartsBrowserbaseSessionCreateParamsProxiesArrayItemBrowserbaseProxyConfig
	} else if !param.IsOmitted(u.OfSessionStartsBrowserbaseSessionCreateParamsProxiesArrayItemExternalProxyConfig) {
		return u.OfSessionStartsBrowserbaseSessionCreateParamsProxiesArrayItemExternalProxyConfig
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SessionStartParamsBrowserbaseSessionCreateParamsProxiesArrayItemUnion) GetGeolocation() *SessionStartParamsBrowserbaseSessionCreateParamsProxiesArrayItemBrowserbaseProxyConfigGeolocation {
	if vt := u.OfSessionStartsBrowserbaseSessionCreateParamsProxiesArrayItemBrowserbaseProxyConfig; vt != nil {
		return &vt.Geolocation
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SessionStartParamsBrowserbaseSessionCreateParamsProxiesArrayItemUnion) GetServer() *string {
	if vt := u.OfSessionStartsBrowserbaseSessionCreateParamsProxiesArrayItemExternalProxyConfig; vt != nil {
		return &vt.Server
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SessionStartParamsBrowserbaseSessionCreateParamsProxiesArrayItemUnion) GetPassword() *string {
	if vt := u.OfSessionStartsBrowserbaseSessionCreateParamsProxiesArrayItemExternalProxyConfig; vt != nil && vt.Password.Valid() {
		return &vt.Password.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SessionStartParamsBrowserbaseSessionCreateParamsProxiesArrayItemUnion) GetUsername() *string {
	if vt := u.OfSessionStartsBrowserbaseSessionCreateParamsProxiesArrayItemExternalProxyConfig; vt != nil && vt.Username.Valid() {
		return &vt.Username.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SessionStartParamsBrowserbaseSessionCreateParamsProxiesArrayItemUnion) GetType() *string {
	if vt := u.OfSessionStartsBrowserbaseSessionCreateParamsProxiesArrayItemBrowserbaseProxyConfig; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfSessionStartsBrowserbaseSessionCreateParamsProxiesArrayItemExternalProxyConfig; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SessionStartParamsBrowserbaseSessionCreateParamsProxiesArrayItemUnion) GetDomainPattern() *string {
	if vt := u.OfSessionStartsBrowserbaseSessionCreateParamsProxiesArrayItemBrowserbaseProxyConfig; vt != nil && vt.DomainPattern.Valid() {
		return &vt.DomainPattern.Value
	} else if vt := u.OfSessionStartsBrowserbaseSessionCreateParamsProxiesArrayItemExternalProxyConfig; vt != nil && vt.DomainPattern.Valid() {
		return &vt.DomainPattern.Value
	}
	return nil
}

// The property Type is required.
type SessionStartParamsBrowserbaseSessionCreateParamsProxiesArrayItemBrowserbaseProxyConfig struct {
	DomainPattern param.Opt[string]                                                                                 `json:"domainPattern,omitzero"`
	Geolocation   SessionStartParamsBrowserbaseSessionCreateParamsProxiesArrayItemBrowserbaseProxyConfigGeolocation `json:"geolocation,omitzero"`
	// This field can be elided, and will marshal its zero value as "browserbase".
	Type constant.Browserbase `json:"type,required"`
	paramObj
}

func (r SessionStartParamsBrowserbaseSessionCreateParamsProxiesArrayItemBrowserbaseProxyConfig) MarshalJSON() (data []byte, err error) {
	type shadow SessionStartParamsBrowserbaseSessionCreateParamsProxiesArrayItemBrowserbaseProxyConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SessionStartParamsBrowserbaseSessionCreateParamsProxiesArrayItemBrowserbaseProxyConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Country is required.
type SessionStartParamsBrowserbaseSessionCreateParamsProxiesArrayItemBrowserbaseProxyConfigGeolocation struct {
	Country string            `json:"country,required"`
	City    param.Opt[string] `json:"city,omitzero"`
	State   param.Opt[string] `json:"state,omitzero"`
	paramObj
}

func (r SessionStartParamsBrowserbaseSessionCreateParamsProxiesArrayItemBrowserbaseProxyConfigGeolocation) MarshalJSON() (data []byte, err error) {
	type shadow SessionStartParamsBrowserbaseSessionCreateParamsProxiesArrayItemBrowserbaseProxyConfigGeolocation
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SessionStartParamsBrowserbaseSessionCreateParamsProxiesArrayItemBrowserbaseProxyConfigGeolocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Server, Type are required.
type SessionStartParamsBrowserbaseSessionCreateParamsProxiesArrayItemExternalProxyConfig struct {
	Server        string            `json:"server,required"`
	DomainPattern param.Opt[string] `json:"domainPattern,omitzero"`
	Password      param.Opt[string] `json:"password,omitzero"`
	Username      param.Opt[string] `json:"username,omitzero"`
	// This field can be elided, and will marshal its zero value as "external".
	Type constant.External `json:"type,required"`
	paramObj
}

func (r SessionStartParamsBrowserbaseSessionCreateParamsProxiesArrayItemExternalProxyConfig) MarshalJSON() (data []byte, err error) {
	type shadow SessionStartParamsBrowserbaseSessionCreateParamsProxiesArrayItemExternalProxyConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SessionStartParamsBrowserbaseSessionCreateParamsProxiesArrayItemExternalProxyConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Client SDK language
type SessionStartParamsXLanguage string

const (
	SessionStartParamsXLanguageTypescript SessionStartParamsXLanguage = "typescript"
	SessionStartParamsXLanguagePython     SessionStartParamsXLanguage = "python"
	SessionStartParamsXLanguagePlayground SessionStartParamsXLanguage = "playground"
)

// Whether to stream the response via SSE
type SessionStartParamsXStreamResponse string

const (
	SessionStartParamsXStreamResponseTrue  SessionStartParamsXStreamResponse = "true"
	SessionStartParamsXStreamResponseFalse SessionStartParamsXStreamResponse = "false"
)
