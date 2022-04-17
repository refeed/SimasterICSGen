package simastertime

import (
	"log"
	"strconv"
	"strings"
	"time"
)

func Parse(sDate string, sTime string) time.Time {
	splitted := strings.Fields(sDate)

	year, err := strconv.Atoi(splitted[2])
	if err != nil {
		log.Println(err)
	}

	month := IDN_MONTH_NUM_MAP[splitted[1]]

	day, err := strconv.Atoi(splitted[0])
	if err != nil {
		log.Println(err)
	}

	timeHr, err := time.Parse("15:04", sTime)
	if err != nil {
		log.Println(err)
	}

	return time.Date(year, month, day, timeHr.Hour(), timeHr.Minute(), 0, 0, TZ)
}
