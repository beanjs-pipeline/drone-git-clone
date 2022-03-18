package drone_git_clone

import (
	"errors"
	"net/url"
	"os"
	"os/exec"
	"path"
)

type Plugin struct {
	Remote   string
	Cache    string
	Branch   string
	Username string
	Password string
}

func (p *Plugin) check() error {
	if p.Remote == "" {
		return errors.New("remote can not nil")
	}

	return nil
}

func (p *Plugin) run(c *exec.Cmd) error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	cacheDir := path.Join(cwd, p.Cache)
	if err := os.MkdirAll(cacheDir, 0700); err != nil {
		return err
	}

	c.Env = os.Environ()
	c.Stderr = os.Stderr
	c.Stdin = os.Stdin
	c.Dir = cacheDir

	return c.Run()
}

func (p *Plugin) Exec() error {
	if err := p.check(); err != nil {
		return err
	}

	u, err := url.Parse(p.Remote)

	if err != nil {
		return err
	}

	if p.Username != "" && p.Password != "" {
		u.User = url.UserPassword(p.Username, p.Password)
	} else if p.Username != "" {
		u.User = url.User(p.Username)
	}

	cloneCmd := exec.Command("git", "clone", u.String(), "-b", p.Branch)

	return p.run(cloneCmd)
}
