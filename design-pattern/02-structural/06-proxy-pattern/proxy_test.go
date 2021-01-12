package proxy

import "testing"

func TestProxyObject(t *testing.T) {
	objectRun := &Object{action: "run"}

	proxyObject := new(ProxyObject)
	proxyObject.object = objectRun
	if proxyObject.ObjDo("run") != "i can run" {
		t.Error("test error")
	}
	if proxyObject.ObjDo("fly") != "i cannot fly" {
		t.Error("test error")
	}
}
