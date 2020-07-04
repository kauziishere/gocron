package main

import (
	"flag"
	"fmt"
)

const (
	EditFlag    = iota
	RestartFlag = iota
	ReadFlag    = iota
	NoFlagSet   = iota
)

func processFlags() int {
	var editFlagSet = flag.Bool("e", false, "Allow user to edit gocrontab")
	var restartFlagSet = flag.Bool("x", false, "Allow user to restart gocrontab")
	var readFlagSet = flag.Bool("r", false, "Allow user read only access to crontab")
	flag.Parse()
	if *editFlagSet {
		return EditFlag
	}
	if *restartFlagSet {
		return RestartFlag
	}
	if *readFlagSet {
		return ReadFlag
	}
	return NoFlagSet
}

func main() {
	flagVal := processFlags()
	if NoFlagSet == flagVal {
		fmt.Print("No Flag is set\n")
		fmt.Print("Usage of gocron:\n")
		flag.PrintDefaults()
		return
	}
}
