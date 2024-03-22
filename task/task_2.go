package task

import "fmt"

// Task2 demonstrates the use of Channel and GoRoutines to roar all felines asynchronously
// The FelineService.RoarAllAsync receive a number of async goroutines
// The felines are sent to a channel and the async goroutines receive the felines from the channel
// We are using wait groups to wait for all async goroutines to finish
func (t *Task) Task2() {
	roaredFelines := t.felineService.RoarAllAsync(4)
	fmt.Printf("Roared %d felines\n", roaredFelines)
}
