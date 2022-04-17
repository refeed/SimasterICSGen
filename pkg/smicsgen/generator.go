package smicsgen

import (
	"SimasterICSGen/internal/simasterclass"
	"SimasterICSGen/internal/simasterexam"
	"fmt"
	"io"

	ics "github.com/arran4/golang-ical"
)

func GenerateExamICS(examSchedHtml io.Reader) (string, int) {
	cal := ics.NewCalendar()
	cal.SetMethod(ics.MethodRequest)

	eventExportedNum := 0

	for _, exam := range simasterexam.Parse(examSchedHtml) {
		eventTitle := fmt.Sprintf("Ujian %v (%v) (Class: %v)", exam.Subject, exam.Code, exam.Class)

		event := cal.AddEvent(eventTitle)
		event.SetStartAt(exam.StartAt)
		event.SetEndAt(exam.EndAt)
		event.SetSummary(eventTitle)
		event.SetLocation(fmt.Sprintf("%v at chair %v", exam.Room, exam.Chair))
		event.SetDescription(fmt.Sprintf("Code: %v\nSKS: %v", exam.Code, exam.Sks))
		eventExportedNum++
	}

	return cal.Serialize(), eventExportedNum
}

func GenerateClassICS(classSchedHtml io.Reader) (string, int) {
	cal := ics.NewCalendar()
	cal.SetMethod(ics.MethodRequest)

	eventExportedNum := 0

	for _, class := range simasterclass.Parse(classSchedHtml) {
		eventTitle := fmt.Sprintf("%v (%v)", class.Subject, class.Class)

		event := cal.AddEvent(eventTitle)
		event.SetStartAt(class.FirstSched[0])
		event.SetEndAt(class.FirstSched[1])
		event.SetSummary(eventTitle)
		event.SetLocation(class.Room)
		event.AddRrule("FREQ=WEEKLY;COUNT=16")
		event.SetDescription(fmt.Sprintf("Code: %v\nSKS: %v\n%v",
			class.Code, class.Sks, class.Lecturer))
		eventExportedNum++
	}

	return cal.Serialize(), eventExportedNum
}
