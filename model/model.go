package model

import (
	"bufio"
	"os"
	"strings"
	"time"

	"github.com/kauziishere/gocron/timer"
	"github.com/rs/zerolog/log"
)

const (
	EditFlag    = iota
	RestartFlag = iota
	ReadFlag    = iota
	NoFlagSet   = iota
)

const cConfigFile = ".gocrontab"

func checkAndCreateFile() error {
	var err error
	if _, err := os.Stat(cConfigFile); os.IsNotExist(err) {
		_, err = os.Create(cConfigFile)
	}
	return err
}

func viewConfigFile(readOnlyFlag bool) {
	err := checkAndCreateFile()
	if nil != err {
		log.Fatal().Msgf("[FATAL] Failed to read file %s: %s\n", cConfigFile, err.Error())
	}

	if readOnlyFlag {
		err = executeCommand([]string{"vi", "-R", cConfigFile})
	} else {
		err = executeCommand([]string{"vi", cConfigFile})
	}

	if nil != err {
		log.Fatal().Msgf("[FATAL] Failed to execute %s\n", err.Error())
	}
}

func restartCron() {
	var err error
	file, err := os.Open(cConfigFile)
	if nil != err {
		log.Fatal().Msgf("[FATAL] %s\n", err.Error())
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		cmdStmt := scanner.Text()
		if "" == cmdStmt {
			continue
		}
		go func(cmdStmt string) {
			for {
				cmdArray := strings.Split(cmdStmt, " ")
				timeArray := cmdArray[:5]
				cmdToExec := cmdArray[5:]

				sleepTimens := timer.GetSleepTime(timeArray)
				log.Debug().Msgf("Sleep time: %f s\n", float64(sleepTimens)/1000000000.0)
				time.Sleep(time.Duration(sleepTimens) * time.Nanosecond)

				err := executeCommand(cmdToExec)
				if nil != err {
					log.Error().Msgf("[ERROR] %s\n", err.Error())
				}
			}
		}(cmdStmt)
	}
	for {
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
		restartCron()
		break
	default:
		log.Fatal().Msg("Invalid metric executed")
	}
}
