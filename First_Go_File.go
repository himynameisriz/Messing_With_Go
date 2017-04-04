package main

import (
	"fmt"
)


func main() {
	// hello world
	hw := "hello world";
	name := "my name is riz"
	fmt.Println(hw, name)

	// concurrency
	fmt.Println("We will now test concurrency")
	defer fmt.Println("we have tested concurrency")
	c := make(chan int)
	go inc(10, c)
	go inc(800, c)
	go inc(60, c)
	fmt.Println(<-c, <-c, <-c)
}

// from learnxinyminutes
// c is a channel, a concurrency-safe communication object.
func inc(i int, c chan int) {
	c <- i + 1 // <- is the "send" operator when a channel appears on the left.
}