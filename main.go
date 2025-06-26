package main

import (
	"fmt"
	"planetTime/planets"
	"time"
)

func main() {
	for true {
		earthTime := time.Now().UTC()
		marsTime := planets.NewMarsTime(&earthTime)
		fmt.Println("\x1b[2J\x1b[H\033[2J" + marsTime.FormatExample("rot 203 Makara 5th started 0 vinquas 1 layers 2 fragments ago"))
		time.Sleep(time.Second / 256)
	}
}
