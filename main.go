package main

import "fmt"

func main() {
	num := 11111110
	res := ""
	revRes := ""

	for num != 0{
		startRune := '0'

		modulus := num % 10

		for range modulus{
			startRune++
		}

		res += string(startRune)

		num = num / 10
	}

	for i:= len(res)-1; i >= 0; i--{
		revRes += string(res[i])
	}

	fmt.Println(revRes)
}


