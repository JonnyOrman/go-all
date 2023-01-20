package command

type Command struct {
	directoryExecutor      DirectoryExecutor
	subDirectoriesExecutor SubdirectoriesExecutor
}

func NewCommand(
	directoryExecutor DirectoryExecutor,
	subDirectoriesExecutor SubdirectoriesExecutor,
) *Command {
	this := new(Command)

	this.directoryExecutor = directoryExecutor
	this.subDirectoriesExecutor = subDirectoriesExecutor

	return this
}

func (this Command) ExecuteAllIn(directory string) {
	ok := this.directoryExecutor.Execute(directory)

	if ok {
		this.subDirectoriesExecutor.Execute(directory, this.directoryExecutor)
	}
}
