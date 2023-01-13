package main

import (
	"fmt"
	"sort"
	"sync"
)

func main() {
	var size int
	fmt.Println("Enter the size of array:")
	fmt.Scan(&size)

	numberArray := make([]float64, size)
	fmt.Println("Enter the elements of the array:")
	for i := 0; i < size; i++ {
		fmt.Scan(&numberArray[i])
	}

	sort.Float64s(numberArray)

	part1 := numberArray[:len(numberArray)/4]
	part2 := numberArray[len(numberArray)/4 : 2*len(numberArray)/4]
	part3 := numberArray[2*len(numberArray)/4 : 3*len(numberArray)/4]
	part4 := numberArray[3*len(numberArray)/4:]

	var wg sync.WaitGroup
	wg.Add(4)

	go func() {
		sort.Float64s(part1)
		fmt.Println("p1: ", part1)
		wg.Done()
	}()
	go func() {
		sort.Float64s(part2)
		fmt.Println("p2: ", part2)
		wg.Done()
	}()
	go func() {
		sort.Float64s(part3)
		fmt.Println("p3: ", part3)
		wg.Done()
	}()
	go func() {
		sort.Float64s(part4)
		fmt.Println("p4: ", part4)
		wg.Done()
	}()

	wg.Wait()

	sortedArray := make([]float64, 0, len(numberArray))
	sortedArray = append(sortedArray, part1...)
	sortedArray = append(sortedArray, part2...)
	sortedArray = append(sortedArray, part3...)
	sortedArray = append(sortedArray, part4...)

	fmt.Println(sortedArray)
}
