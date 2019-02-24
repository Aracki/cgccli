package util

import (
	"gopkg.in/cheggaaa/pb.v1"
)

// WriteCounter counts the number of bytes written to it. It implements to the io.Writer
// interface and we can pass this into io.TeeReader() which will report progress on each
// write cycle.
type WriteCounter struct {
	Total uint64
	Bar   *pb.ProgressBar
}

func (wc *WriteCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.Total += uint64(n)
	wc.incrementProgressBar()
	return n, nil
}

func (wc WriteCounter) incrementProgressBar() {
	wc.Bar.Set64(int64(wc.Total))
}
