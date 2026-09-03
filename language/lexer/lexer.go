package lexer

import (
	"regexp"

	"github.com/graphql-go/graphql/language/source"
)

type TokenKind int

const (
	EOF TokenKind = iota + 1
	BANG
	DOLLAR
	PAREN_L
	PAREN_R
	SPREAD
	COLON
	EQUALS
	AT
	BRACKET_L
	BRACKET_R
	BRACE_L
	PIPE
	BRACE_R
	NAME
	INT
	FLOAT
	STRING
	BLOCK_STRING
	AMP
)

var tokenDescription = map[TokenKind]string{
	EOF:          "EOF",
	BANG:         "!",
	DOLLAR:       "$",
	PAREN_L:      "(",
	PAREN_R:      ")",
	SPREAD:       "...",
	COLON:        ":",
	EQUALS:       "=",
	AT:           "@",
	BRACKET_L:    "[",
	BRACKET_R:    "]",
	BRACE_L:      "{",
	PIPE:         "|",
	BRACE_R:      "}",
	NAME:         "Name",
	INT:          "Int",
	FLOAT:        "Float",
	STRING:       "String",
	BLOCK_STRING: "BlockString",
	AMP:          "&",
}

func (kind TokenKind) String() string { _ = "STUB: not implemented"; return "" }

const (
	FRAGMENT     = "fragment"
	QUERY        = "query"
	MUTATION     = "mutation"
	SUBSCRIPTION = "subscription"
	SCHEMA       = "schema"
	SCALAR       = "scalar"
	TYPE         = "type"
	INTERFACE    = "interface"
	UNION        = "union"
	ENUM         = "enum"
	INPUT        = "input"
	EXTEND       = "extend"
	DIRECTIVE    = "directive"
)

type Token struct {
	Kind  TokenKind
	Start int
	End   int
	Value string
}

type Lexer func(resetPosition int) (Token, error)

func Lex(s *source.Source) Lexer { _ = "STUB: not implemented"; return *new(Lexer) }

func readName(source *source.Source, position, runePosition int) Token {
	_ = "STUB: not implemented"
	return *new(Token)
}

func readNumber(s *source.Source, start int, firstCode rune, codeLength int) (Token, error) {
	_ = "STUB: not implemented"
	return *new(Token), nil
}

func readDigits(s *source.Source, start int, firstCode rune, codeLength int) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func readString(s *source.Source, start int) (Token, error) {
	_ = "STUB: not implemented"
	return *new(Token), nil
}

func readBlockString(s *source.Source, start int) (Token, error) {
	_ = "STUB: not implemented"
	return *new(Token), nil
}

var splitLinesRegex = regexp.MustCompile("\r\n|[\n\r]")

func blockStringValue(in string) string { _ = "STUB: not implemented"; return "" }

func leadingWhitespaceLen(in string) (n int) { _ = "STUB: not implemented"; return 0 }

func lineIsBlank(in string) bool { _ = "STUB: not implemented"; return false }

func uniCharCode(a, b, c, d rune) rune { _ = "STUB: not implemented"; return 0 }

func char2hex(a rune) int { _ = "STUB: not implemented"; return 0 }

func makeToken(kind TokenKind, start int, end int, value string) Token {
	_ = "STUB: not implemented"
	return *new(Token)
}

func printCharCode(code rune) string { _ = "STUB: not implemented"; return "" }

func readToken(s *source.Source, fromPosition int) (Token, error) {
	_ = "STUB: not implemented"
	return *new(Token), nil
}

func runeAt(body []byte, position int) (code rune, charWidth int) {
	_ = "STUB: not implemented"
	return 0, 0
}

func positionAfterWhitespace(body []byte, startPosition int) (position int, runePosition int) {
	_ = "STUB: not implemented"
	return 0, 0
}

func GetTokenDesc(token Token) string { _ = "STUB: not implemented"; return "" }
