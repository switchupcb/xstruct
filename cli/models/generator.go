package models

import (
	"github.com/dave/dst"
)

// Generator represents a code generator.
type Generator struct {
	GoFiles   []string
	ASTFiles  []*dst.File
	ASTDecls  []*dst.GenDecl
	FuncDecls []*dst.FuncDecl
}
