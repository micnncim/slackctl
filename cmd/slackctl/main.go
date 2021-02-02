package main

import (
	"log"

	"github.com/micnncim/slackctl/pkg/cmd"
	"github.com/micnncim/slackctl/pkg/cmd/post"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	c := cmd.NewCommand()

	postCmd, err := post.NewCommand()
	if err != nil {
		return err
	}

	c.AddCommand(
		postCmd,
	)

	return c.Execute()
}
