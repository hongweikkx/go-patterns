package command

import "testing"

func TestInvoker_ExecuteCommand(t *testing.T) {
	//接收者
	receivera := &ReceiverA{}
	receiverb := &ReceiverB{}
	//command
	commanda := CreateCommand(kind, receivera)
	conmandb := CreateCommand(mold, receiverb)
	//调用者
	invoker := new(Invoker)
	invoker.AddCommand(commanda)
	invoker.AddCommand(conmandb)
	//调用： 接收者a执行A的操作，b执行B的操作
	s := invoker.ExecuteCommand()
	if s != "->A action->B action" {
		t.Error("test error:", s)
	}
}
