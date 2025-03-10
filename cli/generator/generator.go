package generator

import (
	"bytes"
	"fmt"
	"sort"

	"github.com/dave/dst"
	"github.com/switchupcb/xstruct/cli/models"
	"golang.org/x/tools/imports"
)

// Generate generates code.
func Generate(gen *models.Generator, pkg string, srt bool) (string, error) {
	if srt {
		sort.Slice(gen.ASTDecls, func(i, j int) bool {
			return gen.ASTDecls[i].Specs[0].(*dst.TypeSpec).Name.Name < gen.ASTDecls[j].Specs[0].(*dst.TypeSpec).Name.Name //nolint:forcetypeassert
		})

		sort.Slice(gen.FuncDecls, func(i, j int) bool {
			return gen.FuncDecls[i].Name.String() < gen.FuncDecls[j].Name.String()
		})
	}

	content := AstWriteDecls(pkg, gen.ASTDecls, gen.FuncDecls)

	// imports
	importsdata, err := imports.Process("", content, nil)
	if err != nil {
		return string(content), fmt.Errorf("an error occurred while formatting the generated code.\n%w", err)
	}

	return string(importsdata), nil
}

// AstWriteDecls writes ast.GenDecl to a file with a specified package.
func AstWriteDecls(pkg string, astDecls []*dst.GenDecl, funcDecls []*dst.FuncDecl) []byte {
	var b bytes.Buffer

	b.WriteString("package " + pkg + "\n\n")
	b.Write(PrintDecls(astDecls, funcDecls))

	return b.Bytes()
}
