package main

import (
	"fmt"
)

type myStructure struct {
	Name    string
	Surname string
	Height  int32
}

func createStruct(n, s string, h int32) *myStructure {
	if h > 300 {
		h = 0
	}
	return &myStructure{n, s, h}
}

func retStructure(n, s string, h int32) myStructure {
	if h > 300 {
		h = 0
	}
	return myStructure{n, s, h}
}

func main() {
	s1 := createStruct("Neil", "Sharma", 118)
	s2 := retStructure("Neil", "Sharma", 118)
	s3 := new(myStructure)
	(*s3).Name = "Diya"

	fmt.Println((*s1).Name)
	fmt.Println(s2.Name)
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(*s3)
}
