package main

import (
	"fmt"
	"sort"
)

type Person struct {
	name    string
	surname string
}

type ByName []Person
type BySurname []Person

func (persons ByName) Less(i, j int) bool {
	return persons[i].name < persons[j].name
}

func (persons BySurname) Less(i, j int) bool {
	return persons[i].surname < persons[j].surname
}

func (persons ByName) Swap(i, j int) {
	persons[i], persons[j] = persons[j], persons[i]
}
func (persons BySurname) Swap(i, j int) {
	persons[i], persons[j] = persons[j], persons[i]
}

func (persons ByName) Len() int {
	return len(persons)
}
func (persons BySurname) Len() int {
	return len(persons)
}

func main() {
	x := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("4:6: %+v\n", x[4:6])

	p := []Person{
		{name: "Jan", surname: "Karelse"},
		{name: "Koos", surname: "Abrahams"},
		{name: "Adam", surname: "Jones"},
	}
	sort.Sort(ByName(p))
	fmt.Printf("byname: %+v\n", p)

	sort.Sort(BySurname(p))
	fmt.Printf("bysurname: %+v\n", p)

	//add entry
	p = append(p, Person{name: "Barend", surname: "Buys"})
	fmt.Printf("a1: %+v\n", p)

	sort.Sort(ByName(p))
	fmt.Printf("a2: %+v\n", p)

	//remove entry [1]
	p = append(p[0:1], p[2:]...)
	sort.Sort(ByName(p))
	fmt.Printf("r: %+v\n", p)
}
