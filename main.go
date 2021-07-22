package main

import (
	"fmt"
	mysort "gitee.com/RocsSun/algorithms/mySort"
)

func main() {
	values := []int{4, 3, 3, 1, 0}
	mysort.BubbleSort(values)
	fmt.Println(values)
}
