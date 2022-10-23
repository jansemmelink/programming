package linkedlists

func NewSingle() *Single {
	return &Single{
		Head: nil,
		Tail: nil,
	}
}

func (s *Single) Add(v interface{}) {
	n := &Node{Next: nil, Value: v}
	if s.Head == nil {
		s.Head = n
	} else {
		s.Tail.Next = n
	}
	s.Tail = n
} //Single.Add()

func (s *Single) Insert(v interface{}, f SearchFunc) {
	var prev *Node
	n := s.Head
	for n != nil && f(v, n.Value) {
		prev = n
		n = n.Next
	}
	newNode := &Node{
		Value: v,
		Next:  nil,
	}
	if n == nil {
		//insert at tail
		if s.Head == nil {
			s.Head = newNode //also head of list that was empty
		} else {
			s.Tail.Next = newNode
		}
		s.Tail = newNode
	} else {
		//insert before existing node n
		newNode.Next = n
		if prev == nil {
			s.Head = newNode //new head
		} else {
			prev.Next = newNode //new in middle
		}
	}
} //single.Insert()

//Trafers until func returns error, then return that node and the error
//or traverse all and return nil,nil
func (s Single) Traverse(f func(interface{}) error) (*Node, error) {
	n := s.Head
	for n != nil {
		if err := f(n.Value); err != nil {
			return n, err
		}
		n = n.Next
	}
	return nil, nil
} //Single.Travers()

type SearchFunc func(key interface{}, value interface{}) bool

//return node if search function returns true
func (s Single) Search(key interface{}, fnc SearchFunc) *Node {
	n := s.Head
	for n != nil {
		if fnc(key, n.Value) {
			return n
		}
		n = n.Next
	}
	return nil //not found
}

type Single struct {
	Head *Node
	Tail *Node
}

type Node struct {
	Next  *Node
	Value interface{}
}
