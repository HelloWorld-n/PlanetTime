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
		fmt.Println("\x1b[2J\x1b[H" + marsTime.Format("\033[2Jrot %R %NM %oD started %V vinquas %L layers %F fragments ago"))
		time.Sleep(time.Second / 16)
	}
}
