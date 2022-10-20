package tui

import (
	"github.com/pterm/pterm"
	"math"
)

func ShowMainHeader() {
	pterm.DefaultCenter.Println(pterm.DefaultHeader.WithFullWidth().WithBackgroundStyle(pterm.NewStyle(pterm.BgBlue)).WithTextStyle(pterm.NewStyle(pterm.FgBlack)).Sprint("FerBench"))
}

func ShowTitle(title string) {
	pterm.DefaultCenter.Println(pterm.DefaultSection.Sprint(title))
}

func ShowBar(title string, runLength float64) (*pterm.ProgressbarPrinter, error) {
	return pterm.DefaultProgressbar.WithMaxWidth(1000).WithTotal(int(runLength * 2)).WithRemoveWhenDone(true).WithTitle(title).WithShowCount(false).WithShowElapsedTime(false).Start()
}

func ShowScore(text string, score float64) {
	pterm.DefaultCenter.Println(pterm.NewRGB(15, 199, 209).Sprint(text, math.Round(score)))
}
