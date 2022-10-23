package avl_test

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/jansemmelink/programming/trees/avl"
)

type Int int

func (i Int) Compare(j avl.Compare) int {
	return int(i) - int(j.(Int))
}

func TestBst(t *testing.T) {
	var avl *avl.Node[Int]
	//for _, i := range []Int{6, 3, 8, 1, 2, 6, 7, 9, 7, 3} {
	//for _, i := range []Int{6, 3, 8, 1, 6, 4, 9} {
	//for _, i := range []Int{95, 72, 110, 43, 87, 102, 120, 2, 50, 85, 108, 105} {
	//for _, v := range []Int{1, 2, 3} {
	// for _, v := range []Int{3, 2, 1} {
	// for _, v := range []Int{5, 2, 8, 6} {
	// 	avl = avl.Add(Int(v))
	// }

	values := []Int{}
	for k := 0; k < 20; k++ {
		v := Int(rand.Intn(100)) + 1
		values = append(values, v)
		avl = avl.Add(v)

		t.Logf("Added: %+v", values)
		traverse(t, fmt.Sprintf("after add %v", v), avl)
		for _, v := range values {
			n := avl.Search(v)
			if n == nil {
				t.Fatalf("NOT FOUND %v", v)
			}
			t.Logf("Found %v: %s", v, n)
		}
	}

}

func traverse(t *testing.T, title string, avl *avl.Node[Int]) {
	t.Logf("BST: %s", title)
	log(t, avl)
}

func log(t *testing.T, node *avl.Node[Int]) {
	if node != nil {
		t.Logf("%s", node.String())
		log(t, node.Left)
		log(t, node.Right)
	}
}
