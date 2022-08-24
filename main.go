package main

import (
	"fmt"
)

type namaTemen struct {
	name string
}

func main() {
	var friends []*namaTemen
	name1 := []string{"anton", "putra", "david", "rafi", "fadli", "putri"}

	for _, a := range name1 {
			v := namaTemen{
			name: a,
			}
			friends = append(friends, &v)

	}

	print := func(friends []*namaTemen) {
		for _, fullNames := range friends {
			fmt.Println(*&fullNames.name)
		}

	}

	print(friends)
}
