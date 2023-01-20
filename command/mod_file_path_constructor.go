package command

type ModFilePathConstructor struct{}

func NewModFilePathConstructor() *ModFilePathConstructor {
	this := new(ModFilePathConstructor)

	return this
}

func (this ModFilePathConstructor) Construct(directory string) string {
	return directory + "/go.mod"
}
