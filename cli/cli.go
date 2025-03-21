package cli

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/switchupcb/xstruct/cli/config"
	"github.com/switchupcb/xstruct/cli/generator"
	"github.com/switchupcb/xstruct/cli/parser"
)

// Environment represents the xstruct environment.
type Environment struct {
	// Dirpath represents the relative path to the directory that is extracted.
	DirPath string

	// Pkg represents the output file's package.
	Pkg string

	// Global represents an option which determines whether to extract global variables and constants.
	Global bool

	// Funcs represents an option which determines whether to extract function declarations.
	Funcs bool

	// Sort represents an option which determines whether to sort the extracted objects.
	Sort bool
}

const (
	osExitCodeSuccess    = 0
	osExitCodeError      = 1
	osExitCodeErrorShell = 2
)

// CLI runs xstruct from a Command Line Interface and returns an exit status.
func CLI() int {
	var env Environment

	if err := env.parseArgs(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)

		return osExitCodeErrorShell
	}

	if _, err := env.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)

		return osExitCodeError
	}

	return osExitCodeSuccess
}

// parseArgs parses the provided command line arguments.
func (e *Environment) parseArgs() error {
	var (
		info   = flag.Bool("a", false, "Use -a to for information about xstruct's author.")
		dir    = flag.String("d", "", "Specify the relative path to the directory that is extracted.")
		pkg    = flag.String("p", "xstruct", "Use -p to set the output file's package.")
		global = flag.Bool("g", false, "Use -g to extract global variables and constants.")
		funcs  = flag.Bool("f", false, "Use -f to extract function declarations.")
		sort   = flag.Bool("s", false, "Use -s to sort the extracted objects.")
	)

	flag.Parse()

	if *info {
		fmt.Println("xstruct developed by switchupcb: https://switchupcb.com/")

		os.Exit(0)
	}

	if *dir == "" {
		return errors.New("you must specify a directory to extract structs from with -d")
	}

	e.DirPath = *dir
	e.Pkg = *pkg
	e.Global = *global
	e.Funcs = *funcs
	e.Sort = *sort

	return nil
}

// Run runs xstruct programmatically using the given Environment's YMLPath.
func (e *Environment) Run() (string, error) {
	gen, err := config.LoadFiles(e.DirPath)
	if err != nil {
		return "", fmt.Errorf("%w", err)
	}

	if err = parser.Parse(gen, e.Global, e.Funcs); err != nil {
		return "", fmt.Errorf("%w", err)
	}

	code, err := generator.Generate(gen, e.Pkg, e.Sort)

	defer fmt.Println(code)

	if err != nil {
		return code, fmt.Errorf("%w", err)
	}

	return code, nil
}
