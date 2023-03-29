package db

import "errors"

var (
	// You are trying to get a next item from the finished iterator.
	ErrIterationFinished = errors.New("iteration finished")
)
