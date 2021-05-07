package main

import (
	"fmt"
	"strconv"
)

func main() {
	number := 1345679
	str := strconv.Itoa(number)
	for i, v := range str {
		s := make([]rune, len(str)-i)
		for k, _ := range s {
			defaultValue := '0'
			if k == 0 {
				defaultValue = v
			}
			s[k] = defaultValue
		}
		fmt.Println(string(s))
	}
}
