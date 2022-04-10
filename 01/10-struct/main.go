package main

import "fmt"

type Person struct {
	Id       string
	FullName string
	Email    string
}


func (p Person) String() string {
	return fmt.Sprintf("%s : %s : %s", p.Id, p.FullName, p.Email)
}

func main() {
	type personRequest struct {
		FullName string
		Email    string
	}

	pRequest := personRequest{
		FullName: "Jackie Chan",
		Email:    "test@abc.com",
	}

	person := Person{
		Id:       "ox-13",
		FullName: pRequest.FullName,
		Email:    pRequest.Email,
	}
	fmt.Println(person)
	fmt.Println(person.String())
}
