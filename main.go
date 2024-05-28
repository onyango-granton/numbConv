package main

import "fmt"

func main() {
	num := 11111110
	numString := intToString(num)

	indexArr := getIndexArr(numString)


	fmt.Println(indexArr)

	fmt.Println(powerOf(2,0))
}

func intToString(num int) string {
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

	return revRes
}

func getIndexArr(numString string) []int {
	indexPow := -1
	indexArr := []int{}

	for  range numString{
		indexPow++
		indexArr = append(indexArr, indexPow)
	}

	for range indexArr{
		for i := range indexArr{
			if i + 1 < len(indexArr) && indexArr[i] < indexArr[i+1]{
				indexArr[i], indexArr[i+1] = indexArr[i+1], indexArr[i]
			}
		}
	}

	return indexArr
}

func powerOf(num, pow int) int {
	if pow == 0{
		return 1
	}
	startNum := num
	for i:=1; i < pow;i++{
		startNum *= num
	}
	return startNum
}