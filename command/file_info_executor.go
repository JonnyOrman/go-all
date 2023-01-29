package command

import "io/fs"

type FileInfoExecutor interface {
	Execute(
		directory string,
		directoryFileInfo fs.FileInfo,
		directoryExecutor DirectoryExecutor,
	)
}
