package ast

import "github.com/jmlattanzi/interp/token"

/*

AST (Abstract Syntax Tree)

The AST is used to turn parsed/tokenized input into individual expressions and statements.
A combination of statements and expressions make up a program. That array of statements and expressions
is ultimately evaluated.

*/

// Node : A single node that implements the TokenLiteral method
type Node interface {
	TokenLiteral() string
}

// Statement : The defining structure for a Statement, such as a LET statement.
// A statement needs to implement everything from Node as well as the statementNode() function.
type Statement interface {
	Node
	statementNode()
}

// Expression : The defining structure for an Expression, such as `x + y`.
// An expression needs to implement everything from Node as well as the expressionNode() function.
type Expression interface {
	Node
	expressionNode()
}

// Program : A program is a list of statements and expressions parsed from source code.
type Program struct {
	Statements []Statement
}

// TokenLiteral : If the amount of parsed statements is greater than 0,
// return the TokenLiteral of the statement. Otherwise return an empty string.
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

// LetStatement : The structure of a LET statement. Contains the token, name, and value of the statement.
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
