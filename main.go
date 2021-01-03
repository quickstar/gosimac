package main

import (
	"github.com/carlmjohnson/versioninfo"
	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
	"github.com/quickstar/wally/internal/cmd"
)

func main() {
	if err := pterm.DefaultBigText.WithLetters(
		putils.LettersFromStringWithStyle("w", pterm.NewStyle(pterm.FgCyan)),
		putils.LettersFromStringWithStyle("al", pterm.NewStyle(pterm.FgLightMagenta)),
		putils.LettersFromStringWithStyle("ly", pterm.NewStyle(pterm.FgLightRed)),
	).Render(); err != nil {
		_ = err
	}

	pterm.Description.Printf("wally %s\n", versioninfo.Short())

	cmd.Execute()
}
