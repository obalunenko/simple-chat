package tracer

import (
	"fmt"
	"io"
	"log"
)

// Tracer is the interface that describes an object capable
// of tracing events
type Tracer interface {
	Trace(a ...interface{})
}

// tracer implements Tracer interface
type tracer struct {
	out io.Writer
}

// New creates a new instance of a Tracer
func New(w io.Writer) Tracer {
	return &tracer{out: w}

}

// Trace write tracing events
func (t *tracer) Trace(a ...interface{}) {
	if _, err := fmt.Fprint(t.out, a...); err != nil {
		log.Fatalf("failed to trace: %v", a...)
	}
	if _, err := fmt.Fprintln(t.out); err != nil {
		log.Fatalf("failed to add line to trace")
	}

}

// nilTracer implements Tracer interface to cover case when tracing is not needed
type nilTracer struct{}

// Trace implements Tracer interface and do nothing
func (nt *nilTracer) Trace(a ...interface{}) {}

// Off creates a Tracer that will ignore calls to Trace
func Off() Tracer {
	return &nilTracer{}
}
