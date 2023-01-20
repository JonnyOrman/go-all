package command

type SubdirectoriesExecutor interface {
	Execute(directory string, directoryExecutor DirectoryExecutor) bool
}
