package http

import (
	"context"
	"errors"

	"github.com/6prod/pullrun/pkg/puller"
)

type HTTP struct {
	Address string
	Dir     string
	Auth    struct {
		Type   string
		Config interface{}
	}
}

func New(plugin puller.Config) (puller.Puller, error) {
	return nil, errors.New("http downloader not implemented yet")
}

func (http HTTP) Pull(ctx context.Context) error {
	return nil
}

type AuthBasic struct {
	Username string
	Password string
}
