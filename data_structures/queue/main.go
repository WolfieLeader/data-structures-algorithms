package main

import (
	"fmt"

	arraydeque "github.com/WolfieLeader/data-structures-algorithms/data_structures/queue/array_deque"
	arrayqueue "github.com/WolfieLeader/data-structures-algorithms/data_structures/queue/array_queue"
	listdeque "github.com/WolfieLeader/data-structures-algorithms/data_structures/queue/linked_list_deque"
	listqueue "github.com/WolfieLeader/data-structures-algorithms/data_structures/queue/linked_list_queue"
)

func main() {
	fmt.Println("Array Queue Example:")
	arrayQueueExample()
	fmt.Println()

	fmt.Println("Linked List Queue Example:")
	linkedListQueueExample()
	fmt.Println()

	fmt.Println("Array Deque Example:")
	arrayDequeExample()
	fmt.Println()

	fmt.Println("Linked List Deque Example:")
	linkedListDequeExample()
	fmt.Println()
}

func arrayQueueExample() {
	q := arrayqueue.New[int]()
	q.Enqueue(1, 2, 3)
	fmt.Printf("- Queue:      %v\n", q)
	q.Dequeue()
	fmt.Printf("- Dequeued:   %v\n", q)
	for i := range 10 {
		q.Enqueue(i)
	}
	fmt.Printf("- Enqueued:   %v\n", q)
	v, _ := q.Peek()
	fmt.Printf("- Peeked:     %v\n", v)
	for range 5 {
		q.Dequeue()
	}
	fmt.Printf("- Dequeued:   %v\n", q)
	v, _ = q.Peek()
	fmt.Printf("- Peeked:     %v\n", v)
}

func linkedListQueueExample() {
	q := listqueue.New[int]()
	q.Enqueue(1, 2, 3)
	fmt.Printf("- Queue:      %v\n", q)
	q.Dequeue()
	fmt.Printf("- Dequeued:   %v\n", q)
	for i := range 10 {
		q.Enqueue(i)
	}
	fmt.Printf("- Enqueued:   %v\n", q)
	v, _ := q.Peek()
	fmt.Printf("- Peeked:     %v\n", v)
	for range 5 {
		q.Dequeue()
	}
	fmt.Printf("- Dequeued:   %v\n", q)
	v, _ = q.Peek()
	fmt.Printf("- Peeked:     %v\n", v)
}

func arrayDequeExample() {
	d := arraydeque.New[int]()
	d.EnqueueLast(1, 2, 3)
	fmt.Printf("- Queue:      %v\n", d)
	d.DequeueFirst()
	fmt.Printf("- Dequeued:   %v\n", d)
	for i := range 10 {
		d.EnqueueFirst(i)
	}
	fmt.Printf("- Enqueued:   %v\n", d)
	v, _ := d.PeekLast()
	fmt.Printf("- Peeked:     %v\n", v)
	for range 5 {
		d.DequeueFirst()
	}
	fmt.Printf("- Dequeued:   %v\n", d)
	v, _ = d.PeekFirst()
	fmt.Printf("- Peeked:     %v\n", v)
}

func linkedListDequeExample() {
	d := listdeque.New[int]()
	d.EnqueueLast(1, 2, 3)
	fmt.Printf("- Queue:      %v\n", d)
	d.DequeueFirst()
	fmt.Printf("- Dequeued:   %v\n", d)
	for i := range 10 {
		d.EnqueueFirst(i)
	}
	fmt.Printf("- Enqueued:   %v\n", d)
	v, _ := d.PeekLast()
	fmt.Printf("- Peeked:     %v\n", v)
	for range 5 {
		d.DequeueFirst()
	}
	fmt.Printf("- Dequeued:   %v\n", d)
	v, _ = d.PeekFirst()
	fmt.Printf("- Peeked:     %v\n", v)
}
