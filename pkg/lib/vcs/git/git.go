package git

import (
	"fmt"
	"net/url"
	"os"
	"os/exec"
)

type Auth interface {
	Auth(cmd *GitCloneCmd) error
	// Close ends the authentication
	Close() error
}

/*
	ssh clone: user@host
	https clonse: https://user:password@host
*/
func Clone(addr string, dst string, ref string, auth Auth) error {
	cmd := GitCloneCmd{
		Address:     addr,
		Destination: dst,
		Env:         make([]string, 0),
		Args:        make([]string, 0),
	}

	if err := auth.Auth(&cmd); err != nil {
		return err
	}
	defer auth.Close()

	cmd.Args = append(cmd.Args, "--branch", ref)

	c := cmd.Cmd()
	fmt.Printf("git clone command: %s\n", c)
	return c.Run()
}

type GitCloneCmd struct {
	Address     string
	Destination string
	Env         []string
	Args        []string
}

func (c GitCloneCmd) Cmd() *exec.Cmd {
	args := []string{"clone"}
	args = append(args, c.Args...)
	args = append(args, c.Address)
	if c.Destination != "" {
		args = append(args, c.Destination)
	}
	cmd := exec.Command("git", args...)
	cmd.Env = c.Env
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd
}

// type AuthSSH struct {
// 	PrivateKey string
// 	sshKeyPath string
// }
//
// func (a *AuthSSH) Auth(cmd *GitCloneCmd) error {
// 	tmpDir := *flags.TmpDir
// 	tmpFile, err := ioutil.TempFile(tmpDir, "ssh-key-*")
// 	if err != nil {
// 		return fmt.Errorf("could not initiate ssh authentication: %v", err)
// 	}
// 	defer tmpFile.Close()
//
// 	a.sshKeyPath = tmpFile.Name()
//
// 	n, err := tmpFile.WriteString(a.PrivateKey)
// 	if n != len(a.PrivateKey) {
// 		tmpFile.Close()
// 		return fmt.Errorf("error git clone: could not write entirely the private key file: lenght: %d, written: %d", len(a.PrivateKey), n)
// 	}
// 	if err != nil {
// 		tmpFile.Close()
// 		return fmt.Errorf("error git clone: could not write private key: %v", err)
// 	}
//
// 	var envKey strings.Builder
// 	envKey.WriteString("GIT_SSH_COMMAND=ssh -i ")
// 	envKey.WriteString(tmpFile.Name())
// 	envKey.WriteString(" -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no")
//
// 	cmd.Env = append(os.Environ(), envKey.String())
//
// 	return nil
// }
//
// func (a *AuthSSH) Close() error {
// 	err := os.Remove(a.sshKeyPath)
// 	if err != nil {
// 		fmt.Printf("warning: %s: could not delete private ssh key file", a.sshKeyPath)
// 	}
// 	return nil
// }

type AuthHTTP struct {
	User, Password string
}

func (a *AuthHTTP) Auth(cmd *GitCloneCmd) error {
	addr, err := url.Parse(cmd.Address)
	if err != nil {
		return fmt.Errorf("could not parse git clone address: %s", err)
	}

	userInfo := url.UserPassword(a.User, a.Password)
	addr.User = userInfo
	cmd.Address = addr.String()
	return nil
}

func (a *AuthHTTP) Close() error {
	return nil
}

type AuthNone struct{}

func (a *AuthNone) Auth(cmd *GitCloneCmd) error {
	return nil
}

func (a *AuthNone) Close() error {
	return nil
}
