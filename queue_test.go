package datastructurego

import (
	"fmt"
	"testing"
)

func TestQueueIsEmpty(t *testing.T) {
	var queue Queue

	status := queue.isEmpty()
	if !status {
		t.Errorf("Queue was incorrect, go %v, expect %v ", status, true)
	}
}

func TestDenqueue(t *testing.T) {
	var queue Queue

	_, isEmpty := queue.Dequeue()

	if !isEmpty {
		t.Errorf("Queue was incorrect, go %v, expect %v ", isEmpty, true)
	}
}

func TestEnqueue(t *testing.T) {
	var queue Queue

	table := []struct {
		data string
	}{
		{"This"}, {"Is"}, {"Sparta"},
	}

	for index, elem := range table {
		queue.Enqueue(elem.data)
		size := queue.Size()
		sizeIndex := index + 1

		if size != sizeIndex {
			t.Errorf("The queue size was incorrect, go %d expected %d", size, sizeIndex)
		}
	}

	fmt.Println(queue.ToString())

	for _, item := range table {
		value, _ := queue.Dequeue()

		if value != item.data {
			t.Errorf("The queue size was incorrect, go %v expected %v", value, item.data)
		}
	}
}
