package dockercompose

type Config struct {
	workingDir string

	// envVas is environment variables.
	envVas map[string]string
}

type Option func(*Config)

func WithWorkingDir(dir string) Option {
	return func(c *Config) { c.workingDir = dir }
}

func WithEnvVars(envVas map[string]string) Option {
	return func(c *Config) { c.envVas = envVas }
}
