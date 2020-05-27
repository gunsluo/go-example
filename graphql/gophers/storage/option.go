package storage

type Option interface {
	apply(*options)
}

func WithLogger(logger int) Option {
	return &fnOption{fn: func(g *options) {
		g.Logger = logger
	}}
}

func WithDB(db int) Option {
	return &fnOption{fn: func(g *options) {
		g.DB = db
	}}
}

type fnOption struct {
	fn func(g *options)
}

func (o *fnOption) apply(g *options) {
	o.fn(g)
}

type options struct {
	Logger   int
	DB       int
	S        int
	Recorder int
	Verifier int
}

func defaultOptions() options {
	return options{Logger: 100}
}
