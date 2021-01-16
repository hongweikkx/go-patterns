package memento

import "testing"

func TestMemento(t *testing.T) {
	ori := NewOriginator("on")
	if ori.GetState() != "on" {
		t.Error("test err")
	}

	// 保存
	mem := ori.CreateMemento()

	ori.SetState("off")
	if ori.GetState() != "off" {
		t.Error("test err")
	}

	// 回档
	ori.SetMemento(mem)
	if ori.GetState() != "on" {
		t.Error("test err")
	}
}
