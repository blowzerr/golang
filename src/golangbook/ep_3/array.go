package main

import "fmt"

func main() {
	array := [5]float64 { 21, 23, 54, 63, 10 }
	for _, value := range array {
		fmt.Println(value)
		fmt.Println(array)	
	}
	slice := make([]float64, 6, 10)
	fmt.Println(slice)

	slice1 := array[:2]
	fmt.Println(slice1)
	slice2 := append(slice1, 2, 0)
	fmt.Println(slice2)
}