package command

type SubdirectoryConstructor struct{}

func NewSubdirectoryConstructor() *SubdirectoryConstructor {
	this := new(SubdirectoryConstructor)

	return this
}

func (this SubdirectoryConstructor) Construct(directory string, subdirectory string) string {
	return directory + "/" + subdirectory
}
