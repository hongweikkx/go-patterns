/* From 大话设计模式: 简单工厂类 vs 工厂方法模式
	简单工厂类的做法是在工厂类中增加了判断的case语句。如果需要增加一个分支条件,那么就要修改原来的简单工厂类。
    这说明简单工厂类不仅对扩展开放，还对修改开放， 这是违反开放-封闭原则的。
    工厂方法模式对简单工厂类进行了进一步的抽象，定义一个用于创建对象的接口，让子类决定实例化哪一个类。
*/
package factory

/*
    A工厂产出A产品，B工厂产出B产品
	设计思想：
		*抽象工厂类
		*抽象产品类
		*具体工厂
		*具体产品
*/
type Factory interface {
	NewProd() Product
}
type Product interface {
	Describe() string
}

type AFactory struct{}

func (*AFactory) NewProd() Product {
	return &AProd{Name: "A"}
}

type BFactory struct{}

func (*BFactory) NewProd() Product {
	return &BProd{Name: "B"}
}

// prod A
type AProd struct {
	Name string
}

func (a *AProd) Describe() string {
	return a.Name
}

// prod B
type BProd struct {
	Name string
}

func (b *BProd) Describe() string {
	return b.Name
}
