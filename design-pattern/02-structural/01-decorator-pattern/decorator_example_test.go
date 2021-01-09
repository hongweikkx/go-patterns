package decorator

import (
	"fmt"
	"testing"
)

/*
	From 大话设计模式-程杰
    实现穿衣程序
*/

type PersonI interface {
	Show() string
}

type Person struct {
	Name string
}

func (p *Person) Show() string {
	return fmt.Sprintf(" on:%+v", p.Name)
}

type Finery struct {
	ComponentT PersonI
}

func (f *Finery) Decorate(i PersonI) {
	f.ComponentT = i
}
func (f *Finery) Show() string {
	return f.ComponentT.Show()
}

type TShirt struct {
	Finery
}

func (t *TShirt) Show() string {
	return fmt.Sprintf("t-shirt ") + t.ComponentT.Show()
}

type BigTrouser struct {
	Finery
}

func (t *BigTrouser) Show() string {
	return fmt.Sprintf("bigTrouser ") + t.ComponentT.Show()
}

func TestShow(t *testing.T) {
	p := &Person{Name: "hwgao"}
	tShirt := &TShirt{}
	bigTrouser := &BigTrouser{}
	bigTrouser.Decorate(p)
	tShirt.Decorate(bigTrouser)
	h := tShirt.Show()
	if h != "t-shirt bigTrouser  on:hwgao" {
		t.Errorf("error in testShow")

	}
}
