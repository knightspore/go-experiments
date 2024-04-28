package main

import (
	"regexp"
)

type TokenType string

const (
	IdentifierToken   TokenType = "IdentifierToken"
	AndToken          TokenType = "AndToken"
	NotToken          TokenType = "NotToken"
	ImpliesToken      TokenType = "ImpliesToken"
	StatementEndToken TokenType = "StatementEndToken"
	UnknownToken      TokenType = "UnknownToken"
)

type Pattern string

const (
	PATTERN_IDENTIFER        Pattern = "[A-Za-z0-9_]"
	PATTERN_OP_AND           Pattern = "[\\.]"
	PATTERN_OP_NOT           Pattern = "[~]"
	PATTERN_OP_IMPLIES       Pattern = "[>]"
	PATTERN_EOL              Pattern = "[;]"
	PATTERN_SPACE_OR_NEWLINE Pattern = "[\\s\\r\\t\\n]"
	PATTERN_NEWLINE          Pattern = "[\\r\\n]"
)

type Token struct {
	Loc   *Loc      `json:"loc"`
	Type_ TokenType `json:"type"`
	Value string    `json:"value"`
}

func NewToken(loc *Loc, type_ TokenType, value string) *Token {
	return &Token{
		loc,
		type_,
		value,
	}
}

func Match(pattern Pattern, s string) bool {
	isMatch, _ := regexp.MatchString(string(pattern), s)
	return isMatch
}

func NewTokenFromString(loc *Loc, s string) *Token {
	switch {
	case Match(PATTERN_EOL, s):
		return NewToken(loc, StatementEndToken, s)
	case Match(PATTERN_OP_AND, s):
		return NewToken(loc, AndToken, s)
	case Match(PATTERN_OP_NOT, s):
		return NewToken(loc, NotToken, s)
	case Match(PATTERN_OP_IMPLIES, s):
		return NewToken(loc, ImpliesToken, s)
	case Match(PATTERN_IDENTIFER, s):
		return NewToken(loc, IdentifierToken, s)
	default:
		return NewToken(loc, UnknownToken, s)
	}
}
