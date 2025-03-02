package main

import "fmt"

func main() {
	s := "kdjfvn"

	for i := range s {
		fmt.Println("i:", i)
	}
}
