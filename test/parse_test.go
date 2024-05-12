package tests

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"sync"
	"strings"
	"testing"

	"github.com/antlr4-go/antlr/v4"
	"github.com/SuneelFreimuth/vinland/src/parser"
)

const (
	NumTests = 3
	Reset = "\033[0m"
	Red = "\033[0;31m"
	Green = "\033[0;32m"
)

func greenPrintf(format string, args ...any) {
	greenFormat := Green + format + Reset
	fmt.Printf(greenFormat, args...)
}

func redPrintf(format string, args ...any) {
	redFormat := Red + format + Reset
	fmt.Printf(redFormat, args...)
}

func parse(input string) parser.IProgramContext {
	is := antlr.NewInputStream(input)
	lexer := parser.NewVinlandLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewVinlandParser(stream)
	return p.Program()
}

type testResult struct {
	TestNumber int
	Succeeded bool
	Expected string
	Actual string
	Diff []string
}

func testParse(testNumber int) testResult {
	expected, _ := ioutil.ReadFile(fmt.Sprintf("parse/test%02d.out", testNumber))
	expected_ := string(expected)

	input, _ := ioutil.ReadFile(fmt.Sprintf("parse/test%02d.vin", testNumber))
	actual := bytes.Buffer{}
	parser.Print(&actual, parse(string(input)))
	actual_ := actual.String()

	fmt.Println("Expected:", expected_)
	fmt.Println("Actual:", actual_)
	diff, textsEqual := Diff(
		strings.Split(expected_, "\n"),
		strings.Split(actual_, "\n"),
	)

	if textsEqual {
		return testResult{
			TestNumber: testNumber,
			Succeeded: true,
		}
	} else {
		return testResult{
			TestNumber: testNumber,
			Succeeded: false,
			Expected: expected_,
			Actual: actual_,
			Diff: diff,
		}
	}
}

func TestParse(t *testing.T) {
	results := make([]testResult, NumTests)
	wg := sync.WaitGroup{}
	for i := 0; i < NumTests; i++ {
		wg.Add(1)
		go func(testNumber int, results []testResult) {
			results[testNumber] = testParse(testNumber)
			wg.Done()
		}(i, results)
	}
	wg.Wait()

	anyFailed := false
	for testNumber, result := range results {
		if result.Succeeded {
			greenPrintf("TEST %d: Success\n", testNumber)
		} else {
			anyFailed = true
			redPrintf("TEST %d: Failure\n", testNumber)

			for _, line := range result.Diff {
				fmt.Println(line)
			}
		}
	}

	if anyFailed {
		t.Fail()
	}
}