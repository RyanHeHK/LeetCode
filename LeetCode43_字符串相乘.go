package main

import (
	"strconv"
)

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	result := "0"
	add := 0
	count := 0
	m, n := len(num1), len(num2)
	for i := m - 1; i >= 0; i-- {
		tmp := ""
		for k := 0; k < count; k++ {
			tmp += "0"
		}
		for j := n - 1; j >= 0; j-- {
			mul := int(num1[i]-'0')*int(num2[j]-'0') + add
			if mul > 9 {
				add = mul / 10
			} else {
				add = 0
			}
			tmp += strconv.Itoa(mul % 10)
		}
		if add > 0 {
			tmp += strconv.Itoa(add)
			add = 0
		}
		result = twoNumAdd(result, tmp)
		count++
	}
	finalRes := ""
	for i := len(result) - 1; i >= 0; i-- {
		finalRes += string(result[i])
	}
	return finalRes
}

func twoNumAdd(num1, num2 string) string {
	res := ""
	add := 0
	for i := 0; i < len(num1); i++ {
		sum := int(num1[i]-'0') + int(num2[i]-'0') + add
		if sum > 9 {
			add = sum / 10
		} else {
			add = 0
		}
		res += strconv.Itoa(sum % 10)
	}
	for j := len(num1); j < len(num2); j++ {
		sum := int(num2[j]-'0') + add
		if sum > 9 {
			add = sum / 10
		} else {
			add = 0
		}
		res += strconv.Itoa(sum % 10)
	}
	if add > 0 {
		res += strconv.Itoa(add)
	}
	return res
}
