package main

import "fmt"

const INT_MAX = int(^uint(0) >> 1)
const INT_MIN = ^INT_MAX

func (t *BST) isValid() bool {
	if t.root != nil {
		return t.root.left._isValid(INT_MIN, t.root.data) &&
			t.root.right._isValid(t.root.data, INT_MAX)
	}
	// empty tree is valid, right?
	return true
}
func (n *Node) _isValid(min, max int) bool {
	if n != nil {
		fmt.Println("isValid called on", n.data, "min", min, "max", max)
		// check within boundary
		if n.data < min || n.data > max {
			fmt.Println("!!! illegal node", n.data, "!!!")
			return false
		}

		// set new min max at current node
		return n.left._isValid(min, n.data) &&
			n.right._isValid(n.data, max)
	}

	// reached end of tree branch and no premature returns
	// means this branch is valid
	return true
}

// helper func for making invalid BSTs
func (t *BST) get(data int) *Node {
	return t.root._get(data)
}
func (n *Node) _get(data int) *Node {
	if n != nil {
		if n.data == data {
			return n
		} else if data > n.data {
			return n.right._get(data)
		} else if data < n.data {
			return n.left._get(data)
		}
	}

	return nil
}

func main() {
	var valid BST
	valid.insert(10)
	valid.insert(5)
	valid.insert(20)
	valid.insert(15)
	valid.print()
	fmt.Println(valid.isValid())

	invalid := valid
	invalid.get(20).left = nil
	invalid.get(5).right = &Node{data: 15}
	invalid.print()
	fmt.Println(invalid.isValid())

	var valid2 BST
	valid2.insert(10)
	valid2.insert(0)
	valid2.insert(12)
	valid2.insert(-1)
	valid2.insert(4)
	valid2.insert(11)
	valid2.insert(20)
	valid2.insert(-20)
	valid2.print()
	fmt.Println(valid2.isValid())

	invalid2 := valid2
	invalid2.get(11).left = &Node{data: 9}
	invalid2.print()
	fmt.Println(invalid2.isValid())

	invalid3 := valid2
	invalid3.insert(7)
	invalid3.insert(5)
	invalid3.get(5).left = &Node{data: 3}
	invalid3.print()
	fmt.Println(invalid3.isValid())
}