# testtools

[![](https://pkg.go.dev/badge/github.com/helbing/testtools)](https://pkg.go.dev/github.com/helbing/testtools)
[![](https://goreportcard.com/badge/github.com/helbing/testtools)](https://goreportcard.com/report/github.com/helbing/testtools)
[![](https://github.com/helbing/testtools/workflows/ci-workflows/badge.svg)](https://github.com/helbing/testtools/actions)
[![](https://codecov.io/gh/helbing/testtools/branch/main/graph/badge.svg?token=1COJEOQ4QQ)](https://codecov.io/gh/helbing/testtools)

The package provides some test tools in Golang programming.

- flow: a tool to help you autoload env runners and run test cases.
- envs: the set of any envirments you want up & down when testing.
  - [x] Docker-Compose
  - [ ] GCP
  - [ ] AWS
  - [ ] Azure
  - [ ] Alibaba Cloud

## Usage

```go
func TestDemo(t *testing.T) {
    // Initial runner
    runner := dockercompose.New("./docker-compose.yml")
    // Run testcases
    flow.New(runner).
        Case("testcase1", func(t *testing.T) {
            // TODO
        }).
        Case("testcase2", func(t *testing.T) {
            // TODO
        }).
        Run(t)
}
```

## Develop your env runner

You just need to implement the Runner interface.

```go
// Runner is to setup the environment required for testing.
type Runner interface {
    // Up runner.
    Up(tb testing.TB) error

    // Down runner.
    Down(tb testing.TB) error
}
```

## Thanks

- [terratest](https://github.com/gruntwork-io/terratest)
