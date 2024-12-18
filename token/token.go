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

	// ==
	ITSGIVING = "ITSGIVING"

	// !=
	CAP = "CAP"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	LOCKIN              = "LOCKIN"              // let
	SIGMA               = "SIGMA"               // True
	SUS                 = "SUS"                 // False
	CHATISTHISREAL      = "CHATISTHISREAL"      // if
	MID                 = "MID"                 // else
	PUTTHEFRIESINTHEBAG = "PUTTHEFRIESINTHEBAG" // return

)
