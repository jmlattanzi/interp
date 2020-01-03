// Package lexer only supports ascii btw
package lexer

import "github.com/jmlattanzi/interp/token"

/*

Lexer

Also known as the tokenizer, takes in string input and turns it into a series of
identifiable objects, or tokens.

A lexer object is created with an input (the source code), a cursor that is the
current character, and a read position that looks forward to the next characters.
The lexer also keeps track of our current character byte.

*/

// Lexer : The lexer contains our input source string, the position of our current
// character, and a cursor pointing forward in the input
type Lexer struct {
	input        string
	position     int // points to the current char
	readPosition int // points to the next char
	ch           byte
}

// New : Create a new instance of a lexer, providing the source code as a string
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// readChar : This function first checks if the readPosition (future cursor) is beyond the length of the input.
// If it is, we will return 0, indicating the end of the file. Otherwise, we will set the current character byte
// to the the current readPosition's respective character. Next, the position is updated to be synced with the readPosition.
// Lastly, readPosition is incremented.
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

// peekChar : This function checks the character in input at the current readPosition
// which is always further along in the string than the position.
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

// readIdentifier : This function returns the identifier that the cursor is currently reading.
// It stores the position locally, then checks if the current character is a letter.
// While it returns true, the next character is read and it increments the positions.
// It then returns the segment of input from the local `position` to the lexer's current `position`
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// readNumber : Does the same thing as readIdentifier expect it checks for a digit.
// It checks for a number and returns the segment of code from the lexer's input.
func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// isLetter : Checks if the current byte is a character
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// isDigit : Checks if the current byte is a digit
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

// newToken : A helper function to return a new Token object.
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// skipWhitespace : Checks if the current character is whitespace and calls readChar
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// NextToken : This function is vital to the lexer's operational existence.
// It returns the next token in the input. It does this by checking if the
// current character is a symbol and returning a new token.
// The default action is to check if the current character is a letter.
// If it is, the token's literal value is set to the output of
// `readIdentifier()` and the type is set to the return result of
// token's `LookupIdentifier()` function.
// The last action this method makes is the read the next character with a
// call to `readChar()`
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.EQ, Literal: literal}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.NOT_EQ, Literal: literal}
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdentifier(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}
