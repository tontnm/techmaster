package main

import "fmt"

type Address struct {
	Location string
	City     string
	Country  string
}
type Person struct {
	Id       string
	FullName string
	Email    string
	Addr     []Address  //Một người có thế có nhiều địa chỉ
}

func (p Person) String() string {
	return fmt.Sprintf("{%s : %s : %s} lives at {%s, %s, %s}",
		p.Id, p.FullName, p.Email, p.Addr[0].Location, p.Addr[0].City, p.Addr[0].Country)

}

func main() {
	type personRequest struct {
		FullName string
		Email    string
		Location string
		City     string
		Country  string
	}

	pRequest := personRequest{
		FullName: "Pom",
		Email:    "abc@abc.vn",
		Location: "123",
		City:     "HCM",
		Country:  "Việt nam",
	}

	person := Person{
		Id:       "ox-13",
		FullName: pRequest.FullName,
		Email:    pRequest.Email,
		Addr: []Address{
			{
				Location: pRequest.Location,
				City:     pRequest.City,
				Country:  pRequest.Country,
			}},
	}
	fmt.Println(person)
}