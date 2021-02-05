package main

import (
	"fmt"
)

func main() {
	nums := []int{16, 8, 42, 23, 15}
	max := nums[0]
	for _, num := range nums {
		if max < num {
			max = num
		}
	}
	fmt.Println(max)

}
