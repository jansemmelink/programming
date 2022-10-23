package linkedlists_test

import (
	"testing"

	linkedlists "github.com/jansemmelink/programming/linked-lists"
)

func TestSingle(t *testing.T) {
	s := linkedlists.NewSingle()
	for _, i := range []int{1, 6, 5, 3, 9, 5, 6} {
		s.Add(i)
	}

	t.Logf("Values in single:")
	s.Traverse(func(v interface{}) error {
		t.Logf("  %d", v)
		return nil
	})

	t.Logf("Find >8:")
	{
		x := s.Search(8, func(key interface{}, v interface{}) bool {
			return v.(int) > key.(int)
		})
		t.Logf("found %v", x)
	}

	//add sorted incrementing order
	s = linkedlists.NewSingle()
	for _, i := range []int{1, 6, 5, 3, 9, 5, 6} {
		s.Insert(i, func(v1, v2 interface{}) bool {
			return v1.(int) > v2.(int) //sort ascending
		}) //insert before if func return true
	}
	t.Logf("Sorted list:")
	s.Traverse(func(v interface{}) error {
		t.Logf("  %d", v)
		return nil
	})

	//add sorted decrementing order
	s = linkedlists.NewSingle()
	for _, i := range []int{1, 6, 5, 3, 9, 5, 6} {
		s.Insert(i, func(v1, v2 interface{}) bool {
			return v1.(int) < v2.(int) //sort descending
		}) //insert before if func return true
	}
	t.Logf("Sorted list:")
	s.Traverse(func(v interface{}) error {
		t.Logf("  %d", v)
		return nil
	})

}
