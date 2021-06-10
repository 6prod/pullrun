package plugins

import (
	"fmt"

	"github.com/6prod/pullrun/pkg/puller"
	"github.com/6prod/pullrun/pkg/puller/plugins/git"
	"github.com/6prod/pullrun/pkg/puller/plugins/http"
)

var register = map[string]func(plugin puller.Config) (puller.Puller, error){
	"git":  git.New,
	"http": http.New,
}

func New(plugin puller.Config) (puller.Puller, error) {
	fn, ok := register[plugin.Type]
	if !ok {
		return nil, fmt.Errorf("%s; unknown puller", plugin.Type)
	}
	return fn(plugin)
}
