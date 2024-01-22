package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	inside := 0

	var points = flag.Int("n", 0, "Number of points to generate")
	flag.Parse()

	// Time taken
	startTime := time.Now()
	defer func() {
		endTime := time.Now()
		elapsedTime := endTime.Sub(startTime)
		fmt.Printf("Time taken: %v\n", elapsedTime)
	}()

	for i := 0; i < *points; i++ {

		x := random.Float64()
		y := random.Float64()

		if x*x+y*y <= 1 {
			inside++
		}
	}
	fmt.Println(inside)
	fmt.Printf("Aprox pi = %f, time = ", 4.0*float64(inside)/float64(*points))
}
