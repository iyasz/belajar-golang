package main

import "fmt"

func main() {

	var value1 = 100
	var value2 = 900

	var result = value1 == value2
	var result1 = value1 > value2
	var result2 = value1 < value2
	var result3 = value1 != value2

	fmt.Println(result)
	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(!result3)
}
