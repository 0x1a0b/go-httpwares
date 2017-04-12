// Copyright 2017 Michal Witkowski. All Rights Reserved.
// See LICENSE for licensing terms.

package httpwares_ctxtags

import (
	"context"
	"net/http"
)

type ctxMarker struct{}

var (
	// ctxMarkerKey is the Context value marker used by *all* logging middleware.
	// The logging middleware object must interf
	ctxMarkerKey = &ctxMarker{}
)

// Tags is the struct used for storing request tags between Context calls.
// This object is *not* thread safe, and should be handled only in the context of the request.
type Tags struct {
	values map[string]interface{}
}

// Set sets the given key in the metadata tags.
func (t *Tags) Set(key string, value interface{}) *Tags {
	t.values[key] = value
	return t
}

// Has checks if the given key exists.
func (t *Tags) Has(key string) bool {
	_, ok := t.values[key]
	return ok
}

// Values returns a map of key to values.
// Do not modify the underlying map, please use Set instead.
func (t *Tags) Values() map[string]interface{} {
	return t.values
}

// Extracts returns a pre-existing Tags object in the request's Context.
// If the context wasn't set in a tag interceptor, a no-op Tag storage is returned that will *not* be propagated in context.
func Extract(req *http.Request) *Tags {
	return ExtractFromContext(req.Context())
}

// Extracts returns a pre-existing Tags object in the request's Context.
// If the context wasn't set in a tag interceptor, a no-op Tag storage is returned that will *not* be propagated in context.
func ExtractFromContext(ctx context.Context) *Tags {
	t, ok := ctx.Value(ctxMarkerKey).(*Tags)
	if !ok {
		return &Tags{values: make(map[string]interface{})}
	}
	return t
}

func setInContext(ctx context.Context, tags *Tags) context.Context {
	return context.WithValue(ctx, ctxMarkerKey, tags)
}