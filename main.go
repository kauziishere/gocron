package main

import (
	"flag"
	"fmt"

	"github.com/kauziishere/gocron/model"
	"github.com/rs/zerolog"
)

func processFlags() int {
	debug := flag.Bool("debug", false, "Allow Debug")
	editFlagSet := flag.Bool("e", false, "Allow user to edit gocrontab")
	restartFlagSet := flag.Bool("x", false, "Allow user to restart gocrontab")
	readFlagSet := flag.Bool("r", false, "Allow user read only access to crontab")
	flag.Parse()

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if *debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	if *editFlagSet {
		return model.EditFlag
	}
	if *restartFlagSet {
		return model.RestartFlag
	}
	if *readFlagSet {
		return model.ReadFlag
	}
	return model.NoFlagSet
}

func main() {
	flagVal := processFlags()
	if model.NoFlagSet == flagVal {
		fmt.Print("No Flag is set\n")
		fmt.Print("Usage of gocron:\n")
		flag.PrintDefaults()
		return
	}
	model.Execute(flagVal)
}
