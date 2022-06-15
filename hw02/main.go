package main

import (
	"fmt"

	"github.com/pmmson/GB.Go.2/hw02/numbers"
)

func main() {
	var n int
	fmt.Print("Введите число до которого будут выведены все простые числа: ")
	fmt.Scan(&n)
	fmt.Println(numbers.SimpleN(n))

}
