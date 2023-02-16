package main

import "math/rand"

type Solution struct {
	Nodes []int
}

func Constructor1(head *ListNode) Solution {
	s := Solution{}
	for node := head; node != nil; node = node.Next {
		s.Nodes = append(s.Nodes, node.Val)
	}
	return s
}

func (this *Solution) GetRandom() int {
	return this.Nodes[rand.Intn(len(this.Nodes))]
}
