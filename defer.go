package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(name string) (lenght int, uppercase string) {
	defer fmt.Println("im done") //defer 함수가 끝나고 실행된다
	lenght = len(name)
	uppercase = strings.ToUpper(name)
	return
}
func main() {
	totalLenght, _ := lenAndUpper("Leedongju haha")
	fmt.Println(totalLenght)

}
