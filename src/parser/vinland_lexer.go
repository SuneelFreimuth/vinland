// Code generated from Vinland.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type VinlandLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var VinlandLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func vinlandlexerLexerInit() {
	staticData := &VinlandLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'{'", "'}'", "'='", "'fn'", "'('", "')'", "','", "'+'", "'-'",
		"'*'", "'/'", "", "", "", "", "", "", "", "", "", "", "'true'", "'false'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "Whitespace", "Identifier",
		"IntLiteral", "DecimalNumber", "HexadecimalNumber", "BinaryNumber",
		"FloatLiteral", "StringLiteral", "CharEscapeSeq", "BooleanLiteral",
		"True", "False",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "Whitespace", "Identifier", "IntLiteral", "DecimalNumber",
		"HexadecimalNumber", "BinaryNumber", "FloatLiteral", "StringLiteral",
		"StringElement", "CharEscapeSeq", "BooleanLiteral", "True", "False",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 23, 158, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 1, 0, 1, 0, 1, 1, 1, 1, 1,
		2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1,
		7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 4, 11, 74, 8, 11, 11, 11,
		12, 11, 75, 1, 11, 1, 11, 1, 12, 1, 12, 5, 12, 82, 8, 12, 10, 12, 12, 12,
		85, 9, 12, 1, 13, 1, 13, 1, 13, 3, 13, 90, 8, 13, 1, 14, 3, 14, 93, 8,
		14, 1, 14, 4, 14, 96, 8, 14, 11, 14, 12, 14, 97, 1, 15, 1, 15, 1, 15, 1,
		15, 4, 15, 104, 8, 15, 11, 15, 12, 15, 105, 1, 16, 1, 16, 1, 16, 1, 16,
		4, 16, 112, 8, 16, 11, 16, 12, 16, 113, 1, 17, 4, 17, 117, 8, 17, 11, 17,
		12, 17, 118, 1, 17, 1, 17, 5, 17, 123, 8, 17, 10, 17, 12, 17, 126, 9, 17,
		1, 18, 1, 18, 5, 18, 130, 8, 18, 10, 18, 12, 18, 133, 9, 18, 1, 18, 1,
		18, 1, 19, 1, 19, 3, 19, 139, 8, 19, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21,
		3, 21, 146, 8, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1,
		23, 1, 23, 1, 23, 1, 23, 0, 0, 24, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6,
		13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31,
		16, 33, 17, 35, 18, 37, 19, 39, 0, 41, 20, 43, 21, 45, 22, 47, 23, 1, 0,
		8, 3, 0, 9, 10, 13, 13, 32, 32, 2, 0, 65, 90, 97, 122, 3, 0, 48, 57, 65,
		90, 97, 122, 1, 0, 48, 57, 2, 0, 48, 57, 65, 70, 1, 0, 48, 49, 2, 0, 32,
		33, 35, 127, 5, 0, 34, 34, 92, 92, 110, 110, 114, 114, 116, 116, 169, 0,
		1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0,
		9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0,
		0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0,
		0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0,
		0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 41, 1,
		0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 1, 49,
		1, 0, 0, 0, 3, 51, 1, 0, 0, 0, 5, 53, 1, 0, 0, 0, 7, 55, 1, 0, 0, 0, 9,
		58, 1, 0, 0, 0, 11, 60, 1, 0, 0, 0, 13, 62, 1, 0, 0, 0, 15, 64, 1, 0, 0,
		0, 17, 66, 1, 0, 0, 0, 19, 68, 1, 0, 0, 0, 21, 70, 1, 0, 0, 0, 23, 73,
		1, 0, 0, 0, 25, 79, 1, 0, 0, 0, 27, 89, 1, 0, 0, 0, 29, 92, 1, 0, 0, 0,
		31, 99, 1, 0, 0, 0, 33, 107, 1, 0, 0, 0, 35, 116, 1, 0, 0, 0, 37, 127,
		1, 0, 0, 0, 39, 138, 1, 0, 0, 0, 41, 140, 1, 0, 0, 0, 43, 145, 1, 0, 0,
		0, 45, 147, 1, 0, 0, 0, 47, 152, 1, 0, 0, 0, 49, 50, 5, 123, 0, 0, 50,
		2, 1, 0, 0, 0, 51, 52, 5, 125, 0, 0, 52, 4, 1, 0, 0, 0, 53, 54, 5, 61,
		0, 0, 54, 6, 1, 0, 0, 0, 55, 56, 5, 102, 0, 0, 56, 57, 5, 110, 0, 0, 57,
		8, 1, 0, 0, 0, 58, 59, 5, 40, 0, 0, 59, 10, 1, 0, 0, 0, 60, 61, 5, 41,
		0, 0, 61, 12, 1, 0, 0, 0, 62, 63, 5, 44, 0, 0, 63, 14, 1, 0, 0, 0, 64,
		65, 5, 43, 0, 0, 65, 16, 1, 0, 0, 0, 66, 67, 5, 45, 0, 0, 67, 18, 1, 0,
		0, 0, 68, 69, 5, 42, 0, 0, 69, 20, 1, 0, 0, 0, 70, 71, 5, 47, 0, 0, 71,
		22, 1, 0, 0, 0, 72, 74, 7, 0, 0, 0, 73, 72, 1, 0, 0, 0, 74, 75, 1, 0, 0,
		0, 75, 73, 1, 0, 0, 0, 75, 76, 1, 0, 0, 0, 76, 77, 1, 0, 0, 0, 77, 78,
		6, 11, 0, 0, 78, 24, 1, 0, 0, 0, 79, 83, 7, 1, 0, 0, 80, 82, 7, 2, 0, 0,
		81, 80, 1, 0, 0, 0, 82, 85, 1, 0, 0, 0, 83, 81, 1, 0, 0, 0, 83, 84, 1,
		0, 0, 0, 84, 26, 1, 0, 0, 0, 85, 83, 1, 0, 0, 0, 86, 90, 3, 29, 14, 0,
		87, 90, 3, 31, 15, 0, 88, 90, 3, 33, 16, 0, 89, 86, 1, 0, 0, 0, 89, 87,
		1, 0, 0, 0, 89, 88, 1, 0, 0, 0, 90, 28, 1, 0, 0, 0, 91, 93, 5, 45, 0, 0,
		92, 91, 1, 0, 0, 0, 92, 93, 1, 0, 0, 0, 93, 95, 1, 0, 0, 0, 94, 96, 7,
		3, 0, 0, 95, 94, 1, 0, 0, 0, 96, 97, 1, 0, 0, 0, 97, 95, 1, 0, 0, 0, 97,
		98, 1, 0, 0, 0, 98, 30, 1, 0, 0, 0, 99, 100, 5, 48, 0, 0, 100, 101, 5,
		120, 0, 0, 101, 103, 1, 0, 0, 0, 102, 104, 7, 4, 0, 0, 103, 102, 1, 0,
		0, 0, 104, 105, 1, 0, 0, 0, 105, 103, 1, 0, 0, 0, 105, 106, 1, 0, 0, 0,
		106, 32, 1, 0, 0, 0, 107, 108, 5, 48, 0, 0, 108, 109, 5, 98, 0, 0, 109,
		111, 1, 0, 0, 0, 110, 112, 7, 5, 0, 0, 111, 110, 1, 0, 0, 0, 112, 113,
		1, 0, 0, 0, 113, 111, 1, 0, 0, 0, 113, 114, 1, 0, 0, 0, 114, 34, 1, 0,
		0, 0, 115, 117, 7, 3, 0, 0, 116, 115, 1, 0, 0, 0, 117, 118, 1, 0, 0, 0,
		118, 116, 1, 0, 0, 0, 118, 119, 1, 0, 0, 0, 119, 120, 1, 0, 0, 0, 120,
		124, 5, 46, 0, 0, 121, 123, 7, 3, 0, 0, 122, 121, 1, 0, 0, 0, 123, 126,
		1, 0, 0, 0, 124, 122, 1, 0, 0, 0, 124, 125, 1, 0, 0, 0, 125, 36, 1, 0,
		0, 0, 126, 124, 1, 0, 0, 0, 127, 131, 5, 34, 0, 0, 128, 130, 3, 39, 19,
		0, 129, 128, 1, 0, 0, 0, 130, 133, 1, 0, 0, 0, 131, 129, 1, 0, 0, 0, 131,
		132, 1, 0, 0, 0, 132, 134, 1, 0, 0, 0, 133, 131, 1, 0, 0, 0, 134, 135,
		5, 34, 0, 0, 135, 38, 1, 0, 0, 0, 136, 139, 7, 6, 0, 0, 137, 139, 3, 41,
		20, 0, 138, 136, 1, 0, 0, 0, 138, 137, 1, 0, 0, 0, 139, 40, 1, 0, 0, 0,
		140, 141, 5, 92, 0, 0, 141, 142, 7, 7, 0, 0, 142, 42, 1, 0, 0, 0, 143,
		146, 3, 45, 22, 0, 144, 146, 3, 47, 23, 0, 145, 143, 1, 0, 0, 0, 145, 144,
		1, 0, 0, 0, 146, 44, 1, 0, 0, 0, 147, 148, 5, 116, 0, 0, 148, 149, 5, 114,
		0, 0, 149, 150, 5, 117, 0, 0, 150, 151, 5, 101, 0, 0, 151, 46, 1, 0, 0,
		0, 152, 153, 5, 102, 0, 0, 153, 154, 5, 97, 0, 0, 154, 155, 5, 108, 0,
		0, 155, 156, 5, 115, 0, 0, 156, 157, 5, 101, 0, 0, 157, 48, 1, 0, 0, 0,
		13, 0, 75, 83, 89, 92, 97, 105, 113, 118, 124, 131, 138, 145, 1, 6, 0,
		0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// VinlandLexerInit initializes any static state used to implement VinlandLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewVinlandLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func VinlandLexerInit() {
	staticData := &VinlandLexerLexerStaticData
	staticData.once.Do(vinlandlexerLexerInit)
}

// NewVinlandLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewVinlandLexer(input antlr.CharStream) *VinlandLexer {
	VinlandLexerInit()
	l := new(VinlandLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &VinlandLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "Vinland.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// VinlandLexer tokens.
const (
	VinlandLexerT__0              = 1
	VinlandLexerT__1              = 2
	VinlandLexerT__2              = 3
	VinlandLexerT__3              = 4
	VinlandLexerT__4              = 5
	VinlandLexerT__5              = 6
	VinlandLexerT__6              = 7
	VinlandLexerT__7              = 8
	VinlandLexerT__8              = 9
	VinlandLexerT__9              = 10
	VinlandLexerT__10             = 11
	VinlandLexerWhitespace        = 12
	VinlandLexerIdentifier        = 13
	VinlandLexerIntLiteral        = 14
	VinlandLexerDecimalNumber     = 15
	VinlandLexerHexadecimalNumber = 16
	VinlandLexerBinaryNumber      = 17
	VinlandLexerFloatLiteral      = 18
	VinlandLexerStringLiteral     = 19
	VinlandLexerCharEscapeSeq     = 20
	VinlandLexerBooleanLiteral    = 21
	VinlandLexerTrue              = 22
	VinlandLexerFalse             = 23
)
