package client

type Options struct {
	HttpClient HttpClient
}

type Option func(*Options)

func WithHttpClient(c HttpClient) Option {
	return func(o *Options) {
		o.HttpClient = c
	}
}
