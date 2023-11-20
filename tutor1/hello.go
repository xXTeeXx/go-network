package main

import (
	"fmt"
)


func main(){
	var t bool = true
	var f bool = false

	var age int =40
	var favNum float64 = 1.6180339
	var complex128 complex128 = 5+5i
	var r rune = 10

	fmt.Println("t is", t)
	fmt.Println("f is", f)

	fmt.Println("age is",age, "favNum is",favNum)
	fmt.Println("complex128",complex128)
	fmt.Println("r",r)
	var myName string ="ไทย"
	fmt.Println(myName)
	fmt.Println(myName[3])
	fmt.Println(len(myName))

	var arry5[5] float64
	arry5[0] = 18.7
	arry5[1] = 88.7
	arry5[2] = 98.7
	arry5[3] = 68.7
	arry5[4] = 78.7
	fmt.Println(arry5)
	fmt.Println(len(arry5))
	fmt.Println(arry5[3])
	

	arry3 := [3]float64 {98,78,66}
	fmt.Println(arry3)

	var s []float64 = arry5[0:2]
	fmt.Println(s)
}