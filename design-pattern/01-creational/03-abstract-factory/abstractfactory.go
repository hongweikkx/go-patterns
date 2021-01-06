// 大话设计模式: 抽象工厂类: 提供一个创建一系列相关或者相互依赖的接口，而无需指定他们具体的类
package abstractfactory

/*
    A工厂生产A1, A2产品, B 工厂生产B1,B2产品
	设计思想：
		*抽象工厂类
		*抽象产品类
		*具体工厂
		*具体产品
*/

type Factory interface {
	NewProd1() Product1
	NewProd2() Product2
}
type Product1 interface {
	Describe1() string
}
type Product2 interface {
	Describe2() string
}

type AFactory struct{}

func (*AFactory) NewProd1() Product1 {
	return &AProd1{Name: "A1"}
}
func (*AFactory) NewProd2() Product2 {
	return &AProd2{Name: "A2"}
}

type BFactory struct{}

func (*BFactory) NewProd1() Product1 {
	return &BProd1{Name: "B1"}
}
func (*BFactory) NewProd2() Product2 {
	return &BProd2{Name: "B2"}
}

// A 工厂产出的产品
type AProd1 struct {
	Name string
}

func (a *AProd1) Describe1() string {
	return a.Name
}

type AProd2 struct {
	Name string
}

func (a *AProd2) Describe2() string {
	return a.Name
}

// prod B
type BProd1 struct {
	Name string
}

func (b *BProd1) Describe1() string {
	return b.Name
}

type BProd2 struct {
	Name string
}

func (b *BProd2) Describe2() string {
	return b.Name
}
