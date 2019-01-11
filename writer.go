package throttlewriter

import (
	"io"
	"io/ioutil"
	"time"
)

type Writer struct {
	writer         io.Writer
	written        uint64
	capacity       uint64
	timeWindow     time.Duration
	lastTimeWindow time.Time
}

func New(writer io.Writer, capacity int, timewindow time.Duration) *Writer {
	return &Writer{
		writer:         writer,
		written:        0,
		capacity:       uint64(capacity),
		timeWindow:     timewindow,
		lastTimeWindow: time.Now(),
	}
}

func (tw *Writer) Write(p []byte) (n int, err error) {
	now := time.Now()
	if time.Now().Sub(tw.lastTimeWindow) > tw.timeWindow {
		// Reset the prev.
		tw.lastTimeWindow = now
		tw.written = 0
	}
	if tw.written >= tw.capacity {
		return ioutil.Discard.Write(p)
	}
	tw.written += uint64(len(p))
	return tw.writer.Write(p)
}
