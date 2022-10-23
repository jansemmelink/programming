package avl

import "fmt"

type Compare interface {
	Compare(v Compare) int //<0 for less than, 0 for equal, >0 for bigger than
}

type Node[ValueType Compare] struct {
	Value ValueType
	Left  *Node[ValueType]
	Right *Node[ValueType]
}

func (avl *Node[ValueType]) String() string {
	var l ValueType
	if avl.Left != nil {
		l = avl.Left.Value
	}
	var r ValueType
	if avl.Right != nil {
		r = avl.Right.Value
	}
	return fmt.Sprintf("{v:%v,l:%v,r:%v,b:%d-%d=%d}", avl.Value, l, r, avl.LeftHeight(), avl.RightHeight(), avl.Balance())
}

func (avl *Node[ValueType]) LeftHeight() int {
	if avl == nil {
		return 0
	}
	if avl.Left != nil {
		return avl.Left.Height() + 1
	}
	return 0
}

func (avl *Node[ValueType]) RightHeight() int {
	if avl == nil {
		return 0
	}
	if avl.Right != nil {
		return avl.Right.Height() + 1
	}
	return 0
}

func (avl *Node[ValueType]) Height() int {
	leftHeight := avl.LeftHeight()
	rightHeight := avl.RightHeight()
	if leftHeight > rightHeight {
		return leftHeight
	} else {
		return rightHeight
	}
}

func (avl *Node[ValueType]) Balance() int {
	return avl.LeftHeight() - avl.RightHeight()
}

//return top node after add
func (avl *Node[ValueType]) Add(v ValueType) *Node[ValueType] {
	newNode := &Node[ValueType]{
		//Parent: nil,
		Value: v,
		Left:  nil,
		Right: nil,
	}

	if avl == nil {
		return newNode
	}
	return avl.add(newNode)
} //Node.Add()

//return top node after add
func (avl *Node[ValueType]) add(newNode *Node[ValueType]) *Node[ValueType] {
	if avl == nil {
		return newNode
	}
	if newNode.Value.Compare(avl.Value) < 0 {
		if avl.Left == nil {
			avl.Left = newNode
			fmt.Printf("avl add %s.Left->%s\n", avl, newNode)
		} else {
			avl.Left = avl.Left.add(newNode)
		}
	} else {
		if avl.Right == nil {
			avl.Right = newNode
			fmt.Printf("avl add %s.Right->%s\n", avl, newNode)
		} else {
			avl.Right = avl.Right.add(newNode)
		}
	}

	//balance this node
	bal := avl.Balance()
	if bal == -2 {
		if avl.Right.Balance() == -1 {
			//single left rotation
			fmt.Printf("avl rot single left %s %s\n", avl, avl.Right)
			a := avl
			b := a.Right
			a.Right = nil
			b.Left = b.Left.add(a)
			return b
		}
		if avl.Right.Balance() == 1 {
			//right left rotation
			fmt.Printf("avl rot right left %s %s %s\n", avl, avl.Right, avl.Right.Left)
			a := avl
			b := a.Right
			a.Right = nil
			c := b.Left
			b.Left = nil
			c.Left = c.Left.add(a)
			c.Right = c.Right.add(b)
			return c
		}
	}

	if bal == 2 {
		if avl.Left.Balance() == 1 {
			//single right rotation
			//e.g. avl rot single right {v:81,l:47,r:87,b:4-2=2} {v:47,l:25,r:59,b:3-2=1}
			fmt.Printf("avl rot single right %s %s\n", avl, avl.Left)
			a := avl    //81
			b := a.Left //47
			a.Left = nil
			b.Right = b.Right.add(a)
			return b
		}
		if avl.Right.Balance() == -1 {
			//left right rotation
			fmt.Printf("avl rot left right %s %s %s\n", avl, avl.Left, avl.Left.Right)
			a := avl
			b := a.Left
			a.Left = nil
			c := b.Right
			b.Right = nil
			c.Left = a
			c.Right = b
			return c
		}
	}
	fmt.Printf("avl --- %s\n", avl)
	return avl
}

func (avl *Node[ValueType]) Search(v ValueType) *Node[ValueType] {
	if avl == nil {
		return nil
	}
	diff := v.Compare(avl.Value)
	if diff == 0 {
		return avl
	} else if diff < 0 {
		return avl.Left.Search(v)
	} else {
		return avl.Right.Search(v)
	}
}
