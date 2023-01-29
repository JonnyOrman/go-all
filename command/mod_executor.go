package command

type ModExecutor struct {
	modLogger  ModLogger
	cmdCreator CmdCreator
	cmdRunner  CmdRunner
}

func NewModExecutor(
	modLogger ModLogger,
	cmdCreator CmdCreator,
	cmdRunner CmdRunner,
) *ModExecutor {
	this := new(ModExecutor)

	this.modLogger = modLogger
	this.cmdCreator = cmdCreator
	this.cmdRunner = cmdRunner

	return this
}

func (this ModExecutor) Execute(modFile string, directory string) {
	this.modLogger.Log(modFile)

	goCmd := this.cmdCreator.Create(modFile, directory)

	this.cmdRunner.Run(goCmd)
}
