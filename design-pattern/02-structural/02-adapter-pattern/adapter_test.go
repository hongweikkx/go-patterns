package adapter

import (
	"testing"
)

func TestAdapter(t *testing.T) {
	adapter := NewAdapter(NewAdaptee())
	req := adapter.Request()
	if req != "special req" {
		t.Error("TestAdapter fail")
	}
}
