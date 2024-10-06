package main

import (
	"fmt"
	"sync"
	"time"
)

// Message represents the structure of a message between processes
type Message struct {
	sender  int
	content string
}

// SnapshotMarker represents a snapshot marker
type SnapshotMarker struct {
	processID int
}

// Process represents a distributed process
type Process struct {
	id              int
	channelIn       chan Message
	channelOut      chan Message
	markerIn        chan SnapshotMarker
	markerOut       chan SnapshotMarker
	localState      string
	receivedMarkers int
	wg              *sync.WaitGroup
}

// StartProcess simulates process behavior
func (p *Process) StartProcess(initiator bool) {
	defer p.wg.Done()
	if initiator {
		time.Sleep(1 * time.Second) // Delay to ensure other processes are running
		fmt.Printf("Process %d initiating snapshot.\n", p.id)
		p.initiateSnapshot()
	}
	for {
		select {
		case msg := <-p.channelIn:
			p.handleMessage(msg)
		case marker := <-p.markerIn:
			p.handleMarker(marker)
		}
	}
}

// handleMessage handles incoming messages
func (p *Process) handleMessage(msg Message) {
	fmt.Printf("Process %d received message from Process %d: %s\n", p.id, msg.sender, msg.content)
	p.localState = fmt.Sprintf("Received message: %s", msg.content)
}

// handleMarker handles snapshot marker messages
func (p *Process) handleMarker(marker SnapshotMarker) {
	fmt.Printf("Process %d received snapshot marker from Process %d.\n", p.id, marker.processID)
	p.receivedMarkers++
	if p.receivedMarkers == 1 {
		p.saveState()
		p.sendMarker()
	}
}

// initiateSnapshot starts the snapshot process
func (p *Process) initiateSnapshot() {
	p.saveState()
	p.sendMarker()
}

// saveState simulates saving the local state
func (p *Process) saveState() {
	fmt.Printf("Process %d saving local state: %s\n", p.id, p.localState)
}

// sendMarker sends a snapshot marker to all connected processes
func (p *Process) sendMarker() {
	marker := SnapshotMarker{processID: p.id}
	p.markerOut <- marker
}

// SendMessages simulates sending messages between processes
func (p *Process) SendMessages() {
	messages := []string{"Hello", "How are you?", "I'm good", "Nice to hear", "Bye", "See you", "Take care", "Later"}
	for i, msg := range messages {
		time.Sleep(time.Duration(i) * time.Second)
		p.channelOut <- Message{sender: p.id, content: msg}
		fmt.Printf("Process %d sent message: %s\n", p.id, msg)
	}
}

func main() {
	// Create channels for communication between processes
	processChannels := make([]chan Message, 4)
	markerChannels := make([]chan SnapshotMarker, 4)

	for i := 0; i < 4; i++ {
		processChannels[i] = make(chan Message)
		markerChannels[i] = make(chan SnapshotMarker)
	}

	// Initialize processes
	var wg sync.WaitGroup
	wg.Add(4)
	processes := make([]*Process, 4)
	for i := 0; i < 4; i++ {
		processes[i] = &Process{
			id:              i + 1,
			channelIn:       processChannels[(i+1)%4],
			channelOut:      processChannels[i],
			markerIn:        markerChannels[(i+1)%4],
			markerOut:       markerChannels[i],
			localState:      fmt.Sprintf("Initial state of Process %d", i+1),
			receivedMarkers: 0,
			wg:              &wg,
		}
		go processes[i].StartProcess(i == 0) // Process 1 initiates the snapshot
	}

	// Simulate message passing between processes
	for i := 0; i < 4; i++ {
		go processes[i].SendMessages()
	}

	// Wait for processes to finish
	wg.Wait()
}
