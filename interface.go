package main

import (
	"fmt"
	"strconv"
)

// Converter interface to convert int to string
type Converter interface {
	ConvertToString(number int) string
}

// SimpleConverter simple direct converter
type SimpleConverter struct {
}

// ConvertToString implement Converter interface method ConvertToString
func (SimpleConverter) ConvertToString(number int) string {
	return strconv.Itoa(number)
}

func main() {
	fmt.Println("will convert some ints to string")
	table := []int{0, 99, 1234, 4321}

	converter := SimpleConverter{}

	for i, v := range table {
		result := GetString(v, converter)
		fmt.Println(i, ": convert int", v, "into string:", result)
	}
}

// GetString convert a int into string using a converter interface
func GetString(number int, converter Converter) string {
	return converter.ConvertToString(number)
}
