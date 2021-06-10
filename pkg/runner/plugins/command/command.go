package command

import (
	"context"
	"os"
	"os/exec"
	"strings"

	"github.com/6prod/pullrun/pkg/runner"
)

type Command runner.Config

func New(plugin runner.Config) (runner.Runner, error) {
	return Command(plugin), nil
}

func (c Command) Run(ctx context.Context) error {
	fields := strings.Fields(c.Command)
	cmd := exec.CommandContext(ctx, fields[0], fields[1:]...)
	cmd.Env = c.Env
	cmd.Dir = c.Dir
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
