package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	
)

var db []*Person
var PORT = ":8080"
var person Person
var personSrvc = NewPersonService(db)


type Person struct {
	name string `json:"name"`
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

func main() {
	http.HandleFunc("/test", greet)
	http.HandleFunc("/register", RegisterWeb)
	
	fmt.Println("server is listen in port", PORT)
	http.ListenAndServe(PORT, nil)
}

func greet(w http.ResponseWriter, r *http.Request) {
	msg := "Hello World"
	fmt.Fprintln(w, msg)
}

func (u *personService) Register(in *Person) string {
	u.db = append(u.db, in)
	return in.name + "berhasil ditambahkan"
}

func (u *personService) GetUser() []*Person {
	return u.db
}

func RegisterWeb(w http.ResponseWriter, r *http.Request) {
	decode := json.NewDecoder(r.Body)
	decode.Decode(&person)
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &person)
	fmt.Println(person.name)
	personSrvc.Register(&person)
	fmt.Fprintln(w, "berhasil")
}

func GetWeb(w http.ResponseWriter, r *http.Request) {
	allUser := personSrvc.GetUser()
	data, err := json.Marshal(allUser)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(data)
	}
	json.NewEncoder(w).Encode(allUser)
}