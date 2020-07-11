package model

import (
	"os"
	"os/exec"
)

func executeCommand(args []string) error {
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stdout
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if nil != err {
		return err
	}
	return nil
}
