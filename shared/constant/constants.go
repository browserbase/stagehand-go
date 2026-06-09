// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package constant

import (
	shimjson "github.com/browserbase/stagehand-go/v3/internal/encoding/json"
)

type Constant[T any] interface {
	Default() T
}

// ValueOf gives the default value of a constant from its type. It's helpful when
// constructing constants as variants in a one-of. Note that empty structs are
// marshalled by default. Usage: constant.ValueOf[constant.Foo]()
func ValueOf[T Constant[T]]() T {
	var t T
	return t.Default()
}

type Azure string                // Always "azure"
type AzureEntraID string         // Always "azureEntraId"
type Browserbase string          // Always "browserbase"
type External string             // Always "external"
type GoogleServiceAccount string // Always "googleServiceAccount"
type Running string              // Always "running"
type Vertex string               // Always "vertex"

func (c Azure) Default() Azure                               { return "azure" }
func (c AzureEntraID) Default() AzureEntraID                 { return "azureEntraId" }
func (c Browserbase) Default() Browserbase                   { return "browserbase" }
func (c External) Default() External                         { return "external" }
func (c GoogleServiceAccount) Default() GoogleServiceAccount { return "googleServiceAccount" }
func (c Running) Default() Running                           { return "running" }
func (c Vertex) Default() Vertex                             { return "vertex" }

func (c Azure) MarshalJSON() ([]byte, error)                { return marshalString(c) }
func (c AzureEntraID) MarshalJSON() ([]byte, error)         { return marshalString(c) }
func (c Browserbase) MarshalJSON() ([]byte, error)          { return marshalString(c) }
func (c External) MarshalJSON() ([]byte, error)             { return marshalString(c) }
func (c GoogleServiceAccount) MarshalJSON() ([]byte, error) { return marshalString(c) }
func (c Running) MarshalJSON() ([]byte, error)              { return marshalString(c) }
func (c Vertex) MarshalJSON() ([]byte, error)               { return marshalString(c) }

type constant[T any] interface {
	Constant[T]
	*T
}

func marshalString[T ~string, PT constant[T]](v T) ([]byte, error) {
	var zero T
	if v == zero {
		v = PT(&v).Default()
	}
	return shimjson.Marshal(string(v))
}
