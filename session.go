// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stagehand

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"slices"

	shimjson "github.com/browserbase/stagehand-go/internal/encoding/json"
	"github.com/browserbase/stagehand-go/internal/requestconfig"
	"github.com/browserbase/stagehand-go/option"
	"github.com/browserbase/stagehand-go/packages/param"
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
func (r *SessionService) Act(ctx context.Context, id any, params SessionActParams, opts ...option.RequestOption) (res *SessionActResponse, err error) {
	if !param.IsOmitted(params.XLanguage) {
		opts = append(opts, option.WithHeader("x-language", fmt.Sprintf("%s", params.XLanguage)))
	}
	if !param.IsOmitted(params.XSDKVersion) {
		opts = append(opts, option.WithHeader("x-sdk-version", fmt.Sprintf("%s", params.XSDKVersion)))
	}
	if !param.IsOmitted(params.XSentAt) {
		opts = append(opts, option.WithHeader("x-sent-at", fmt.Sprintf("%s", params.XSentAt)))
	}
	if !param.IsOmitted(params.XStreamResponse) {
		opts = append(opts, option.WithHeader("x-stream-response", fmt.Sprintf("%s", params.XStreamResponse)))
	}
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("sessions/%v/act", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Terminates the browser session and releases all associated resources.
func (r *SessionService) End(ctx context.Context, id any, body SessionEndParams, opts ...option.RequestOption) (res *SessionEndResponse, err error) {
	if !param.IsOmitted(body.XLanguage) {
		opts = append(opts, option.WithHeader("x-language", fmt.Sprintf("%s", body.XLanguage)))
	}
	if !param.IsOmitted(body.XSDKVersion) {
		opts = append(opts, option.WithHeader("x-sdk-version", fmt.Sprintf("%s", body.XSDKVersion)))
	}
	if !param.IsOmitted(body.XSentAt) {
		opts = append(opts, option.WithHeader("x-sent-at", fmt.Sprintf("%s", body.XSentAt)))
	}
	if !param.IsOmitted(body.XStreamResponse) {
		opts = append(opts, option.WithHeader("x-stream-response", fmt.Sprintf("%s", body.XStreamResponse)))
	}
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("sessions/%v/end", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Runs an autonomous AI agent that can perform complex multi-step browser tasks.
func (r *SessionService) ExecuteAgent(ctx context.Context, id any, params SessionExecuteAgentParams, opts ...option.RequestOption) (res *SessionExecuteAgentResponse, err error) {
	if !param.IsOmitted(params.XLanguage) {
		opts = append(opts, option.WithHeader("x-language", fmt.Sprintf("%s", params.XLanguage)))
	}
	if !param.IsOmitted(params.XSDKVersion) {
		opts = append(opts, option.WithHeader("x-sdk-version", fmt.Sprintf("%s", params.XSDKVersion)))
	}
	if !param.IsOmitted(params.XSentAt) {
		opts = append(opts, option.WithHeader("x-sent-at", fmt.Sprintf("%s", params.XSentAt)))
	}
	if !param.IsOmitted(params.XStreamResponse) {
		opts = append(opts, option.WithHeader("x-stream-response", fmt.Sprintf("%s", params.XStreamResponse)))
	}
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("sessions/%v/agentExecute", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Extracts structured data from the current page using AI-powered analysis.
func (r *SessionService) Extract(ctx context.Context, id any, params SessionExtractParams, opts ...option.RequestOption) (res *SessionExtractResponse, err error) {
	if !param.IsOmitted(params.XLanguage) {
		opts = append(opts, option.WithHeader("x-language", fmt.Sprintf("%s", params.XLanguage)))
	}
	if !param.IsOmitted(params.XSDKVersion) {
		opts = append(opts, option.WithHeader("x-sdk-version", fmt.Sprintf("%s", params.XSDKVersion)))
	}
	if !param.IsOmitted(params.XSentAt) {
		opts = append(opts, option.WithHeader("x-sent-at", fmt.Sprintf("%s", params.XSentAt)))
	}
	if !param.IsOmitted(params.XStreamResponse) {
		opts = append(opts, option.WithHeader("x-stream-response", fmt.Sprintf("%s", params.XStreamResponse)))
	}
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("sessions/%v/extract", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Navigates the browser to the specified URL.
func (r *SessionService) Navigate(ctx context.Context, id any, params SessionNavigateParams, opts ...option.RequestOption) (res *SessionNavigateResponse, err error) {
	if !param.IsOmitted(params.XLanguage) {
		opts = append(opts, option.WithHeader("x-language", fmt.Sprintf("%s", params.XLanguage)))
	}
	if !param.IsOmitted(params.XSDKVersion) {
		opts = append(opts, option.WithHeader("x-sdk-version", fmt.Sprintf("%s", params.XSDKVersion)))
	}
	if !param.IsOmitted(params.XSentAt) {
		opts = append(opts, option.WithHeader("x-sent-at", fmt.Sprintf("%s", params.XSentAt)))
	}
	if !param.IsOmitted(params.XStreamResponse) {
		opts = append(opts, option.WithHeader("x-stream-response", fmt.Sprintf("%s", params.XStreamResponse)))
	}
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("sessions/%v/navigate", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Identifies and returns available actions on the current page that match the
// given instruction.
func (r *SessionService) Observe(ctx context.Context, id any, params SessionObserveParams, opts ...option.RequestOption) (res *SessionObserveResponse, err error) {
	if !param.IsOmitted(params.XLanguage) {
		opts = append(opts, option.WithHeader("x-language", fmt.Sprintf("%s", params.XLanguage)))
	}
	if !param.IsOmitted(params.XSDKVersion) {
		opts = append(opts, option.WithHeader("x-sdk-version", fmt.Sprintf("%s", params.XSDKVersion)))
	}
	if !param.IsOmitted(params.XSentAt) {
		opts = append(opts, option.WithHeader("x-sent-at", fmt.Sprintf("%s", params.XSentAt)))
	}
	if !param.IsOmitted(params.XStreamResponse) {
		opts = append(opts, option.WithHeader("x-stream-response", fmt.Sprintf("%s", params.XStreamResponse)))
	}
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("sessions/%v/observe", id)
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
		opts = append(opts, option.WithHeader("x-sdk-version", fmt.Sprintf("%s", params.XSDKVersion)))
	}
	if !param.IsOmitted(params.XSentAt) {
		opts = append(opts, option.WithHeader("x-sent-at", fmt.Sprintf("%s", params.XSentAt)))
	}
	if !param.IsOmitted(params.XStreamResponse) {
		opts = append(opts, option.WithHeader("x-stream-response", fmt.Sprintf("%s", params.XStreamResponse)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "sessions/start"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type SessionActResponse = any

type SessionEndResponse = any

type SessionExecuteAgentResponse = any

type SessionExtractResponse = any

type SessionNavigateResponse = any

type SessionObserveResponse = any

type SessionStartResponse = any

type SessionActParams struct {
	Body            any
	XLanguage       any `header:"x-language,omitzero" json:"-"`
	XSDKVersion     any `header:"x-sdk-version,omitzero" json:"-"`
	XSentAt         any `header:"x-sent-at,omitzero" json:"-"`
	XStreamResponse any `header:"x-stream-response,omitzero" json:"-"`
	paramObj
}

func (r SessionActParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *SessionActParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

type SessionEndParams struct {
	XLanguage       any `header:"x-language,omitzero" json:"-"`
	XSDKVersion     any `header:"x-sdk-version,omitzero" json:"-"`
	XSentAt         any `header:"x-sent-at,omitzero" json:"-"`
	XStreamResponse any `header:"x-stream-response,omitzero" json:"-"`
	paramObj
}

type SessionExecuteAgentParams struct {
	Body            any
	XLanguage       any `header:"x-language,omitzero" json:"-"`
	XSDKVersion     any `header:"x-sdk-version,omitzero" json:"-"`
	XSentAt         any `header:"x-sent-at,omitzero" json:"-"`
	XStreamResponse any `header:"x-stream-response,omitzero" json:"-"`
	paramObj
}

func (r SessionExecuteAgentParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *SessionExecuteAgentParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

type SessionExtractParams struct {
	Body            any
	XLanguage       any `header:"x-language,omitzero" json:"-"`
	XSDKVersion     any `header:"x-sdk-version,omitzero" json:"-"`
	XSentAt         any `header:"x-sent-at,omitzero" json:"-"`
	XStreamResponse any `header:"x-stream-response,omitzero" json:"-"`
	paramObj
}

func (r SessionExtractParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *SessionExtractParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

type SessionNavigateParams struct {
	Body            any
	XLanguage       any `header:"x-language,omitzero" json:"-"`
	XSDKVersion     any `header:"x-sdk-version,omitzero" json:"-"`
	XSentAt         any `header:"x-sent-at,omitzero" json:"-"`
	XStreamResponse any `header:"x-stream-response,omitzero" json:"-"`
	paramObj
}

func (r SessionNavigateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *SessionNavigateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

type SessionObserveParams struct {
	Body            any
	XLanguage       any `header:"x-language,omitzero" json:"-"`
	XSDKVersion     any `header:"x-sdk-version,omitzero" json:"-"`
	XSentAt         any `header:"x-sent-at,omitzero" json:"-"`
	XStreamResponse any `header:"x-stream-response,omitzero" json:"-"`
	paramObj
}

func (r SessionObserveParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *SessionObserveParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

type SessionStartParams struct {
	Body            any
	XLanguage       any `header:"x-language,omitzero" json:"-"`
	XSDKVersion     any `header:"x-sdk-version,omitzero" json:"-"`
	XSentAt         any `header:"x-sent-at,omitzero" json:"-"`
	XStreamResponse any `header:"x-stream-response,omitzero" json:"-"`
	paramObj
}

func (r SessionStartParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *SessionStartParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}
