package tryparser

import (
	"fmt"
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
	tryTokenType := lb.RegisterTokenType("try")
	lb.UseTokenInterceptor(func(l *lexer.Lexer, next func() token.Token) token.Token {
		ret := next()
		if ret.Literal == "try" {
			ret.Type = tryTokenType
		}
		return ret
	})

	pb.UseStatementInterceptor(func(p *parser.Parser, next func() ast.Statement) ast.Statement {
		tok := p.CurrentToken()
		if tok.Type != tryTokenType {
			return next()
		}
		if !p.ExpectToken(token.LBRACE) {
			p.AddError(fmt.Sprintf("Expected %s, found %s", token.LBRACE, p.CurrentToken().Literal))
			return nil
		}
		peekTok := p.PeekToken()
		stmt := &TryStatement{}
		stmt.TryBlock = p.ParseBlockStatement()
		if peekTok.Literal == "catch" {
			p.NextToken() // consumes catch
			if p.ExpectToken(token.LPAREN) {
				p.NextToken() // consumes (
				tok := p.CurrentToken()
				stmt.CatchParameter = &ast.Identifier{Token: tok, Value: tok.Literal}
			}
			if !p.ExpectToken(token.LBRACE) {
				p.AddError(fmt.Sprintf("Expected %s, found %s", token.LBRACE, p.CurrentToken().Literal))
				return nil
			}
			stmt.CatchBlock = p.ParseBlockStatement()
		}
		if peekTok.Literal == "finally" {
			p.NextToken() // consumes finally
			if !p.ExpectToken(token.LBRACE) {
				return nil
			}
			stmt.FinallyBlock = p.ParseBlockStatement()
		}
		return stmt
	})
}
