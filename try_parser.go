package tryparser

import (
	"strings"

	"github.com/xjslang/xjs/ast"
	"github.com/xjslang/xjs/parser"
	"github.com/xjslang/xjs/token"
)

type TryStatement struct {
	TryBlock       *ast.BlockStatement
	CatchParameter *ast.Identifier
	CatchBlock     *ast.BlockStatement
	FinallyBlock   *ast.BlockStatement
}

func (ts *TryStatement) WriteTo(b *strings.Builder) {
	b.WriteString(" try ")
	ts.TryBlock.WriteTo(b)
	if ts.CatchBlock != nil {
		b.WriteString(" catch ")
		if ts.CatchParameter != nil {
			b.WriteRune('(')
			ts.CatchParameter.WriteTo(b)
			b.WriteRune(')')
		}
		ts.CatchBlock.WriteTo(b)
	}
	if ts.FinallyBlock != nil {
		b.WriteString(" finally ")
		ts.FinallyBlock.WriteTo(b)
	}
}

func ParseTryStatement(p *parser.Parser, next func() ast.Statement) ast.Statement {
	if p.CurrentToken.Type != token.IDENT || p.CurrentToken.Literal != "try" {
		return next()
	}
	if !p.ExpectToken(token.LBRACE) {
		return nil
	}
	stmt := &TryStatement{}
	stmt.TryBlock = p.ParseBlockStatement()
	if p.PeekToken.Type == token.IDENT && p.PeekToken.Literal == "catch" {
		p.NextToken() // consumes catch
		if p.ExpectToken(token.LPAREN) {
			p.NextToken() // consumes (
			stmt.CatchParameter = &ast.Identifier{Token: p.CurrentToken, Value: p.CurrentToken.Literal}
		}
		if !p.ExpectToken(token.LBRACE) {
			return nil
		}
		stmt.CatchBlock = p.ParseBlockStatement()
	}
	if p.PeekToken.Type == token.IDENT && p.PeekToken.Literal == "finally" {
		p.NextToken() // consumes finally
		if !p.ExpectToken(token.LBRACE) {
			return nil
		}
		stmt.FinallyBlock = p.ParseBlockStatement()
	}
	return stmt
}
