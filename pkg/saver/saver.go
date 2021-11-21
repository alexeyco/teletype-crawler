package saver

import (
	"fmt"
	"os"

	"github.com/spf13/afero"
)

type Saver struct {
	fs afero.Fs
}

func NewSaver(root string, opts ...Option) *Saver {
	o := Options{
		Fs: afero.NewOsFs(),
	}

	for _, opt := range opts {
		opt(&o)
	}

	return &Saver{
		fs: afero.NewBasePathFs(o.Fs, root),
	}
}

func (s *Saver) Clean() error {
	if err := s.fs.RemoveAll("."); err != nil {
		return fmt.Errorf(`error removing root directory: %w`, err)
	}

	if err := s.fs.MkdirAll(".", os.ModePerm); err != nil {
		return fmt.Errorf(`error creating root directory: %w`, err)
	}

	return nil
}

func (s *Saver) Save(fileName, content string) error {
	f, err := s.fs.OpenFile(fileName, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return fmt.Errorf(`error open file: %w`, err)
	}
	defer func() { _ = f.Close() }()

	if _, err = f.WriteString(content); err != nil {
		return fmt.Errorf(`error writing file: %w`, err)
	}

	return nil
}
