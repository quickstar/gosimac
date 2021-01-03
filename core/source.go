package core

import "io"

// Source represents source for image background.
type Source interface {
	Init() (int, error)                             // call once on source and return number of available images to fetch
	Name() string                                   // name of source in string format
	Fetch(index int) (string, io.ReadCloser, error) // fetch image from source
}
