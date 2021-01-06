package prototype

import (
	"testing"
)

func TestExample_Clone(t *testing.T) {
	shallowOrigin := New()
	shallowCurrent := shallowOrigin.ShallowCopy()
	shallowCurrent.M["1"] = 2
	shallowCurrent.Desc = "copyed"
	if shallowOrigin.Desc == "copyed" || shallowOrigin.M["1"] == 1 {
		t.Errorf("shallow copy err")
	}

	deepOrigin := New()
	deepCurrent := deepOrigin.DeepClone()
	deepCurrent.M["1"] = 2
	deepCurrent.Desc = "copyed"
	if deepOrigin.Desc == "copyed" || deepOrigin.M["1"] == 2 {
		t.Errorf("deep copy err")
	}

}
