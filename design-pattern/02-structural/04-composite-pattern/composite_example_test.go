package composite

import (
	"container/list"
	"errors"
	"fmt"
	"testing"
)

type Company interface {
	Add(company Company) error
	Remove(company Company) error
	Display(depth int) string
	LineOfDuty() string
}

type ConcreteCompany struct {
	Name     string
	Children *list.List
}

func NewConCreteCompany(name string) *ConcreteCompany {
	return &ConcreteCompany{
		Name:     name,
		Children: list.New(),
	}
}
func (c *ConcreteCompany) Add(company Company) error {
	c.Children.PushBack(company)
	return nil
}

func (c *ConcreteCompany) Remove(company Company) error {
	for i := c.Children.Front(); i != nil; i = i.Next() {
		if i.Value.(Company) == company {
			c.Children.Remove(i)
			break
		}
	}
	return nil
}

func (c *ConcreteCompany) Display(depth int) string {
	pre := fmt.Sprintf("depth:%d, name:%s, children Len:%d\n", depth, c.Name, c.Children.Len())
	for i := c.Children.Front(); i != nil; i = i.Next() {
		pre += i.Value.(Company).Display(depth + 1)
	}
	return pre
}

func (c *ConcreteCompany) LineOfDuty() string {
	duty := ""
	for i := c.Children.Front(); i != nil; i = i.Next() {
		duty += i.Value.(Company).LineOfDuty()
	}
	return duty
}

type HRDepartment struct {
	Name string
}

func NewHRDepartment(name string) *HRDepartment {
	return &HRDepartment{Name: name}
}
func (hr *HRDepartment) Add(company Company) error {
	return errors.New("hr department can not add")
}

func (hr *HRDepartment) Remove(company Company) error {
	return errors.New("hr department can not remove")
}

func (hr *HRDepartment) Display(depth int) string {
	return fmt.Sprintf("depth:%d, hr department name:%s\n", depth, hr.Name)
}

func (hr *HRDepartment) LineOfDuty() string {
	return fmt.Sprintf("hr department name:%s line of duty\n", hr.Name)
}

type FinanceDepartment struct {
	Name string
}

func NewFinaceDepartment(name string) *FinanceDepartment {
	return &FinanceDepartment{Name: name}
}

func (fw *FinanceDepartment) Add(company Company) error {
	return errors.New("finance department can not add")
}

func (fw *FinanceDepartment) Remove(company Company) error {
	return errors.New("finance department can not remove")
}

func (fw *FinanceDepartment) Display(depth int) string {
	return fmt.Sprintf("depth:%d, finance department name:%s\n", depth, fw.Name)
}

func (fw *FinanceDepartment) LineOfDuty() string {
	return fmt.Sprintf("finance department name:%s line of duty\n", fw.Name)
}

func TestCompositeExample(t *testing.T) {
	bj := NewConCreteCompany("北京总公司")
	bj.Add(NewHRDepartment("总公司人力资源部门"))
	bj.Add(NewFinaceDepartment("总公司财务部门"))

	sh := NewConCreteCompany("上海华东办事处")
	sh.Add(NewHRDepartment("华东人力资源部门"))
	sh.Add(NewFinaceDepartment("华东财务部门"))
	bj.Add(sh)

	nj := NewConCreteCompany("南京办事处")
	nj.Add(NewHRDepartment("南京人力资源部门"))
	nj.Add(NewFinaceDepartment("南京财务部门"))
	sh.Add(nj)

	hz := NewConCreteCompany("杭州办事处")
	hz.Add(NewHRDepartment("杭州人力资源部门"))
	hz.Add(NewFinaceDepartment("杭州财务部门"))
	sh.Add(hz)

	t.Log(bj.Display(1))
	t.Log(bj.LineOfDuty())
}
