package command

import "os/exec"

type CmdCreator interface {
	Create(modFile string, directory string) *exec.Cmd
}
