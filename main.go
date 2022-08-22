package main

import "fmt"

func main() {
	fmt.Println("hello world")

	for i := 0; i <=10; i++ {
		if i % 2 == 1 {
			fmt.Println(i, "ganjil")
		} else if i % 2  == 0 {
			fmt.Println(i, "genap")
		}
	}
}