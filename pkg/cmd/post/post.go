package post

import (
	"context"
	"errors"
	"os"

	"github.com/spf13/cobra"

	"github.com/micnncim/slackctl/pkg/logging"
	"github.com/micnncim/slackctl/pkg/slack"
)

const (
	envSlackTokenKey = "SLACK_TOKEN"
)

type runner struct {
	slackClient *slack.Client

	logger *logging.Logger
}

func NewCommand() (*cobra.Command, error) {
	r, err := newRunner()
	if err != nil {
		return nil, err
	}

	cmd := &cobra.Command{
		Use:   "post <channel_id> <message>",
		Short: "Post a message to Slack channel",
		RunE: func(_ *cobra.Command, args []string) error {
			return r.run(context.Background(), args)
		},
	}

	return cmd, nil
}

func newRunner() (*runner, error) {
	token := os.Getenv(envSlackTokenKey)

	client, err := slack.NewClient(token)
	if err != nil {
		return nil, err
	}

	return &runner{
		slackClient: client,
		logger:      logging.NewLogger(os.Stdout),
	}, nil
}

func (r *runner) run(ctx context.Context, args []string) error {
	if len(args) != 2 {
		return errors.New("arguments must be <channel_id> <message>")
	}

	channelID := args[0]
	message := args[1]

	if err := r.slackClient.PostMessage(ctx, channelID, message); err != nil {
		r.logger.Errorf("Failed to post a message: %s\n", err)
		return nil
	}

	r.logger.Info("Successfully posted a message")

	return nil
}
