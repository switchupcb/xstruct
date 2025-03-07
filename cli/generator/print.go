package generator

import (
	"bytes"
	"strings"

	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
)

// dst is sadge
// https://github.com/dave/dst/issues/24
func printDecls(astDecls []*dst.GenDecl, funcDecls []*dst.FuncDecl) []byte {
	var output bytes.Buffer

	// create a new ast file, print it, then remove its package declaration.
	prelude := `package main`
	for _, decl := range astDecls {
		newFile := &dst.File{
			Name: dst.NewIdent("main"),
			Decls: []dst.Decl{
				decl,
			},
		}

		var newBuffer strings.Builder
		decorator.Fprint(&newBuffer, newFile)
		newSrc := newBuffer.String()
		output.WriteString(strings.TrimSpace(newSrc[len(prelude):]))
		output.WriteString("\n\n")
	}

	for _, decl := range funcDecls {
		newFile := &dst.File{
			Name: dst.NewIdent("main"),
			Decls: []dst.Decl{
				decl,
			},
		}

		var newBuffer strings.Builder
		decorator.Fprint(&newBuffer, newFile)
		newSrc := newBuffer.String()
		output.WriteString(strings.TrimSpace(newSrc[len(prelude):]))
		output.WriteString("\n\n")
	}

	return output.Bytes()
}
