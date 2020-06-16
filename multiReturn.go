package main


import (
	"fmt"
	"strings"
)

func lenAndUpper(name string) (lenght int, uppercase string) {
	lenght = len(name)
	uppercase = strings.ToUpper(name)
	return
}
func main() {
	totalLenght, _ := lenAndUpper("Leedongju haha")
	fmt.Println(totalLenght)

}
