package simastertime

import (
	"log"
	"strconv"
	"strings"
	"time"
)

var (
	IDN_MONTH_NUM_MAP = map[string]time.Month{
		"Januari":   1,
		"Februari":  2,
		"Maret":     3,
		"April":     4,
		"Mei":       5,
		"Juni":      6,
		"Juli":      7,
		"Agustus":   8,
		"September": 9,
		"Oktober":   10,
		"November":  11,
		"Desember":  12,
	}
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

	// Simaster seemingly always uses UTC+7
	loc := time.FixedZone("UTC+7", 7*60*60)

	return time.Date(year, month, day, timeHr.Hour(), timeHr.Minute(), 0, 0, loc)
}
