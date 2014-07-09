package test2doc

import (
	"go/ast"
	"go/parser"
	"go/token"
	"regexp"
	"sort"
)

type TestNameExtractor struct {
	nameMatcher *regexp.Regexp
}

func NewTestNameExtractor() TestNameExtractor {
	r, _ := regexp.Compile("Test[^a-z].*")
	return TestNameExtractor{nameMatcher: r}
}

func (e TestNameExtractor) ExtractFromFile(path string) ([]string, error) {
	astFile, err := parser.ParseFile(token.NewFileSet(), path, nil, 0)
	if err != nil {
		return []string{}, err
	}
	namedEntites := toList(astFile.Scope.Objects)
	sort.Sort(namedEntites)
	names := make([]string, 0, len(namedEntites))
	for _, entity := range namedEntites {
		if e.isTestFunction(entity) {
			names = append(names, entity.Name)
		}
	}
	return names, nil
}

func (e TestNameExtractor) isTestFunction(obj *ast.Object) bool {
	decl, isDecl := obj.Decl.(*ast.FuncDecl)
	if isDecl && obj.Kind == ast.Fun {
		return e.isTestFunctionDeclaration(decl)
	}
	return false
}

func (e TestNameExtractor) isTestFunctionDeclaration(decl *ast.FuncDecl) bool {
	return e.isTestName(decl.Name.Name) &&
		doesNotReturnAnything(decl) &&
		hasExactlyOneParameter(decl) &&
		parametersIsOfTestType(decl)
}

func (e TestNameExtractor) isTestName(name string) bool {
	return name == "Test" || e.nameMatcher.MatchString(name)
}

func doesNotReturnAnything(decl *ast.FuncDecl) bool {
	return decl.Type.Results == nil
}

func hasExactlyOneParameter(decl *ast.FuncDecl) bool {
	return decl.Type.Params.NumFields() == 1
}

func parametersIsOfTestType(decl *ast.FuncDecl) bool {
	return isTestType(decl.Type.Params.List[0].Type)
}

func isTestType(e ast.Expr) bool {
	pointer, ok := e.(*ast.StarExpr)
	if !ok {
		return false
	}
	selector, ok := pointer.X.(*ast.SelectorExpr)
	if !ok {
		return false
	}
	leftOfDot, ok := selector.X.(*ast.Ident)
	if !ok {
		return false
	}
	rightOfDot := selector.Sel.Name
	return leftOfDot.Name == "testing" && rightOfDot == "T"
}

type objectList []*ast.Object

func toList(objs map[string]*ast.Object) objectList {
	list := make([]*ast.Object, 0, len(objs))
	for _, obj := range objs {
		list = append(list, obj)
	}
	return list
}

func (objs objectList) Len() int           { return len(objs) }
func (objs objectList) Less(i, j int) bool { return objs[i].Pos() < objs[j].Pos() }
func (objs objectList) Swap(i, j int)      { objs[i], objs[j] = objs[j], objs[i] }
