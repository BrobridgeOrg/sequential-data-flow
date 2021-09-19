package sdf

import "sync"

type Task struct {
	Flow      *Flow
	Payload   interface{}
	Completed bool
	ID        int
	available sync.WaitGroup
	completed sync.WaitGroup
}

func NewTask(flow *Flow) *Task {
	t := &Task{
		Flow: flow,
	}

	t.completed.Add(1)

	return t
}

func (t *Task) Reset() {
	t.available.Add(1)
}

func (t *Task) Wait() {
	t.completed.Wait()
}

func (t *Task) Available() {
	t.available.Wait()
}

func (t *Task) Release() {
	t.completed.Add(1)
	t.Completed = false
	t.available.Done()
}

func (t *Task) Done(result interface{}) {
	t.Payload = result

	if !t.Completed {
		t.Completed = true
		t.completed.Done()
	}
}
