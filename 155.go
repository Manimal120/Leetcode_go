package main

type MinStack struct {
	stack, min []int
	l          int
}

func Constructor() MinStack {
	return MinStack{make([]int, 0), make([]int, 0), 0}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	if this.l == 0 {
		this.min = append(this.min, val)
	} else {
		min := this.GetMin()
		if val < min {
			this.min = append(this.min, val)
		} else {
			this.min = append(this.min, min)
		}
	}
	this.l++

}

func (this *MinStack) Pop() {
	this.l--
	this.min = this.min[:this.l]
	this.stack = this.stack[:this.l]
}

func (this *MinStack) Top() int {
	return this.stack[this.l-1]
}

func (this *MinStack) GetMin() int {
	return this.min[this.l-1]
}
