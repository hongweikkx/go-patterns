package bridge

/*
	分离抽象部分和实现部分， 桥接设计的核心在于抽象接口和组合抽象接口的结构体
    form me: 抽象和实现分离
	设计思想：
		1. 抽象接口，(实现该接口的具体struct，可扩展多个struct)
		2. 属性为抽象接口的struct Phone（桥接层）
		3. 与Phone组合模式的具体struct（可以是多个struct）
*/

type Implementor interface {
	Operation() string
}

type ConcreateImplementorA struct{}

func (a *ConcreateImplementorA) Operation() string {
	return "implementor a"
}

type ConcreateImplementorB struct{}

func (b *ConcreateImplementorB) Operation() string {
	return "implementor b"
}

type Abstraction struct {
	Imp Implementor
}

func NewAbstraction() *Abstraction {
	return &Abstraction{}
}
func (abs *Abstraction) SetImplementor(imp Implementor) {
	abs.Imp = imp
}
func (abs *Abstraction) Operation() string {
	return abs.Imp.Operation()
}
