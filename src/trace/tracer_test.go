package trace

import (
	"bytes"
	"testing"
)

func TestNew(t *testing.T) {
	var buf bytes.Buffer
	tracer := New(&buf)
	if tracer == nil {
		t.Error("Return value is nil from \"New\"")
	} else {
		tracer.Trace("Hello, trace package.")
		if buf.String() != "Hello, trace package.\n" {
			t.Errorf("'%s'という誤った文字列が出力されました", buf.String())
		}
	}
}

func TestOFF(t *testing.T) {
	var silientTracer Tracer = Off()
	silientTracer.Trace("Data")
}
