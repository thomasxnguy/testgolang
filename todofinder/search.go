package todofinder

import (
	"fmt"
	"go/ast"
	"go/build"
	"go/parser"
	"go/token"
	"strings"
	. "todofinder/todofinder/error"
	"path/filepath"
)

//search result structure
type SearchResult struct {
	fileName string `json:"file"`
	position int    `json:"pos"`
	comment  string `json:"com"`
	Error    *Error
}

//Return a verbose message corresponding to the search result
func (s *SearchResult) ToString() string {
	return fmt.Sprintf("%s:%v:\n%s\n", s.fileName, s.position, s.comment)
}

func ImportPkg(path, dir string) (*build.Package, *Error) {

	//TODO Optimisation
	p, err := build.Import(path, dir, build.ImportComment)
	if err != nil {
		return nil, &Error{PACKAGE_NOT_FOUND, err}
	}
	//if p.BinaryOnly &&
	if p.IsCommand() {
		return nil, &Error{NO_SOURCE, err}
	}

	return p, nil
}

func ExtractPattern(p *build.Package, pattern string, resultChannel chan *SearchResult) {
	for _, f := range p.GoFiles {
		fname := filepath.Join(p.Dir, f)
		fset := token.NewFileSet()
		f, err := parser.ParseFile(fset, fname, nil, parser.ParseComments)
		if err != nil {
			resultChannel <- &SearchResult{"", 0, "", &Error{SOURCE_NOT_READABLE, err}}
			return
		}
		cmap := ast.NewCommentMap(fset, f, f.Comments)
		for n, cgs := range cmap {
			f := fset.File(n.Pos())
			for _, cg := range cgs {
				t := cg.Text()
				if strings.Contains(t, pattern) {
					resultChannel <- &SearchResult{fname, f.Position(cg.Pos()).Line, t, nil}
				}
			}
		}
	}
	//End of function, nil object will terminate the routine
	resultChannel <- nil
}
