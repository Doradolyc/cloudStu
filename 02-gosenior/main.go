package main

import (
	"fmt"
	"sync"
	"time"
)

type Queue struct {
	queue []int
	cond  *sync.Cond
}

func main() {
	queue := &Queue{
		queue: []int{},
		cond:  sync.NewCond(&sync.Mutex{}),
	}

	var wg sync.WaitGroup

	wg.Add(10)

	go func() {
		for i := 0; i < 10; i++ {
			queue.Consume()
			wg.Done()
		}
	}()

	for i := 0; i < 10; i++ {
		go queue.Produce(i)
		time.Sleep(1 * time.Second)
	}

	wg.Wait()
}

func (q *Queue) Produce(item int) {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	q.queue = append(q.queue, item)
	fmt.Printf("%d in queue, notify all\n", item)
	q.cond.Broadcast()
}

func (q *Queue) Consume() {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	if len(q.queue) == 0 {
		fmt.Println("no data available, wait")
		q.cond.Wait()
	}
	result := q.queue[0]
	q.queue = q.queue[1:]
	fmt.Println(result)
}
