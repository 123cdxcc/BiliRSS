package pool

import (
	"sync"
)

type Worker struct {
	name  string
	flag  chan bool
	tasks chan *Task
	wg    *sync.WaitGroup
	sync.Mutex
}

func NewWork(name string, tasks chan *Task, wg *sync.WaitGroup) *Worker {
	return &Worker{
		name:  name,
		flag:  make(chan bool),
		tasks: tasks,
		wg:    wg,
	}
}
func (w *Worker) Start() {
	go func(w *Worker) {
		for {
			select {
			case task := <-w.tasks:
				task.Run(task.params...)
				w.wg.Done()
			case <-w.flag:
				return
			}
		}
	}(w)
}
func (w *Worker) Stop() {
	w.flag <- false
}
