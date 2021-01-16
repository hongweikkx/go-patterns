package memento

// 发起人
type Originator struct {
	State string // 需要保存的属性
}

func NewOriginator(state string) *Originator {
	return &Originator{State: state}
}

func (ori *Originator) GetState() string {
	return ori.State
}

func (ori *Originator) SetState(state string) {
	ori.State = state
}

func (ori *Originator) CreateMemento() *Memento {
	return NewMemento(ori.State)
}

func (ori *Originator) SetMemento(memento *Memento) {
	ori.State = memento.GetState()
}

// 备忘录
type Memento struct {
	State string
}

func NewMemento(state string) *Memento {
	return &Memento{State: state}
}

func (mem *Memento) GetState() string {
	return mem.State
}
