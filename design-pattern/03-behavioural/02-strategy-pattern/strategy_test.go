package strategy

import "testing"

func TestStrategy(t *testing.T) {
	c := NewContext(&ConcreteAlgorithmB{})
	if c.contextInterface() != "algorithm b" {
		t.Error("test strategy error")
	}
}
