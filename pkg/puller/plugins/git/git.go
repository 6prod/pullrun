package git

import (
	"context"
	"fmt"

	"github.com/6prod/pullrun/pkg/lib/vcs/git"
	"github.com/6prod/pullrun/pkg/puller"
	"github.com/mitchellh/mapstructure"
)

type Git struct {
	Address string
	Ref     string
	Dir     string
	Auth    struct {
		Type   string
		Config interface{}
	}
}

func New(plugin puller.Config) (puller.Puller, error) {
	var git Git
	if err := mapstructure.Decode(plugin.Conf, &git); err != nil {
		return nil, err
	}

	if err := setDefault(&git); err != nil {
		return nil, err
	}

	return git, nil
}

func setDefault(g *Git) error {
	if g.Ref == "" {
		g.Ref = "main"
	}
	return nil
}

func (g Git) Pull(ctx context.Context) error {
	auth, err := Auth(g.Auth.Type, g.Auth.Config)
	if err != nil {
		return err
	}

	return git.Clone(g.Address, g.Dir, g.Ref, auth)
}

type AuthAccessToken struct {
	Username string
	Token    string
}

func NewAuthAccessToken(iconf interface{}) (git.Auth, error) {
	var c AuthAccessToken
	if err := mapstructure.Decode(iconf, &c); err != nil {
		return nil, err
	}

	auth := git.AuthHTTP{User: c.Username, Password: c.Token}
	return &auth, nil
}

type AuthSSH struct {
	PublicKey  string
	PrivateKey string
}

func NewAuthSSH(iconf interface{}) (git.Auth, error) {
	var c AuthSSH
	if err := mapstructure.Decode(iconf, &c); err != nil {
		return nil, err
	}

	//auth := git.AuthSSH{PrivateKey: c.PrivateKey}
	//return &auth, nil
	return nil, nil
}

type AuthDeployKey struct {
	PublicKey  string
	PrivateKey string
}

func NewAuthDeployKey(iconf interface{}) (git.Auth, error) {
	var c AuthDeployKey
	if err := mapstructure.Decode(iconf, &c); err != nil {
		return nil, err
	}
	//auth := git.AuthSSH{PrivateKey: c.PrivateKey}
	//return &auth, nil
	return nil, nil
}

func Auth(_type string, iconf interface{}) (git.Auth, error) {
	switch _type {
	case "access_token":
		return NewAuthAccessToken(iconf)
	case "ssh":
		return NewAuthSSH(iconf)
	case "deploy_key":
		return NewAuthDeployKey(iconf)
	case "":
		return &git.AuthNone{}, nil
	}
	return nil, fmt.Errorf("%s: unknown authentication type", _type)
}
