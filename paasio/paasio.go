package paasio

import (
	"io"
	"sync"
)

type counter struct {
	count int64
	ops   int
	mutex *sync.Mutex
}

type paasReadCounter struct {
	counter
	read func(p []byte) (n int, err error)
}

type paasWriteCounter struct {
	counter
	write func(p []byte) (n int, err error)
}

type paasReadWriteCounter struct {
	paasReadCounter
	paasWriteCounter
}

func (rc paasReadCounter) ReadCount() (n int64, nops int) {
	return rc.count, rc.ops
}

func (wc paasWriteCounter) WriteCount() (n int64, nops int) {
	return wc.count, wc.ops
}

// NewReadCounter creates a new ReadCounter
func NewReadCounter(r io.Reader) ReadCounter {
	rc := &paasReadCounter{read: r.Read}
	rc.mutex = &sync.Mutex{}
	return rc
}

// NewWriteCounter returns a new WriteCounter
func NewWriteCounter(w io.Writer) WriteCounter {
	wc := &paasWriteCounter{write: w.Write}
	wc.mutex = &sync.Mutex{}
	return wc
}

// NewReadWriteCounter returns a new ReadWriteCounter
func NewReadWriteCounter(rw io.ReadWriter) ReadWriteCounter {
	rwc := &paasReadWriteCounter{}
	rwc.paasReadCounter.read = rw.Read
	rwc.paasReadCounter.mutex = &sync.Mutex{}
	rwc.paasWriteCounter.write = rw.Write
	rwc.paasWriteCounter.mutex = &sync.Mutex{}
	return rwc
}

func (rc *paasReadCounter) Read(p []byte) (n int, err error) {
	rc.mutex.Lock()
	defer rc.mutex.Unlock()

	n, err = rc.read(p)
	if err != nil {
		return
	}

	rc.ops++
	rc.count += int64(n)

	return
}

func (wc *paasWriteCounter) Write(p []byte) (n int, err error) {
	wc.mutex.Lock()
	defer wc.mutex.Unlock()

	n, err = wc.write(p)
	if err != nil {
		return
	}

	wc.ops++
	wc.count += int64(n)

	return
}
