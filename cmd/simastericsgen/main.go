package main

import (
	"SimasterICSGen/pkg/smicsgen"
	"flag"
	"log"
	"os"
	"path"
)

var (
	htmlFilePath   = flag.String("input", "", "(Mandatory) The HTML file of Simaster Jadwal Ujian page")
	outputFilePath = flag.String("output", "result.ics", "The ICS output")
)

func main() {
	flag.Parse()
	if *htmlFilePath == "" {
		flag.Usage()
		os.Exit(1)
	}

	f, err := os.Open(*htmlFilePath)
	checkErr(err)
	defer f.Close()

	result, eventExportedNum := smicsgen.Generate(f)

	cwd, err := os.Getwd()
	checkErr(err)

	outputPath := path.Join(cwd, *outputFilePath)
	log.Printf("Writing the ICS file to: %v\n", outputPath)

	outF, err := os.Create(outputPath)
	checkErr(err)
	defer outF.Close()

	outF.WriteString(result)
	log.Printf("Successfully exported %v event(s)", eventExportedNum)
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
