package command

import (
	"os"
)

type DirectoryExecutor struct {
	config                 *Config
	modFilePathConstructor *ModFilePathConstructor
	modExecutor            *ModExecutor
}

func NewDirectoryExecutor(
	config *Config,
	modFilePathConstructor *ModFilePathConstructor,
	modExecutor *ModExecutor,
) *DirectoryExecutor {
	this := new(DirectoryExecutor)

	this.config = config
	this.modFilePathConstructor = modFilePathConstructor
	this.modExecutor = modExecutor

	return this
}

func (this DirectoryExecutor) Execute(directory string) {
	modFile := this.modFilePathConstructor.Construct(directory)
	_, modFileErr := os.Stat(modFile)

	if modFileErr == nil {
		this.modExecutor.Execute(modFile, directory)
	}
}
