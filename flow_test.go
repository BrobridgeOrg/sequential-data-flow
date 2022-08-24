package sdf

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Payload struct {
	String   string   `json:"string,omitempty"`
	Number   int      `json:"number,omitempty"`
	Elements []string `json:"elements,omitempty"`
}

func TestPush(t *testing.T) {

	opts := NewOptions()

	flow := NewFlow(opts)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)

		// Prepare payload
		payload := &Payload{
			Number: i,
		}

		flow.Push(payload)
	}

	go func() {
		for i := 0; i < 10; i++ {
			data := <-flow.Output()
			payload := data.(*Payload)

			if payload.Number != i {
				t.Fail()
				continue
			}

			wg.Done()
		}
	}()

	wg.Wait()

	flow.Close()
}

func TestFullBuffer(t *testing.T) {

	opts := NewOptions()
	opts.BufferSize = 10

	flow := NewFlow(opts)

	lastSeq := 0
	for i := 0; i < opts.BufferSize+1; i++ {

		// Prepare payload
		payload := &Payload{
			Number: i,
		}

		select {
		case flow.input <- payload:
			lastSeq = i
		default:
			// Failed to push data because it is full
			assert.Equal(t, opts.BufferSize, i)
		}
	}

	assert.Equal(t, opts.BufferSize-1, lastSeq)

	// Pull data from buffer
	data := <-flow.Output()

	// First record
	assert.Equal(t, 0, data.(*Payload).Number)

	// Push records to make it full again
	target := lastSeq + 2
	for i := lastSeq + 1; i < target; i++ {

		// Prepare payload
		payload := &Payload{
			Number: i,
		}

		select {
		case flow.input <- payload:
			lastSeq = i
		default:
			// Failed to push data because it is full
			assert.Equal(t, target, i)
		}
	}
	t.Log(lastSeq)

	assert.Equal(t, opts.BufferSize, lastSeq)

	for i := 1; i <= lastSeq; i++ {
		data := <-flow.Output()
		payload := data.(*Payload)

		assert.Equal(t, i, payload.Number)
	}

	flow.Close()
}

func TestRunWithFullBuffer(t *testing.T) {

	opts := NewOptions()
	opts.BufferSize = 10

	flow := NewFlow(opts)

	go func() {
		for i := 0; i < 10000; i++ {

			// Prepare payload
			payload := &Payload{
				Number: i,
			}

			flow.Push(payload)
		}
	}()

	for i := 0; i < 10000; i++ {
		data := <-flow.Output()
		payload := data.(*Payload)

		if payload.Number != i {
			t.Fail()
			continue
		}
	}

	flow.Close()
}

func Test1MBuffer(t *testing.T) {

	opts := NewOptions()
	opts.BufferSize = 1000

	n := 10000000

	flow := NewFlow(opts)

	go func() {
		for i := 0; i < n; i++ {

			// Prepare payload
			payload := &Payload{
				Number: i,
			}

			flow.Push(payload)
		}
	}()

	for i := 0; i < n; i++ {
		data := <-flow.Output()
		payload := data.(*Payload)

		if payload.Number != i {
			t.Fail()
			continue
		}
	}

	flow.Close()
}
