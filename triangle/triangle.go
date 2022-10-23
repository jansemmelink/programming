package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	n := 0
	if i64, err := strconv.ParseInt(os.Args[1], 10, 64); err != nil {
		panic(err)
	} else {
		n = int(i64)
	}
	if n < 1 || n > 100 {
		panic(fmt.Sprintf("%d must be 1..100", n))
	}

	space := "                                                                                                              "
	for i := 1; i <= n; i++ {
		fmt.Printf("%.*s", n-i, space)
		for j := 0; j < i; j++ {
			fmt.Printf("* ")
		}
		fmt.Printf("\n")
	}

}
