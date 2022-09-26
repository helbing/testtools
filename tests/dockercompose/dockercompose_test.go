package dockercompose

import (
	"net/http"
	"testing"

	http_helper "github.com/gruntwork-io/terratest/modules/http-helper"
	"github.com/stretchr/testify/assert"

	"github.com/helbing/testtools/envs/dockercompose"
	"github.com/helbing/testtools/flow"
)

func TestDockerCompose(t *testing.T) {
	runner := dockercompose.New("./docker-compose.yml")
	flow.New(runner).
		Case("CURL get success", func(t *testing.T) {
			url := "http://127.0.0.1:8080"
			statusCode, body := http_helper.HttpGet(t, url, nil)
			assert.Equal(t, http.StatusOK, statusCode)
			assert.Equal(t, "hello", body)
		}).
		Run(t)
}
