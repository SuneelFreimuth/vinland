package tests

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"sync"
	"testing"

	"github.com/SuneelFreimuth/vinland/src/parser"
	"github.com/antlr4-go/antlr/v4"
)

const (
	NumParseTests = 3
	NumASTTests = 3
	Reset = "\033[0m"
	Red = "\033[0;31m"
	Green = "\033[0;32m"
)

func testParse(testNumber int) testResult {
	expected, _ := os.ReadFile(fmt.Sprintf("parse/test%02d.out", testNumber))
	expected_ := string(expected)

	input, _ := os.ReadFile(fmt.Sprintf("parse/test%02d.vin", testNumber))
	actual := bytes.Buffer{}
	parser.Print(&actual, parse(string(input)))
	actual_ := actual.String()

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

func TestAllStages(t *testing.T) {
	newTestRunner().RunAllStages(t)
}


type testRunner struct {
	numParseTests int
	numASTTests int
}


func newTestRunner() *testRunner {
	return &testRunner{
		numParseTests: NumParseTests,
		numASTTests: NumASTTests,
	}
}

func (tr *testRunner) RunAllStages(t *testing.T) {
	tr.TestParse(t)
	tr.TestASTBuild(t)
}

func (tr *testRunner) TestParse(t *testing.T) {
	results := make([]testResult, tr.numParseTests)
	wg := sync.WaitGroup{}
	for i := 0; i < tr.numParseTests; i++ {
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

			fmt.Println("======= BEGIN EXPECTED =========")
			fmt.Println(result.Expected)
			fmt.Println("======= END EXPECTED =========")
			fmt.Println("======= BEGIN ACTUAL =========")
			fmt.Println(result.Actual)
			fmt.Println("======= END ACTUAL =========")
			// for _, line := range result.Diff {
			// 	fmt.Println(line)
			// }
		}
	}

	if anyFailed {
		t.Fail()
	}
}

func (tr *testRunner) TestASTBuild(*testing.T) {}


type testResult struct {
	TestNumber int
	Succeeded bool
	Expected string
	Actual string
	Diff []string
}


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
