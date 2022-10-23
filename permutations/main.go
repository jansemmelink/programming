package main

import "fmt"

func main() {
	s := "abc"
	Perm([]rune(s), func(v []rune) {
		fmt.Println(string(v))
	})
}

func Perm(a []rune, f func([]rune)) {
	perm(a, f, 0)
}

func perm(a []rune, f func([]rune), i int) {
	if i > len(a) {
		f(a) //this is a permutation - call user func
	} else {
		perm(a, f, i+1)
		for j := i + 1; j < len(a); j++ {
			a[i], a[j] = a[j], a[i]
			perm(a, f, i+1)
			a[i], a[j] = a[j], a[i]
		}
	}
}
