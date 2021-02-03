package slack

import (
	"context"
)

type ctxKey struct{}

func WithClient(ctx context.Context, client *Client) context.Context {
	return context.WithValue(ctx, ctxKey{}, client)
}

func ClientFromContext(ctx context.Context) *Client {
	client, ok := ctx.Value(ctxKey{}).(*Client)
	if !ok {
		return &Client{}
	}
	return client
}
