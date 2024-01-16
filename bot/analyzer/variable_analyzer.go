package analyzer

import (
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
	"unicode"

	"github.com/mrinalxdev/CodeSage/models"
)

type VariableNameAnalyzer struct{}

func (a *VariableNameAnalyzer) Analyze(code string) []models.Comment {
	var comments []models.Comment

	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "source.go", code, parser.ParseComments)
	if err != nil {
		return comments
	}

	ast.Inspect(file, func(n ast.Node) bool {
		if n == nil {
			return false
		}

		switch x := n.(type) {
		case *ast.Ident:
			if isShortName(x.Name){
				comments = append(comments, models.Comment{
					Message : "Short variable name detected :" + x.Name,
					Line : fset.Position(x.Pos()).Line, 
				})
			}

			if !isMeaningfulName(x.Name) {
				comments = append(comments, models.Comment{
					Message: "Non-meaningful variable name detected : " + x.Name,
					Line : fset.Position(x.Pos()).Line,
				})
			}
		}

		return true
	})

	return comments
}

func isShortName(name string) bool {
	return len(name) <= 2
}

func isMeaningfulName(name string) bool {
	hasUppercase := strings.IndexFunc(name, unicode.IsUpper) >= 0
	hasLowercase := strings.IndexFunc(name, unicode.IsLower) >= 0

	return hasUppercase && hasLowercase
}