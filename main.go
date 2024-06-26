package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	// hello,_ := fmt.Scan()
	// fmt.Println(hello)
	fmt.Print("Enter byte:	")
	input := ""
	fmt.Scanln(&input)
	decimal, err := byteToDecimal(input)

	if err != nil{
		fmt.Println(err.Error())
		os.Exit(1)
	} else {
		fmt.Printf("Decimal Value:	%v\n",decimal)
	}
}

func intToString(num int) string {
	res := ""
	revRes := ""
	for num != 0 {
		startRune := '0'

		modulus := num % 10

		for range modulus {
			startRune++
		}

		res += string(startRune)

		num = num / 10
	}

	for i := len(res) - 1; i >= 0; i-- {
		revRes += string(res[i])
	}

	return revRes
}

func getIndexArr(numString string) []int {
	indexPow := -1
	indexArr := []int{}

	for range numString {
		indexPow++
		indexArr = append(indexArr, indexPow)
	}

	for range indexArr {
		for i := range indexArr {
			if i+1 < len(indexArr) && indexArr[i] < indexArr[i+1] {
				indexArr[i], indexArr[i+1] = indexArr[i+1], indexArr[i]
			}
		}
	}

	return indexArr
}

func powerOf(num, pow int) int {
	if pow == 0 {
		return 1
	}
	startNum := num
	for i := 1; i < pow; i++ {
		startNum *= num
	}
	return startNum
}

func byteToDecimal(numString string) (int, error) {
	indexArr := getIndexArr(numString)
	res := 0

	for i, ch := range numString {
		if ch == '1' {
			res += powerOf(2, indexArr[i]) * 1
		} else if ch == '0' {
			res += powerOf(2, indexArr[i]) * 0
		} else {
			errorMsg := "invalid bit numeral " + string(ch)
			return 0, errors.New(errorMsg)
		}
	}

	return res, nil
}
