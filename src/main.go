package main

import (
	"fmt"
	"planetTime/planets"
	"time"
)

func main() {
	earthTime := time.Now().UTC()
	marsTime := planets.NewMarsTime(&earthTime)
	fmt.Println(marsTime.Format("rot %R %NM %Dth started %V vinquas %L layers %F fragments ago"))
}
