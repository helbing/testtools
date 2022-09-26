package layotto

import "testing"

type Runner struct{}

func New() *Runner {
	return &Runner{}
}

func (r *Runner) Install(tb testing.TB) error { return nil }

func (r *Runner) Uninstall(tb testing.TB) error { return nil }
