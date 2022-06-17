package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	ID      string `json:"id"`
	Name    string
	Gender  string
	Age     int
	Married bool
	test    string
}

func main() {
	originalUser := user{
		ID:      "user1",
		Name:    "kenny",
		Gender:  "male",
		Age:     23,
		Married: false,
		test:    "test",
	}
	b, err := json.Marshal(originalUser)
	if err != nil {
		panic(err)
	}
	fmt.Println(b)
	fmt.Println(string(b))

	newUser := user{}
	err = json.Unmarshal(b, &newUser)
	if err != nil {
		panic(err)
	}

	fmt.Println(newUser)

}
