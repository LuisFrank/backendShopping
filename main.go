package main

import (
	"fmt"
	"strconv"
	"time"
)

type Employee struct {
	id   int
	name string
}

func (e *Employee) setId(id int) {
	e.id = id
}

func (e *Employee) setName(name string) {
	e.name = name
}

func (e *Employee) getId() int {
	return e.id
}

func (e *Employee) getName() string {
	return e.name
}

func main() {
	var x int
	x = 8
	y := 7

	fmt.Println(x)
	fmt.Println(y)

	myValue, err := strconv.ParseInt("7", 0, 64)
	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		fmt.Println(myValue)
	}

	m := make(map[string]int)
	m["luis"] = 20
	fmt.Println(m["luis"])

	slice := []int{1, 2, 3}

	for index, value := range slice {
		fmt.Println(index)
		fmt.Println(value)
	}

	// c := make(chan int)
	// go doSomething(c)
	// <-c

	g := 25
	fmt.Println(g)
	h := &g
	fmt.Println(h)
	fmt.Println(*h)

	empleyee := Employee{}
	fmt.Printf("%v", empleyee)
	empleyee.id = 2
	empleyee.name = "nameeeeeluis"
	fmt.Printf("%v", empleyee)

	empleyee.setId(5)
	empleyee.setName("frank")
}

func doSomething(c chan int) {
	time.Sleep(3 * time.Second)
	fmt.Println("Done")
	c <- 1
}
