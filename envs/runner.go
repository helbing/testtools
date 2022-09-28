package envs

import (
	"testing"
)

//go:generate mockgen -source=runner.go -destination=./mock/runner_mock.go -package=mock

// Runner is a executor to up & down the environment required for testing.
type Runner interface {
	// Up runner to setup environment.
	Up(tb testing.TB) error

	// Down runner to teardown envirment.
	Down(tb testing.TB) error
}

// Runners is alias of array of Runner and support functions to operation
// easily.
//
//	runners := New()
//	runners = runners.Add(runner1, runner2)
//	err := runners.Up(t)
//	assert.Nil(t, err)
//	err = runners.Down(t)
//	assert.Nil(t, err)
type Runners []Runner

func New() Runners {
	return make(Runners, 0, 32)
}

// Add non-nil runners.
func (rs Runners) Add(runners ...Runner) Runners {
	for _, runner := range runners {
		if runner != nil {
			rs = append(rs, runners...)
		}
	}
	return rs
}

func (rs Runners) Up(tb testing.TB) error {
	for _, runner := range rs {
		if err := runner.Up(tb); err != nil {
			return err
		}
	}
	return nil
}

func (rs Runners) Down(tb testing.TB) error {
	for _, runner := range rs {
		if err := runner.Down(tb); err != nil {
			return err
		}
	}
	return nil
}
