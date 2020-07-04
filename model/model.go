package model

import (
	"log"
	"os"
)

const (
	EditFlag    = iota
	RestartFlag = iota
	ReadFlag    = iota
	NoFlagSet   = iota
)

const configFile = ".gocrontab"

func checkAndCreateFile() error {
	var err error
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		_, err = os.Create(configFile)
	}
	return err
}

func viewConfigFile(readOnlyFlag bool) {
	err := checkAndCreateFile()
	if nil != err {
		log.Fatalf("Failed to read file %s: %s\n", configFile, err.Error())
	}

	if readOnlyFlag {
		err = executeCommandWithArgs([]string{"vi", "-R", configFile})
	} else {
		err = executeCommandWithArgs([]string{"vi", configFile})
	}

	if nil != err {
		log.Fatal(err.Error())
	}
}

func Execute(cmdToExec int) {
	switch cmdToExec {
	case ReadFlag:
		viewConfigFile(true)
		break
	case EditFlag:
		viewConfigFile(false)
		break
	case RestartFlag:
		break
	default:
		log.Fatalf("Invalid metric executed")
	}
}
