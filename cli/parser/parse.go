package parser

import (
	"fmt"
	"go/parser"
	"go/token"

	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
	"github.com/switchupcb/xstruct/cli/models"
)

// Parse parses structs from .go files.
func Parse(gen *models.Generator, global bool, funcs bool) error {
	gen.ASTFiles = make([]*dst.File, len(gen.GoFiles))
	for _, path := range gen.GoFiles {
		astFile, err := astParseFilepath(path)
		if err != nil {
			return fmt.Errorf("parse: %w", err)
		}

		gen.ASTFiles = append(gen.ASTFiles, astFile)
		gen.ASTDecls = append(gen.ASTDecls, astParseDecls(astFile, global)...)

		if funcs {
			gen.FuncDecls = append(gen.FuncDecls, astParseFuncs(astFile)...)
		}
	}

	return nil
}

// astParseFilepath parses a filepath to an *ast.File (with comments intact).
func astParseFilepath(path string) (*dst.File, error) {
	fileset := token.NewFileSet()
	file, err := decorator.ParseFile(fileset, path, nil, parser.ParseComments)

	if err != nil {
		return nil, fmt.Errorf("astParseFilepath: %w", err)
	}

	return file, nil
}

// astParseDecls parses type declarations from an ast.File (with comments).
func astParseDecls(f *dst.File, global bool) []*dst.GenDecl {
	decls := make([]*dst.GenDecl, 0, len(f.Decls))

	for _, node := range f.Decls {
		switch decl := node.(type) {
		case *dst.GenDecl:
			if decl.Tok == token.TYPE {
				decls = append(decls, decl)
			} else if global && (decl.Tok == token.VAR || decl.Tok == token.CONST) {
				decls = append(decls, decl)
			}
		}
	}

	return decls
}

// astParseFuncs parses func declarations from an ast.File (with comments).
func astParseFuncs(f *dst.File) []*dst.FuncDecl {
	decls := make([]*dst.FuncDecl, 0, len(f.Decls))

	for _, node := range f.Decls {
		switch decl := node.(type) {
		case *dst.FuncDecl:
			decls = append(decls, decl)
		}
	}

	return decls
}
