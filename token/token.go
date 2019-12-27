package token

// TokenType is basically the enum value
type TokenType string

// Token is the definition of what a token is
// you can optionally store metadata about the token
// like the line and column number
type Token struct {
	Type    TokenType // the type, kinda like the enum
	Literal string    // the actual character itself
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Indentifiers & literals
	IDENT = "IDENT" // add, foobar, x, y
	INT   = "INT"

	// operators
	ASSIGN   = "="
	PLUS     = "+"
	BANG     = "!"
	ASTERISK = "*"
	MINUS    = "-"
	SLASH    = "/"

	GT = ">"
	LT = "<"
	EQ = "=="
	NOT_EQ = "!="

	// delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"if":     IF,
	"else":   ELSE,
	"true":   TRUE,
	"false":  FALSE,
	"return": RETURN,
}

// LookupIdentifier finds the keyword used in our keyword map
func LookupIdentifier(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
