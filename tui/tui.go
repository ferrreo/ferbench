package tui

import (
	osRelease "github.com/acobaugh/osrelease"
	"github.com/pterm/pterm"
	cpuinfo "github.com/shirou/gopsutil/v3/cpu"
	hostinfo "github.com/shirou/gopsutil/v3/host"
	"math"
	"runtime"
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

func ShowOSInfo() error {
	if runtime.GOOS != "linux" {
		platform, family, version, err := hostinfo.PlatformInformation()
		if err != nil {
			return err
		}
		pterm.DefaultCenter.Println(pterm.NewRGB(15, 199, 209).Sprint(platform + " " + family + " " + version))
		return nil
	}

	osrelease, err := osRelease.Read()
	if err != nil {
		return err
	}
	pterm.DefaultCenter.Println(pterm.NewRGB(15, 199, 209).Sprint(osrelease["PRETTY_NAME"]))
	return nil
}

func ShowCPUInfo() error {
	info, err := cpuinfo.Info()
	if err != nil {
		return err
	}
	pterm.DefaultCenter.Println(pterm.NewRGB(15, 199, 209).Sprint(info[0].ModelName))
	return nil
}
