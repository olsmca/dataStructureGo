package datastructurego

import (
	"fmt"
	"testing"
)

func TestStackIsEmpty(t *testing.T) {
	var stack Stack

	if !stack.isEmpty() {
		t.Errorf("Stack was incorrect, go %v expected %v", stack.isEmpty(), true)
	}
}

func TestStackPopEmpty(t *testing.T) {
	var stack Stack

	_, isEmpty := stack.Pop()

	if !isEmpty {
		t.Errorf("Stack was incorrect, go %v expected %v", stack.isEmpty(), true)
	}
}

func TestStackPush(t *testing.T) {
	var stack Stack

	tables := []struct {
		data string
	}{
		{"this"},
		{"is"},
		{"sparta"},
	}

	for index, item := range tables {
		stack.Push(item.data)
		var size = stack.Size()
		var sizeIndex = index + 1

		if size != sizeIndex {
			t.Errorf("The stack size was incorrect, go %d expected %d", size, sizeIndex)
		}
	}

	fmt.Println(stack.ToString())

	newSize := stack.Size()

	for i := 1; i <= newSize; i++ {
		stack.Pop()
	}

	if !stack.isEmpty() {
		t.Errorf("Stack was incorrect, go %v expected %v", stack.isEmpty(), true)
	}
}
