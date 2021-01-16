package visitor

import "testing"

func TestVisitor(t *testing.T) {
	obj := NewObjectSub()
	obj.Attach(&ConcreteElementA{Name: "EA"})
	obj.Attach(&ConcreteElementB{Name: "EB"})

	v1 := NewVisitor1("v1")
	v2 := NewVisitor2("v2")

	if obj.Accept(v1) != "v1EAv1EB" {
		t.Error("test error")
	}
	if obj.Accept(v2) != "v2EAv2EB" {
		t.Error("test error")
	}
}
