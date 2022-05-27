package lab2

import (
	"strconv"
	"strings"
)

type StackNode struct {
	data string
	next *StackNode
}

func newNode(data string, top *StackNode) *StackNode {
	return &StackNode{
		data,
		top,
	}
}

type MyStack struct {
	top   *StackNode
	count int
}

func initStack() *MyStack {
	return &MyStack{
		nil,
		0,
	}
}

func (this MyStack) size() int {
	return this.count
}

func (this MyStack) isEmpty() bool {
	if this.size() > 0 {
		return false
	} else {
		return true
	}
}

func (this *MyStack) push(data string) {
	this.top = newNode(data, this.top)
	this.count++
}

func (this *MyStack) pop() string {
	var temp string = ""
	if this.isEmpty() == false {
		temp = this.top.data
		this.top = this.top.next
		this.count--
	}
	return temp
}

func isOperator(symbol byte) bool {
	if symbol == '+' || symbol == '-' ||
		symbol == '*' || symbol == '/' ||
		symbol == '^' || symbol == '%' {
		return true
	}
	return false
}

func isOperands(text string) bool {
	if _, err := strconv.Atoi(text); err == nil {
		return true
	}
	return false
}

func PrefixToPostfix(prefix string) (string, error) {
	var prefixarr []string = strings.Split(prefix, " ")
	var s *MyStack = initStack()
	var temp string = ""
	var op1 string = ""
	var op2 string = ""
	var postfix string
	var errorMsg error

	for i := len(prefixarr) - 1; i >= 0; i-- {
		if isOperator(prefixarr[i][0]) {
			if s.size() > 1 {
				op1 = s.pop()
				op2 = s.pop()
				temp = op1 + " " + op2 + " " + string(prefixarr[i][0])
				s.push(temp)
			}
		} else if isOperands(prefixarr[i]) {
			temp = string(prefixarr[i])
			s.push(temp)
		}
	}

	postfix = s.pop()
	errorMsg = nil

	return postfix, errorMsg
}
