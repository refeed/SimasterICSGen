package simasterexam

import (
	"SimasterICSGen/internal/simastertime"
	"io"
	"log"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type Exam struct {
	Code    string
	Subject string
	Sks     string
	Class   string
	Date    string
	StartAt time.Time
	EndAt   time.Time
	Room    string
	Chair   string
}

func Parse(page io.Reader) []Exam {
	doc, err := goquery.NewDocumentFromReader(page)
	if err != nil {
		log.Fatalln(err)
	}

	exams := []Exam{}

	doc.Find("table > tbody > tr").Each(func(i int, s *goquery.Selection) {
		exam := Exam{}

		s.Children().Each(func(i int, s *goquery.Selection) {
			switch i {
			case 1:
				exam.Code = s.Text()
			case 2:
				exam.Subject = s.Text()
			case 3:
				exam.Sks = s.Text()
			case 4:
				exam.Class = s.Text()
			case 5:
				exam.Date = s.Text()
			case 6:
				if exam.Date == "" {
					return
				}

				startEndTime := strings.Split(
					strings.Trim(s.Text(), " "), "-")
				exam.StartAt = simastertime.Parse(exam.Date, startEndTime[0])
				exam.EndAt = simastertime.Parse(exam.Date, startEndTime[1])
			case 7:
				exam.Room = s.Text()
			case 8:
				exam.Chair = s.Text()
			}
		})

		if exam.Date == "" {
			return
		}

		exams = append(exams, exam)
	})

	return exams
}
