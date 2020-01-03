package token

// TokenType : Essentially an enum value that's mapped up to the constants below
type TokenType string

// Token is the definition of what a token is
// you can optionally store metadata about the token
// like the line and column number
type Token struct {
	Type    TokenType // the type, kinda like the enum
	Literal string    // the actual character itself
}

// A list of contants that is essentially an enum
const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Indentifiers & literals
	IDENT = "IDENT"
	INT   = "INT"

	// operators
	ASSIGN   = "="
	PLUS     = "+"
	BANG     = "!"
	ASTERISK = "*"
	MINUS    = "-"
	SLASH    = "/"
	GT       = ">"
	LT       = "<"
	EQ       = "=="
	NOT_EQ   = "!="

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

// Keywords are special words reserved for the interpreter and not
// to be used in a variable name
var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"if":     IF,
	"else":   ELSE,
	"true":   TRUE,
	"false":  FALSE,
	"return": RETURN,
}

// LookupIdentifier : This function checks if the input is a reserved word,
// if it isn't, we assum that it's an IDENT
func LookupIdentifier(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
