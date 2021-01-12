package facade

import (
	"fmt"
	"testing"
)

/*
	设计思想:
		通过组合模式来实现外观模式, 为子系统实现统一的访问api
		1. Facade struct: 属性为子系统的结构体
		2. 子系统结构体
*/
//示例：微服务框架： 音乐服务、视频服务
//创建子服务struct
//music struct
type Music struct {
	Name string
}

func (m *Music) GetMusic() string {
	return m.Name
}

//Video struct
type Video struct {
	Id int64
}

func (v *Video) GetVideoId() int64 {
	return v.Id
}

//count struct
type Count struct {
	Comment int64
	Praise  int64
	Collect int64
}

func (c *Count) GetComment() int64 {
	return c.Comment
}

//外观结构Facade
type FacadeAll struct {
	music Music
	count Count
	video Video
}

//对外访问接口
func (f *FacadeAll) PrintServerInfo() string {
	ret := fmt.Sprintf("%s %d %d", f.music.GetMusic(), f.video.GetVideoId(), f.count.GetComment())
	return ret
}

func NewFacadeAll(music Music, count Count, video Video) *FacadeAll {
	return &FacadeAll{
		music: music,
		video: video,
		count: count,
	}
}

func TestNewFacade(t *testing.T) {
	music := Music{"love"}
	video := Video{1}
	count := Count{12, 30, 5}
	facade := NewFacadeAll(music, count, video)
	s := facade.PrintServerInfo()
	if s != "love 1 12" {
		t.Error("test new facade error")
	}
}
