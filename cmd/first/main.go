package main

import (
	"fmt"
)

func main() {
	slice := make([]int, 0, 10)
	for i := 0; i < 10; i++ {
		slice = append(slice, i*2)
	}
	for _, item := range slice {
		fmt.Println(item)
	}
	fmt.Printf("%v => %d => %d\n", slice, len(slice), cap(slice))

	mp := map[int]string{
		0: "zero",
		1: "one",
		2: "two",
	}
	mp[0] = "three"
	for key, value := range mp {
		fmt.Printf("%d => %s\n", key, value)
	}
	strConverter := func(number int) string {
		return fmt.Sprintf("%d - %d", number, len(mp))
	}
	for _, item := range slice {
		fmt.Println(strConverter(item))
	}
}
