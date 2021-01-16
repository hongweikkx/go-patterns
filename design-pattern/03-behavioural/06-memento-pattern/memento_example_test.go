package memento

import "testing"

/*
	游戏备忘录。
    form 大话设计模式-程杰
*/

type Role struct {
	Vit int
	Atk int
	Def int
}

func NewRole(vit, atk, def int) *Role {
	return &Role{
		Vit: vit,
		Atk: atk,
		Def: def,
	}
}

func (role *Role) GetState() (int, int, int) {
	return role.Vit, role.Atk, role.Def
}

func (role *Role) SetState(vit, atk, def int) {
	role.Vit = vit
	role.Atk = atk
	role.Def = def
}

func (role *Role) CreateMementoRole() *MementoRole {
	return NewMementoRole(role.Vit, role.Atk, role.Def)
}

func (role *Role) SetMementoRole(memento *MementoRole) {
	role.Vit = memento.Vit
	role.Def = memento.Def
	role.Atk = memento.Atk
}

func (role *Role) Fight() {
	role.Vit = 0
	role.Atk = 101
	role.Def = 101
}

type MementoRole struct {
	Vit int
	Atk int
	Def int
}

func NewMementoRole(vit, atk, def int) *MementoRole {
	return &MementoRole{
		Vit: vit,
		Atk: atk,
		Def: def,
	}
}
func (mem *MementoRole) GetSate() (int, int, int) {
	return mem.Vit, mem.Atk, mem.Def
}

func TestRole(t *testing.T) {
	role := NewRole(100, 100, 100)
	mRole := role.CreateMementoRole()
	role.Fight()
	vit, atk, def := role.GetState()
	if vit != 0 || atk != 101 || def != 101 {
		t.Error("test err")
	}
	role.SetMementoRole(mRole)
	vit, atk, def = role.GetState()
	if vit != 100 || atk != 100 || def != 100 {
		t.Error("test err")
	}
}
