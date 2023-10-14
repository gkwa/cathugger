package main

import (
	"log/slog"

	"github.com/taylormonacelli/cathugger"
	"github.com/taylormonacelli/eachgoose"
	"github.com/taylormonacelli/goldbug"
)

func main() {
	goldbug.SetDefaultLoggerText(slog.LevelDebug)

	resources := eachgoose.ParseArgs()

	cathugger.Execute(resources)
}
