package state

type State interface {
	GetName() string
	Handle(context *Context)
}

type Context struct {
	State State
}

func NewContext(state State) *Context {
	return &Context{State: state}
}
func (context *Context) GetState() string {
	return context.State.GetName()
}
func (context *Context) Request() {
	context.State.Handle(context)
}

type StateA struct {
	Name string
}

func NewStateA() *StateA {
	return &StateA{Name: "state A"}
}
func (state *StateA) GetName() string {
	return state.Name
}
func (state *StateA) Handle(context *Context) {
	context.State = NewStateB()
}

type StateB struct {
	Name string
}

func NewStateB() *StateB {
	return &StateB{Name: "state B"}
}
func (state *StateB) GetName() string {
	return state.Name
}
func (state *StateB) Handle(context *Context) {
	context.State = NewStateA()
}
