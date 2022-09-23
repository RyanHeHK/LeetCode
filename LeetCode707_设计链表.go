package main

type node struct {
	Value int
	Pre   *node
	Next  *node
}

type MyLinkedList struct {
	Prev *node
	Next *node
	Size int
}

func Constructor() MyLinkedList {
	return MyLinkedList{
		Prev: nil,
		Next: nil,
		Size: 0,
	}
}

func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index > this.Size-1 {
		return -1
	}
	next := this.Next
	for i := 0; i < index; i++ {
		next = next.Next
	}
	return next.Value
}

func (this *MyLinkedList) AddAtHead(val int) {
	n := &node{
		Pre:   this.Prev,
		Next:  this.Next,
		Value: val,
	}
	if this.Next != nil {
		this.Next.Pre = n
	}
	if this.Prev != nil {
		this.Prev.Next = n
	}
	this.Next = n
	if this.Prev == nil {
		this.Prev = n
	}
	this.Size++
}

func (this *MyLinkedList) AddAtTail(val int) {
	n := &node{
		Pre:   this.Prev,
		Next:  this.Next,
		Value: val,
	}
	if this.Next != nil {
		this.Next.Pre = n
	}
	if this.Prev != nil {
		this.Prev.Next = n
	}
	this.Prev = n
	if this.Next == nil {
		this.Next = n
	}
	this.Size++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.Size {
		return
	}
	if index == 0 {
		this.AddAtHead(val)
		return
	}
	if index == this.Size {
		this.AddAtTail(val)
		return
	}
	n := &node{
		Pre:   nil,
		Next:  nil,
		Value: val,
	}
	next := this.Next
	for i := 0; i < index; i++ {
		next = next.Next
	}
	n.Pre = next.Pre
	n.Next = next
	next.Pre.Next = n
	next.Pre = n
	this.Size++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if 0 <= index && index <= this.Size-1 {
		if this.Size == 1 {
			this.Next = nil
			this.Prev = nil
			this.Size--
			return
		}
		next := this.Next
		for i := 0; i < index; i++ {
			next = next.Next
		}
		this.Size--
		if next == this.Next {
			this.Next = next.Next
			this.Prev = next.Pre
		}
		if next == this.Prev {
			this.Next = next.Next
			this.Prev = next.Pre
		}
		if this.Size == 1 {
			next.Pre.Next = nil
			next.Next.Pre = nil
			return
		}
		next.Pre.Next = next.Next
		next.Next.Pre = next.Pre
	}
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
