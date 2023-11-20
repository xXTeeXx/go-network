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

	type person struct {
		name string
		age int
	}
	var p person
	p.name ="woramet"
	p.age=40
	fmt.Println(p)

	var x int =5 
	var xptr *int= &x
	fmt.Println(xptr)

	for i:=0; i < 5; i++ {
		fmt.Println(i)
	  }

	  var a = make(map[string]string)
	  var b map[string]string
	
	  fmt.Println(a == nil)
	  fmt.Println(b == nil)

	  month := 5

    switch month {
    case 1:
        fmt.Println("January")
    case 2:
        fmt.Println("February")
    case 3:
        fmt.Println("March")
    case 4:
        fmt.Println("April")
    case 5:
        fmt.Println("May")
    case 6:
        fmt.Println("June")
    case 7:
        fmt.Println("July")
    case 8:
        fmt.Println("August")
    case 9:
        fmt.Println("September")
    case 10:
        fmt.Println("October")
    case 11:
        fmt.Println("November")
    case 12:
        fmt.Println("December")
    default:
        fmt.Println("Invalid month")
    }

	q:= 20
  i:= 18
  if q > i {
    fmt.Println("x is greater than y")
  }

  time := 5
  if time < 10 {
    fmt.Println("Good morning.")
  } else if time < 20 {
    fmt.Println("Good day.")
  } else {
    fmt.Println("Good evening.")
  }
  
  var k = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
  l := map[string]int{"Oslo": 1, "Bergen": 2, "Trondheim": 3, "Stavanger": 4}

  fmt.Printf("k\t%v\n", k)
  fmt.Printf("l\t%v\n", l)
}