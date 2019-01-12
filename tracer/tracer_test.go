package tracer

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	var buf bytes.Buffer

	tracer := New(&buf)
	assert.NotNil(t, tracer, "Return from New() should not be nil")

	if tracer != nil {
		tracer.Trace("Hello from tracer package")
		assert.Equal(t, "Hello from tracer package\n", buf.String(), "Trace should not write '%s'.", buf.String())
	}

}

func TestOff(t *testing.T) {
	var silentTracer = Off()
	silentTracer.Trace("something")
}
