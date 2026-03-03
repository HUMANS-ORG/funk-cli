package main

import (
	"errors"
	"fmt"

)

func main() {
	// a1 :=[...]int{3,4,3}
	var a1[]int = []int{4,5,3}
	a1 = append(a1,4)
	var a int = 20
	b := 0
	var main_result, second_result , err = hello(a, b)
	if err!=nil{
		fmt.Println(err.Error())
	}
	fmt.Println(main_result, second_result)
	fmt.Println(a1)

}
func hello(a int, b int) (int, int,error){
	var err error;
	if b==0{
		err = errors.New("0 can be divided")
		return 0,0,err
	}
	var result int = a + b     
	var results int = a - b 
	return result, results ,err
}
