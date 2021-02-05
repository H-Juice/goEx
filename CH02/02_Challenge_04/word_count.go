package main

import (
	"fmt"
	"strings"
)

func main() {
	text := `Needles and pins
	Needles and pins
	Sew me a sail
	To Catch me the wind
	`
	list := []string{}
	list = strings.Fields(text)
	fmt.Println(list)
	maps := map[string]int{}

	for _, key := range list {
		maps[strings.ToLower(key)]++
	}

	fmt.Println(maps)
}
