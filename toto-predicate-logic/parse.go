package main

import "regexp"

type StatementPattern string

const (
	CallStatementPattern StatementPattern = "^IdentifierTokenStatementEndToken$"
)

type StatementType string

const (
	DeclarationStatement StatementType = "DeclarationStatement"
	CallStatement        StatementType = "CallStatement"
)

type Statement struct {
	Type_   StatementType `json:"type"`
	Tokens  []*Token      `json:"tokens"`
	Pattern string        `json:"pattern"`
}

func createSignature(t []*Token) string {
	var pattern string
	for _, tok := range t {
		pattern += string(tok.Type_)
	}
	return pattern
}

func parseStatement(l *Lexer) *Statement {
	tokens := l.NextLine()
	if tokens == nil {
		return nil
	}

	signature := createSignature(tokens)

	isCallStmt, _ := regexp.MatchString(signature, string(CallStatementPattern))

	if isCallStmt {
		return &Statement{
			CallStatement,
			tokens,
			signature,
		}
	}

	return &Statement{
		DeclarationStatement,
		tokens,
		signature,
	}
}

func ParseStatements(l *Lexer) []*Statement {
	var stmts []*Statement
	for l.IsNotEmpty() {
		stmt := parseStatement(l)
		if stmt == nil {
			break
		}
		stmts = append(stmts, stmt)
	}
	return stmts
}
