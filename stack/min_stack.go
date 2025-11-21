package stack

import "math"

// 155. 最小栈
//
// https://leetcode.cn/problems/min-stack/
//
// 设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。
//
// 实现 MinStack 类:
//
//	MinStack() 初始化堆栈对象。
//	void push(int val) 将元素val推入堆栈。
//	void pop() 删除堆栈顶部的元素。
//	int top() 获取堆栈顶部的元素。
//	int getMin() 获取堆栈中的最小元素。
//
// 使用两个栈实现最小栈
// 一个存放栈元素
// 另一个存放对应的栈中最小元素
type MinStack struct {
	stack    []int
	minStack []int
}

func Constructor() MinStack {
	return MinStack{
		stack:    []int{},
		minStack: []int{math.MaxInt64},
	}
}

func (ms *MinStack) Push(x int) {
	ms.stack = append(ms.stack, x)
	top := ms.minStack[len(ms.minStack)-1]
	ms.minStack = append(ms.minStack, min(x, top))
}

func (ms *MinStack) Pop() {
	ms.stack = ms.stack[:len(ms.stack)-1]
	ms.minStack = ms.minStack[:len(ms.minStack)-1]
}

func (ms *MinStack) Top() int {
	return ms.stack[len(ms.stack)-1]
}

func (ms *MinStack) GetMin() int {
	return ms.minStack[len(ms.minStack)-1]
}
