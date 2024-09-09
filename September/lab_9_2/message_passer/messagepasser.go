package messagepasser
import (
	"fmt"
)

func Messagepasser() {
	// Create a channel for communication
	messages := make(chan string)

	// Create 3 processes as Go routines
	go sender(messages)
	go receiver1(messages)
	go receiver2(messages)

	// Wait for all processes to finish
	for i := 0; i < 3; i++ {
			<-messages
	}
}

func sender(messages chan<- string) {
	for i := 0; i < 3; i++ {
			messages <- fmt.Sprintf("Message %d from sender", i)
	}
}

func receiver1(messages <-chan string) {
	for message := range messages {
			fmt.Println("Receiver 1 received:", message)
	}
}

func receiver2(messages <-chan string) {
	for message := range messages {
			fmt.Println("Receiver 2 received:", message)
	}
}
/* 
You can use buffered channels to control message flow and prevent blocking.
For more complex scenarios, consider using select statements to handle multiple channels simultaneously.
If you need to synchronize processes, you can use techniques like waitgroups or semaphores.
*/