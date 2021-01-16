package observer

import "container/list"

type Subject struct {
	Observers *list.List
}

func NewSubject() *Subject {
	return &Subject{Observers: list.New()}
}

func (subject *Subject) Attach(obs ...Observer) {
	for _, ob := range obs {
		subject.Observers.PushBack(ob)
	}
}

func (subject *Subject) Detach(ob Observer) {
	for e := subject.Observers.Front(); e != nil; e = e.Next() {
		if e.Value.(Observer) == ob {
			subject.Observers.Remove(e)
		}
	}
}

func (subject *Subject) Notify() {
	for e := subject.Observers.Front(); e != nil; e = e.Next() {
		e.Value.(Observer).Update()
	}
}

type Observer interface {
	Get() bool
	Update()
}

type ConcreteObserver struct {
	Name     string
	IsUpdate bool
}

func NewConcreteObserver(name string) Observer {
	return &ConcreteObserver{IsUpdate: false, Name: name}

}
func (con *ConcreteObserver) Update() {
	con.IsUpdate = true
}

func (con *ConcreteObserver) Get() bool {
	return con.IsUpdate
}
