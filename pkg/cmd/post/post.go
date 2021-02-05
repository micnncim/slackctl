package post

import (
	"context"
	"errors"

	"github.com/spf13/cobra"

	"github.com/micnncim/slackctl/pkg/logging"
	"github.com/micnncim/slackctl/pkg/slack"
)

type runner struct {
	client *slack.Client

	logger *logging.Logger
}

func NewCommand(ctx context.Context) *cobra.Command {
	r := &runner{
		client: slack.FromContext(ctx),
		logger: logging.FromContext(ctx),
	}

	cmd := &cobra.Command{
		Use:   "post <channel_id> <message>",
		Short: "Post a message to Slack channel",
		RunE: func(_ *cobra.Command, args []string) error {
			return r.run(context.Background(), args)
		},
	}

	return cmd
}

func (r *runner) run(ctx context.Context, args []string) error {
	if len(args) != 2 {
		return errors.New("arguments must be <channel_id> <message>")
	}

	channelID := args[0]
	message := args[1]

	if err := r.client.PostMessage(ctx, channelID, message); err != nil {
		r.logger.Errorf("Failed to post a message: %s\n", err)
		return nil
	}

	r.logger.Info("Successfully posted a message")

	return nil
}
