package main

import (
	"flag"
	"fmt"
	"log/slog"

	"github.com/taylormonacelli/cathugger"
	"github.com/taylormonacelli/goldbug"
)

var (
	verbose   bool
	logFormat string
)

func main() {
	flag.BoolVar(&verbose, "verbose", false, "Enable verbose output")
	flag.BoolVar(&verbose, "v", false, "Enable verbose output (shorthand)")
	flag.StringVar(&logFormat, "log-format", "", "Log format (text or json)")

	flag.Parse()

	if verbose || logFormat != "" {
		if logFormat == "json" {
			goldbug.SetDefaultLoggerJson(slog.LevelDebug)
		} else {
			goldbug.SetDefaultLoggerText(slog.LevelDebug)
		}
	}

	args := flag.Args()
	if len(args) != 2 {
		fmt.Printf("Usage: %s <service> <region>\n", "cathugger")
		return
	}
	service := args[0]
	region := args[1]

	slog.Debug("args", "service", service)
	slog.Debug("args", "region", region)

	url := cathugger.GetAWSConsoleUrl(region, service)
	if url != "" {
		cathugger.RunCmdOpenUrl(url)
	}
}
