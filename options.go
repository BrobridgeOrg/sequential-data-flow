package sdf

type Options struct {
	BufferSize  int
	WorkerCount int
	Handler     func(interface{}, func(interface{}))
}

func NewOptions() *Options {
	return &Options{
		BufferSize:  1024,
		WorkerCount: 4,
		Handler: func(data interface{}, done func(interface{})) {
			done(data)
		},
	}
}
