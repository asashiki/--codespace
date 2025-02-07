package main

import (
	"fmt"
)

type User struct {	
	Name string
	Age int
}

func main() {
	lst := []User{
		{"Alice", 25},
		{"Bob", 30},
		{"Charlie", 35},
		{"David", 40},
		{"Eve", 45},
}

var old []User
var young []User

for _, u := range lst {
	if u.Age > 40 {
		old = append(old, u)
	}
}

for _, u := range lst {
	if u.Age <= 35 {
		young = append(young, u)
	}
}

fmt.Println("Old:", old)
fmt.Println("Young:", young)

}