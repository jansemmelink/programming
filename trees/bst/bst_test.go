package bst_test

import (
	"fmt"
	"testing"

	"github.com/jansemmelink/programming/trees/bst"
)

type Int int

func (i Int) LessThan(j bst.Compare) bool {
	if i < j.(Int) {
		return true
	} else {
		return false
	}
}

func TestBst(t *testing.T) {
	var bst *bst.Node[Int]
	// for _, i := range []Int{6, 3, 8, 1, 2, 6, 7, 9, 7, 3} {
	for _, i := range []Int{6, 3, 8, 1, 6, 4, 9} {
		bst = bst.Add(i)
	}
	traverse(t, "new", bst)
}

func traverse(t *testing.T, title string, bst *bst.Node[Int]) {
	h := bst.Height()  //e.g. 3 (=top + 3 levels)
	lineTot := h*2 + 1 //e.g. 7 (lineNr 0..6)
	t.Logf("BST: %s (height=%d)", title, h)
	for lineNr := 0; lineNr < lineTot; lineNr++ {
		fmt.Printf("%2d/%2d: ", lineNr, lineTot)
		log(t, bst, lineNr, lineNr, lineTot)
		fmt.Printf("\n")
	}
}

//height=4:
//lines:
//0:                   +-----------------000(000)----------------+
//1:                   |                                         |
//2:         +------000(000)------+                    +------000(000)------+
//3:         |                    |                    |                    |
//4:    +-000(000)-+         +-000(000)-+         +-000(000)-+         +-000(000)-+
//5:    |          |         |          |         |          |         |          |
//6: 000(000)   000(000)  000(000)   000(000)  000(000)   000(000)  000(000)   000(000)

//0:                     +---------------------000(000)-------------------+
//   012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789
//1:                     |                                                |
//   012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789
//2:          +-------000(000)------+                       +-------000(000)--------+
//   012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789
//3:          |                     |                       |                        |
//   012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789
//4:     +-000(000)-+          +-000(000)-+           +-000(000)--+           +-000(000)--+
//   012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789
//5:     |          |          |          |           |           |           |           |
//   012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789
//6: -000(000)- -000(000)- -000(000)- -000(000)-  -000(000)- -000(000)- -000(000)- -000(000)-

//last line items are " 000(000) " = 10 wide with no down lines

const space = "                                                   "
const line = "---------------------------------------------------"

func log(t *testing.T, node *bst.Node[Int], relLineNr int, lineNr int, lineTot int) {
	//own height:
	//h := node.Height() + 1

	//width of each level:
	//10 for   6 tot 7: (1<<((7-6)/2))*10 = 10
	//20 for 4,5 tot 7: 2^1*10 = 1<<1
	//40 for 2,3 tot 7: 2^2*10 = 1<<2
	//80 for 0,1 tot 7: 2^4*10 = 1<<3
	width := (1 << ((lineTot - lineNr) / 2)) * 10

	//nr=0 = own value, 1=down lines, 2=next value, 3=down, ...
	valueStr := fmt.Sprintf("%d(%d)", node.Value, node.Balance()) //len=8
	if relLineNr == 0 {
		//own value line with connected left/right
		if node.Left != nil {
			fmt.Printf("%.*s+%.*s", (width/4)+2, space, (width/2-8)/2, line)
		} else {
			fmt.Printf("%.*s", (width-8)/2, ".........")
		}
		fmt.Printf("%8.8s", valueStr)
		if node.Right != nil {
			fmt.Printf("%.*s+%.*s", (width/2-8)/2+1, line, (width / 4), space)
		} else {
			fmt.Printf("%.*s", (width-8)/2, "=========")
		}
	}
	if relLineNr == 1 {
		fmt.Printf("%.*s|%.*s|%.*s", width/4+2, space, width/2+1, space, width/4, space)
		// //fmt.Printf("(nr=%d,node=%d)", nr, node.Value)
		// if h > 1 {
		// 	//down lines for left and right
		// 	fmt.Printf("[%.*s|%.*s|%.*s]", h*2, space, h*5*2+len(valueStr), space, h*2, space)
		// } else {
		// 	//no child nodes
		// 	fmt.Printf("[%.*s]", 10+len(valueStr), space)
		// }
	}
	if relLineNr > 1 {
		//print left and right
		if node.Left != nil {
			log(t, node.Left, relLineNr-2, lineNr, lineTot)
		} else {
			fmt.Printf("%.*s", width, space)
		}
		if node.Right != nil {
			log(t, node.Right, relLineNr-2, lineNr, lineTot)
		} else {
			fmt.Printf("%.*s", width, space)
		}
	}
}
