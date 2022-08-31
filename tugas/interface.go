package main

import (
	"fmt"
	"sync"
)

type Person struct {
	name string
}

type personService struct {
	db []*Person
}

type UserIface interface {
	Register(u *Person) string
	GetUser() []*Person
}

func NewPersonService(db []*Person) UserIface {
	return &personService{
		db: db,
	}
}

func (u *personService) Register(in *Person) string {
	u.db = append(u.db, in)
	return in.name + "berhasil ditambahkan"
}

func (u *personService) GetUser() []*Person {
	return u.db
}

func main() {
	var db []*Person

	personSrvc := NewPersonService(db)

	names := []string{"putra", "aji", "fika", "rizky"}
	var wg sync.WaitGroup
	wg.Add(len(names))
	for _, v := range names {
		go func (name string)  {
				res := personSrvc.Register(&Person{name: name})
				fmt.Println(res)
				wg.Done()
		}(v)
		
	}
	wg.Wait()

	resGetUser := personSrvc.GetUser()

	fmt.Println("-----Hasil Get User-----")
	for _, n := range resGetUser {
		fmt.Println(n.name)
	}
}