package main

import "fmt"

//Задача 1
func srezsum(slice []int) {
   var sum int  
    for _, i := range slice {
       sum += i
    }
    fmt.Println(sum)
} 

//Задача 2

func even_or_odd(number int) (string) {
    i := number % 2
    if (i == 0){
        return "even"
    } else {
        return "odd"}    
}

func swap (xPtr *int, yPtr *int){
    x := *xPtr
    y := *yPtr
    *xPtr = y
    *yPtr = x
}

func main() {
    //Задача 1 
    //asd := []int{2,5,4,5,4,2,2}
    //srz := asd[:3]
    //srezsum(srz)

    //Задача 2
    //fmt.Println("Input a number:")
    //var input int
    //fmt.Scanf("%d",&input)
    //fmt.Println("Your number is",even_or_odd(input))

    x := 213
    y := 21
    swap(&x, &y)
    fmt.Println(x,y)
    
}