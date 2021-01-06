package factory

import "testing"

/*
   From 大话设计模式-程杰
   实现一个简单的计算器
*/
type OPFac interface {
	NewOp() Operation
}

type Operation interface {
	FillArgs(...int)
	GetResult() int
}
type AddOpFac struct{}

func (*AddOpFac) NewOp() Operation {
	return &AddST{}
}

type SubOpFac struct{}

func (*SubOpFac) NewOp() Operation {
	return &SubST{}
}

// add
type AddST struct {
	Args []int
}

func (st *AddST) FillArgs(args ...int) {
	st.Args = args
}
func (st *AddST) GetResult() int {
	sum := 0
	for _, e := range st.Args {
		sum += e
	}
	return sum
}

// sub
type SubST struct {
	Args []int
}

func (st *SubST) FillArgs(args ...int) {
	st.Args = args
}
func (st *SubST) GetResult() int {
	if len(st.Args) == 0 {
		return 0
	}
	ret := st.Args[0]
	for i := 1; i < len(st.Args); i++ {
		ret -= st.Args[i]
	}
	return ret
}

func TestCal(t *testing.T) {
	addOp := (&AddOpFac{}).NewOp()
	addOp.FillArgs(1, 2, 3, 4, 5, 6)
	if addOp.GetResult() != 21 {
		t.Error("test cal add error")
	}
	subOp := (&SubOpFac{}).NewOp()
	subOp.FillArgs(6, 1, 1, 1, 2)
	if subOp.GetResult() != 1 {
		t.Error("test cal sub error")
	}
}
