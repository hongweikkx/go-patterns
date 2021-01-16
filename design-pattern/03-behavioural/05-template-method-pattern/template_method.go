package template

type TmpT interface {
	PrimitiveOP1() string
	PrimitiveOP2() string
	Method() string
}

type Tmp struct {
	Name   string
	SubTmp TmpT
}

func (tmp *Tmp) PrimitiveOP1() string {
	return "Temp:" + tmp.Name
}
func (tmp *Tmp) PrimitiveOP2() string {
	return "Temp:" + tmp.Name
}
func (tmp *Tmp) Method() string {
	return tmp.SubTmp.PrimitiveOP1() + tmp.SubTmp.PrimitiveOP2()
}

type ConcreteElementA struct {
	Tmp
}

func (c *ConcreteElementA) PrimitiveOP1() string {
	return "A1:" + c.Name
}

type ConcreteElementB struct {
	Tmp
}

func (c *ConcreteElementB) PrimitiveOP2() string {
	return "B2:" + c.Name
}
