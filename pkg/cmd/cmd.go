package cmd

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/micnncim/slackctl/pkg/cmd/post"
	"github.com/micnncim/slackctl/pkg/logging"
	"github.com/micnncim/slackctl/pkg/slack"
	"github.com/spf13/cobra"
)

const (
	envSlackTokenKey = "SLACK_TOKEN"
)

func NewCommand(ctx context.Context) (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use: "slackctl",
	}

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-ch
		cancel()
	}()

	logger := logging.NewLogger(os.Stderr)
	ctx = logging.WithLogger(ctx, logger)

	token := os.Getenv(envSlackTokenKey)
	client, err := slack.NewClient(token)
	if err != nil {
		return nil, err
	}
	ctx = slack.WithClient(ctx, client)

	cmd.AddCommand(post.NewCommand(ctx))

	return cmd, nil
}
