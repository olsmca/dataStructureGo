package datastructurego

import (
	"fmt"
)

type Queue []string

func (q *Queue) isEmpty() bool {
	return q.Size() == 0
}

func (q *Queue) Enqueue(str string) {
	*q = append(*q, str)
}

func (q *Queue) Dequeue() (string, bool) {
	if q.isEmpty() {
		return "", true
	} else {
		element := (*q)[0]
		*q = (*q)[1:]
		return element, false
	}
}

func (q *Queue) Size() int {
	return len(*q)
}

func (q *Queue) ToString() string {
	return fmt.Sprintf("Element enqueue: %v", *q)
}
