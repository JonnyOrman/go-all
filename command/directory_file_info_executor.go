package command

import "io/fs"

type DirectoryFileInfoExecutor struct {
	subdirectoryConstructor DirectoryConstructor
}

func NewDirectoryFileInfoExecutor(
	subdirectoryConstructor DirectoryConstructor,
) *DirectoryFileInfoExecutor {
	this := new(DirectoryFileInfoExecutor)

	this.subdirectoryConstructor = subdirectoryConstructor

	return this
}

func (this DirectoryFileInfoExecutor) Execute(
	directory string,
	directoryFileInfo fs.FileInfo,
	directoryExecutor DirectoryExecutor,
) {
	subdirectory := this.subdirectoryConstructor.Construct(directory, directoryFileInfo.Name())
	directoryExecutor.Execute(subdirectory)
}
