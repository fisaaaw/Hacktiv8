package main

import (
	"fmt"
	"os"
	"strconv"
)

type friends struct {
	nama      string
	alamat    string
	pekerjaan string
	alasan    string
}

func main() {
	student := []friends{
		{
			nama:      "fadli",
			alamat:    "bekasi",
			pekerjaan: "developer",
			alasan:    "nambah ilmu",
		},
		{
			nama:      "fajar",
			alamat:    "jakarta",
			pekerjaan: "developer",
			alasan:    "nambah ilmu",
		},
		{
			nama:      "putra",
			alamat:    "bogor",
			pekerjaan: "pns",
			alasan:    "iseng",
		},
	}
	args := os.Args
	indexArgs := args[1]
	index, _ := strconv.Atoi(indexArgs)

	for i, teman := range student{
		if i == index - 1 {
			fmt.Println(teman)
		}
	}
}