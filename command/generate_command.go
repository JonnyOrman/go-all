package command

func GenerateCommand(action string, args string) *Command {
	config := NewConfig(action, args)

	modFilePathConstructor := NewModFilePathConstructor()

	fileModNameProvider := NewFileModNameProvider()

	modLogger := NewActionModLogger(
		fileModNameProvider,
		config,
	)

	modCmdCreator := NewModCmdCreator(
		config,
	)

	modCmdRunner := NewModCmdRunner()

	modExecutor := NewModExecutor(
		modLogger,
		modCmdCreator,
		modCmdRunner,
	)

	directoryExecutor := NewModDirectoryExecutor(
		config,
		modFilePathConstructor,
		modExecutor,
	)

	ioutilFileInfoProvider := NewIoutilFileInfoProvider()

	subdirectoryConstructor := NewSubdirectoryConstructor()

	directoryFileInfoExecutor := NewDirectoryFileInfoExecutor(
		subdirectoryConstructor,
	)

	subdirectoriesExecutor := NewModSubdirectoriesExecutor(
		ioutilFileInfoProvider,
		directoryFileInfoExecutor,
	)

	return NewCommand(
		directoryExecutor,
		subdirectoriesExecutor,
	)
}
