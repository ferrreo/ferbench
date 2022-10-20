package cpu

import (
	"ferbench/tui"
	"math"
	"strconv"
	"time"
)

const (
	scoreMultiplier     = 57000.0
	runLengthMultiplier = 1000000000
)

func Bench(runLength float64, numThreads int) error {

	multiScore := 0.0
	count := 0.0
	title, barTitle, scoreTitle := setupText(numThreads)

	tui.ShowTitle(title)

	var channels []chan float64
	for i := 0; i < numThreads; i++ {
		channel := make(chan float64)
		channels = append(channels, channel)
		go work(channel)
	}

	ticker := time.NewTicker(500 * time.Millisecond)
	bar, _ := tui.ShowBar(barTitle, runLength)

	for multiScore == 0.0 {
		<-ticker.C
		bar.Increment()
		count += 0.5
		if count >= runLength {
			ticker.Stop()
			bar.Stop()
			for j := range channels {
				channels[j] <- -1
				multiScore += <-channels[j]
			}
		}
	}

	score := (multiScore / (runLengthMultiplier * runLength)) * scoreMultiplier
	tui.ShowScore(scoreTitle, score)

	return nil
}

func work(channel chan float64) {
	k := 0.0
	f := 3.0

	for {
		select {
		case x := <-channel:
			if x == -1 {
				channel <- k
				break
			}
		default:
			f += 4 * math.Pow(-1, k) / ((2*k + 2) * (2*k + 3) * (2*k + 4))
			k += 1
		}
	}
}

func setupText(numThreads int) (string, string, string) {
	if numThreads == 1 {
		return "Single thread benchmark", "Single thread test progress", "Single thread score: "
	} else {
		return "Multi thread benchmark - Threads: " + strconv.Itoa(numThreads), "Multi thread test progress", "Multi thread score: "
	}
}
