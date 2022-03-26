package Lesson_3

import (
	"fmt"
)

func main() {
	defer func() {
		if v := recover(); v != nil {
			fmt.Println("recovered after panic error :", v)
		}
	}()
	ErrorWithPanic()
}

func ErrorWithPanic() {
	arr := []int{1, 2}

	for i := 0; i <= 4; i++ {
		fmt.Printf("%d %d\n", i, arr[i])
	}
}
