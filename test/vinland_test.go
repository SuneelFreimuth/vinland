package tests

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"sync"
	"testing"

	"github.com/SuneelFreimuth/vinland/src/ast"
	"github.com/SuneelFreimuth/vinland/src/eval"
	"github.com/SuneelFreimuth/vinland/src/parser"
	"github.com/antlr4-go/antlr/v4"
)

const (
	NumTests = 3
	Reset    = "\033[0m"
	Red      = "\033[0;31m"
	Green    = "\033[0;32m"
)

func TestAllStages(t *testing.T) {
	newTestRunner().RunAllStages(t)
}

type testRunner struct {
	numTests   int
	parseTrees []parser.IProgramContext
	asts       []ast.Node
}

func newTestRunner() *testRunner {
	return &testRunner{
		numTests:   NumTests,
		parseTrees: make([]parser.IProgramContext, NumTests),
		asts:       make([]ast.Node, NumTests),
	}
}

func (tr *testRunner) RunAllStages(t *testing.T) {
	tr.TestParse(t)
	tr.TestASTBuild(t)
}

func (tr *testRunner) TestParse(t *testing.T) {
	fmt.Println("====== TEST PARSE ======")
	results := make([]testResult, tr.numTests)
	wg := sync.WaitGroup{}
	for i := 0; i < tr.numTests; i++ {
		wg.Add(1)
		go func(testNumber int, results []testResult) {
			results[testNumber] = tr.testParse(testNumber)
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
	fmt.Println()
}

func (tr *testRunner) testParse(testNumber int) testResult {
	expected, _ := os.ReadFile(fmt.Sprintf("parse/test%02d.out", testNumber))
	expected_ := string(expected)

	input, _ := os.ReadFile(fmt.Sprintf("parse/test%02d.vin", testNumber))
	tr.parseTrees[testNumber] = parse(string(input))

	actual := bytes.Buffer{}
	parser.Print(&actual, tr.parseTrees[testNumber])
	actual_ := actual.String()

	diff, textsEqual := Diff(
		strings.Split(expected_, "\n"),
		strings.Split(actual_, "\n"),
	)

	if textsEqual {
		return testResult{
			TestNumber: testNumber,
			Succeeded:  true,
		}
	} else {
		return testResult{
			TestNumber: testNumber,
			Succeeded:  false,
			Expected:   expected_,
			Actual:     actual_,
			Diff:       diff,
		}
	}
}

func (tr *testRunner) TestASTBuild(t *testing.T) {
	fmt.Println("====== TEST AST ======")
	results := make([]testResult, tr.numTests)
	wg := sync.WaitGroup{}
	for i := 0; i < tr.numTests; i++ {
		wg.Add(1)
		go func(testNumber int, results []testResult) {
			results[testNumber] = tr.testASTBuild(testNumber)
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
	fmt.Println()
}

func (tr *testRunner) testASTBuild(testNumber int) testResult {
	expected, _ := os.ReadFile(fmt.Sprintf("ast/test%02d.out", testNumber))
	expected_ := string(expected)

	tr.asts[testNumber] = ast.Build(tr.parseTrees[testNumber])
	actual := bytes.Buffer{}
	ast.Print(&actual, tr.asts[testNumber])
	actual_ := actual.String()

	diff, textsEqual := Diff(
		strings.Split(expected_, "\n"),
		strings.Split(actual_, "\n"),
	)

	if textsEqual {
		return testResult{
			TestNumber: testNumber,
			Succeeded:  true,
		}
	} else {
		return testResult{
			TestNumber: testNumber,
			Succeeded:  false,
			Expected:   expected_,
			Actual:     actual_,
			Diff:       diff,
		}
	}
}

func (tr *testRunner) TestEval(t *testing.T) {
	fmt.Println("====== TEST EVAL ======")
	results := make([]testResult, tr.numTests)
	wg := sync.WaitGroup{}
	for i := 0; i < tr.numTests; i++ {
		wg.Add(1)
		go func(testNumber int, results []testResult) {
			results[testNumber] = tr.testEval(testNumber)
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
	fmt.Println()
}

func (tr *testRunner) testEval(testNumber int) testResult {
	expected, _ := os.ReadFile(fmt.Sprintf("ast/test%02d.out", testNumber))
	expected_ := string(expected)

	tr.asts[testNumber] = ast.Build(tr.parseTrees[testNumber])
	actual := bytes.Buffer{}
	ast.Print(&actual, tr.asts[testNumber])
	actual_ := actual.String()

	diff, textsEqual := Diff(
		strings.Split(expected_, "\n"),
		strings.Split(actual_, "\n"),
	)

	if textsEqual {
		return testResult{
			TestNumber: testNumber,
			Succeeded:  true,
		}
	} else {
		return testResult{
			TestNumber: testNumber,
			Succeeded:  false,
			Expected:   expected_,
			Actual:     actual_,
			Diff:       diff,
		}
	}
}

type testResult struct {
	TestNumber int
	Succeeded  bool
	Expected   string
	Actual     string
	Diff       []string
}

type evalTestResult struct {
	TestNumber int
	Succeeded bool
	Expected eval.Value
	Actual eval.Value
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
