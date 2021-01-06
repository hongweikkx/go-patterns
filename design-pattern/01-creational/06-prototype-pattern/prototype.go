package prototype

import (
	"bytes"
	"encoding/gob"
)

/*
	通过将已经存在的实例赋值给新的变量来完成clone, 可定制clone对象
    这里涉及到一个深拷贝和浅拷贝.

*/
type Example struct {
	Desc string
	M    map[string]int
}

//实现深拷贝
func (e *Example) DeepClone() *Example {
	var res Example
	deepCopy(&res, e)
	return &res
}

// 实现浅拷贝
func (e *Example) ShallowCopy() *Example {
	res := *e
	return &res
}

func New() *Example {
	return &Example{
		Desc: "string to int",
		M: map[string]int{
			"1": 1,
			"2": 2,
		},
	}
}

func deepCopy(dst, src interface{}) error {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(src); err != nil {
		return err
	}
	return gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(dst)
}
