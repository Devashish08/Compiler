package ast

import "compiler/token"

// Node represents a node in the abstract syntax tree (AST).
type Node interface {
	TokenLiteral() string
}

// Statement represents a statement node in the AST.
type Statement interface {
	Node
	statementNode()
}

// Expression represents an expression node in the AST.
type Expression interface {
	Node
	expressionNode()
}

// Program is a wrapper around a slice of statements in the AST.
type Program struct {
	Statements []Statement
}

type ReturnStatement struct {
	Token       token.Token // the token.RETURN token (to hold the identifier of the binding)
	ReturnValue Expression  // the value the function is returning (to hold the expression that produces the value)
}

// TokenLiteral returns the literal value of the token associated with the Program node.
// This method is used for debugging and testing purposes.
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type LetStatement struct {
	Token token.Token // the token.LET token
	Name  *Identifier // the name of the variable (to hold the identifier of the binding)
	Value Expression  // the value the variable is being assigned (to hold the expression that produces the value)
}

type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

// statementNode is a marker for the Statement interface
// TokenLiteral returns the literal value of the token associated with the LetStatement node.

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

// the value of the identifier
// expressionNode is a marker for the Expression interface
// TokenLiteral returns the literal value of the token associated with the Identifier node.

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

// ReturnStatement represents a return statement in the AST.
// It holds the token.RETURN token and the expression that produces the value the function is returning.
// TokenLiteral returns the literal value of the token associated with the ReturnStatement node.

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
