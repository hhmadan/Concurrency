package main

import (
	"fmt"
	"sync"
)

var num int
var wg sync.WaitGroup

func func1() {
	var replica = num
	fmt.Println("This is func1 reading num ", replica)
	wg.Done()
}
func func2() {
	//time.Sleep(time.Second)
	num += 2
	fmt.Println("This is func2 writing to num ", num)
	wg.Done()
}
func main() {
	wg.Add(2)
	go func1()
	go func2()

	wg.Wait()
	fmt.Println("Exiting main function..!")
}
