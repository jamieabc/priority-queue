package main

import (
	"fmt"
	"github.com/jamieabc/priority-queue/pkg/priority_queue"
)

func main() {
	h := priority_queue.New(10)
	h.Offer(5)
	h.Offer(3)
	h.Offer(-2)
	h.Offer(10)
	fmt.Printf("polled: %d\n", h.Poll().(int))

	h.Offer(-1)
	fmt.Println("peek: ", h.Peek().(int))

	_ = h.Poll()
	fmt.Println("peek: ", h.Peek().(int))
}
