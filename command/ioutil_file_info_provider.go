package command

import (
	"io/fs"
	"io/ioutil"
)

type IoutilFileInfoProvider struct{}

func NewIoutilFileInfoProvider() *IoutilFileInfoProvider {
	this := new(IoutilFileInfoProvider)

	return this
}

func (this IoutilFileInfoProvider) Get(directory string) []fs.FileInfo {
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		panic(err)
	}

	return files
}
