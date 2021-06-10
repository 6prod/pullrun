package config

import (
	"encoding/json"
	"io"
	"os"

	"github.com/6prod/pullrun/pkg/puller"
	"github.com/6prod/pullrun/pkg/runner"
)

type Conf struct {
	Pull []puller.Config
	Run  []runner.Config
}

func LoadFromFile(name string) (Conf, error) {
	if name == "-" {
		return Load(os.Stdin)
	}

	file, err := os.Open(name)
	if err != nil {
		return Conf{}, err
	}
	defer file.Close()
	return Load(file)
}

func Load(r io.Reader) (Conf, error) {
	var conf Conf
	decoder := json.NewDecoder(r)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&conf); err != nil {
		return conf, err
	}
	if err := setDefault(&conf); err != nil {
		return conf, err
	}
	return conf, nil
}

const (
	defaultRunner = "command"
)

func setDefault(conf *Conf) error {
	for i := range conf.Run {
		if conf.Run[i].Type == "" {
			conf.Run[i].Type = defaultRunner
		}
	}
	return nil
}
