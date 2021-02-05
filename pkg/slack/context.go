package slack

import (
	"context"
)

type ctxKey struct{}

func NewContext(ctx context.Context, c *Client) context.Context {
	return context.WithValue(ctx, ctxKey{}, c)
}

func FromContext(ctx context.Context) *Client {
	if c, ok := ctx.Value(ctxKey{}).(*Client); ok {
		return c
	}

	return nil
}
