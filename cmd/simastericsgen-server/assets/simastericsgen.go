// This package's functions are intended to be called from JS side

package main

import (
	"SimasterICSGen/pkg/smicsgen"
	"strings"
	"syscall/js"
)

func convertICSfromHTML() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) != 2 {
			return "Invalid number of arguments passed"
		}

		inputSMHTML := args[0].String()
		htmlType := args[1].String()

		var (
			ics       string
			numEvents int
		)

		switch htmlType {
		case "exam":
			ics, numEvents = smicsgen.GenerateExamICS(strings.NewReader(inputSMHTML))
		case "class":
			ics, numEvents = smicsgen.GenerateClassICS(strings.NewReader(inputSMHTML))
		}

		return []interface{}{ics, numEvents}
	})
}

func main() {
	// convertICSfromHTML(inputSMHTML string, htmlType string) string, int
	js.Global().Set("convertICSfromHTML", convertICSfromHTML())
	<-make(chan int)
}
