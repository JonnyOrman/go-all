package command

import "io/ioutil"

type SubdirectoriesExecutor struct {
	subdirectoryConstructor *SubdirectoryConstructor
}

func NewSubdirectoriesExecutor(
	subdirectoryConstructor *SubdirectoryConstructor,
) *SubdirectoriesExecutor {
	this := new(SubdirectoriesExecutor)

	this.subdirectoryConstructor = subdirectoryConstructor

	return this
}

func (this SubdirectoriesExecutor) Execute(directory string, directoryExecutor *DirectoryExecutor) {
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
}
