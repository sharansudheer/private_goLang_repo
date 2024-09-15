package time_clock

import (
	"fmt"
	"sync"
)

type VectorClock struct {
	clocks map[int]int
	mu     sync.Mutex
}

func NewVectorClock() *VectorClock {
	return &VectorClock{
			clocks: make(map[int]int),
	}
}

func (vc *VectorClock) Increment(processID int) {
	vc.mu.Lock()
	defer vc.mu.Unlock()

	vc.clocks[processID]++
}

func (vc *VectorClock) Update(otherVC *VectorClock) {
	vc.mu.Lock()
	defer vc.mu.Unlock()

	for processID, clock := range otherVC.clocks {
			if vc.clocks[processID] < clock {
					vc.clocks[processID] = clock
			}
	}
}

func (vc *VectorClock) String() string {
	var s string
	for processID, clock := range vc.clocks {
			s += fmt.Sprintf("%d:%d ", processID, clock)
	}
	return s
}

