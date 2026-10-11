package token

// TokenType reprezentuje typ tokenu (napr. INT, IDENT, PLUS)
type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	// Specialni typy
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifikatory a literaly
	IDENT = "IDENT" // pr. x, y, add, mujVysledek
	INT   = "INT"   // pr. 1343456

	// Operatory
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*" // Sorry for breaking the line i know its bad for people with ocd but idk why its like this
	SLASH    = "/"

	LT = "<"
	GT = ">"

	EQ     = "=="
	NOT_EQ = "!="

	// Oddelovace
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Klicova slova
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

// keywords mapuje rezervovana slova na jejich TokenType
var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

// LookupIdent kontroluje, zda identifikator neni klicovym slovem
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
