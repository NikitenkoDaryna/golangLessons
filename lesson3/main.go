package main

import (
	"fmt"
	"strconv"
)

type Address struct {
	City    string `json:"city"`
	Country string `json:"country"`
}
type Person struct {
	Name      string  `json:"name"`
	BirthDate string  `json:"birth_date"`
	Address   Address `json:"address"`
	Age       int     `json:"age"`
}
type User struct {
	Person           Person `json:"person"`
	UserName         string `json:"user_name"`
	Email            string `json:"email"`
	RegistrationDate string `json:"registration_date"`
	//call panic
	hh *User
	p  *Person
}

func (u *User) Hello() string {
	return "world"
}
func (p *Person) InfoPerson() string {
	str := p.Name + " " + p.BirthDate + "\n" + strconv.Itoa(p.Age) + " " + p.Address.City + " " + p.Address.Country
	return str
}
func (p *Person) New(name string, birthDate string, age int, city string, country string) {
	p.Name = name
	p.BirthDate = birthDate
	p.Age = age
	p.Address.City = city
	p.Address.Country = country

}

var p *Person
var a *Address
var u *User

func main() {
	fmt.Print("Lesson 3")

	person := Person{
		Name:      "Dylan",
		BirthDate: "4.08.1992",
		Address:   Address{Country: "Italy", City: "Toscana"},
		Age:       27,
	}
	user := User{
		Person:           person,
		UserName:         "SnowFlake",
		Email:            "Dylan29@yahoo.com",
		RegistrationDate: "13.09.2001",
	}
	var p2 Person
	p2.New("Lila", "15.07.2002", 18, "Paris", "France")





	fmt.Println(p) //return nil.If used *,the panic: runtime error: invalid memory address or nil pointer dereference
	//!panic: runtime error: invalid memory address or nil pointer dereference
	//fmt.Println(u.hh.Hello())
	//!panic
	//fmt.Println(p.InfoPerson())
	user.hh = &user
	user.p = &p2
	fmt.Println(*user.p)
	fmt.Println(person.InfoPerson())
	worker1 := worker{
		basicHuman: &human{"Kevin", "Macaulay", 27},
	}
	chef := chef{
		w:                     &worker1,
		bakingAbility:         "perfect",
		ordinaryCuisineSkills: "mastered",
		senseOfTaste:          "good",
		senseOfSmell:          "excellent",
	}
	fmt.Println(chef)
	fmt.Println(chef.w.basicHuman.HappyBirthday())
}
