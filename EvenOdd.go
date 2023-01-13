package main

import "fmt"

func even(slice []int, ch chan []int) {
	var evenSlice []int
	for _, num := range slice {
		if num%2 == 0 {
			evenSlice = append(evenSlice, num)
		}
	}
	ch <- evenSlice
}

func odd(slice []int, ch chan []int) {
	var oddSlice []int
	for _, num := range slice {
		if num%2 != 0 {
			oddSlice = append(oddSlice, num)
		}
	}
	ch <- oddSlice
}

func main() {
	inputSlice := []int{2, 0, 44, 99, 32, 57, 87, 34, 29}
	ch := make(chan []int)

	go even(inputSlice, ch)
	go odd(inputSlice, ch)

	slice1 := <-ch
	slice2 := <-ch

	fmt.Println(slice1)
	fmt.Println(slice2)
}
