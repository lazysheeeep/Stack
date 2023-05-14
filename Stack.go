package main

import "fmt"

type Stack struct {
	Element []any
	Top     int
	Size    int
}

func NewStack(size int) *Stack { //初始化栈
	return &Stack{
		Size:    size,
		Top:     0,
		Element: make([]any, size),
	}
}

func (temp *Stack) Length() int { //返回栈内元素个数
	return temp.Top
}

func (temp *Stack) IsFull() bool { //判断栈是否已满
	return temp.Top == temp.Size
}

func (temp *Stack) IsEmpty() bool { //判断栈是否为空
	return temp.Top == 0
}

func (temp *Stack) Push(element any) bool { //压栈
	if temp.Size == temp.Top {
		fmt.Println("栈满")
		return false
	}
	temp.Element[temp.Top] = element
	temp.Top++
	return true
}

func (temp *Stack) Pop() any {
	if temp.IsEmpty() {
		fmt.Println("栈已空")
		return nil //在写这里的时候没有写return nil，可能会让函数出不来
	}
	temp.Top--
	return temp.Element[temp.Top] //这里top不用加一，因为栈顶本来指向的就是一个空位置
}

func (temp *Stack) Peek() any {
	if temp.IsEmpty() {
		fmt.Println("栈为空")
		return nil
	}
	return temp.Element[temp.Top-1]
}
