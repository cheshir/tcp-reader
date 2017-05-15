package main

import (
	"io"
)

func newOutput(writer io.Writer, filter filterer) io.Writer {
	return &output{
		filter: filter,
		writer: writer,
	}
}

type output struct {
	filter filterer
	writer io.Writer
}

func (output output) Write(message []byte) (int, error) {
	return output.writer.Write(output.filter.Filter(message))
}
