package runner

import "context"

type Config struct {
	Type    string
	Command string
	Env     []string
	Dir     string
	Conf    interface{}
}

type Runner interface {
	Run(ctx context.Context) error
}
