# worker
```
package main

import (
	"fmt"
	"time"
	"utils/worker"
)

func main() {

	wq := worker.New()

	wq.Start()
	for i := 0; i < 100; i++ {
		fmt.Printf("init is %v \n", i)
		func(data int) {
			wq.Invoke(func() {
				fmt.Printf("get params is %v \n", data)
			})
		}(i)
	}
	time.After(time.Second * 3)
	wq.Stop()

}
```
