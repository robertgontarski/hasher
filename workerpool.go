package main

import (
	"fmt"
	"log/slog"
	"sync"
)

type Worker interface {
	worker()
	Run()
}

type DefaultWorker struct {
	Tasks       []Task
	Concurrency int
	taskch      chan Task
	wg          sync.WaitGroup
	Worker
}

func NewDefaultWorker(tasks []Task, concurrency int) *DefaultWorker {
	return &DefaultWorker{
		Tasks:       tasks,
		Concurrency: concurrency,
	}
}

func (cw *DefaultWorker) worker() {
	for task := range cw.taskch {
		if err := task.Process(); err != nil {
			slog.Error(fmt.Sprintf("error while processing task: %v", err))
		}

		cw.wg.Done()
	}
}

func (cw *DefaultWorker) Run() {
	cw.taskch = make(chan Task, len(cw.Tasks))
	defer close(cw.taskch)

	for i := 0; i < cw.Concurrency; i++ {
		go cw.worker()
	}

	cw.wg = sync.WaitGroup{}
	cw.wg.Add(len(cw.Tasks))

	for _, task := range cw.Tasks {
		cw.taskch <- task
	}

	cw.wg.Wait()
}
