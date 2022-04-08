// This package's functions are intended to be called from JS side

package main

import (
	"SimasterICSGen/pkg/smicsgen"
	"strings"
	"syscall/js"
)

func convertICSfromHTML() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) != 1 {
			return "Invalid number of arguments passed"
		}

		inputSMHTML := args[0].String()

		ics, numEvents := smicsgen.Generate(strings.NewReader(inputSMHTML))

		return []interface{}{ics, numEvents}
	})
}

func main() {
	// convertICSfromHTML(inputSMHTML string) string, int
	js.Global().Set("convertICSfromHTML", convertICSfromHTML())
	<-make(chan int)
}
