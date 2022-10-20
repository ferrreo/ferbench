package main

import (
	"ferbench/cpu"
	"ferbench/tui"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"runtime"
)

// env GOOS=windows GOARCH=amd64 go build
// env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w"
func main() {
	runLength := 30.0
	numThreads := runtime.NumCPU()
	mt := false
	st := false

	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"v"},
		Usage:   "Version number",
	}

	app := &cli.App{
		Name:                 "ferbench",
		Usage:                "A simple commandline benchmark, with no flags runs both a single thread and multi thread benchmark",
		Version:              "v0.11",
		EnableBashCompletion: true,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "singlethread",
				Aliases:     []string{"st"},
				Usage:       "Runs a single thread cpu benchmark only",
				Destination: &st,
			},
			&cli.BoolFlag{
				Name:        "multithread",
				Aliases:     []string{"mt"},
				Usage:       "Runs a multi thread cpu benchmark only",
				Destination: &mt,
			},
			&cli.Float64Flag{
				Name:        "length",
				Value:       30,
				Usage:       "Sets how long to run each benchmarks for",
				Destination: &runLength,
				Aliases:     []string{"l"},
			},
			&cli.IntFlag{
				Name:        "threads",
				Value:       runtime.NumCPU(),
				Usage:       "Sets the number of threads to use for multithread benchmark",
				Destination: &numThreads,
				Aliases:     []string{"t"},
			},
		},
		Action: func(ctx *cli.Context) error {
			all := !st && !mt
			if st || all {
				err := cpu.Bench(runLength, 1)
				if err != nil {
					return err
				}
			}
			if mt || all {
				err := cpu.Bench(runLength, numThreads)
				return err
			}
			return nil
		},
		Before: func(ctx *cli.Context) error {
			tui.ShowMainHeader()
			return nil
		},
	}
	app.Suggest = true

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
