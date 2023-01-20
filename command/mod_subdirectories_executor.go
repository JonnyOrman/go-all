package command

import "io/ioutil"

type ModSubdirectoriesExecutor struct {
	subdirectoryConstructor *SubdirectoryConstructor
}

func NewModSubdirectoriesExecutor(
	subdirectoryConstructor *SubdirectoryConstructor,
) *ModSubdirectoriesExecutor {
	this := new(ModSubdirectoriesExecutor)

	this.subdirectoryConstructor = subdirectoryConstructor

	return this
}

func (this ModSubdirectoriesExecutor) Execute(directory string, directoryExecutor DirectoryExecutor) bool {
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		panic(err)
	}

	for _, fileInfo := range files {
		if fileInfo.IsDir() {
			subdirectory := this.subdirectoryConstructor.Construct(directory, fileInfo.Name())
			directoryExecutor.Execute(subdirectory)
		}
	}

	return true
}
