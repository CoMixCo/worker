package worker

import "fmt"

type WorkQueue struct {
	work    chan func()
	quit    chan bool
	Working bool
}

func New() *WorkQueue {
	return &WorkQueue{
		work:    make(chan func(), 10),
		quit:    make(chan bool),
		Working: false,
	}
}

func (wq *WorkQueue) Invoke(work func()) {
	wq.work <- work
}

func (wq *WorkQueue) Stop() {
	fmt.Println("stoping...")
	wq.quit <- true
}

func (wq *WorkQueue) Start() {
	fmt.Println("starting...")
	if wq.Working {
		return
	}
	go func() {
		for {
			select {
			case <-wq.quit:
				fmt.Println("stoped!")
				wq.Working = false
				return
			case f := <-wq.work:
				f()
			}
		}
	}()
	wq.Working = true
	fmt.Println("started!")
}
