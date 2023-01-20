package command

import (
	"testing"

	"github.com/stretchr/testify/mock"
)

type DirectoryExecutorMock struct {
	mock.Mock
}

func (this DirectoryExecutorMock) Execute(directory string) bool {
	args := this.Called(directory)
	return args.Bool(0)
}

type SubdirectoriesExecutorMock struct {
	mock.Mock
}

func (this SubdirectoriesExecutorMock) Execute(directory string, directoryExecutor DirectoryExecutor) bool {
	args := this.Called(directory, directoryExecutor)
	return args.Bool(0)
}

func TestExecuteAllIn(t *testing.T) {
	directory := "some/directory"

	directoryExecutor := new(DirectoryExecutorMock)
	directoryExecutor.On("Execute", directory).Return(true)

	subdirectoriesExecutor := new(SubdirectoriesExecutorMock)
	subdirectoriesExecutor.On("Execute", directory, directoryExecutor).Return(true)

	sut := NewCommand(directoryExecutor, subdirectoriesExecutor)

	sut.ExecuteAllIn(directory)
}
