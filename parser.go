package tryparser

import (
	"github.com/xjslang/xjs/ast"
	"github.com/xjslang/xjs/parser"
	"github.com/xjslang/xjs/token"
)

type TryStatement struct {
	ast.Node

	TryBlock       *ast.BlockStatement
	CatchParameter *ast.Identifier
	CatchBlock     *ast.BlockStatement
	FinallyBlock   *ast.BlockStatement
}

func (ts *TryStatement) String() string {
	out := "try " + ts.TryBlock.String()
	if ts.CatchBlock != nil {
		out += "catch"
		if ts.CatchParameter != nil {
			out += "(" + ts.CatchParameter.String() + ")"
		}
		out += ts.CatchBlock.String()
	}
	if ts.FinallyBlock != nil {
		out += "finally " + ts.FinallyBlock.String()
	}
	return out
}

func ParseStatement(p *parser.Parser, next func(p *parser.Parser) ast.Statement) ast.Statement {
	if p.CurrentToken.Type != token.IDENT || p.CurrentToken.Literal != "try" {
		return next(p)
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

		if !p.ExpectToken(token.RPAREN) {
			return nil
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
