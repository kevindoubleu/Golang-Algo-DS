package main

type Node struct {
	// example node representing dishes to be washed
	meal string
	count int
}

type Stack []Node

func (s *Stack) push(n Node) {
	*s = append(*s, n)
}

func (s *Stack) pop() *Node {
	if len(*s) > 0 {
		tmp := (*s)[len(*s)-1]
		*s = (*s)[:len(*s)-1]
		return &tmp
	}

	return nil
}

func (s Stack) peek() *Node {
	if len(s) > 0 {
		return s.pop()
	}

	return nil
}

func (s Stack) lookup(meal string) *Node {
	for s.peek() != nil {
		if s.peek().meal == meal {
			return s.pop()
		}
		s.pop()
	}

	return nil
}