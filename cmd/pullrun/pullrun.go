package main

import (
	"bytes"
	"context"
	"errors"
	"flag"
	"io/ioutil"
	"log"
	"os"

	"github.com/6prod/pullrun/pkg/config"
	"github.com/6prod/pullrun/pkg/flags"
	pullerplugin "github.com/6prod/pullrun/pkg/puller/plugins"
	runnerplugin "github.com/6prod/pullrun/pkg/runner/plugins"
)

func main() {
	flag.Usage = flags.Usage
	flags.Parse()

	logger := log.New(os.Stderr, "", log.LstdFlags|log.LUTC)

	flagConfig := *flags.Config
	dir := *flags.Dir
	dirPerm := *flags.DirPerm
	ctx := context.Background()

	if flagConfig == "" {
		logger.Fatalf("flag -%s must be set\n", flags.FlagConfig)
	}

	conf, err := LoadConfig(flagConfig)
	if err != nil {
		logger.Fatal(err)
	}

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, os.FileMode(dirPerm)); err != nil {
			logger.Fatal(err)
		}
	}

	if err := os.Chdir(dir); err != nil {
		logger.Fatal(err)
	}

	download, err := IsDirEmpty(dir)
	if err != nil {
		log.Fatal(err)
	}

	if download {
		for _, pullerConf := range conf.Pull {
			puller, err := pullerplugin.New(pullerConf)
			if err != nil {
				logger.Fatal(err)
			}
			if err := puller.Pull(ctx); err != nil {
				logger.Fatal(err)
			}
		}
	}

	for _, runnerConf := range conf.Run {
		runner, err := runnerplugin.New(runnerConf)
		if err != nil {
			logger.Fatal(err)
		}
		if err := runner.Run(ctx); err != nil {
			logger.Fatal(err)
		}
	}
}

func IsDirEmpty(dir string) (bool, error) {
	if _, err := os.Stat(dir); errors.Is(err, os.ErrNotExist) {
		return true, nil
	}

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return false, err
	}

	return len(files) == 0, nil
}

func LoadConfig(flagConfig string) (config.Conf, error) {
	if flagConfig[0] == '{' {
		return config.Load(bytes.NewBufferString(flagConfig))
	}

	if _, err := os.Stat(flagConfig); os.IsNotExist(err) {
		return config.Conf{}, err
	}

	return config.LoadFromFile(flagConfig)
}
