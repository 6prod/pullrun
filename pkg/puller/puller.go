package puller

import "context"

type Config struct {
	Type string
	Conf interface{}
}

type Puller interface {
	Pull(ctx context.Context) error
}
