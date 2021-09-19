package sdf

import (
	"encoding/json"
	"testing"
)

var smallPayload []byte = []byte("{\"sample\":\"sample\"}")

func BenchmarkBaseline(b *testing.B) {

	b.Run("Small", func(b *testing.B) {

		payload, _ := json.Marshal(&Payload{
			String: "string",
			Number: 99999,
		})

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			var result map[string]interface{}
			json.Unmarshal(payload, &result)
		}
	})

	b.Run("Large", func(b *testing.B) {

		data := Payload{
			String:   "string",
			Number:   99999,
			Elements: make([]string, 0, 1000),
		}

		for i := 0; i < 1000; i++ {
			data.Elements = append(data.Elements, "sample")
		}

		payload, _ := json.Marshal(&data)

		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			var result map[string]interface{}
			json.Unmarshal(payload, &result)
		}
	})
}

func BenchmarkLowBufferSize(b *testing.B) {

	// Create Options object
	options := NewOptions()
	options.BufferSize = 1024
	options.Handler = func(data interface{}, output func(interface{})) {
		var result map[string]interface{}
		json.Unmarshal(data.([]byte), &result)
		output(result)
	}

	b.Run("Small", func(b *testing.B) {

		// Create flow with options
		flow := NewFlow(options)

		// Prepare json
		payload, _ := json.Marshal(&Payload{
			String: "string",
			Number: 99999,
		})

		b.ResetTimer()
		go func() {
			for i := 0; i < b.N; i++ {
				flow.Push(payload)
			}
		}()

		for i := 0; i < b.N; i++ {
			<-flow.Output()
		}

		flow.Close()
	})

	b.Run("Large", func(b *testing.B) {

		// Create flow with options
		flow := NewFlow(options)

		// Prepare json
		data := Payload{
			String:   "string",
			Number:   99999,
			Elements: make([]string, 0, 1000),
		}

		for i := 0; i < 1000; i++ {
			data.Elements = append(data.Elements, "sample")
		}

		payload, _ := json.Marshal(&data)

		b.ResetTimer()
		go func() {
			for i := 0; i < b.N; i++ {
				flow.Push(payload)
			}
		}()

		for i := 0; i < b.N; i++ {
			<-flow.Output()
		}

		flow.Close()
	})
}

func BenchmarkHighBufferSize(b *testing.B) {

	// Create Options object
	options := NewOptions()
	options.BufferSize = 1024000
	options.Handler = func(data interface{}, output func(interface{})) {
		var result map[string]interface{}
		json.Unmarshal(data.([]byte), &result)
		output(result)
	}

	b.Run("Small", func(b *testing.B) {

		// Create flow with options
		flow := NewFlow(options)

		// Prepare json
		payload, _ := json.Marshal(&Payload{
			String: "string",
			Number: 99999,
		})

		b.ResetTimer()
		go func() {
			for i := 0; i < b.N; i++ {
				flow.Push(payload)
			}
		}()

		for i := 0; i < b.N; i++ {
			<-flow.Output()
		}

		flow.Close()
	})

	b.Run("Large", func(b *testing.B) {

		// Create flow with options
		flow := NewFlow(options)

		// Prepare json
		data := Payload{
			String:   "string",
			Number:   99999,
			Elements: make([]string, 0, 1000),
		}

		for i := 0; i < 1000; i++ {
			data.Elements = append(data.Elements, "sample")
		}

		payload, _ := json.Marshal(&data)

		b.ResetTimer()
		go func() {
			for i := 0; i < b.N; i++ {
				flow.Push(payload)
			}
		}()

		for i := 0; i < b.N; i++ {
			<-flow.Output()
		}

		flow.Close()
	})
}

func BenchmarkHighWorkerCount(b *testing.B) {

	// Create Options object
	options := NewOptions()
	options.WorkerCount = 8
	options.BufferSize = 1024
	options.Handler = func(data interface{}, output func(interface{})) {
		var result map[string]interface{}
		json.Unmarshal(data.([]byte), &result)
		output(result)
	}

	b.Run("Small", func(b *testing.B) {

		// Create flow with options
		flow := NewFlow(options)

		// Prepare json
		payload, _ := json.Marshal(&Payload{
			String: "string",
			Number: 99999,
		})

		b.ResetTimer()
		go func() {
			for i := 0; i < b.N; i++ {
				flow.Push(payload)
			}
		}()

		for i := 0; i < b.N; i++ {
			<-flow.Output()
		}

		flow.Close()
	})

	b.Run("Large", func(b *testing.B) {

		// Create flow with options
		flow := NewFlow(options)

		// Prepare json
		data := Payload{
			String:   "string",
			Number:   99999,
			Elements: make([]string, 0, 1000),
		}

		for i := 0; i < 1000; i++ {
			data.Elements = append(data.Elements, "sample")
		}

		payload, _ := json.Marshal(&data)

		b.ResetTimer()
		go func() {
			for i := 0; i < b.N; i++ {
				flow.Push(payload)
			}
		}()

		for i := 0; i < b.N; i++ {
			<-flow.Output()
		}

		flow.Close()
	})
}

func BenchmarkHighBufferSizeAndWorkerCount(b *testing.B) {

	// Create Options object
	options := NewOptions()
	options.WorkerCount = 8
	options.BufferSize = 1024000
	options.Handler = func(data interface{}, output func(interface{})) {
		var result map[string]interface{}
		json.Unmarshal(data.([]byte), &result)
		output(result)
	}

	b.Run("Small", func(b *testing.B) {

		// Create flow with options
		flow := NewFlow(options)

		// Prepare json
		payload, _ := json.Marshal(&Payload{
			String: "string",
			Number: 99999,
		})

		b.ResetTimer()
		go func() {
			for i := 0; i < b.N; i++ {
				flow.Push(payload)
			}
		}()

		for i := 0; i < b.N; i++ {
			<-flow.Output()
		}

		flow.Close()
	})

	b.Run("Large", func(b *testing.B) {

		// Create flow with options
		flow := NewFlow(options)

		// Prepare json
		data := Payload{
			String:   "string",
			Number:   99999,
			Elements: make([]string, 0, 1000),
		}

		for i := 0; i < 1000; i++ {
			data.Elements = append(data.Elements, "sample")
		}

		payload, _ := json.Marshal(&data)

		b.ResetTimer()
		go func() {
			for i := 0; i < b.N; i++ {
				flow.Push(payload)
			}
		}()

		for i := 0; i < b.N; i++ {
			<-flow.Output()
		}

		flow.Close()
	})
}
