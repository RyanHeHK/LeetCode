package main

type CQueue struct {
	in, out []int
}

func Constructor2() CQueue {
	return CQueue{}
}

func (this *CQueue) AppendTail(value int) {
	this.in = append(this.in, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.out) == 0 {
		if len(this.in) == 0 {
			return -1
		}
		this.inToOut()
	}
	value := this.out[len(this.out)-1]
	this.out = this.out[:len(this.out)-1]
	return value
}

func (this *CQueue) inToOut() {
	n := len(this.in)
	for i := 0; i < n; i++ {
		this.out = append(this.out, this.in[n-1-i])
		this.in = this.in[:len(this.in)-1]
	}
}
