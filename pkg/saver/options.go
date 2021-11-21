package saver

import "github.com/spf13/afero"

type Options struct {
	Fs afero.Fs
}

type Option func(*Options)

func WithFs(fs afero.Fs) Option {
	return func(o *Options) {
		o.Fs = fs
	}
}
