package command

type FilePathConstructor interface {
	Construct(directory string) string
}
