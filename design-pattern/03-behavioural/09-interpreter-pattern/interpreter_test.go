package interpreter

import (
	"container/list"
	"testing"
)

func TestInterpreter(t *testing.T) {
	context := Context{}
	listC := list.New()
	listC.PushBack(&NonTerminalEexpression{})
	listC.PushBack(&TerminalExpression{})
	ret := "pre"
	for v := listC.Front(); v != nil; v = v.Next() {
		ret += "->" + v.Value.(AbstractExpression).Interpret(context)
	}
	if ret != "pre->nonterminal->terminal" {
		t.Error("test err:", ret)
	}
}
