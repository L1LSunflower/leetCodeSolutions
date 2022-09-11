package main

//import "fmt"

//func main() {
//	num1, num2 := "11", "123"
//	fmt.Println(addStrings(num1, num2))
//}
//
//func addStrings(num1, num2 string) string {
//	if num1 == "0" && num2 == "0" {
//		return "0"
//	}
//	result := ""
//	tempRes := convertStrToInt(num1, 0) + convertStrToInt(num2, 0)
//	temp := convertIntToStr(tempRes, "")
//	for i := len(temp) - 1; i >= 0; i-- {
//		result += string(temp[i])
//	}
//	return result
//}

func convertIntToStr(number uint, result string) string {
	if number == 0 {
		return result
	}
	result += string(number%10 + '0')
	return convertIntToStr(number/10, result)
}

func convertStrToInt(number string, result uint) uint {
	if len(number) <= 0 {
		return result
	}
	tempVariable := 1
	for i := 0; i < len(number)-1; i++ {
		tempVariable *= 10
	}
	result += uint(number[0]-'0') * uint(tempVariable)
	return convertStrToInt(number[1:], result)
}
