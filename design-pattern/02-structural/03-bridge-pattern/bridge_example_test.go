package bridge

import (
	"fmt"
	"testing"
)

//抽象接口
type SoftWare interface {
	Run() string
}

//具体类型CPU和Storage
type Cpu struct{}

func (c *Cpu) Run() string {
	return fmt.Sprint("this is cpu run")
}

type Storage struct{}

func (s *Storage) Run() string {
	return fmt.Sprint("this is storage run")
}

//桥接层结构体
type Phone struct {
	software SoftWare
}

//赋值具体software
func (s *Phone) SetSoftWare(soft SoftWare) {
	s.software = soft
}

//Apple结构体
type Apple struct {
	phone Phone
}

func (p *Apple) SetShape(soft SoftWare) {
	p.phone.SetSoftWare(soft)
}

func (p *Apple) Print() string {
	return p.phone.software.Run()
}

//HuaWei结构体
type HuaWei struct {
	phone Phone
}

func (p *HuaWei) SetShape(soft SoftWare) {
	p.phone.SetSoftWare(soft)
}

func (p *HuaWei) Print() string {
	return p.phone.software.Run()
}

func TestCpuRun(t *testing.T) {
	cpu := &Cpu{}
	apple := new(Apple)
	apple.SetShape(cpu)
	if apple.Print() != "this is cpu run" {
		t.Error("TestCpuRun fail")
	}
}

func TestStorageRun(t *testing.T) {
	storage := &Storage{}
	hw := new(HuaWei)
	hw.SetShape(storage)
	if hw.Print() != "this is storage run" {
		t.Error("TestStorageRun fail")
	}
}
