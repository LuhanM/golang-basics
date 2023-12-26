package main

import (
	"fmt"
	"sort"
)

func bubbleSort(n []int32) []int32 {
	// for i, _ := range n {
	// 	changes := 0
	// 	for j := 0; j < len(n)-i-1; j++ {
	// 		if n[j] > n[j+1] {
	// 			n[j], n[j+1] = n[j+1], n[j]
	// 			changes++
	// 		}
	// 	}
	// 	if changes <= 1 { //isso serve para reduzir o numero de iterações devido a ordenação da lista
	// 		break
	// 	}
	// }
	sort.Slice(n, func(i, j int) bool { return n[i] < n[j] })
	return n
}

func medianValue(n []int32) int32 {
	return int32(len(n) / 2)
}

func findMedian(arr []int32) int32 {
	// Write your code here
	sort.Slice(arr, func(f, s int) bool { return arr[f] < arr[s] })
	return arr[int32(len(arr)/2)]
}

func main() {
	unsorted := []int32{0, 1, 2, 4, 6, 5, 3}

	fmt.Println(unsorted)
	fmt.Println(bubbleSort(unsorted))
	fmt.Println(unsorted)
	fmt.Println(medianValue(unsorted))
}
