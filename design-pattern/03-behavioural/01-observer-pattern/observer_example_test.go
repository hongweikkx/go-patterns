package observer

import (
	"container/list"
	"testing"
)

/*
   From 大话设计模式-程杰
   让前台观察boss是否进门
*/

type ReceptSub struct {
	StaffList *list.List
}

func NewReceptSub() *ReceptSub {
	return &ReceptSub{StaffList: list.New()}
}
func (sub *ReceptSub) AttachStaff(staffs ...Staff) {
	for _, staff := range staffs {
		sub.StaffList.PushBack(staff)
	}
}

func (sub *ReceptSub) DetachStaff(staff Staff) {
	for e := sub.StaffList.Front(); e != nil; e = e.Next() {
		if e.Value.(Staff) == staff {
			sub.StaffList.Remove(e)
		}
	}
}

func (sub *ReceptSub) NotifyBossComming() {
	for e := sub.StaffList.Front(); e != nil; e = e.Next() {
		e.Value.(Staff).UpdateWhenBossComming()
	}
}

func (sub *ReceptSub) NotifyBossLeave() {
	for e := sub.StaffList.Front(); e != nil; e = e.Next() {
		e.Value.(Staff).UpdateWhenBossLeave()
	}
}

type Staff interface {
	GetStatus() string
	UpdateWhenBossComming()
	UpdateWhenBossLeave()
}

type SleepStaff struct {
	Name   string
	Status string
}

func NewSleepStaff(name string) *SleepStaff {
	return &SleepStaff{
		Name:   name,
		Status: "work",
	}
}

func (staff *SleepStaff) GetStatus() string {
	return staff.Status
}

func (staff *SleepStaff) UpdateWhenBossComming() {
	staff.Status = "work hard"
}

func (staff *SleepStaff) UpdateWhenBossLeave() {
	staff.Status = "sleep"
}

func TestStaff(t *testing.T) {
	sub := NewReceptSub()
	joy := NewSleepStaff("joy")
	sub.AttachStaff(joy)
	sub.NotifyBossLeave()
	if joy.GetStatus() != "sleep" {
		t.Error("test error")
	}
	sub.DetachStaff(joy)
	sub.NotifyBossComming()
	if joy.GetStatus() != "sleep" {
		t.Error("test error:", joy.GetStatus())
	}

}
