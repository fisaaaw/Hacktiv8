package main

import "fmt"

func print(names ...string) []map[string]string {
	var result []map[string]string

	for i, v := range names {
		key := fmt.Sprintf("student%d", i+1)
		temp := map[string]string{
			key: v,
		}
		result = append(result, temp)
	}
	return result
}

func summ(number ...int) (out map[string]int, kalimat string) {
	total := 0

	for _, v := range number {
		total += v
	}

	out = map[string]int{
		"total": total,
	}
	
	return
}

func sum(numbers ...int) (total int) {
	total = 0

	for _, v := range numbers {
		total += v
	}
	return 
}

func main() {

	numberList := sum(1, 2, 3, 4, 5, 6, 7, 8)

	//numberList := []int{1,2,3,4,5,6,7,8}
	//result, kalimat := summ(1, 2, 3, 4, 5, 6, 7, 8)

	fmt.Println("Result:", numberList)
}
