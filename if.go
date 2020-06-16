package main

import "fmt"

func canIDrink(age int) bool {

	// if age < 18 {
	// 	return false
	// } else {
	// 	return true
	// }

	//this is ture syntax
	if koreanAge := age + 1; koreanAge < 18 {
		return false
	} else {
		return true
	}

	//this is false syntax
	// if koreanAge := age +2; koreanAge < 18{
	// 	return false
	// }
	// else {
	// 	return true
	// }
}

func main() {
	fmt.Println(canIDrink(16))
}
