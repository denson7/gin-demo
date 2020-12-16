package utils

import "fmt"

func testSlice() {
	slice := make([]int, 5)
	//slice := make([]int, 3, 5)
	//slice := []string{"Red", "Blue", "Green", "Yellow", "Pink"}
	//intSlice:= []int{10, 20, 30}
	fmt.Println(len(slice)) // 打印结果 5
	fmt.Println(cap(slice)) // 打印结果 5
}
