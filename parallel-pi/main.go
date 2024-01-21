package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func worker(inside_ptr *int, times int, routines int, wg_ptr *sync.WaitGroup) {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	fmt.Print("hola")

	for i := 0; i < (times / routines); i++ {
		wg_ptr.Add(1)

		x := random.Float64()
		y := random.Float64()

		if x*x+y*y <= 1 {
			*inside_ptr = *inside_ptr + 1
		}
	}
	wg_ptr.Done()
}

func main() {
	inside := 0
	inside_ptr := &inside

	var wg sync.WaitGroup
	wg_ptr := &wg

	var points = flag.Int("n", 0, "Number of points to generate")
	var routines = flag.Int("r", 0, "Number of routines to use")
	flag.Parse()

	// Time taken
	startTime := time.Now()
	defer func() {
		endTime := time.Now()
		elapsedTime := endTime.Sub(startTime)
		fmt.Printf("Time taken: %v\n", elapsedTime)
	}()

	for i := 0; i < *routines; i++ {
		go worker(inside_ptr, *points, *routines, wg_ptr)
	}
	wg.Wait()
	fmt.Printf("Aprox pi = %f, time = ", 4.0*float64(inside)/float64(*points))
}
