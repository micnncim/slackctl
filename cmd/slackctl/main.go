package main

import (
	"context"
	"log"
	"time"

	"github.com/micnncim/slackctl/pkg/cmd"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	c, err := cmd.NewCommand(ctx)
	if err != nil {
		return err
	}
	if err := c.Execute(); err != nil {
		return err
	}

	return nil
}
