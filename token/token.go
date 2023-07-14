package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

var keyword = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keyword[ident]; ok {
		return tok
	}
	return IDENT
}

const (
	ILLEGAL   TokenType = "ILLEGAL"
	EOF       TokenType = "EOF"
	IDENT     TokenType = "IDENT"
	INT       TokenType = "INT"
	ASSIGN    TokenType = "="
	PLUS      TokenType = "+"
	MINUS     TokenType = "-"
	BANG      TokenType = "!"
	ASTERISK  TokenType = "*"
	SLASH     TokenType = "/"
	LT        TokenType = "<"
	GT        TokenType = ">"
	EQ        TokenType = "=="
	NOT_EQ    TokenType = "!="
	LTE       TokenType = "<="
	GTE       TokenType = ">="
	COMMA     TokenType = ","
	SEMICOLON TokenType = ";"
	LPAREN    TokenType = "("
	RPAREN    TokenType = ")"
	LBRACE    TokenType = "{"
	RBRACE    TokenType = "}"
	FUNCTION  TokenType = "FUNCTION"
	LET       TokenType = "LET"
	TRUE      TokenType = "TRUE"
	FALSE     TokenType = "FALSE"
	IF        TokenType = "IF"
	ELSE      TokenType = "ELSE"
	RETURN    TokenType = "RETURN"
)
