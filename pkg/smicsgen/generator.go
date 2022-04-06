package smicsgen

import (
	"SimasterICSGen/internal/simasterexam"
	"fmt"
	"io"

	ics "github.com/arran4/golang-ical"
)

func Generate(examSchedHtml io.Reader) string {
	cal := ics.NewCalendar()
	cal.SetMethod(ics.MethodRequest)

	for _, exam := range simasterexam.Parse(examSchedHtml) {
		eventTitle := fmt.Sprintf("Ujian %v (%v) (Class: %v)", exam.Subject, exam.Code, exam.Class)

		event := cal.AddEvent(eventTitle)
		event.SetStartAt(exam.StartAt)
		event.SetEndAt(exam.EndAt)
		event.SetSummary(eventTitle)
		event.SetLocation(fmt.Sprintf("%v at chair %v", exam.Room, exam.Chair))
		event.SetDescription(fmt.Sprintf("Code: %v\nSKS: %v", exam.Code, exam.Sks))
	}

	return cal.Serialize()
}
