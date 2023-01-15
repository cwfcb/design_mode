package singleton

import "testing"

func TestGetInstanceSafety(t *testing.T) {
	goroutineNum := 100

	receiveChan := make(chan *SingleInstance, goroutineNum)

	for i := 0; i < goroutineNum; i++ {
		go func() {
			receiveChan <- GetInstanceSafety()
		}()
	}

	for i := 0; i < goroutineNum; i++ {
		_ = <-receiveChan
	}

}
