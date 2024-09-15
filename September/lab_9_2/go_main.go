package main

import (
	"fmt"
	"time"
	"lab_9_2/time_clock"
)

func main(){
	processes := make(map[int]*time_clock.VectorClock)
	numProcesses := 3

	for i := 0; i < numProcesses; i++ {
			processes[i] = NewVectorClock()
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