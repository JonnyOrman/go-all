package command

import (
	"io/ioutil"

	"golang.org/x/mod/modfile"
)

type FileModNameProvider struct{}

func NewFileModNameProvider() *FileModNameProvider {
	this := new(FileModNameProvider)

	return this
}

func (this FileModNameProvider) Get(modFile string) string {
	modFileBytes, err := ioutil.ReadFile(modFile)
	if err != nil {
		panic(err)
	}

	return modfile.ModulePath(modFileBytes)
}
