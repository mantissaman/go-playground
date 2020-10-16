package main

import (
	"fmt"
	"sort"
)

type aStructure struct {
	person string
	height int
	weight int
}

func main() {
	mySlice := make([]aStructure, 0)
	mySlice = append(mySlice, aStructure{"Neil", 108, 18})
	mySlice = append(mySlice, aStructure{"Atul", 166, 90})
	mySlice = append(mySlice, aStructure{"Diya", 120, 23})
	mySlice = append(mySlice, aStructure{"Priya", 158, 70})
	mySlice = append(mySlice, aStructure{"Merlin", 50, 5})

	fmt.Println("0:", mySlice)

	sort.Slice(mySlice, func(i, j int) bool {
		return mySlice[i].height < mySlice[j].height
	})
	fmt.Println("<:", mySlice)

	sort.Slice(mySlice, func(i, j int) bool {
		return mySlice[i].height > mySlice[j].height
	})
	fmt.Println(">:", mySlice)
}
