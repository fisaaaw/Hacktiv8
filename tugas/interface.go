package main

import "fmt"

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

	resRegister := personSrvc.Register(&Person{name: "nama1"})
	fmt.Println(resRegister)

	resRegister = personSrvc.Register(&Person{name: "nama2"})
	fmt.Println(resRegister)

	resGetUser := personSrvc.GetUser()

	fmt.Println("-------Hasil get user-----")
	for _, v :=range resGetUser {
		fmt.Println(v.name)
	}
}