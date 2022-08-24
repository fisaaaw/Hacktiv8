package main

import (
	"fmt"
	
)

func main() {

	name1 := "fadli"
	name2 := "anton"

	friends :=[]*string{&name1, &name2}

	print := func(friends []*string){
		for _, fullNames := range friends {
			fmt.Println(*fullNames)
		}
		
	}
	
	print(friends)
}