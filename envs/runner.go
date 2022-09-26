package envs

import (
	"testing"
)

//go:generate mockgen -source=runner.go -destination=./mock/runner_mock.go -package=mock

// Runner is to setup the environment required for the test.
type Runner interface {
	// Install runner.
	Install(tb testing.TB) error

	// Uninstall runner.
	Uninstall(tb testing.TB) error
}

// Runners is array of Runner.
type Runners []Runner

func New() Runners {
	return make(Runners, 0, 32)
}

// Add runner.
func (rs Runners) Add(runners ...Runner) Runners {
	for _, runner := range runners {
		if runner != nil {
			rs = append(rs, runners...)
		}
	}
	return rs
}

// Install all runners.
func (rs Runners) Install(tb testing.TB) error {
	for _, runner := range rs {
		if err := runner.Install(tb); err != nil {
			return err
		}
	}
	return nil
}

// Uninstall all runners.
func (rs Runners) Uninstall(tb testing.TB) error {
	for _, runner := range rs {
		if err := runner.Uninstall(tb); err != nil {
			return err
		}
	}
	return nil
}
