package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"

	app "github.com/gjbranham/Text-Finder/internal/application"
	"github.com/gjbranham/Text-Finder/internal/args"
	"github.com/gjbranham/Text-Finder/internal/concurrency"
	out "github.com/gjbranham/Text-Finder/internal/output"
)

func main() {
	out.SetPrinter(&out.Stdout{})

	args, outp, err := args.ProcessArgs(os.Args[0], os.Args[1:])
	if err == flag.ErrHelp {
		out.Print(outp)
		os.Exit(2)
	} else if err != nil {
		out.Print(fmt.Sprintf("Failed to parse command-line arguments: %v", err))
		out.Print(fmt.Sprintf("Info: %v\n", outp))
		os.Exit(1)
	}
	app := app.TextFinder{Args: args, MatchInfo: new(concurrency.MatchInfo)}

	absPath, err := filepath.Abs(app.Args.RootPath)
	if err != nil {
		out.Print(fmt.Sprintf("Fatal error: could not resolve absolute path for '%v'\n", app.Args.RootPath))
	}

	info, err := os.Stat(absPath)
	if err != nil {
		out.Print(fmt.Sprintf("Fatal error: could not get info for path '%v'\n", absPath))
	}

	start := time.Now()

	if info.IsDir() {
		app.FindFiles(absPath)
	} else {
		app.CheckFileForMatch(absPath)
	}

	out.PrintResults(start, app.Args.SearchTerms, app.MatchInfo)
}
