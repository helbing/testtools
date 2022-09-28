// Package dockercompose is a runner to up & down environment through
// Docker-Compose.
package dockercompose

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/docker"
)

type Runner struct {
	config   *Config
	filename string
}

// New docker-compose runner with filename of docker-compose.yml and options.
//
//	runner := New("./docker-compose.yml")
func New(filename string, opts ...Option) *Runner {
	config := &Config{
		envVas: make(map[string]string),
	}
	for _, opt := range opts {
		opt(config)
	}

	return &Runner{
		config:   config,
		filename: filename,
	}
}

func (r *Runner) Up(tb testing.TB) error {
	options := &docker.Options{
		WorkingDir: r.config.workingDir,
		EnvVars:    r.config.envVas,
	}

	output, err := docker.RunDockerComposeE(tb, options,
		"-f", r.filename,
		"up", "-d",
		"--remove-orphans")

	tb.Log(output)

	return err
}

func (r *Runner) Down(tb testing.TB) error {
	options := &docker.Options{
		WorkingDir: r.config.workingDir,
		EnvVars:    r.config.envVas,
	}

	output, err := docker.RunDockerComposeE(tb, options,
		"-f", r.filename,
		"down", "-v")

	tb.Log(output)

	return err
}
