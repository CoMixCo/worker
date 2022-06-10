package worker

import (
	"fmt"
	"testing"
	"time"
)

func TestRun(t *testing.T) {

	wq := New()

	wq.Start()
	for i := 0; i < 5; i++ {
		fmt.Printf("init is %v \n", i)
		func(args ...any) {
			wq.Invoke(func() {
				fmt.Printf("get params is %v \n", args[0])
			})
		}(i)
	}

	go func() {
		for {
			select {
			case <-time.After(time.Second * 3):
				wq.Stop()
				return
			}
		}
	}()

	for {
	}

}
