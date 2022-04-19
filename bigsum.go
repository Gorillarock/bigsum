package main

import (
	"fmt"
)

func main() {
	var add1 int = 0
	var sum int = 0
	arr1 := "4002009001892"
	arr2 := "4002009001892"
	arr1len := len(arr1)
	arr2len := len(arr2)
	arr1runes := make([]rune, arr1len)
	arr2runes := make([]rune, arr2len)

	for i := 0; i < arr1len; i++ {
		arr1runes[i] = rune(arr1[i] - '0')
	}
	for i := 0; i < arr2len; i++ {
		arr2runes[i] = rune(arr2[i] - '0')
	}

	var diff int = 0

	if arr1len > arr2len {
		diff = arr1len - arr2len
		arr2runes = make([]rune, arr1len)

		for i, j := diff, 0; i < arr1len; i, j = i+1, j+1 {
			arr2runes[i] = rune(arr2[j] - '0')
		}

		arr2len = len(arr2runes)

	} else if arr2len > arr1len {
		diff = arr2len - arr1len
		arr1runes = make([]rune, arr2len)

		for i, j := diff, 0; i < arr2len; i, j = i+1, j+1 {
			arr1runes[i] = rune(arr1[j] - '0')
		}

		arr1len = len(arr1runes)
	}

	result := make([]rune, arr1len+1)

	for i := (arr1len - 1); i >= 0; i-- {
		sum = int(arr1runes[i]) + int(arr2runes[i]) + add1
		if sum > 9 {
			add1 = 1
		} else {
			add1 = 0
		}
		result[i+1] = rune(sum % 10)
		sum = 0
	}

	if add1 > 0 {
		result[0] = rune(1)
	} else {
		tmp := make([]rune, len(result)-1)
		copy(tmp, result[1:])
		result = tmp
	}
	for _, j := range result {
		fmt.Printf("%d", j)
	}
	fmt.Printf("\n")
}
