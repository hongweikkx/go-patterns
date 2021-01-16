package visitor

/*
	TODO 不太懂
*/

import "container/list"

// 访问者接口
type Visitor interface {
	VisitConcreteElementA(*ConcreteElementA) string
	VisitConcreteElementB(*ConcreteElementB) string
}

// 具体访问者1
type Visitor1 struct {
	Name string
}

func NewVisitor1(name string) Visitor {
	return &Visitor1{Name: name}
}

func (visit *Visitor1) VisitConcreteElementA(conA *ConcreteElementA) string {
	return visit.Name + conA.Name
}
func (visit *Visitor1) VisitConcreteElementB(conB *ConcreteElementB) string {
	return visit.Name + conB.Name
}

// 具体访问者2
type Visitor2 struct {
	Name string
}

func NewVisitor2(name string) Visitor {
	return &Visitor2{Name: name}
}

func (visit *Visitor2) VisitConcreteElementA(conA *ConcreteElementA) string {
	return visit.Name + conA.Name
}
func (visit *Visitor2) VisitConcreteElementB(conB *ConcreteElementB) string {
	return visit.Name + conB.Name
}

type Element interface {
	Accept(visitor Visitor) string
}

type ConcreteElementA struct {
	Name string
}

func (conA *ConcreteElementA) Accept(visitor Visitor) string {
	return visitor.VisitConcreteElementA(conA)
	// other func
}

type ConcreteElementB struct {
	Name string
}

func (conB *ConcreteElementB) Accept(visitor Visitor) string {
	return visitor.VisitConcreteElementB(conB)
	// other func
}

type ObjectSub struct {
	IList *list.List
}

func NewObjectSub() *ObjectSub {
	return &ObjectSub{IList: list.New()}

}
func (o *ObjectSub) Attach(e Element) {
	o.IList.PushBack(e)
}
func (o *ObjectSub) Detach(e Element) {
	for v := o.IList.Front(); v != nil; v = v.Next() {
		if v.Value.(Element) == e {
			o.IList.Remove(v)
		}
	}
}

func (o *ObjectSub) Accept(visitor Visitor) string {
	sum := ""
	for v := o.IList.Front(); v != nil; v = v.Next() {
		sum += v.Value.(Element).Accept(visitor)
	}
	return sum
}
