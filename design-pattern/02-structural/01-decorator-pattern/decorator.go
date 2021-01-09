package decorator

/*
	装饰模式使用对象组合的方式动态改变或增加对象行为， 在原对象的基础上增加功能
    它和建造者模式的区别在于建造者模式的建造过程是稳定的。
*/
type Component interface {
	GetCount() int
}

//基础结构体
type Fruit struct {
	Count       int
	Description string
}

func (f *Fruit) GetCount() int {
	return f.Count
}

type Decorator struct {
	ComponentT Component
}

func (d *Decorator) SetComponent(com Component) {
	d.ComponentT = com
}

func (d *Decorator) GetCount() int {
	return d.ComponentT.GetCount()
}

//装饰结构体
type AppleDecorator struct {
	Num int
	Decorator
}

func (apple *AppleDecorator) GetCount() int {
	return apple.ComponentT.GetCount() + apple.Num
}

func CreateAppleDecorator(n int) *AppleDecorator {
	return &AppleDecorator{Num: n}
}

type OrangeDecorator struct {
	Num int
	Decorator
}

func (orange *OrangeDecorator) GetCount() int {
	return orange.ComponentT.GetCount() + orange.Num
}

func CreateOrangeDecorator(n int) *OrangeDecorator {
	return &OrangeDecorator{Num: n}
}
