package bst

type Compare interface {
	LessThan(v Compare) bool
}

type Node[ValueType Compare] struct {
	Parent *Node[ValueType]
	Value  ValueType
	Left   *Node[ValueType]
	Right  *Node[ValueType]
}

func (bst *Node[ValueType]) LeftHeight() int {
	if bst.Left != nil {
		return bst.Left.Height() + 1
	}
	return 0
}

func (bst *Node[ValueType]) RightHeight() int {
	if bst.Right != nil {
		return bst.Right.Height() + 1
	}
	return 0
}

func (bst *Node[ValueType]) Height() int {
	leftHeight := bst.LeftHeight()
	rightHeight := bst.RightHeight()
	if leftHeight > rightHeight {
		return leftHeight
	} else {
		return rightHeight
	}
}

func (bst *Node[ValueType]) Balance() int {
	return bst.LeftHeight() - bst.RightHeight()
}

func (bst *Node[ValueType]) Add(v ValueType) *Node[ValueType] {
	newNode := &Node[ValueType]{
		Parent: nil,
		Value:  v,
		Left:   nil,
		Right:  nil,
	}
	if bst == nil {
		return newNode
	}
	if v.LessThan(bst.Value) {
		if bst.Left == nil {
			bst.Left = newNode
		} else {
			bst.Left = bst.Left.Add(v)
		}
	} else {
		if bst.Right == nil {
			bst.Right = newNode
		} else {
			bst.Right = bst.Right.Add(v)
		}
	}

	//balance this node
	bal := bst.Balance()
	if bal < -1 {
		// X.bal() = (LH:1) - (RH:2) = -1 -> left rotate
		//      X(-1)
		//    /   \
		//  A(0)   Y(0)
		//       /   \
		//      B(0)  C(0)
		//
		//
		// Y = X.R
		// X.R = Y.L
		// Y.L = X
		Y := bst.Right
		bst.Right = Y.Left
		Y.Left = bst
		bst = Y
	}
	if bal > 1 {
		// X.bal() = (LH:2) - (RH:1) = +1 -> right rotate
		//         X(+1)
		//       /   \
		//     Y(0)   C(0)
		//    /   \
		//   B(0)  C(0)
		//
		// Y = X.L
		// X.L = Y.R
		// Y.R = X
		Y := bst.Left
		bst.Left = Y.Right
		Y.Right = bst
		bst = Y
	}
	return bst
}
