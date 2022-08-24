package main

import (
	"fmt"
)

func main() {

	name1 := []string{"anton", "putra", "david", "rafi", "fadli", "putri"}

	var friends []*string

	for i := 0; i < len(name1); i++ {
		friends = append(friends, &name1[i])
	}

	print := func(friends []*string) {
		for _, fullNames := range friends {
			fmt.Println(*fullNames)
		}

	}

	print(friends)
}
