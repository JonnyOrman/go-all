package command

type Command struct {
	config                 *Config
	directoryExecutor      *DirectoryExecutor
	subDirectoriesExecutor *SubdirectoriesExecutor
}

func NewCommand(
	config *Config,
	directoryExecutor *DirectoryExecutor,
	subDirectoriesExecutor *SubdirectoriesExecutor,
) *Command {
	this := new(Command)

	this.config = config
	this.directoryExecutor = directoryExecutor
	this.subDirectoriesExecutor = subDirectoriesExecutor

	return this
}

func (this Command) ExecuteAllIn(directory string) {
	this.directoryExecutor.Execute(directory)

	this.subDirectoriesExecutor.Execute(directory, this.directoryExecutor)
}
