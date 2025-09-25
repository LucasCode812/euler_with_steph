package main

import "fmt"

func main() {
	first := 1
	second := 2
	evenSum := 0
	evenSum += second

	for {
		next := first + second

		if next > 4000000 {
			break
		}

		if next%2 == 0 {
			evenSum += next
		}

		first = second
		second = next
	}

	fmt.Println("Answer:", evenSum)
}
