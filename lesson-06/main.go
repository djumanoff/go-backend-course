package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	FirstName string `json:"firstName,omitempty"`
	LastName string `json:"lastName,omitempty"`
	Age int `json:"age,omitempty"`
	Size Size
	SizeUnit SizeUnit
	Married *bool `json:"married"`
}

const (
	jsonPerson = `{"FirstName":"Jaslan","LastName":"Jasik","Age":25,"Size":"100x90x79:m", "SizeUnit": "kg", "asdad": "adasd", "asdasd": "aaa"}`

	arrayJson = `[]`
)

var (
	married = true
	notMarried = false
)

func main() {
	person1 := &Person{"Igor", "Tin", 36, Size{90, 60, 60, CM}, CM, &married}
	person2 := &Person{"Egor", "Tin", 30, Size{100, 80, 80, M}, M, nil}

	data, err := json.Marshal(person1)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))

	data2, err := json.Marshal(person2)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data2))

	person3 := &Person{}
	err = json.Unmarshal([]byte(jsonPerson), person3)
	if err != nil {
		panic(err)
	}
	fmt.Println(person3)
}
