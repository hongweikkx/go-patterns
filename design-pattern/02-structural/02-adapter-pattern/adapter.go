package adapter

/*
	If the Target and Adaptee has similarities then the adapter has just to delegate
	the requests from the Target to the Adaptee.
	If Target and Adaptee are not similar, then the adapter might have to convert the
	data structures between those and to implement the operations required by the Target
	but not implemented by the Adaptee
*/

type Target interface {
	Request() string
}

type Adapter struct {
	Adaptee
}

func NewAdapter(adaptee Adaptee) *Adapter {
	return &Adapter{adaptee}

}
func (a *Adapter) Request() string {
	return a.SpecialRequest()
}

type Adaptee struct {
}

func NewAdaptee() Adaptee {
	return Adaptee{}
}
func (a *Adaptee) SpecialRequest() string {
	return "special req"
}
