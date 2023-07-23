package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strings"
)

type Person struct {
	Name  string `json:"full_name"`
	Age   int    `json:"-"`
	Email string `json:"email,omitempty" random:"something"`

	// both Faulty and DupliateFaulty will be ignored during serialization
	Faulty          string `json:"omitempty"`
	DuplicateFaulty string `json:"omitempty"`
}

func main() {
	printMarshalExample()

	fmt.Println(getFieldTags(reflect.TypeOf(Person{})))
}

func printMarshalExample() {
	// marshal
	pp := []Person{
		{
			Name:   "first person",
			Age:    20,
			Email:  "first@person.com",
			Faulty: "faulty",
		},
		{
			Name:  "first person",
			Age:   20,
			Email: "",
		},
	}

	b, _ := json.MarshalIndent(pp, "", "\t")
	fmt.Println(string(b))

	// unmarshal
	f, _ := os.ReadFile("person.json")
	pp2 := []Person{}
	_ = json.Unmarshal(f, &pp2)
	fmt.Printf("%+v\n", pp2) // both "Faulty" and "DuplicateFaulty" are empty
}

func getFieldTags(t reflect.Type) map[string][]string {
	if t.Kind() != reflect.Struct {
		panic(fmt.Sprintf("t should be struct but is %s", t))
	}

	tags := make(map[string][]string)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		jTags := strings.SplitN(f.Tag.Get("json"), ",", -1)
		tags[f.Name] = jTags
	}

	return tags
}
