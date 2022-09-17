package main

func romanToInt(s string) int {
	preValue := getValue(string(s[0]))
	sum := 0
	for i := 1; i < len(s); i++ {
		nowValue := getValue(string(s[i]))
		if preValue < nowValue {
			sum -= preValue
		} else {
			sum += preValue
		}
		preValue = nowValue
	}
	sum += preValue
	return sum
}

func getValue(c string) int {
	switch c {
	case "I":
		return 1
	case "V":
		return 5
	case "X":
		return 10
	case "L":
		return 50
	case "C":
		return 100
	case "D":
		return 500
	case "M":
		return 1000
	default:
		return 0
	}
}
