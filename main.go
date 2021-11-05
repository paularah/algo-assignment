package main

import (
	"fmt"
)

func main() {
	fmt.Println(countSockPair([]int{1, 2, 1, 3, 4, 2, 5, 4, 1, 3})) // result -> 4
	fmt.Println(countSockPair([]int{1, 2, 1, 2, 1, 3, 2}))          // result -> 2

}

func countSockPair(sockArr []int) int {
	result := 0
	countMap := freqCount(sockArr)
	for _, sock := range countMap {
		if sock >= 2 {
			result = result + sock/2
		}
	}
	return result
}

func freqCount(arr []int) map[int]int {
	freqMap := make(map[int]int)
	for _, sock := range arr {
		freqMap[sock]++
	}
	return freqMap
}
