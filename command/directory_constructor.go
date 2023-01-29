package command

type DirectoryConstructor interface {
	Construct(directory string, subdirectory string) string
}
