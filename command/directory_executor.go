package command

type DirectoryExecutor interface {
	Execute(directory string) bool
}
