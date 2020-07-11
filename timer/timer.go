package timer

import (
	"strconv"
	"strings"
	"time"
)

const (
	MINUTE  = iota
	HOUR    = iota
	DAY     = iota
	MONTH   = iota
	WEEKDAY = iota
)

func getTimeDimension(timeDimension int, timeStruct time.Time) int {
	switch timeDimension {
	case MINUTE:
		return timeStruct.Minute()
	case HOUR:
		return timeStruct.Hour()
	case DAY:
		return timeStruct.Day()
	case MONTH:
		return int(timeStruct.Month())
	case WEEKDAY:
		return int(timeStruct.Weekday())
	}
	return -1
}

func compareTimeDimen(timeFormat string, timeDimension int, timeNext time.Time) bool {
	timeFormatList := strings.Split(timeFormat, ",")
	currVal := getTimeDimension(timeDimension, timeNext)
	if -1 == currVal {
		return false
	}

	for _, val := range timeFormatList {
		if val == "*" {
			return false
		}
		if strings.Contains(val, "/") {
			divFactor, _ := strconv.Atoi(strings.Split(val, "/")[1])
			if currVal%divFactor == 0 {
				return false
			}
		}
		value, _ := strconv.Atoi(val)
		if value == currVal {
			return false
		}
	}

	return true
}

func setTime(year int, month time.Month, day int, hour int, minute int, second int, nanosecond int, location *time.Location) time.Time {
	return time.Date(year,
		month,
		day,
		hour,
		minute,
		second,
		nanosecond,
		location)
}

func GetSleepTime(timeArr []string) int {
	if len(timeArr) != 5 {
		return -1
	}

	timeNow := time.Now()
	timeNext := setTime(timeNow.Year(), timeNow.Month(), timeNow.Day(), timeNow.Hour(),
		timeNow.Minute(), 0, 0, timeNow.Location()).Add(time.Minute)

	for {
		if compareTimeDimen(timeArr[0], MINUTE, timeNext) {
			timeNext = timeNext.Add(time.Minute)
			continue
		}
		if compareTimeDimen(timeArr[1], HOUR, timeNext) {
			timeNext = timeNext.Add(time.Hour)
			timeNext = setTime(timeNext.Year(), timeNext.Month(), timeNext.Day(), timeNext.Hour(),
				0, 0, 0, timeNext.Location())
			continue
		}
		if compareTimeDimen(timeArr[2], DAY, timeNext) ||
			compareTimeDimen(timeArr[4], WEEKDAY, timeNext) {
			timeNext = timeNext.AddDate(0, 0, 1)
			timeNext = setTime(timeNext.Year(), timeNext.Month(), timeNext.Day(), 0,
				0, 0, 0, timeNext.Location())
			continue
		}
		if compareTimeDimen(timeArr[3], MONTH, timeNext) {
			timeNext = timeNext.AddDate(0, 1, 0)
			timeNext = time.Date(timeNext.Year(), timeNext.Month(), 0, 0,
				0, 0, 0, timeNext.Location())
			continue
		}
		if compareTimeDimen(timeArr[2], DAY, timeNext) ||
			compareTimeDimen(timeArr[4], WEEKDAY, timeNext) {
			timeNext = timeNext.AddDate(0, 0, 1)
			continue
		}
		if compareTimeDimen(timeArr[1], HOUR, timeNext) {
			timeNext = timeNext.Add(time.Hour)
			continue
		}
		if compareTimeDimen(timeArr[0], MINUTE, timeNext) {
			timeNext = timeNext.Add(time.Minute)
			continue
		}
		break
	}

	sleepTime := timeNext.Sub(timeNow)
	return int(sleepTime.Nanoseconds())
}
