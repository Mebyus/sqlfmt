package parser

import (
	"fmt"

	"github.com/mebyus/sqlfmt/syntax/ast"
)

func (p *Parser) parseTablespaceClause() (*ast.TablespaceClause, error) {
	tablespaceKeyword := p.tok
	p.advance()

	if !p.isIdent() {
		return nil, fmt.Errorf("expected identifier, got [ %v ]", p.tok)
	}
	name := ast.Identifier{
		Token: p.tok,
	}
	p.advance()

	clause := &ast.TablespaceClause{
		Keywords: ast.TablespaceClauseKeywords{
			Tablespace: tablespaceKeyword,
		},
		Name: name,
	}
	return clause, nil
}
