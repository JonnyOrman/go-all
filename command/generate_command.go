package command

func GenerateCommand(action string, args string) *Command {
	config := NewConfig(action, args)

	modFilePathConstructor := NewModFilePathConstructor()

	modExecutor := NewModExecutor(config)

	directoryExecutor := NewDirectoryExecutor(config, modFilePathConstructor, modExecutor)

	subdirectoryConstructor := NewSubdirectoryConstructor()

	subdirectoriesExecutor := NewSubdirectoriesExecutor(subdirectoryConstructor)

	return NewCommand(config, directoryExecutor, subdirectoriesExecutor)
}
