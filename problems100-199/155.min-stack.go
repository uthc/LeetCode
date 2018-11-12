/* 最小栈
设计一个支持 push，pop，top 操作，并能在常数时间内检索到最小元素的栈。

push(x) -- 将元素 x 推入栈中。
pop() -- 删除栈顶的元素。
top() -- 获取栈顶元素。
getMin() -- 检索栈中的最小元素。
示例:

MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin();   --> 返回 -3.
minStack.pop();
minStack.top();      --> 返回 0.
minStack.getMin();   --> 返回 -2. */

type MinStack struct {
	Head  int
	stack []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	m := new(MinStack)
	m.Head = 0
	m.stack = []int{}
	return *m
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	this.Head++
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:this.Head-1]
	this.Head--
}

func (this *MinStack) Top() int {
	return this.stack[this.Head-1]
}

func (this *MinStack) GetMin() int {
	if this.Head == 0 {
		return 0
	}
	min := this.stack[0]
	for i := 0; i < this.Head; i++ {
		if min > this.stack[i] {
			min = this.stack[i]
		}
	}
	return min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */