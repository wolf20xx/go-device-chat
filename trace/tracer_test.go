package trace

import (
	"bytes"
	"testing"
)

func TestNew(t *testing.T) {
	var buf bytes.Buffer
	tracer := New(&buf)
	if tracer == nil {
		t.Error("New戻り値nil")
	} else {
		tracer.Trace("traceパッケージ")
		if buf.String() != "traceパッケージ\n" {
			t.Errorf("'%s'が出力", buf.String())
		}
	}
}

func TestOff(t *testing.T) {
	var silentTracer Tracer = Off()
	silentTracer.Trace("データ")
}
