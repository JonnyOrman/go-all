package command

import (
	"io/fs"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
)

type DirectoryConstructorMock struct {
	mock.Mock
}

func (this DirectoryConstructorMock) Construct(directory string, subdirectory string) string {
	args := this.Called(directory, subdirectory)
	return args.Get(0).(string)
}

type FileInfoMock struct {
	mock.Mock
}

func (this FileInfoMock) Name() string {
	args := this.Called()
	return args.Get(0).(string)
}

func (this FileInfoMock) Size() int64 {
	args := this.Called()
	return args.Get(0).(int64)
}

func (this FileInfoMock) Mode() fs.FileMode {
	args := this.Called()
	return args.Get(0).(fs.FileMode)
}

func (this FileInfoMock) ModTime() time.Time {
	args := this.Called()
	return args.Get(0).(time.Time)
}

func (this FileInfoMock) IsDir() bool {
	args := this.Called()
	return args.Get(0).(bool)
}

func (this FileInfoMock) Sys() any {
	args := this.Called()
	return args.Get(0).(any)
}

func TestDirectoryFileInfoExecutor(t *testing.T) {
	directory := "/my-project"
	subdirectory := "my-module"

	constructedDirectory := "/my-project/my-module"

	directoryFileInfo := new(FileInfoMock)
	directoryFileInfo.On("Name").Return(subdirectory)

	directoryConstructor := new(DirectoryConstructorMock)
	directoryConstructor.On("Construct", directory, subdirectory).Return(constructedDirectory)

	directoryExecutor := new(DirectoryExecutorMock)
	directoryExecutor.On("Execute", constructedDirectory).Return(true)

	directoryFileInfoExecutor := NewDirectoryFileInfoExecutor(
		directoryConstructor,
	)

	directoryFileInfoExecutor.Execute(
		directory,
		directoryFileInfo,
		directoryExecutor,
	)
}
