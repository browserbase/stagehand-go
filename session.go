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

type SessionStartResponse = any

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
