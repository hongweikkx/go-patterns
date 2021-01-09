package decorator

import (
	"testing"
)

func TestCreateAppleDecorator(t *testing.T) {
	fruit := &Fruit{Count: 0, Description: "水果计数"}
	apple := CreateAppleDecorator(20)
	orange := CreateOrangeDecorator(30)
	apple.SetComponent(fruit)
	orange.SetComponent(apple)
	re := orange.GetCount()
	if re != 50 {
		t.Errorf("装饰错误，期待结果为%d, 实际为:%d", 50, re)
	}
}
