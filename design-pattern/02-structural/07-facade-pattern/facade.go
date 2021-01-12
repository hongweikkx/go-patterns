package facade

type SystemOne struct{}

func (sys *SystemOne) MethodOne() string {
	return "one"
}

type SystemTwo struct{}

func (sys *SystemTwo) MethodTwo() string {
	return "two"
}

type SystemThree struct{}

func (sys *SystemThree) MethodThree() string {
	return "three"
}

type SystemFour struct{}

func (sys *SystemFour) MethodFour() string {
	return "four"
}

type Facade struct {
	SysOne   SystemOne
	SysTwo   SystemTwo
	SysThree SystemThree
	SysFour  SystemFour
}

func NewFacade() *Facade {
	return &Facade{
		SysOne:   SystemOne{},
		SysTwo:   SystemTwo{},
		SysThree: SystemThree{},
		SysFour:  SystemFour{},
	}
}

func (facade *Facade) MethodA() string {
	return facade.SysOne.MethodOne() + " " +
		facade.SysTwo.MethodTwo() + " " +
		facade.SysFour.MethodFour()
}

func (facade *Facade) MethodB() string {
	return facade.SysTwo.MethodTwo() + " " +
		facade.SysFour.MethodFour()
}
