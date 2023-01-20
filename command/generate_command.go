package command

func GenerateCommand(action string, args string) *Command {
	config := NewConfig(action, args)

	modFilePathConstructor := NewModFilePathConstructor()

	modExecutor := NewModExecutor(config)

	directoryExecutor := NewModDirectoryExecutor(config, modFilePathConstructor, modExecutor)

	subdirectoryConstructor := NewSubdirectoryConstructor()

	subdirectoriesExecutor := NewModSubdirectoriesExecutor(subdirectoryConstructor)

	return NewCommand(directoryExecutor, subdirectoriesExecutor)
}
