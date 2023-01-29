package command

import "io/fs"

type FileInfoProvider interface {
	Get(directory string) []fs.FileInfo
}
