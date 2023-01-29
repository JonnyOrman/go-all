package command

type ModSubdirectoriesExecutor struct {
	fileInfoProvider          FileInfoProvider
	directoryFileInfoExecutor FileInfoExecutor
}

func NewModSubdirectoriesExecutor(
	fileInfoProvider FileInfoProvider,
	directoryFileInfoExecutor FileInfoExecutor,
) *ModSubdirectoriesExecutor {
	this := new(ModSubdirectoriesExecutor)

	this.fileInfoProvider = fileInfoProvider
	this.directoryFileInfoExecutor = directoryFileInfoExecutor

	return this
}

func (this ModSubdirectoriesExecutor) Execute(
	directory string,
	directoryExecutor DirectoryExecutor,
) bool {
	files := this.fileInfoProvider.Get(directory)

	for _, fileInfo := range files {
		if fileInfo.IsDir() {
			this.directoryFileInfoExecutor.Execute(
				directory,
				fileInfo,
				directoryExecutor,
			)
		}
	}

	return true
}
