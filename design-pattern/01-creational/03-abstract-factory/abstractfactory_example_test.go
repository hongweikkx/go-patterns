package abstractfactory

import (
	"fmt"
	"testing"
)

/*
   From 大话设计模式-程杰
   支持不同DB
*/

// 首先先写一个工厂模式的 支持多个db的user表的访问.
// 然后再增加其他表，已形成抽象工厂模式
type IFactory interface {
	NewUser() IUser
	NewDepartment() IDepartment
}

type IUser interface {
	Insert(user) string
	Get(int) string
}
type IDepartment interface {
	Insert(department) string
	GetDepartment(int) string
}
type user struct {
	Id   int
	Name string
}

type department struct {
	Id   int
	Name string
	Desc string
}

// sqlServer
type SqlServer struct{}

func (s *SqlServer) NewUser() IUser {
	return &SqlServerUser{}
}

func (s *SqlServer) NewDepartment() IDepartment {
	return &SqlServerDepartment{}
}

// access server
type AccessServer struct{}

func (s *AccessServer) NewUser() IUser {
	return &AccessServerUser{}
}
func (s *AccessServer) NewDepartment() IDepartment {
	return &AccessServerDepartment{}
}

type SqlServerUser struct{}

func (s *SqlServerUser) Insert(u user) string {
	return fmt.Sprintf("sqlServer user insert:%+v\n", u)
}

func (s *SqlServerUser) Get(id int) string {
	return fmt.Sprintf("sqlServer user find:%+v\n", id)
}

type SqlServerDepartment struct{}

func (s *SqlServerDepartment) Insert(d department) string {
	return fmt.Sprintf("sqlServer department insert:%+v\n", d)
}

func (s *SqlServerDepartment) GetDepartment(id int) string {
	return fmt.Sprintf("sqlServer department get:%+v\n", id)
}

type AccessServerUser struct{}

func (a *AccessServerUser) Insert(u user) string {
	return fmt.Sprintf("access Server user insert:%+v\n", u)
}

func (a *AccessServerUser) Get(id int) string {
	return fmt.Sprintf("access Server user find:%+v\n", id)
}

type AccessServerDepartment struct{}

func (s *AccessServerDepartment) Insert(d department) string {
	return fmt.Sprintf("access Server department insert:%+v\n", d)
}

func (s *AccessServerDepartment) GetDepartment(id int) string {
	return fmt.Sprintf("access Server department get:%+v\n", id)
}

func TestExample(t *testing.T) {
	//server := &SqlServer{}
	server := &AccessServer{} // 需指明是access server。 可以通过配置来优化
	u := server.NewUser()
	t.Log(u.Insert(user{Id: 1, Name: "mark"}))
	t.Log(u.Get(2))

	d := server.NewDepartment()
	t.Log(d.Insert(department{
		Id:   20,
		Name: "department20",
		Desc: "department20 desc",
	}))
	t.Log(d.GetDepartment(18))
}
