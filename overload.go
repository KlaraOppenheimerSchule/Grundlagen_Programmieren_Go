package main

import (
	"fmt"
)

//test whether function overload is possible in go or not
func main() {

	fmt.Println(test(","))
	fmt.Println(test(2))

}

func test(comma string) {
	fmt.Println("Hello" + comma)
}

func test(number int) {
	fmt.Println("World", number)
}
