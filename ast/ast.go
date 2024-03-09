package ast

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

// TokenLiteral returns the literal value of the token associated with the Program node.
// This method is used for debugging and testing purposes.
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}
