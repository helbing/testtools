package flow

import (
	"testing"

	"github.com/helbing/testtools/envs"
)

type testcase struct {
	name string
	fn   func(t *testing.T)
}

type Client struct {
	// runners is to setup the environment.
	runners envs.Runners

	// testcase related.
	testcases  []testcase
	setupFn    func(t *testing.T, name string)
	teardownFn func(t *testing.T, name string)
}

func New(runners ...envs.Runner) *Client {
	return &Client{
		runners:    runners,
		testcases:  make([]testcase, 0, 32),
		setupFn:    func(t *testing.T, name string) {},
		teardownFn: func(t *testing.T, name string) {},
	}
}

// Case of test.
func (c *Client) Case(name string, fn func(t *testing.T)) *Client {
	if fn != nil {
		c.testcases = append(c.testcases, testcase{name: name, fn: fn})
	}
	return c
}

// Setup for every case. You can determine which case is currently through
// name.
func (c *Client) Setup(fn func(t *testing.T, name string)) *Client {
	if fn != nil {
		c.setupFn = fn
	}
	return c
}

// Teardown for every case. You can determine which case is currently through
// name.
func (c *Client) Teardown(fn func(t *testing.T, name string)) *Client {
	if fn != nil {
		c.teardownFn = fn
	}
	return c
}

// Run all testcases.
func (c *Client) Run(t *testing.T) {
	if err := c.runE(t); err != nil {
		t.Errorf("[unittest] Run caught error: %+v", err)
	}
}

// runE is run all testcases that will return err.
func (c *Client) runE(t *testing.T) (err error) {
	if err = c.runners.Up(t); err != nil {
		return err
	}

	for _, testcase := range c.testcases {
		c.setupFn(t, testcase.name)
		t.Run(testcase.name, testcase.fn)
		c.teardownFn(t, testcase.name)
	}

	return c.runners.Down(t)
}
