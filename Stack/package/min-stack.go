package _package

// 设置一个支持push pop top 操作,在常数时间内找到的最小元素的栈

type MinStack struct {
	min []int
	stack []int
}

func Construction() MinStack {
	return MinStack{
		min:   make([]int, 0),
		stack: make([]int, 0),
	}
}

func (this *MinStack) Push(x int) {
	// 普通栈直接加入
	this.stack = append(this.stack, x)
	// 从min栈取出第一个元素
	result := this.FindMin(x)
	// 比较加入元素
	if result < x {
		this.min = append(this.min, result)
	} else {
		this.min = append(this.min, x)
	}

}

// 从min栈上取出第一个, 若为空返回负值
func (this *MinStack) FindMin(x int) int {
	if len(this.min) == 0 {
		return 1 << 32
	}
	return this.min[len(this.min) - 1]
}

// 弹出元素
func (this *MinStack) Pop() {
	if len(this.stack) == 0 {
		return
	}
	this.stack = this.stack[:len(this.stack)-1]
	this.min = this.min[:len(this.min)-1]
}

func (this *MinStack) Top() int {
	if len(this.stack) == 0 {
		return 1 << 32
	}
	return this.stack[len(this.stack)-1]
}

