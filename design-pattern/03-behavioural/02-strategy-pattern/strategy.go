package strategy

// 抽象算法类
type Strategy interface {
	AlgorithmInterface() string
}

// 具体算法类
type ConcreteAlgorithmA struct{}

func (a *ConcreteAlgorithmA) AlgorithmInterface() string {
	return "algorithm a"
}

type ConcreteAlgorithmB struct{}

func (b *ConcreteAlgorithmB) AlgorithmInterface() string {
	return "algorithm b"
}

type Context struct {
	strategy Strategy
}

func NewContext(strategy Strategy) *Context {
	return &Context{strategy: strategy}
}

func (context *Context) contextInterface() string {
	return context.strategy.AlgorithmInterface()
}
