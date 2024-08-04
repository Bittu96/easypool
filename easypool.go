package easypool

import (
	"sync"
	"time"
)

type WaitCondition int

const (
	WaitTillNoTasks WaitCondition = iota // default
	WaitForever
	WaitRelease
)

type pool struct {
	wg            *sync.WaitGroup
	mux           *sync.Mutex
	id            int
	inflow        *chan interface{}
	outflow       *chan interface{}
	taskFunc      func(interface{}) interface{}
	waitCondition WaitCondition
	tasksRuntime  time.Duration
	poolRuntime   time.Duration
}

var poolIndex = 0

func New(taskFunc func(interface{}) interface{}) *pool {
	poolId := poolIndex + 1
	return &pool{
		wg:            &sync.WaitGroup{},
		mux:           &sync.Mutex{},
		taskFunc:      taskFunc,
		waitCondition: 0,
		id:            poolId,
	}
}

func (p *pool) AddInflow(inflow *chan interface{}) *pool {
	p.inflow = inflow
	return p
}

func (p *pool) AddOutflow(outflow *chan interface{}) *pool {
	p.outflow = outflow
	return p
}

func (p *pool) AddWaitCondition(waitCondition WaitCondition) *pool {
	p.waitCondition = waitCondition
	return p
}

// func (p *pool) getPoolBenefit() float64 {
// 	r, t := float64(p.poolRuntime.Nanoseconds()), float64(p.tasksRuntime.Nanoseconds())
// 	runtimeBenefit := (t - r) / r * 100
// 	return runtimeBenefit
// }

func (p *pool) Deploy(poolSize int) {
	startTime := time.Now()
	p.wg.Add(poolSize)
	for id := 1; id <= poolSize; id++ {
		if p.inflow != nil {
			go p.flowBot(id)
		} else {
			go p.fastBot(id)
		}
	}
	p.wg.Wait()
	p.poolRuntime = time.Since(startTime)

	// fmt.Println("easypool runtime:", p.poolRuntime)
	// fmt.Println("easypool task runtime:", p.tasksRuntime)
	// fmt.Println("easypool benefit:", p.getPoolBenefit(), "%")
}

func (p *pool) flowBot(_ int) {
	// log.Println("flow bot started", id)
	defer p.wg.Done()

	for {
		select {
		case msg := <-*p.inflow:
			if p.waitCondition == WaitRelease {
				return
			}
			p.easyBot(msg)
		default:
			if p.waitCondition == WaitTillNoTasks || p.waitCondition == WaitRelease {
				return
			} else if p.waitCondition == WaitForever {
				continue
			}
		}
	}
}

func (p *pool) fastBot(id int) {
	// log.Println("fast bot started", id)
	defer p.wg.Done()

	p.easyBot(id)
}

func (p *pool) easyBot(msg interface{}) {
	taskStartTime := time.Now()
	res := p.taskFunc(msg)
	taskDuration := time.Since(taskStartTime)

	p.mux.Lock()
	if p.outflow != nil {
		*p.outflow <- res
	}
	p.tasksRuntime += taskDuration
	p.mux.Unlock()
}
