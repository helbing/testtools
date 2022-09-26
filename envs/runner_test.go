package envs

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/helbing/testtools/envs/mock"
)

func TestRunners_Add(t *testing.T) {
	t.Run("Add nil runner", func(t *testing.T) {
		runners := New()
		runners = runners.Add(nil)
		assert.Equal(t, 0, len(runners))
	})

	t.Run("Add runner success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mock := mock.NewMockRunner(ctrl)

		runners := New()
		runners = runners.Add(mock)
		assert.Equal(t, 1, len(runners))
	})
}

func TestRunners_Install(t *testing.T) {
	t.Run("Install failed", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockErr := errors.New("runner install error")

		mock := mock.NewMockRunner(ctrl)
		mock.EXPECT().Install(gomock.Any()).Return(mockErr)

		err := New().Add(mock).Install(t)
		assert.Equal(t, mockErr, err)
	})

	t.Run("Install success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mock := mock.NewMockRunner(ctrl)
		mock.EXPECT().Install(gomock.Any()).Return(nil)

		err := New().Add(mock).Install(t)
		assert.Nil(t, err)
	})
}

func TestRunners_Uninstall(t *testing.T) {
	t.Run("Uninstall failed", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockErr := errors.New("runner uninstall error")

		mock := mock.NewMockRunner(ctrl)
		mock.EXPECT().Uninstall(gomock.Any()).Return(mockErr)

		err := New().Add(mock).Uninstall(t)
		assert.Equal(t, mockErr, err)
	})

	t.Run("Uninstall success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mock := mock.NewMockRunner(ctrl)
		mock.EXPECT().Uninstall(gomock.Any()).Return(nil)

		err := New().Add(mock).Uninstall(t)
		assert.Nil(t, err)
	})
}
