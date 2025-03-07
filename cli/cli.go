package cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/swichupcb/xstruct/cli/config"
	"github.com/swichupcb/xstruct/cli/generator"
	"github.com/swichupcb/xstruct/cli/parser"
)

// Environment represents the xstruct environment.
type Environment struct {
	DirPath string // The relative path to the directory that will be extracted.
	Pkg     string // The output file's package.
	Global  bool   // Whether to extract global variables and constants.
	Funcs   bool   // Whether to extract function declarations.
	Sort    bool   // Whether to sort the structs.
}

// CLI runs xstruct from a Command Line Interface and returns the exit status.
func CLI() int {
	var env Environment

	if err := env.parseArgs(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)

		return 2
	}

	if _, err := env.run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)

		return 1
	}

	return 0
}

// parseArgs parses the provided command line arguments.
func (e *Environment) parseArgs() error {
	var (
		info   = flag.Bool("a", false, "Use -a to receive information about xstruct's author.")
		dir    = flag.String("d", "", "The relative path to the directory that will be extracted.")
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
		return fmt.Errorf("you must specify a directory to extract structs from.")
	}

	e.DirPath = *dir
	e.Pkg = *pkg
	e.Global = *global
	e.Funcs = *funcs
	e.Sort = *sort

	return nil
}

// Run runs xstruct programmatically using the given Environment's YMLPath.
func (e *Environment) run() (string, error) {
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
