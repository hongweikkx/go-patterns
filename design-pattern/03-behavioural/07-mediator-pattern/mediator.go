package mediator

type Mediator interface {
	Send(message string, colleague *Colleague) string
}

type Colleague struct {
	Name string
	Mediator
}

func NewColleague(name string, mediator Mediator) *Colleague {
	return &Colleague{Name: name, Mediator: mediator}
}

func (c *Colleague) Notify(message string) string {
	return c.Name + ":" + message
}
func (c *Colleague) Send(message string) string {
	return c.Mediator.Send(message, c)
}

type ConcreteMediator struct {
	Colleague1 *Colleague
	Colleague2 *Colleague
}

func NewConcreteMediator() *ConcreteMediator {
	return &ConcreteMediator{}
}

func (c *ConcreteMediator) Send(message string, colleague *Colleague) string {
	if colleague.Name != c.Colleague1.Name {
		return c.Colleague1.Notify(message)
	}
	if colleague.Name != c.Colleague2.Name {
		return c.Colleague2.Notify(message)
	}
	return ""
}
