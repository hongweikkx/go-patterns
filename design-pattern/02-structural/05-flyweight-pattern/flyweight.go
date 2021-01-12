package flyweight

import "fmt"

type Flyweight interface {
	Operation(int) string
}

type ConcreteFlyweight struct {
}

// 增加内部存储空间
func NewConcreteFlyweight() *ConcreteFlyweight {
	return &ConcreteFlyweight{}
}
func (con *ConcreteFlyweight) Operation(i int) string {
	return fmt.Sprint("concrete:", i)
}

func NewUnsharedConcreteFlyweight() *UnsharedConcreteFlyweight {
	return &UnsharedConcreteFlyweight{}
}

type UnsharedConcreteFlyweight struct{}

func (unshared *UnsharedConcreteFlyweight) Operation(i int) string {
	return fmt.Sprint("unshared:", i)
}

// 享元工厂, 用于创建并且管理Flyweight状态
type FlyweightFactory struct {
	Flyweights map[string]Flyweight
}

func NewFlyweightFactory() *FlyweightFactory {
	flyweights := map[string]Flyweight{}
	flyweights["X"] = NewConcreteFlyweight()
	flyweights["Y"] = NewConcreteFlyweight()
	flyweights["Z"] = NewConcreteFlyweight()
	return &FlyweightFactory{Flyweights: flyweights}
}

func (fac *FlyweightFactory) GetFlyweight(key string) Flyweight {
	if v, ok := fac.Flyweights[key]; ok {
		return v
	}
	return nil
}
