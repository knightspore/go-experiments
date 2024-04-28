package main

import (
	"os"
	"regexp"
)

type Lexer struct {
	filePath *string
	source   *string
	cur      int
	bol      int
	row      int
}

func NewLexer(filePath string) *Lexer {
	data, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	source := string(data)

	return &Lexer{
		filePath: &filePath,
		source:   &source,
		cur:      0,
		bol:      0,
		row:      0,
	}
}

func (l *Lexer) Reset() {
	l.cur = 0
	l.bol = 0
	l.row = 0
}

func (l *Lexer) col() int {
	return l.cur - l.bol
}

func (l *Lexer) IsNotEmpty() bool {
	return l.cur < len(*l.source)
}

func (l *Lexer) Match(pattern Pattern, s string) bool {
	isMatch, _ := regexp.MatchString(string(pattern), s)
	return isMatch
}

func (l *Lexer) MatchCurrent(pattern Pattern) bool {
	return l.Match(pattern, string((*l.source)[l.cur]))
}

func (l *Lexer) ChopChar() {
	if l.IsNotEmpty() {
		ch := (*l.source)[l.cur]
		l.cur++
		if l.Match(PATTERN_NEWLINE, string(ch)) {
			l.bol = l.cur
			l.row++
		}
	}
}

func (l *Lexer) TrimLeft() {
	if l.IsNotEmpty() {
		if l.MatchCurrent(PATTERN_SPACE_OR_NEWLINE) {
			l.ChopChar()
			l.TrimLeft()
		}
	}
}

func (l *Lexer) DropLine() {
	if l.IsNotEmpty() && !l.MatchCurrent(PATTERN_SPACE_OR_NEWLINE) {
		l.ChopChar()
		l.DropLine()
	} else if l.IsNotEmpty() {
		l.ChopChar()
	}
}

func (l *Lexer) NextToken() *Token {
	l.TrimLeft()
	if !l.IsNotEmpty() {
		return nil
	}

	loc := NewLoc(l.row, l.col(), l.filePath)
	start := l.cur
	l.ChopChar()

	if l.IsNotEmpty() {
		for !l.MatchCurrent(PATTERN_SPACE_OR_NEWLINE) && !l.MatchCurrent(PATTERN_EOL) {
			l.ChopChar()
		}
		value := (*l.source)[start:l.cur]
		return NewTokenFromString(loc, value)
	}

	return nil
}

func (l *Lexer) Tokens() []*Token {
	var tokens []*Token
	for l.IsNotEmpty() {
		tok := l.NextToken()
		if tok == nil {
			break
		}
		tokens = append(tokens, tok)
	}
	l.Reset()
	return tokens
}

func (l *Lexer) NextLine() []*Token {
	var tokens []*Token
	for l.IsNotEmpty() {
		tok := l.NextToken()
		if tok == nil {
			break
		}
		if tok.Type_ == StatementEndToken {
			tokens = append(tokens, tok)
			break
		}
		tokens = append(tokens, tok)
	}
	return tokens
}

func (l *Lexer) Lines() [][]*Token {
	var lines [][]*Token
	for l.IsNotEmpty() {
		line := l.NextLine()
		if line == nil {
			break
		}
		lines = append(lines, line)
	}
	return lines
}
