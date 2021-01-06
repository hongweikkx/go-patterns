package simplefactory

/*
    工厂生产A or B 产品
	设计思想
		1. 简单工厂函数(分支case判断)
		3. 抽象产品类
		4. 具体产品对象
*/
type ProdTyp int

const (
	ProdTypA ProdTyp = iota
	ProdTypB
)

func SimpleFactory(typ ProdTyp) Product {
	switch typ {
	case ProdTypA:
		return NewAProd()
	default:
		return NewBProd()
	}
}

type Product interface {
	Describe() string
}

// prod A
type AProd struct {
	Name string
}

func NewAProd() *AProd {
	return &AProd{Name: "A"}
}

func (a *AProd) Describe() string {
	return a.Name
}

// prod B
type BProd struct {
	Name string
}

func NewBProd() *BProd {
	return &BProd{Name: "B"}
}

func (b *BProd) Describe() string {
	return b.Name
}
