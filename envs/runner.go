package envs

import (
	"testing"
)

//go:generate mockgen -source=runner.go -destination=./mock/runner_mock.go -package=mock

// Runner is to setup the environment required for testing.
type Runner interface {
	// Up runner.
	Up(tb testing.TB) error

	// Down runner.
	Down(tb testing.TB) error
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

// Up all runners.
func (rs Runners) Up(tb testing.TB) error {
	for _, runner := range rs {
		if err := runner.Up(tb); err != nil {
			return err
		}
	}
	return nil
}

// Down all runners.
func (rs Runners) Down(tb testing.TB) error {
	for _, runner := range rs {
		if err := runner.Down(tb); err != nil {
			return err
		}
	}
	return nil
}
