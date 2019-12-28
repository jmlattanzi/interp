package ast

import "github.com/jmlattanzi/interp/token"

// Node is needed by all of our interfaces
type Node interface {
	TokenLiteral() string
}

// Statement is the structure of a statement
type Statement interface {
	Node
	statementNode()
}

// Expression is the structure of our expression
type Expression interface {
	Node
	expressionNode()
}

// Program consists of a list of statements
type Program struct {
	Statements []Statement
}

// TokenLiteral returns the token literal of the expression
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// LetStatement is the structure of a let statment
type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}

// TokenLiteral returns the token literal of the statement
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

// Identifier is essentially a token
type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode() {}

// TokenLiteral returns the token literal of the statement
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
