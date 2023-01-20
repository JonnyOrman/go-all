package command

import (
	"os"
)

type ModDirectoryExecutor struct {
	config                 *Config
	modFilePathConstructor FilePathConstructor
	modExecutor            *ModExecutor
}

func NewModDirectoryExecutor(
	config *Config,
	modFilePathConstructor FilePathConstructor,
	modExecutor *ModExecutor,
) *ModDirectoryExecutor {
	this := new(ModDirectoryExecutor)

	this.config = config
	this.modFilePathConstructor = modFilePathConstructor
	this.modExecutor = modExecutor

	return this
}

func (this ModDirectoryExecutor) Execute(directory string) bool {
	modFile := this.modFilePathConstructor.Construct(directory)
	_, modFileErr := os.Stat(modFile)

	if modFileErr == nil {
		this.modExecutor.Execute(modFile, directory)
	}

	return modFileErr == nil
}
