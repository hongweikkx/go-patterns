package template

import (
	"testing"
)

func TestTemplate(t *testing.T) {
	ea := &ConcreteElementA{}
	tmp := Tmp{Name: "A", SubTmp: ea}
	ea.Tmp = tmp
	if ea.Method() != "A1:ATemp:A" {
		t.Error("test error")
	}

	eb := &ConcreteElementB{}
	tmp = Tmp{Name: "B", SubTmp: eb}
	eb.Tmp = tmp
	if eb.Method() != "Temp:BB2:B" {
		t.Error("test error:")
	}
}
