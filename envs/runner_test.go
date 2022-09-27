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

func TestRunners_Up(t *testing.T) {
	t.Run("Up failed", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockErr := errors.New("runner error")

		mock := mock.NewMockRunner(ctrl)
		mock.EXPECT().Up(gomock.Any()).Return(mockErr)

		err := New().Add(mock).Up(t)
		assert.Equal(t, mockErr, err)
	})

	t.Run("Up success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mock := mock.NewMockRunner(ctrl)
		mock.EXPECT().Up(gomock.Any()).Return(nil)

		err := New().Add(mock).Up(t)
		assert.Nil(t, err)
	})
}

func TestRunners_Down(t *testing.T) {
	t.Run("Down failed", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockErr := errors.New("runner error")

		mock := mock.NewMockRunner(ctrl)
		mock.EXPECT().Down(gomock.Any()).Return(mockErr)

		err := New().Add(mock).Down(t)
		assert.Equal(t, mockErr, err)
	})

	t.Run("Down success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mock := mock.NewMockRunner(ctrl)
		mock.EXPECT().Down(gomock.Any()).Return(nil)

		err := New().Add(mock).Down(t)
		assert.Nil(t, err)
	})
}
