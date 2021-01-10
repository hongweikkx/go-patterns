package composite

import (
	"container/list"
	"errors"
	"fmt"
)

/*
	设计思想：
		struct不依赖interface
		1. 包含角色：
			1). 共同的接口MenuComponent，为root和leaf结构体共有的方法
			2). root结构体(包含leaf列表)和leaf结构体
			3). 将结构体中共同部分的数据抽离，使用匿名组合的方式实现2中的两类结构体
*/

type Component interface {
	Add(component Component) error
	Remove(component Component) error
	Display(int) string
}

// 树枝节点
type Composite struct {
	Name     string
	Children *list.List
}

func NewComposite(name string) *Composite {
	return &Composite{
		Name:     name,
		Children: list.New(),
	}
}

func (composite *Composite) Add(c Component) error {
	composite.Children.PushBack(c)
	return nil
}

func (composite *Composite) Remove(c Component) error {
	for i := composite.Children.Front(); i != nil; i = i.Next() {
		if i.Value.(Component) == c {
			composite.Children.Remove(i)
			break
		}
	}
	return nil
}

func (composite *Composite) Display(depth int) string {
	pre := fmt.Sprintf("depth:%d, name:%s, children Len:%d\n", depth, composite.Name, composite.Children.Len())
	for i := composite.Children.Front(); i != nil; i = i.Next() {
		pre += i.Value.(Component).Display(depth + 1)
	}
	return pre
}

// 树叶
type Leaf struct {
	Name string
}

func NewLeaf(name string) Component {
	return &Leaf{Name: name}
}

func (leaf *Leaf) Add(c Component) error {
	return errors.New("Cannot add to a leaf")
}

func (leaf *Leaf) Remove(c Component) error {
	return errors.New("Cannot remove from a leaf")
}

func (leaf *Leaf) Display(depth int) string {
	return fmt.Sprintf("depth:%d, name:%s\n", depth, leaf.Name)
}
