package main

import "fmt"

func main() {
	p := new(int)
	*p = 2
	v := 1
	p = &v
	fmt.Println(*p)
	fmt.Println(&p) // address of P
	fmt.Println(p)  // address of p is pointing
}
