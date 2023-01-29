package command

import "os/exec"

type CmdRunner interface {
	Run(cmd *exec.Cmd)
}
