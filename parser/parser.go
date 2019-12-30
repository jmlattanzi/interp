package parser

import (
	"github.com/jmlattanzi/interp/token"

	"github.com/jmlattanzi/interp/ast"
	"github.com/jmlattanzi/interp/lexer"
)

/*

Parser

The parser takes the token's created by our Lexer and turns them into statements defined by the AST
package. It does this in a very similar way to how the lexer works. Meaning it uses two cursors.
One pointing to the current token, and the another pointing to the future.

The parser generates this AST and returns it for further evaluation.

*/

// Parser : An object containing the lexer and current/future tokens.
type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
}

// New : This function creates a new instance of Parser using our existing lexer as input
func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	// read two tokens so curToken and peekToken are set
	p.nextToken()
	p.nextToken()

	return p
}

// ParseProgram : Parses the program into a slice of statements.
// First, we create a new, empty instance of `ast.Program{}` and set the
// program's statements to be an empty slice of `ast.Statement`
// While the current token is not of the type `EOF`, the current statement is parsed
// via the `parseStatement` function. The returned statement is appened to the
// programs statements and the `nextToken()` is called.
// Lastly the program object itself is returned.
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

// nextToken : Shifts the cursors. Sets the current token to the peek token
// and sets the peak token the result of the lexer's `NextToken()` function
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

// parseStatement : This function checks the current token's type and returns the result
// if it's appropriate parse function.
func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	default:
		return nil
	}
}

// parseLetStatement : This function sets a local statement to the be of the AST's type for
// LET statements. It then goes on to check if the next token is an identifier. If the statement
// is really a LET, the next token should always be an identifier. We do this by calling the parser's
// `expectPeek()` method. Next it checks to make sure the next token is of the ASSIGN type.
// Lastly we return the full statement after hitting a semicolon.
func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curToken}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	// TODO: we're skipping the expression until we encoutner a semicolon
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

// curTokenIs : Checks if the current token's type is what we were expecting
func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

// peekTokenIs : Checks if the next token's type is what we were expecting
func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

// expectPeek : Checks if the peek token is of a certain type. If it evaluates to
// true, we advance to the next token and return true.
func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	}
	return false
}
