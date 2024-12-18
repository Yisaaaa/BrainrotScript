package token

// Defining Tokens

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers
	IDENT = "IDENT"
	INT   = "INT"

	// Operators
	PLUS     = "+"
	MINUS    = "-"
	ASSIGN   = "="
	SLASH    = "/"
	ASTERISK = "*"
	BANG     = "!"

	LT = "<"
	GT = ">"

	EQ     = "EQ"
	NOT_EQ = "NOT_EQ"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	LET      = "LET"
	FUNCTION = "FUNCTION"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

var keywords = map[string]TokenType{
	"lock in":                  LET,
	"slay":                     FUNCTION,
	"sigma":                    TRUE,
	"cap":                      FALSE,
	"chat is this real?":       IF,
	"only in ohio":             ELSE,
	"put the fries in the bag": RETURN,
	"its giving":               EQ,
	"feels different":          NOT_EQ,
}

func LookUpIdent(ident string) TokenType {
	if tokenType, ok := keywords[ident]; ok {
		return tokenType
	}
	return IDENT
}
