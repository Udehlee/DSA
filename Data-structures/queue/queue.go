package main

import "fmt"

type Queue struct {
	data []int
}

func (q *Queue) Enqueue(i int) {
	q.data = append(q.data, i)
}

func (q *Queue) Dequeue() int {
	RemoveVal := q.data[0]
	q.data = q.data[1:]
	return RemoveVal
}

func main() {

	Q := Queue{}
	fmt.Println(Q)
	Q.Enqueue(1)
	Q.Enqueue(2)
	Q.Enqueue(3)
	fmt.Println(Q)

	Q.Dequeue()
	fmt.Println(Q)

}
