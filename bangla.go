package golib

import (
	"fmt"
)

func ENNumToBNNum[number int | float64](num number) string {
	bnNum := ""
	for _, v := range fmt.Sprint(num) {
		bnNum += getBnNum(string(v))
	}
	return bnNum
}

func getBnNum(char string) string {
	switch char {
	case "0":
		return "০"
	case "1":
		return "১"
	case "2":
		return "২"
	case "3":
		return "৩"
	case "4":
		return "৪"
	case "5":
		return "৫"
	case "6":
		return "৬"
	case "7":
		return "৭"
	case "8":
		return "৮"
	case "9":
		return "৯"
	case ".":
		return "."
	default:
		return ""
	}
}
