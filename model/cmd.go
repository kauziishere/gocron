package model

import (
	"os"
	"os/exec"
)

const viPath = "/usr/bin/vi"

func executeCommandWithArgs(args []string) error {
	cmd := &exec.Cmd{
		Path:   viPath,
		Args:   args,
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stdout,
	}

	err := cmd.Run()
	if nil != err {
		return err
	}
	return nil
}
