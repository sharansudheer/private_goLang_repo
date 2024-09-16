package time_clock

import (
	"fmt"
)

type Process struct {
	Id       int
	Events   []int
	Messages map[int]chan int
}

func NewProcess(id int) *Process {
	return &Process{
			Id:       id,
			Events:   []int{},
			Messages: make(map[int]chan int),
	}
}

func (p *Process) ExecuteEvent(event int) {
	p.Events = append(p.Events, event)
	fmt.Printf("Process %d executed event %d\n", p.Id, event)
}

func (p *Process) SendMessage(to int, event int) {
	p.Messages[to] <- event
}

func (p *Process) ReceiveMessage() int {
	for _, ch := range p.Messages {
			if msg := <-ch; msg != 0 {
					return msg
			}
	}
	return 0
}




// type ScalarClock struct {
// 	time int
// }

// func ScalarTime() {
// 	// Create 3 processes with unique starting times
// 	process1 := ScalarClock{time: 1}
// 	process2 := ScalarClock{time: 2}
// 	process3 := ScalarClock{time: 3}

// 	// Create channels for communication
// 	messages1to2 := make(chan string)
// 	messages2to1 := make(chan string)
// 	messages2to3 := make(chan string)
// 	messages3to2 := make(chan string)

// 	// Create Go routines for the processes
// 	go process(process1, messages1to2, messages2to1)
// 	go process(process2, messages2to1, messages1to2, messages2to3, messages3to2)
// 	go process(process3, messages3to2, messages2to3)

// 	// Simulate message passing and event logging
// 	// ... (add your specific message passing and logging logic here)
// }

// func process(p ScalarClock, sendTo1, receiveFrom1, sendTo2, receiveFrom2 chan string) {
// 	// ... (your process logic, including message sending, receiving, and clock updates)
// }

/*
By following these steps and considering the additional considerations,
you can effectively implement a scalar clock with 3 processes and 4 events
(message sending and receiving) using GoLang to analyze 
causality and concurrency in your distributed system. 
*/

/*
You can customize the message structure to include additional information if needed.
Consider using a more advanced data structure for the scalar clock
if you need to handle more complex scenarios.
For larger-scale systems, you may want to explore distributed
clock synchronization algorithms like Vector Clocks or Lamport Clocks.
*/