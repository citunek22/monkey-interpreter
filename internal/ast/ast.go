package ast

import (
	"bytes"
	"monkey/internal/token"
)

// Node je zakladni rozhrani pro kazdy uzel v AST
type Node interface {
	TokenLiteral() string
	String() string
}

// Statement reprezentuje prikaz (nevraci hodnotu, napr. "let x = 5;")
type Statement interface {
	Node
	statementNode()
}

// Expression reprezentujje vyraz (vraci hodnotu, napr. "5 + 5", "x")
type Expression interface {
	Node
	expressionNode()
}

// Program je korenovy uzel kazdeho programu v nasem jazyce
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

func (p *Program) String() string {
	var out bytes.Buffer
	for _, s := range p.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

// LetStatement reprezentuje prikaz deklerace promenne: let x = 5;
type LetStatement struct {
	Token token.Token // token.LET
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }
func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")
	return out.String()
}

// Identifier reprezentuje identifikator (nazev promenne)
type Identifier struct {
	Token token.Token // token.IDENT
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string       { return i.Value }

// ReturnStatement reprezentuje prikaz return: return 5;
type ReturnStatement struct {
	Token       token.Token // token.RETURN
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(lsTokenLiteralSafe(rs) + " ")
	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")
	return out.String()
}

func lsTokenLiteralSafe(rs *ReturnStatement) string {
	return rs.Token.Literal
}

// ExpressionStatement slouzi k obaleni vyrazy, aby mohl stat jako samostatny prikaz
type ExpressionStatement struct {
	Token      token.Token // prvni token vyrazu
	Expression Expression
}

func (es *ExpressionStatement) statementNode()       {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

// IntegerLiteral reprezentuje ciselny literal, napr. "6"
type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) expressionNode()      {} // I accidentally wrote explosion instead of expression
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *IntegerLiteral) String() string

// PrefixExpression reprezentuje unarni operaci, napr. "!5" nebo "-15"
type PrefixExpression struct {
	Token    token.Token // Prefix token, napr, '!' nebo '-'
	Operator string
	Right    Expression
}

func (pe *PrefixExpression) expressionNode()      {}
func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.Literal }
func (pe *PrefixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")
	return out.String()
}

// InfixExpression reprezentuje binarni operaci, napr. "22 + 13"
type InfixExpression struct {
	Token    token.Token // Operatorovy token, napr. '+'
	Left     Expression
	Operator string
	Right    Expression
}

func (ie *InfixExpression) expressionNode()      {}
func (ie *InfixExpression) TokenLiteral() string { return ie.Token.Literal } // I wrote literature atleast it looks like i still go to school even if there are typos but im from czechia so
func (ie *InfixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString(" " + ie.Operator + " ") // It seems like i want to be a doctor because operation instead of operator but sike i want to be a pilot or atleast thats my dream but remeber like our teacher said go big or go home, or aim for the stars
	out.WriteString(ie.Right.String())
	out.WriteString(")")
	return out.String()
}
