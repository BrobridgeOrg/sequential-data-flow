package sdf

import (
	"sync"
	"testing"
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
	opts.BufferSize = 1

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
