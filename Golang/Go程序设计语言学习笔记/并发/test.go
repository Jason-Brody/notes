package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)

	for {
		go fmt.Print(0)
		fmt.Print(1)

	}
	// ca, cb := make(chan string), make(chan string)

	// count := 0
	// ga := func() {
	// 	for m := range ca {

	// 		count++
	// 		cb <- m
	// 	}
	// }
	// gb := func() {
	// 	for m := range cb {

	// 		ca <- m
	// 	}
	// }
	// go ga()
	// go gb()
	// ca <- "1"
	// lastCount := 0
	// for range time.Tick(1 * time.Second) {
	// 	fmt.Println(count - lastCount)
	// 	lastCount = count
	// }

}

func process(i int) {

	du, _ := time.ParseDuration(fmt.Sprintf("%ds", i))
	time.Sleep(du)
	fmt.Println(i)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
