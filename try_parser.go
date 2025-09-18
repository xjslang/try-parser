package tryparser

import (
	"strings"

	"github.com/xjslang/xjs/ast"
	"github.com/xjslang/xjs/lexer"
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

func Plugin(pb *parser.Builder) {
	lb := pb.LexerBuilder
	tryTokenType := lb.RegisterTokenType("TryStatement")
	lb.UseInterceptor(func(l *lexer.Lexer, next func() token.Token) token.Token {
		ret := next()
		if ret.Literal == "try" {
			ret.Type = tryTokenType
		}
		return ret
	})

	pb.UseStatementInterceptor(func(p *parser.Parser, next func() ast.Statement) ast.Statement {
		if p.CurrentToken.Type != tryTokenType {
			return next()
		}
		stmt := &TryStatement{}
		if !p.ExpectToken(token.LBRACE) {
			return nil
		}
		stmt.TryBlock = p.ParseBlockStatement()
		if p.PeekToken.Literal == "catch" {
			p.NextToken()
			if !p.ExpectToken(token.LBRACE) {
				return nil
			}
			stmt.CatchBlock = p.ParseBlockStatement()
		}
		if p.PeekToken.Literal == "finally" {
			p.NextToken()
			if !p.ExpectToken(token.LBRACE) {
				return nil
			}
			stmt.FinallyBlock = p.ParseBlockStatement()
		}
		return stmt
	})
}
