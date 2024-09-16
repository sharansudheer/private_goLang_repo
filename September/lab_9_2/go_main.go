package main

import (
	"fmt"
	"time"
	"lab_9_2/time_clock"
	"sync"
)

func main(){
	var check bool = true
	var quest int
	fmt.Println("Input")
	fmt.Scanln(&quest)
	for check {

		if quest == 0 {

			fmt.Println("Nice Day")
			check = false

		} else if quest> 2 {

			fmt.Println("Wrong Input")

			fmt.Println("Input")
			fmt.Scanln(&quest)
		
		} else {
			switch quest {

			case 1:{
				processes := []*time_clock.Process{
					time_clock.NewProcess(1),
					time_clock.NewProcess(2),
					time_clock.NewProcess(3),
				}
			
				// Connect processes

				
				processes[0].Messages[1] = make(chan int)
				processes[1].Messages[0] = make(chan int)
				processes[1].Messages[2] = make(chan int)
				processes[2].Messages[1] = make(chan int)
				
				// Simulate events and messages
				var wg sync.WaitGroup
				wg.Add(3)
				
				for i := 0; i < 3; i++ {
					go func(p *time_clock.Process) {
							defer wg.Done()
			
							p.ExecuteEvent(1)
							p.ExecuteEvent(2)
							p.ExecuteEvent(3)
			
							// Process P1 specific message
							if p.Id == 1 {
									p.SendMessage(2, 2)
									p.ExecuteEvent(8)
							}
			
							// Process P2 specific messages
							if p.Id == 2 {
									p.SendMessage(1, 3)
									p.SendMessage(3, 4)
									p.ExecuteEvent(4)
									p.ExecuteEvent(5)
									p.ExecuteEvent(7)
									p.ExecuteEvent(11)
							}
			
							// Process P3 specific message
							if p.Id == 3 {
									p.SendMessage(2, 7)
									p.ExecuteEvent(5)
							}
			
							for {
									msg := p.ReceiveMessage()
									if msg == 0 {
											break
									}
									p.ExecuteEvent(msg)
							}
					}(processes[i])
			}
			wg.Wait()
			
		}	
	
	case 2:{

	
	processes := make(map[int]*time_clock.VectorClock)
	numProcesses := 3

	for i := 0; i < numProcesses; i++ {
			processes[i] = time_clock.NewVectorClock()
	}

	// Simulate events
	for i := 0; i < 10; i++ {
			processID := i % numProcesses
			processes[processID].Increment(processID)
			fmt.Printf("Process %d: %s\n", processID, processes[processID].String())

			// Simulate message passing
			if i%2 == 0 {
					senderID := i % numProcesses
					receiverID := (i + 1) % numProcesses
					processes[receiverID].Update(processes[senderID])
					fmt.Printf("Message sent from %d to %d: %s\n", senderID, receiverID, processes[receiverID].String())
			}

			time.Sleep(100 * time.Millisecond)
				}
			}
			}
		}
	}
}

/* 
Import the package: This is the simplest approach if you only need to use the Process type directly.
Use a shared interface: This method provides more flexibility if you need to work with different types 
that implement the same interface,
 or if you want to hide the implementation details of the Process type
*/

/* 
for i := 0; i < 3; i++ {
					go func(p *time_clock.Process) {
							defer wg.Done()
	
							p.ExecuteEvent(1)
							p.ExecuteEvent(2)
							p.ExecuteEvent(3)
	
							// Process 1 is the initiator, sending messages first
							if p.Id == 1 {
									p.SendMessage(2, 2)
									p.SendMessage(3, 4)
							}
	
							// Process P2 specific messages
							if p.Id == 2 {
									p.ExecuteEvent(4)
									p.ExecuteEvent(5)
									msg := <-p.Messages[0] // Wait for a message from process 1 (after sending)
									p.ExecuteEvent(msg)
									p.ExecuteEvent(7)
									p.ExecuteEvent(11)
							}
	
							// Process P3 specific message
							if p.Id == 3 {
									msg := <-p.Messages[1] // Wait for a message from process 1 (after sending)
									p.ExecuteEvent(msg)
									p.SendMessage(2, 7)
									p.ExecuteEvent(5)
							}
	
							for {
									msg := p.ReceiveMessage()
									if msg == 0 {
											break
									}
									p.ExecuteEvent(msg)
							}
					}(processes[i])
				}
*/