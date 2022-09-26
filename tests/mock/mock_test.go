package mock

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/helbing/testtools/envs/mock"
	"github.com/helbing/testtools/flow"
)

func TestMock(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	runner := mock.NewMockRunner(ctrl)
	runner.EXPECT().Install(gomock.Any()).Return(nil)
	runner.EXPECT().Uninstall(gomock.Any()).Return(nil)

	flow.New(runner).
		Case("test mock success", func(t *testing.T) {
			assert.Nil(t, nil)
		}).
		Run(t)
}
