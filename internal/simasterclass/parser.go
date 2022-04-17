package simasterclass

import (
	"SimasterICSGen/internal/simastertime"
	"io"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const (
	ODD_SEM_MONTH_START  = time.August
	EVEN_SEM_MONTH_START = time.February
)

type Class struct {
	Code        string
	Class       string
	Subject     string
	Sks         string
	Lecturer    string
	Room        string
	FirstSched  [2]time.Time
	RepeatUntil time.Time // Only the day and the month are used
}

func Parse(page io.Reader) []Class {
	doc, err := goquery.NewDocumentFromReader(page)
	if err != nil {
		log.Fatalln(err)
	}

	// Get info about what semester is this
	semesterStr := doc.Find("#select2-sesiId-container").Text()
	semStartTime := getSemStartTime(semesterStr)

	classes := []Class{}
	doc.Find("table > tbody > tr").Each(func(_ int, s *goquery.Selection) {
		class := Class{}

		s.Children().Each(func(i int, s *goquery.Selection) {
			switch i {
			case 1:
				class.Code = s.Text()
			case 2:
				subjectAndClass := strings.Split(s.Text(), "\n")
				class.Subject = strings.Trim(subjectAndClass[1], " ")
				class.Class = strings.Trim(subjectAndClass[2], " ")
			case 3:
				class.Sks = s.Text()
			case 5:
				class.Lecturer = s.Text()
			case 6:
				if s.Text() == "" {
					return
				}

				firstStart, firstEnd, room := getStartEndScheduleAndPlace(s.Text(), semStartTime)

				class.Room = room
				class.FirstSched = [2]time.Time{
					firstStart, firstEnd}
				class.RepeatUntil = getScheduleEndTime(firstStart)
			}
		})

		if class.Room == "" {
			return
		}
		classes = append(classes, class)
	})

	return classes
}

// Helper function to parse the string "Jumat, 13:00-16:40 Ruang D.TEDI DARING 10".
// Returns startTime, endTime, place
func getStartEndScheduleAndPlace(s string, startSem time.Time) (time.Time, time.Time, string) {
	re := regexp.MustCompile(`(?P<Day>[a-zA-z]+), (?P<TimeStart>\d{2}:\d{2})-(?P<TimeEnd>\d{2}:\d{2}) (?P<Room>.+)`)
	match := re.FindStringSubmatch(s)

	day := match[re.SubexpIndex("Day")]
	timeStart := parseHour(match[re.SubexpIndex("TimeStart")])
	timeEnd := parseHour(match[re.SubexpIndex("TimeEnd")])
	duration := timeEnd.Sub(timeStart)
	room := match[re.SubexpIndex("Room")]

	mondayDiffNum := simastertime.IDN_DAY_NUM_MAP[day] - time.Monday
	startSchedTime := startSem.Add(24 * time.Hour * time.Duration(mondayDiffNum)).Add(
		time.Hour*time.Duration(timeStart.Hour()) +
			time.Minute*time.Duration(timeStart.Minute()))
	endSchedTime := startSchedTime.Add(duration)

	return startSchedTime, endSchedTime, room
}

func parseHour(tS string) time.Time {
	t, _ := time.Parse("15:04", tS)
	return t
}

// Get the first Monday 0.00 AM of the semester.
// semStr is something like "Semester Genap 2021/2022"
func getSemStartTime(semStr string) time.Time {
	re := regexp.MustCompile(`Semester (?P<Sem>(Gasal|Genap)) (?P<Year1>\d{4})\/(?P<Year2>\d{4})`)
	match := re.FindStringSubmatch(semStr)

	semester := match[re.SubexpIndex("Sem")]
	year1, _ := strconv.Atoi(match[re.SubexpIndex("Year1")])
	year2, _ := strconv.Atoi(match[re.SubexpIndex("Year2")])

	var startSched time.Time

	if semester == "Gasal" {
		firstMonday := firstMonday(year1, ODD_SEM_MONTH_START)
		startSched = time.Date(year1, ODD_SEM_MONTH_START, firstMonday, 0, 0, 0, 0, simastertime.TZ)
	} else {
		firstMonday := firstMonday(year2, EVEN_SEM_MONTH_START)
		startSched = time.Date(year2, EVEN_SEM_MONTH_START, firstMonday, 0, 0, 0, 0, simastertime.TZ)
	}

	return startSched
}

// Get the last schedule from adding the first schedule by 16 weeks
func getScheduleEndTime(startSched time.Time) time.Time {
	return startSched.Add(time.Hour * 24 * 7 * 16) // 16 weeks after the start
}

// Return the day of the first Monday in the given month.
func firstMonday(year int, month time.Month) int {
	t := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	return (8-int(t.Weekday()))%7 + 1
}
