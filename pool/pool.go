package pool

import (
	"strconv"
	"sync"
)

type Pool struct {
	tasks   chan *Task
	works   []*Worker
	running int
	cap     int
	wg      *sync.WaitGroup
	sync.Mutex
}

func NewPool(cap int) *Pool {
	tasks := make(chan *Task, cap)
	return &Pool{
		tasks:   tasks,
		works:   make([]*Worker, cap),
		running: 0,
		cap:     cap,
		wg:      &sync.WaitGroup{},
	}
}
func (p *Pool) Go(fun func(v ...interface{}), params ...interface{}) {
	p.Lock()
	defer p.Unlock()
	p.wg.Add(1)
	t := &Task{
		Run:    fun,
		params: params,
	}
	if p.running < p.cap {
		p.works[p.running] = NewWork(strconv.Itoa(p.running), p.tasks, p.wg)
		p.works[p.running].Start()
		p.running++
	}
	p.tasks <- t
}
func (p *Pool) Wait() {
	p.wg.Wait()
}
