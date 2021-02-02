package slack

import (
	"context"
	"errors"

	"github.com/slack-go/slack"
)

type Client struct {
	client *slack.Client
}

func NewClient(token string) (*Client, error) {
	if token == "" {
		return nil, errors.New("slack token is missing")
	}

	return &Client{
		client: slack.New(token),
	}, nil
}

func (c *Client) PostMessage(ctx context.Context, channelID, message string) error {
	_, _, err := c.client.PostMessageContext(ctx, channelID, slack.MsgOptionText(message, false))

	return err
}
