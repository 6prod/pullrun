package plugins

import (
	"fmt"

	"github.com/6prod/pullrun/pkg/runner"
	"github.com/6prod/pullrun/pkg/runner/plugins/command"
)

var register = map[string]func(plugin runner.Config) (runner.Runner, error){
	"command": command.New,
}

func New(plugin runner.Config) (runner.Runner, error) {
	fn, ok := register[plugin.Type]
	if !ok {
		return nil, fmt.Errorf("%s; unknown runner", plugin.Type)
	}
	return fn(plugin)
}
