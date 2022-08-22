package main

import "fmt"

func main() {
	fmt.Println("hello world")

	for i := 0; i <=10; i++ {
		if i % 2 == 1 {
			fmt.Println(i, "ganjil")
		} else {
			fmt.Println(i, "genap")
		}
	}

	name := []string{"teguh", "fadli", "rafi", "david", "fajar", "san", "kadek", "stevanus", "iqbal", "kevin"}
	for _, a := range name {
	fmt.Println(a)
	}
}