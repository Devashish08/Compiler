package parser

/*
A parser is a program that takes input data (frequently text) and builds a data structure
– often some kind of parse tree, abstract syntax tree or other hierarchical structure,
giving a structural representation of the input while checking for correct syntax.

A parser is often used by the compiler to break the input into smaller elements (tokens)
and then build a parse tree. The parser is a key component in most language implementations,
which transform the source code into executable code.
*/

import (
	"compiler/ast"
	"compiler/lexer"
	"compiler/token"
	"fmt"
)

type Parser struct {
	l *lexer.Lexer // pointer to an instance of the lexer that the parser uses to get tokens from the input
	// curToken and peekToken are used to look at current and next tokens in the input
	// This is a common pattern in parsers: you look at the current token and make decisions based on it,
	// and then you advance the current and peek tokens and look at them to make further decisions.
	errors    []string
	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l, errors: []string{}}
	// Read two tokens, so curToken and peekToken are both set
	// This is a common pattern in parsers: you read two tokens, so both curToken and peekToken are set.
	// This allows you to make decisions based on the current and next tokens.
	// For example, you might look at the current token and decide what to do based on the next token.
	// Then you advance both curToken and peekToken, and repeat the process.
	// This is a common pattern in parsers, and it's called a two-token lookahead.
	// It's a simple and effective way to parse complex languages.
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

// func (p *Parser) ParseProgram() *ast.Program {
// 	return nil
// }

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}
	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	default:
		return nil
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curToken}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}
