// Code generated from grulev3.g4 by ANTLR 4.13.2. DO NOT EDIT.

package grulev3

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

type grulev3Lexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var Grulev3LexerLexerStaticData struct {
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

func grulev3lexerLexerInit() {
	staticData := &Grulev3LexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "','", "'+'", "'-'", "'/'", "'*'", "'%'", "'.'", "';'", "'{'", "'}'",
		"'('", "')'", "'['", "']'", "", "", "", "'&&'", "'||'", "", "", "",
		"'!'", "", "'=='", "'='", "'+='", "'-='", "'/='", "'*='", "'>'", "'<'",
		"'>='", "'<='", "'!='", "'&'", "'|'",
	}
	staticData.SymbolicNames = []string{
		"", "", "PLUS", "MINUS", "DIV", "MUL", "MOD", "DOT", "SEMICOLON", "LR_BRACE",
		"RR_BRACE", "LR_BRACKET", "RR_BRACKET", "LS_BRACKET", "RS_BRACKET",
		"RULE", "WHEN", "THEN", "AND", "OR", "TRUE", "FALSE", "NIL_LITERAL",
		"NEGATION", "SALIENCE", "EQUALS", "ASSIGN", "PLUS_ASIGN", "MINUS_ASIGN",
		"DIV_ASIGN", "MUL_ASIGN", "GT", "LT", "GTE", "LTE", "NOTEQUALS", "BITAND",
		"BITOR", "SIMPLENAME", "DQUOTA_STRING", "SQUOTA_STRING", "DECIMAL_FLOAT_LIT",
		"DECIMAL_EXPONENT", "HEX_FLOAT_LIT", "HEX_EXPONENT", "DEC_LIT", "HEX_LIT",
		"OCT_LIT", "SPACE", "COMMENT", "LINE_COMMENT",
	}
	staticData.RuleNames = []string{
		"T__0", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L",
		"M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
		"ISC", "IC", "PLUS", "MINUS", "DIV", "MUL", "MOD", "DOT", "SEMICOLON",
		"LR_BRACE", "RR_BRACE", "LR_BRACKET", "RR_BRACKET", "LS_BRACKET", "RS_BRACKET",
		"RULE", "WHEN", "THEN", "AND", "OR", "TRUE", "FALSE", "NIL_LITERAL",
		"NEGATION", "SALIENCE", "EQUALS", "ASSIGN", "PLUS_ASIGN", "MINUS_ASIGN",
		"DIV_ASIGN", "MUL_ASIGN", "GT", "LT", "GTE", "LTE", "NOTEQUALS", "BITAND",
		"BITOR", "SIMPLENAME", "DQUOTA_STRING", "SQUOTA_STRING", "DECIMAL_FLOAT_LIT",
		"DECIMAL_EXPONENT", "HEX_FLOAT_LIT", "HEX_MANTISA", "HEX_EXPONENT",
		"DEC_LIT", "HEX_LIT", "OCT_LIT", "HEX_DIGITS", "DEC_DIGITS", "OCT_DIGITS",
		"DEC_DIGIT", "OCT_DIGIT", "HEX_DIGIT", "SPACE", "COMMENT", "LINE_COMMENT",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 50, 484, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46,
		2, 47, 7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2,
		52, 7, 52, 2, 53, 7, 53, 2, 54, 7, 54, 2, 55, 7, 55, 2, 56, 7, 56, 2, 57,
		7, 57, 2, 58, 7, 58, 2, 59, 7, 59, 2, 60, 7, 60, 2, 61, 7, 61, 2, 62, 7,
		62, 2, 63, 7, 63, 2, 64, 7, 64, 2, 65, 7, 65, 2, 66, 7, 66, 2, 67, 7, 67,
		2, 68, 7, 68, 2, 69, 7, 69, 2, 70, 7, 70, 2, 71, 7, 71, 2, 72, 7, 72, 2,
		73, 7, 73, 2, 74, 7, 74, 2, 75, 7, 75, 2, 76, 7, 76, 2, 77, 7, 77, 2, 78,
		7, 78, 2, 79, 7, 79, 2, 80, 7, 80, 2, 81, 7, 81, 2, 82, 7, 82, 2, 83, 7,
		83, 2, 84, 7, 84, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4,
		1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10,
		1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1,
		15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 20, 1, 20,
		1, 21, 1, 21, 1, 22, 1, 22, 1, 23, 1, 23, 1, 24, 1, 24, 1, 25, 1, 25, 1,
		26, 1, 26, 1, 27, 1, 27, 1, 28, 1, 28, 3, 28, 230, 8, 28, 1, 29, 1, 29,
		1, 30, 1, 30, 1, 31, 1, 31, 1, 32, 1, 32, 1, 33, 1, 33, 1, 34, 1, 34, 1,
		35, 1, 35, 1, 36, 1, 36, 1, 37, 1, 37, 1, 38, 1, 38, 1, 39, 1, 39, 1, 40,
		1, 40, 1, 41, 1, 41, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 43, 1, 43, 1,
		43, 1, 43, 1, 43, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 1, 45, 1, 45, 1, 45,
		1, 46, 1, 46, 1, 46, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 48, 1, 48, 1,
		48, 1, 48, 1, 48, 1, 48, 1, 49, 1, 49, 1, 49, 1, 49, 1, 50, 1, 50, 1, 51,
		1, 51, 1, 51, 1, 51, 1, 51, 1, 51, 1, 51, 1, 51, 1, 51, 1, 52, 1, 52, 1,
		52, 1, 53, 1, 53, 1, 54, 1, 54, 1, 54, 1, 55, 1, 55, 1, 55, 1, 56, 1, 56,
		1, 56, 1, 57, 1, 57, 1, 57, 1, 58, 1, 58, 1, 59, 1, 59, 1, 60, 1, 60, 1,
		60, 1, 61, 1, 61, 1, 61, 1, 62, 1, 62, 1, 62, 1, 63, 1, 63, 1, 64, 1, 64,
		1, 65, 1, 65, 5, 65, 341, 8, 65, 10, 65, 12, 65, 344, 9, 65, 1, 66, 1,
		66, 1, 66, 1, 66, 1, 66, 1, 66, 5, 66, 352, 8, 66, 10, 66, 12, 66, 355,
		9, 66, 1, 66, 1, 66, 1, 67, 1, 67, 1, 67, 1, 67, 1, 67, 1, 67, 5, 67, 365,
		8, 67, 10, 67, 12, 67, 368, 9, 67, 1, 67, 1, 67, 1, 68, 1, 68, 1, 68, 1,
		68, 3, 68, 376, 8, 68, 1, 68, 1, 68, 1, 68, 1, 68, 1, 68, 1, 68, 3, 68,
		384, 8, 68, 3, 68, 386, 8, 68, 1, 69, 1, 69, 1, 69, 3, 69, 391, 8, 69,
		1, 69, 1, 69, 1, 70, 1, 70, 1, 70, 1, 70, 1, 70, 1, 71, 1, 71, 1, 71, 3,
		71, 403, 8, 71, 1, 71, 1, 71, 1, 71, 1, 71, 3, 71, 409, 8, 71, 1, 72, 1,
		72, 1, 72, 3, 72, 414, 8, 72, 1, 72, 1, 72, 1, 73, 1, 73, 1, 73, 3, 73,
		421, 8, 73, 3, 73, 423, 8, 73, 1, 74, 1, 74, 1, 74, 1, 74, 1, 75, 1, 75,
		1, 75, 1, 76, 4, 76, 433, 8, 76, 11, 76, 12, 76, 434, 1, 77, 4, 77, 438,
		8, 77, 11, 77, 12, 77, 439, 1, 78, 4, 78, 443, 8, 78, 11, 78, 12, 78, 444,
		1, 79, 1, 79, 1, 80, 1, 80, 1, 81, 1, 81, 1, 82, 4, 82, 454, 8, 82, 11,
		82, 12, 82, 455, 1, 82, 1, 82, 1, 83, 1, 83, 1, 83, 1, 83, 5, 83, 464,
		8, 83, 10, 83, 12, 83, 467, 9, 83, 1, 83, 1, 83, 1, 83, 1, 83, 1, 83, 1,
		84, 1, 84, 1, 84, 1, 84, 5, 84, 478, 8, 84, 10, 84, 12, 84, 481, 9, 84,
		1, 84, 1, 84, 1, 465, 0, 85, 1, 1, 3, 0, 5, 0, 7, 0, 9, 0, 11, 0, 13, 0,
		15, 0, 17, 0, 19, 0, 21, 0, 23, 0, 25, 0, 27, 0, 29, 0, 31, 0, 33, 0, 35,
		0, 37, 0, 39, 0, 41, 0, 43, 0, 45, 0, 47, 0, 49, 0, 51, 0, 53, 0, 55, 0,
		57, 0, 59, 2, 61, 3, 63, 4, 65, 5, 67, 6, 69, 7, 71, 8, 73, 9, 75, 10,
		77, 11, 79, 12, 81, 13, 83, 14, 85, 15, 87, 16, 89, 17, 91, 18, 93, 19,
		95, 20, 97, 21, 99, 22, 101, 23, 103, 24, 105, 25, 107, 26, 109, 27, 111,
		28, 113, 29, 115, 30, 117, 31, 119, 32, 121, 33, 123, 34, 125, 35, 127,
		36, 129, 37, 131, 38, 133, 39, 135, 40, 137, 41, 139, 42, 141, 43, 143,
		0, 145, 44, 147, 45, 149, 46, 151, 47, 153, 0, 155, 0, 157, 0, 159, 0,
		161, 0, 163, 0, 165, 48, 167, 49, 169, 50, 1, 0, 36, 2, 0, 65, 65, 97,
		97, 2, 0, 66, 66, 98, 98, 2, 0, 67, 67, 99, 99, 2, 0, 68, 68, 100, 100,
		2, 0, 69, 69, 101, 101, 2, 0, 70, 70, 102, 102, 2, 0, 71, 71, 103, 103,
		2, 0, 72, 72, 104, 104, 2, 0, 73, 73, 105, 105, 2, 0, 74, 74, 106, 106,
		2, 0, 75, 75, 107, 107, 2, 0, 76, 76, 108, 108, 2, 0, 77, 77, 109, 109,
		2, 0, 78, 78, 110, 110, 2, 0, 79, 79, 111, 111, 2, 0, 80, 80, 112, 112,
		2, 0, 81, 81, 113, 113, 2, 0, 82, 82, 114, 114, 2, 0, 83, 83, 115, 115,
		2, 0, 84, 84, 116, 116, 2, 0, 85, 85, 117, 117, 2, 0, 86, 86, 118, 118,
		2, 0, 87, 87, 119, 119, 2, 0, 88, 88, 120, 120, 2, 0, 89, 89, 121, 121,
		2, 0, 90, 90, 122, 122, 13, 0, 65, 90, 97, 122, 192, 214, 216, 246, 248,
		767, 880, 893, 895, 8191, 8204, 8205, 8304, 8591, 11264, 12271, 12289,
		55295, 63744, 64975, 65008, 65533, 5, 0, 48, 57, 95, 95, 183, 183, 768,
		879, 8255, 8256, 2, 0, 34, 34, 92, 92, 2, 0, 39, 39, 92, 92, 1, 0, 49,
		57, 1, 0, 48, 57, 1, 0, 48, 55, 3, 0, 48, 57, 65, 70, 97, 102, 3, 0, 9,
		10, 13, 13, 32, 32, 2, 0, 10, 10, 13, 13, 475, 0, 1, 1, 0, 0, 0, 0, 59,
		1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0,
		67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0,
		0, 75, 1, 0, 0, 0, 0, 77, 1, 0, 0, 0, 0, 79, 1, 0, 0, 0, 0, 81, 1, 0, 0,
		0, 0, 83, 1, 0, 0, 0, 0, 85, 1, 0, 0, 0, 0, 87, 1, 0, 0, 0, 0, 89, 1, 0,
		0, 0, 0, 91, 1, 0, 0, 0, 0, 93, 1, 0, 0, 0, 0, 95, 1, 0, 0, 0, 0, 97, 1,
		0, 0, 0, 0, 99, 1, 0, 0, 0, 0, 101, 1, 0, 0, 0, 0, 103, 1, 0, 0, 0, 0,
		105, 1, 0, 0, 0, 0, 107, 1, 0, 0, 0, 0, 109, 1, 0, 0, 0, 0, 111, 1, 0,
		0, 0, 0, 113, 1, 0, 0, 0, 0, 115, 1, 0, 0, 0, 0, 117, 1, 0, 0, 0, 0, 119,
		1, 0, 0, 0, 0, 121, 1, 0, 0, 0, 0, 123, 1, 0, 0, 0, 0, 125, 1, 0, 0, 0,
		0, 127, 1, 0, 0, 0, 0, 129, 1, 0, 0, 0, 0, 131, 1, 0, 0, 0, 0, 133, 1,
		0, 0, 0, 0, 135, 1, 0, 0, 0, 0, 137, 1, 0, 0, 0, 0, 139, 1, 0, 0, 0, 0,
		141, 1, 0, 0, 0, 0, 145, 1, 0, 0, 0, 0, 147, 1, 0, 0, 0, 0, 149, 1, 0,
		0, 0, 0, 151, 1, 0, 0, 0, 0, 165, 1, 0, 0, 0, 0, 167, 1, 0, 0, 0, 0, 169,
		1, 0, 0, 0, 1, 171, 1, 0, 0, 0, 3, 173, 1, 0, 0, 0, 5, 175, 1, 0, 0, 0,
		7, 177, 1, 0, 0, 0, 9, 179, 1, 0, 0, 0, 11, 181, 1, 0, 0, 0, 13, 183, 1,
		0, 0, 0, 15, 185, 1, 0, 0, 0, 17, 187, 1, 0, 0, 0, 19, 189, 1, 0, 0, 0,
		21, 191, 1, 0, 0, 0, 23, 193, 1, 0, 0, 0, 25, 195, 1, 0, 0, 0, 27, 197,
		1, 0, 0, 0, 29, 199, 1, 0, 0, 0, 31, 201, 1, 0, 0, 0, 33, 203, 1, 0, 0,
		0, 35, 205, 1, 0, 0, 0, 37, 207, 1, 0, 0, 0, 39, 209, 1, 0, 0, 0, 41, 211,
		1, 0, 0, 0, 43, 213, 1, 0, 0, 0, 45, 215, 1, 0, 0, 0, 47, 217, 1, 0, 0,
		0, 49, 219, 1, 0, 0, 0, 51, 221, 1, 0, 0, 0, 53, 223, 1, 0, 0, 0, 55, 225,
		1, 0, 0, 0, 57, 229, 1, 0, 0, 0, 59, 231, 1, 0, 0, 0, 61, 233, 1, 0, 0,
		0, 63, 235, 1, 0, 0, 0, 65, 237, 1, 0, 0, 0, 67, 239, 1, 0, 0, 0, 69, 241,
		1, 0, 0, 0, 71, 243, 1, 0, 0, 0, 73, 245, 1, 0, 0, 0, 75, 247, 1, 0, 0,
		0, 77, 249, 1, 0, 0, 0, 79, 251, 1, 0, 0, 0, 81, 253, 1, 0, 0, 0, 83, 255,
		1, 0, 0, 0, 85, 257, 1, 0, 0, 0, 87, 262, 1, 0, 0, 0, 89, 267, 1, 0, 0,
		0, 91, 272, 1, 0, 0, 0, 93, 275, 1, 0, 0, 0, 95, 278, 1, 0, 0, 0, 97, 283,
		1, 0, 0, 0, 99, 289, 1, 0, 0, 0, 101, 293, 1, 0, 0, 0, 103, 295, 1, 0,
		0, 0, 105, 304, 1, 0, 0, 0, 107, 307, 1, 0, 0, 0, 109, 309, 1, 0, 0, 0,
		111, 312, 1, 0, 0, 0, 113, 315, 1, 0, 0, 0, 115, 318, 1, 0, 0, 0, 117,
		321, 1, 0, 0, 0, 119, 323, 1, 0, 0, 0, 121, 325, 1, 0, 0, 0, 123, 328,
		1, 0, 0, 0, 125, 331, 1, 0, 0, 0, 127, 334, 1, 0, 0, 0, 129, 336, 1, 0,
		0, 0, 131, 338, 1, 0, 0, 0, 133, 345, 1, 0, 0, 0, 135, 358, 1, 0, 0, 0,
		137, 385, 1, 0, 0, 0, 139, 387, 1, 0, 0, 0, 141, 394, 1, 0, 0, 0, 143,
		408, 1, 0, 0, 0, 145, 410, 1, 0, 0, 0, 147, 422, 1, 0, 0, 0, 149, 424,
		1, 0, 0, 0, 151, 428, 1, 0, 0, 0, 153, 432, 1, 0, 0, 0, 155, 437, 1, 0,
		0, 0, 157, 442, 1, 0, 0, 0, 159, 446, 1, 0, 0, 0, 161, 448, 1, 0, 0, 0,
		163, 450, 1, 0, 0, 0, 165, 453, 1, 0, 0, 0, 167, 459, 1, 0, 0, 0, 169,
		473, 1, 0, 0, 0, 171, 172, 5, 44, 0, 0, 172, 2, 1, 0, 0, 0, 173, 174, 7,
		0, 0, 0, 174, 4, 1, 0, 0, 0, 175, 176, 7, 1, 0, 0, 176, 6, 1, 0, 0, 0,
		177, 178, 7, 2, 0, 0, 178, 8, 1, 0, 0, 0, 179, 180, 7, 3, 0, 0, 180, 10,
		1, 0, 0, 0, 181, 182, 7, 4, 0, 0, 182, 12, 1, 0, 0, 0, 183, 184, 7, 5,
		0, 0, 184, 14, 1, 0, 0, 0, 185, 186, 7, 6, 0, 0, 186, 16, 1, 0, 0, 0, 187,
		188, 7, 7, 0, 0, 188, 18, 1, 0, 0, 0, 189, 190, 7, 8, 0, 0, 190, 20, 1,
		0, 0, 0, 191, 192, 7, 9, 0, 0, 192, 22, 1, 0, 0, 0, 193, 194, 7, 10, 0,
		0, 194, 24, 1, 0, 0, 0, 195, 196, 7, 11, 0, 0, 196, 26, 1, 0, 0, 0, 197,
		198, 7, 12, 0, 0, 198, 28, 1, 0, 0, 0, 199, 200, 7, 13, 0, 0, 200, 30,
		1, 0, 0, 0, 201, 202, 7, 14, 0, 0, 202, 32, 1, 0, 0, 0, 203, 204, 7, 15,
		0, 0, 204, 34, 1, 0, 0, 0, 205, 206, 7, 16, 0, 0, 206, 36, 1, 0, 0, 0,
		207, 208, 7, 17, 0, 0, 208, 38, 1, 0, 0, 0, 209, 210, 7, 18, 0, 0, 210,
		40, 1, 0, 0, 0, 211, 212, 7, 19, 0, 0, 212, 42, 1, 0, 0, 0, 213, 214, 7,
		20, 0, 0, 214, 44, 1, 0, 0, 0, 215, 216, 7, 21, 0, 0, 216, 46, 1, 0, 0,
		0, 217, 218, 7, 22, 0, 0, 218, 48, 1, 0, 0, 0, 219, 220, 7, 23, 0, 0, 220,
		50, 1, 0, 0, 0, 221, 222, 7, 24, 0, 0, 222, 52, 1, 0, 0, 0, 223, 224, 7,
		25, 0, 0, 224, 54, 1, 0, 0, 0, 225, 226, 7, 26, 0, 0, 226, 56, 1, 0, 0,
		0, 227, 230, 3, 55, 27, 0, 228, 230, 7, 27, 0, 0, 229, 227, 1, 0, 0, 0,
		229, 228, 1, 0, 0, 0, 230, 58, 1, 0, 0, 0, 231, 232, 5, 43, 0, 0, 232,
		60, 1, 0, 0, 0, 233, 234, 5, 45, 0, 0, 234, 62, 1, 0, 0, 0, 235, 236, 5,
		47, 0, 0, 236, 64, 1, 0, 0, 0, 237, 238, 5, 42, 0, 0, 238, 66, 1, 0, 0,
		0, 239, 240, 5, 37, 0, 0, 240, 68, 1, 0, 0, 0, 241, 242, 5, 46, 0, 0, 242,
		70, 1, 0, 0, 0, 243, 244, 5, 59, 0, 0, 244, 72, 1, 0, 0, 0, 245, 246, 5,
		123, 0, 0, 246, 74, 1, 0, 0, 0, 247, 248, 5, 125, 0, 0, 248, 76, 1, 0,
		0, 0, 249, 250, 5, 40, 0, 0, 250, 78, 1, 0, 0, 0, 251, 252, 5, 41, 0, 0,
		252, 80, 1, 0, 0, 0, 253, 254, 5, 91, 0, 0, 254, 82, 1, 0, 0, 0, 255, 256,
		5, 93, 0, 0, 256, 84, 1, 0, 0, 0, 257, 258, 3, 37, 18, 0, 258, 259, 3,
		43, 21, 0, 259, 260, 3, 25, 12, 0, 260, 261, 3, 11, 5, 0, 261, 86, 1, 0,
		0, 0, 262, 263, 3, 47, 23, 0, 263, 264, 3, 17, 8, 0, 264, 265, 3, 11, 5,
		0, 265, 266, 3, 29, 14, 0, 266, 88, 1, 0, 0, 0, 267, 268, 3, 41, 20, 0,
		268, 269, 3, 17, 8, 0, 269, 270, 3, 11, 5, 0, 270, 271, 3, 29, 14, 0, 271,
		90, 1, 0, 0, 0, 272, 273, 5, 38, 0, 0, 273, 274, 5, 38, 0, 0, 274, 92,
		1, 0, 0, 0, 275, 276, 5, 124, 0, 0, 276, 277, 5, 124, 0, 0, 277, 94, 1,
		0, 0, 0, 278, 279, 3, 41, 20, 0, 279, 280, 3, 37, 18, 0, 280, 281, 3, 43,
		21, 0, 281, 282, 3, 11, 5, 0, 282, 96, 1, 0, 0, 0, 283, 284, 3, 13, 6,
		0, 284, 285, 3, 3, 1, 0, 285, 286, 3, 25, 12, 0, 286, 287, 3, 39, 19, 0,
		287, 288, 3, 11, 5, 0, 288, 98, 1, 0, 0, 0, 289, 290, 3, 29, 14, 0, 290,
		291, 3, 19, 9, 0, 291, 292, 3, 25, 12, 0, 292, 100, 1, 0, 0, 0, 293, 294,
		5, 33, 0, 0, 294, 102, 1, 0, 0, 0, 295, 296, 3, 39, 19, 0, 296, 297, 3,
		3, 1, 0, 297, 298, 3, 25, 12, 0, 298, 299, 3, 19, 9, 0, 299, 300, 3, 11,
		5, 0, 300, 301, 3, 29, 14, 0, 301, 302, 3, 7, 3, 0, 302, 303, 3, 11, 5,
		0, 303, 104, 1, 0, 0, 0, 304, 305, 5, 61, 0, 0, 305, 306, 5, 61, 0, 0,
		306, 106, 1, 0, 0, 0, 307, 308, 5, 61, 0, 0, 308, 108, 1, 0, 0, 0, 309,
		310, 5, 43, 0, 0, 310, 311, 5, 61, 0, 0, 311, 110, 1, 0, 0, 0, 312, 313,
		5, 45, 0, 0, 313, 314, 5, 61, 0, 0, 314, 112, 1, 0, 0, 0, 315, 316, 5,
		47, 0, 0, 316, 317, 5, 61, 0, 0, 317, 114, 1, 0, 0, 0, 318, 319, 5, 42,
		0, 0, 319, 320, 5, 61, 0, 0, 320, 116, 1, 0, 0, 0, 321, 322, 5, 62, 0,
		0, 322, 118, 1, 0, 0, 0, 323, 324, 5, 60, 0, 0, 324, 120, 1, 0, 0, 0, 325,
		326, 5, 62, 0, 0, 326, 327, 5, 61, 0, 0, 327, 122, 1, 0, 0, 0, 328, 329,
		5, 60, 0, 0, 329, 330, 5, 61, 0, 0, 330, 124, 1, 0, 0, 0, 331, 332, 5,
		33, 0, 0, 332, 333, 5, 61, 0, 0, 333, 126, 1, 0, 0, 0, 334, 335, 5, 38,
		0, 0, 335, 128, 1, 0, 0, 0, 336, 337, 5, 124, 0, 0, 337, 130, 1, 0, 0,
		0, 338, 342, 3, 55, 27, 0, 339, 341, 3, 57, 28, 0, 340, 339, 1, 0, 0, 0,
		341, 344, 1, 0, 0, 0, 342, 340, 1, 0, 0, 0, 342, 343, 1, 0, 0, 0, 343,
		132, 1, 0, 0, 0, 344, 342, 1, 0, 0, 0, 345, 353, 5, 34, 0, 0, 346, 347,
		5, 92, 0, 0, 347, 352, 9, 0, 0, 0, 348, 349, 5, 34, 0, 0, 349, 352, 5,
		34, 0, 0, 350, 352, 8, 28, 0, 0, 351, 346, 1, 0, 0, 0, 351, 348, 1, 0,
		0, 0, 351, 350, 1, 0, 0, 0, 352, 355, 1, 0, 0, 0, 353, 351, 1, 0, 0, 0,
		353, 354, 1, 0, 0, 0, 354, 356, 1, 0, 0, 0, 355, 353, 1, 0, 0, 0, 356,
		357, 5, 34, 0, 0, 357, 134, 1, 0, 0, 0, 358, 366, 5, 39, 0, 0, 359, 360,
		5, 92, 0, 0, 360, 365, 9, 0, 0, 0, 361, 362, 5, 39, 0, 0, 362, 365, 5,
		39, 0, 0, 363, 365, 8, 29, 0, 0, 364, 359, 1, 0, 0, 0, 364, 361, 1, 0,
		0, 0, 364, 363, 1, 0, 0, 0, 365, 368, 1, 0, 0, 0, 366, 364, 1, 0, 0, 0,
		366, 367, 1, 0, 0, 0, 367, 369, 1, 0, 0, 0, 368, 366, 1, 0, 0, 0, 369,
		370, 5, 39, 0, 0, 370, 136, 1, 0, 0, 0, 371, 372, 3, 147, 73, 0, 372, 373,
		3, 69, 34, 0, 373, 375, 3, 155, 77, 0, 374, 376, 3, 139, 69, 0, 375, 374,
		1, 0, 0, 0, 375, 376, 1, 0, 0, 0, 376, 386, 1, 0, 0, 0, 377, 378, 3, 147,
		73, 0, 378, 379, 3, 139, 69, 0, 379, 386, 1, 0, 0, 0, 380, 381, 3, 69,
		34, 0, 381, 383, 3, 155, 77, 0, 382, 384, 3, 139, 69, 0, 383, 382, 1, 0,
		0, 0, 383, 384, 1, 0, 0, 0, 384, 386, 1, 0, 0, 0, 385, 371, 1, 0, 0, 0,
		385, 377, 1, 0, 0, 0, 385, 380, 1, 0, 0, 0, 386, 138, 1, 0, 0, 0, 387,
		390, 3, 11, 5, 0, 388, 391, 3, 59, 29, 0, 389, 391, 3, 61, 30, 0, 390,
		388, 1, 0, 0, 0, 390, 389, 1, 0, 0, 0, 390, 391, 1, 0, 0, 0, 391, 392,
		1, 0, 0, 0, 392, 393, 3, 155, 77, 0, 393, 140, 1, 0, 0, 0, 394, 395, 5,
		48, 0, 0, 395, 396, 3, 49, 24, 0, 396, 397, 3, 143, 71, 0, 397, 398, 3,
		145, 72, 0, 398, 142, 1, 0, 0, 0, 399, 400, 3, 153, 76, 0, 400, 402, 3,
		69, 34, 0, 401, 403, 3, 153, 76, 0, 402, 401, 1, 0, 0, 0, 402, 403, 1,
		0, 0, 0, 403, 409, 1, 0, 0, 0, 404, 409, 3, 153, 76, 0, 405, 406, 3, 69,
		34, 0, 406, 407, 3, 153, 76, 0, 407, 409, 1, 0, 0, 0, 408, 399, 1, 0, 0,
		0, 408, 404, 1, 0, 0, 0, 408, 405, 1, 0, 0, 0, 409, 144, 1, 0, 0, 0, 410,
		413, 3, 33, 16, 0, 411, 414, 3, 59, 29, 0, 412, 414, 3, 61, 30, 0, 413,
		411, 1, 0, 0, 0, 413, 412, 1, 0, 0, 0, 413, 414, 1, 0, 0, 0, 414, 415,
		1, 0, 0, 0, 415, 416, 3, 155, 77, 0, 416, 146, 1, 0, 0, 0, 417, 423, 5,
		48, 0, 0, 418, 420, 7, 30, 0, 0, 419, 421, 3, 155, 77, 0, 420, 419, 1,
		0, 0, 0, 420, 421, 1, 0, 0, 0, 421, 423, 1, 0, 0, 0, 422, 417, 1, 0, 0,
		0, 422, 418, 1, 0, 0, 0, 423, 148, 1, 0, 0, 0, 424, 425, 5, 48, 0, 0, 425,
		426, 3, 49, 24, 0, 426, 427, 3, 153, 76, 0, 427, 150, 1, 0, 0, 0, 428,
		429, 5, 48, 0, 0, 429, 430, 3, 157, 78, 0, 430, 152, 1, 0, 0, 0, 431, 433,
		3, 163, 81, 0, 432, 431, 1, 0, 0, 0, 433, 434, 1, 0, 0, 0, 434, 432, 1,
		0, 0, 0, 434, 435, 1, 0, 0, 0, 435, 154, 1, 0, 0, 0, 436, 438, 3, 159,
		79, 0, 437, 436, 1, 0, 0, 0, 438, 439, 1, 0, 0, 0, 439, 437, 1, 0, 0, 0,
		439, 440, 1, 0, 0, 0, 440, 156, 1, 0, 0, 0, 441, 443, 3, 161, 80, 0, 442,
		441, 1, 0, 0, 0, 443, 444, 1, 0, 0, 0, 444, 442, 1, 0, 0, 0, 444, 445,
		1, 0, 0, 0, 445, 158, 1, 0, 0, 0, 446, 447, 7, 31, 0, 0, 447, 160, 1, 0,
		0, 0, 448, 449, 7, 32, 0, 0, 449, 162, 1, 0, 0, 0, 450, 451, 7, 33, 0,
		0, 451, 164, 1, 0, 0, 0, 452, 454, 7, 34, 0, 0, 453, 452, 1, 0, 0, 0, 454,
		455, 1, 0, 0, 0, 455, 453, 1, 0, 0, 0, 455, 456, 1, 0, 0, 0, 456, 457,
		1, 0, 0, 0, 457, 458, 6, 82, 0, 0, 458, 166, 1, 0, 0, 0, 459, 460, 5, 47,
		0, 0, 460, 461, 5, 42, 0, 0, 461, 465, 1, 0, 0, 0, 462, 464, 9, 0, 0, 0,
		463, 462, 1, 0, 0, 0, 464, 467, 1, 0, 0, 0, 465, 466, 1, 0, 0, 0, 465,
		463, 1, 0, 0, 0, 466, 468, 1, 0, 0, 0, 467, 465, 1, 0, 0, 0, 468, 469,
		5, 42, 0, 0, 469, 470, 5, 47, 0, 0, 470, 471, 1, 0, 0, 0, 471, 472, 6,
		83, 0, 0, 472, 168, 1, 0, 0, 0, 473, 474, 5, 47, 0, 0, 474, 475, 5, 47,
		0, 0, 475, 479, 1, 0, 0, 0, 476, 478, 8, 35, 0, 0, 477, 476, 1, 0, 0, 0,
		478, 481, 1, 0, 0, 0, 479, 477, 1, 0, 0, 0, 479, 480, 1, 0, 0, 0, 480,
		482, 1, 0, 0, 0, 481, 479, 1, 0, 0, 0, 482, 483, 6, 84, 0, 0, 483, 170,
		1, 0, 0, 0, 22, 0, 229, 342, 351, 353, 364, 366, 375, 383, 385, 390, 402,
		408, 413, 420, 422, 434, 439, 444, 455, 465, 479, 1, 6, 0, 0,
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

// grulev3LexerInit initializes any static state used to implement grulev3Lexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// Newgrulev3Lexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func Grulev3LexerInit() {
	staticData := &Grulev3LexerLexerStaticData
	staticData.once.Do(grulev3lexerLexerInit)
}

// Newgrulev3Lexer produces a new lexer instance for the optional input antlr.CharStream.
func Newgrulev3Lexer(input antlr.CharStream) *grulev3Lexer {
	Grulev3LexerInit()
	l := new(grulev3Lexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &Grulev3LexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "grulev3.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// grulev3Lexer tokens.
const (
	grulev3LexerT__0              = 1
	grulev3LexerPLUS              = 2
	grulev3LexerMINUS             = 3
	grulev3LexerDIV               = 4
	grulev3LexerMUL               = 5
	grulev3LexerMOD               = 6
	grulev3LexerDOT               = 7
	grulev3LexerSEMICOLON         = 8
	grulev3LexerLR_BRACE          = 9
	grulev3LexerRR_BRACE          = 10
	grulev3LexerLR_BRACKET        = 11
	grulev3LexerRR_BRACKET        = 12
	grulev3LexerLS_BRACKET        = 13
	grulev3LexerRS_BRACKET        = 14
	grulev3LexerRULE              = 15
	grulev3LexerWHEN              = 16
	grulev3LexerTHEN              = 17
	grulev3LexerAND               = 18
	grulev3LexerOR                = 19
	grulev3LexerTRUE              = 20
	grulev3LexerFALSE             = 21
	grulev3LexerNIL_LITERAL       = 22
	grulev3LexerNEGATION          = 23
	grulev3LexerSALIENCE          = 24
	grulev3LexerEQUALS            = 25
	grulev3LexerASSIGN            = 26
	grulev3LexerPLUS_ASIGN        = 27
	grulev3LexerMINUS_ASIGN       = 28
	grulev3LexerDIV_ASIGN         = 29
	grulev3LexerMUL_ASIGN         = 30
	grulev3LexerGT                = 31
	grulev3LexerLT                = 32
	grulev3LexerGTE               = 33
	grulev3LexerLTE               = 34
	grulev3LexerNOTEQUALS         = 35
	grulev3LexerBITAND            = 36
	grulev3LexerBITOR             = 37
	grulev3LexerSIMPLENAME        = 38
	grulev3LexerDQUOTA_STRING     = 39
	grulev3LexerSQUOTA_STRING     = 40
	grulev3LexerDECIMAL_FLOAT_LIT = 41
	grulev3LexerDECIMAL_EXPONENT  = 42
	grulev3LexerHEX_FLOAT_LIT     = 43
	grulev3LexerHEX_EXPONENT      = 44
	grulev3LexerDEC_LIT           = 45
	grulev3LexerHEX_LIT           = 46
	grulev3LexerOCT_LIT           = 47
	grulev3LexerSPACE             = 48
	grulev3LexerCOMMENT           = 49
	grulev3LexerLINE_COMMENT      = 50
)
