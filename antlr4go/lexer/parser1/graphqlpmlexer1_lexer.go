// Generated from GraphQLPMLexer1.g4 by ANTLR 4.6.

package parser1

import (
	"fmt"
	"unicode"

	"github.com/millergarym/antlr4-go"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 1072, 54993, 33286, 44333, 17431, 44785, 36224, 43741, 2, 14, 70, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4,
	13, 9, 13, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 6,
	3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3,
	11, 7, 11, 50, 10, 11, 12, 11, 14, 11, 53, 11, 11, 3, 12, 3, 12, 7, 12,
	57, 10, 12, 12, 12, 14, 12, 60, 11, 12, 3, 12, 3, 12, 3, 13, 6, 13, 65,
	10, 13, 13, 13, 14, 13, 66, 3, 13, 3, 13, 2, 2, 14, 3, 3, 5, 4, 7, 5, 9,
	6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 3, 2, 6,
	5, 2, 67, 92, 97, 97, 99, 124, 7, 2, 48, 48, 50, 59, 67, 92, 97, 97, 99,
	124, 3, 2, 36, 36, 5, 2, 11, 12, 15, 15, 34, 34, 72, 2, 3, 3, 2, 2, 2,
	2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2,
	2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2,
	2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 3, 27, 3, 2,
	2, 2, 5, 29, 3, 2, 2, 2, 7, 31, 3, 2, 2, 2, 9, 33, 3, 2, 2, 2, 11, 36,
	3, 2, 2, 2, 13, 39, 3, 2, 2, 2, 15, 41, 3, 2, 2, 2, 17, 43, 3, 2, 2, 2,
	19, 45, 3, 2, 2, 2, 21, 47, 3, 2, 2, 2, 23, 54, 3, 2, 2, 2, 25, 64, 3,
	2, 2, 2, 27, 28, 7, 66, 2, 2, 28, 4, 3, 2, 2, 2, 29, 30, 7, 42, 2, 2, 30,
	6, 3, 2, 2, 2, 31, 32, 7, 43, 2, 2, 32, 8, 3, 2, 2, 2, 33, 34, 7, 126,
	2, 2, 34, 35, 7, 126, 2, 2, 35, 10, 3, 2, 2, 2, 36, 37, 7, 40, 2, 2, 37,
	38, 7, 40, 2, 2, 38, 12, 3, 2, 2, 2, 39, 40, 7, 46, 2, 2, 40, 14, 3, 2,
	2, 2, 41, 42, 7, 60, 2, 2, 42, 16, 3, 2, 2, 2, 43, 44, 7, 125, 2, 2, 44,
	18, 3, 2, 2, 2, 45, 46, 7, 127, 2, 2, 46, 20, 3, 2, 2, 2, 47, 51, 9, 2,
	2, 2, 48, 50, 9, 3, 2, 2, 49, 48, 3, 2, 2, 2, 50, 53, 3, 2, 2, 2, 51, 49,
	3, 2, 2, 2, 51, 52, 3, 2, 2, 2, 52, 22, 3, 2, 2, 2, 53, 51, 3, 2, 2, 2,
	54, 58, 7, 36, 2, 2, 55, 57, 10, 4, 2, 2, 56, 55, 3, 2, 2, 2, 57, 60, 3,
	2, 2, 2, 58, 56, 3, 2, 2, 2, 58, 59, 3, 2, 2, 2, 59, 61, 3, 2, 2, 2, 60,
	58, 3, 2, 2, 2, 61, 62, 7, 36, 2, 2, 62, 24, 3, 2, 2, 2, 63, 65, 9, 5,
	2, 2, 64, 63, 3, 2, 2, 2, 65, 66, 3, 2, 2, 2, 66, 64, 3, 2, 2, 2, 66, 67,
	3, 2, 2, 2, 67, 68, 3, 2, 2, 2, 68, 69, 8, 13, 2, 2, 69, 26, 3, 2, 2, 2,
	6, 2, 51, 58, 66, 3, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'@'", "'('", "')'", "'||'", "'&&'", "','", "':'", "'{'", "'}'",
}

var lexerSymbolicNames = []string{
	"", "AT", "LP", "RP", "OR", "AND", "COMMA", "COLON", "LC", "RC", "NAME",
	"STRING", "WS",
}

var lexerRuleNames = []string{
	"AT", "LP", "RP", "OR", "AND", "COMMA", "COLON", "LC", "RC", "NAME", "STRING",
	"WS",
}

type GraphQLPMLexer1 struct {
	*antlr.BaseLexer
	modeNames []string
	// TODO: EOF string
}

func NewGraphQLPMLexer1(input antlr.CharStream) *GraphQLPMLexer1 {
	var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}

	l := new(GraphQLPMLexer1)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "GraphQLPMLexer1.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// GraphQLPMLexer1 tokens.
const (
	GraphQLPMLexer1AT     = 1
	GraphQLPMLexer1LP     = 2
	GraphQLPMLexer1RP     = 3
	GraphQLPMLexer1OR     = 4
	GraphQLPMLexer1AND    = 5
	GraphQLPMLexer1COMMA  = 6
	GraphQLPMLexer1COLON  = 7
	GraphQLPMLexer1LC     = 8
	GraphQLPMLexer1RC     = 9
	GraphQLPMLexer1NAME   = 10
	GraphQLPMLexer1STRING = 11
	GraphQLPMLexer1WS     = 12
)
