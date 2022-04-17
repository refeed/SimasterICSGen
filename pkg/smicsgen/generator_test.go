package smicsgen

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateExam(t *testing.T) {
	f, err := os.Open("testdata/exam_input.html")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	fics, err := os.ReadFile("testdata/exam_expected.ics")
	if err != nil {
		fmt.Println(err)
	}

	expectedICS := string(fics)
	actualICS, actualEventNum := GenerateExamICS(f)

	assert.Equal(t, expectedICS, actualICS)
	assert.Equal(t, 8, actualEventNum)
}

func TestGenerateClassICS(t *testing.T) {
	f, err := os.Open("testdata/class_input.html")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	fics, err := os.ReadFile("testdata/class_expected.ics")
	if err != nil {
		fmt.Println(err)
	}

	expectedICS := string(fics)
	actualICS, actualEventNum := GenerateClassICS(f)

	assert.Equal(t, expectedICS, actualICS)
	assert.Equal(t, 8, actualEventNum)
}
