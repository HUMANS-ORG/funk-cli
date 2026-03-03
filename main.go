package main

import (
	"errors"
	"fmt"
)

type firststruct struct {
	name string
	age  int
}



func usestruct(e firststruct) string {
	return e.name
}

func main() {
	// a1 :=[...]int{3,4,3}
	var a1 []int = []int{4, 5, 3}
	a1 = append(a1, 4)
	var one_string string = "helloworld"

	var indexing = one_string[0]

	fmt.Printf("%v ,%T ", indexing, indexing)
	var a int = 20
	b := 0
	var main_result, second_result, err = hello(a, b)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(main_result, second_result)
	fmt.Println(a1)

	// use the struck funcation

	demo := firststruct{
		name: "gokulraj",
		age:  21,
	}

	fmt.Println("name:", demo.name, "age:", demo.age)
}
func hello(a int, b int) (int, int, error) {
	var err error
	if b == 0 {
		err = errors.New("0 can be divided")
		return 0, 0, err
	}
	var result int = a + b
	var results int = a - b
	return result, results, err
}
