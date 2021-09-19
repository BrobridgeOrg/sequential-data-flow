package sdf

import (
	"container/ring"
)

type Flow struct {
	options   *Options
	input     chan interface{}
	taskQueue chan *Task
	output    chan interface{}
	buffer    *ring.Ring
	cursor    *ring.Ring
}

func NewFlow(options *Options) *Flow {

	f := &Flow{
		options:   options,
		input:     make(chan interface{}, options.BufferSize),
		taskQueue: make(chan *Task, options.BufferSize),
		output:    make(chan interface{}, options.BufferSize),
		buffer:    ring.New(options.BufferSize),
	}

	// Setup cursor
	f.cursor = f.buffer

	f.initialize()

	return f
}

func (f *Flow) initialize() {

	// Initial buffer
	cur := f.buffer
	for i := 0; i < f.options.BufferSize; i++ {
		cur.Value = NewTask(f)
		cur = cur.Next()
	}

	// Start workers
	for i := 0; i < f.options.WorkerCount; i++ {
		go f.startWorker()
	}

	go f.startPublisher()
	go f.startReceiver()
}

func (f *Flow) startReceiver() {

	for data := range f.input {

		var task *Task
		if f.buffer.Value == nil {
			task = NewTask(f)
		} else {
			task = f.buffer.Value.(*Task)
		}

		task.Available()
		task.Reset()

		// set task status
		task.Payload = data

		// Put data into buffer
		f.buffer.Value = task
		f.buffer = f.buffer.Next()

		f.taskQueue <- task
	}
}

func (f *Flow) startWorker() {

	for task := range f.taskQueue {
		f.options.Handler(task.Payload, task.Done)
	}
}

func (f *Flow) startPublisher() {

	for {

		if f.cursor.Value == nil {
			continue
		}

		task := f.cursor.Value.(*Task)
		task.Wait()

		f.done(task.Payload)

		task.Release()
		f.cursor = f.cursor.Next()
	}
}

func (f *Flow) done(data interface{}) {
	f.output <- data
}

func (f *Flow) Push(data interface{}) {
	f.input <- data
}

func (f *Flow) Output() chan interface{} {
	return f.output
}

func (f *Flow) Close() {
	close(f.input)
	close(f.taskQueue)
	close(f.output)
}
