package parser

import (
	"github.com/mebyus/sqlfmt/syntax/ast"
)

func (p *Parser) parseTablespaceClause() (*ast.TablespaceClause, error) {
	tablespaceKeyword := p.tok
	p.advance() // consume TABLESPACE

	name, err := p.consumeIdentifier()
	if err != nil {
		return nil, err
	}

	clause := &ast.TablespaceClause{
		Keywords: ast.TablespaceClauseKeywords{
			Tablespace: tablespaceKeyword,
		},
		Name: name,
	}
	return clause, nil
}
