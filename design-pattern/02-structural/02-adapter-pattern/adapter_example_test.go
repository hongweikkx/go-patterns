package adapter

/*
	实现用视频播放器播放音乐
*/

import (
	"fmt"
	"testing"
)

//音乐播放器
type Player interface {
	PlayMusic()
}

type MusicPlayer struct {
	Src string
}

func (music *MusicPlayer) PlayMusic() {
	fmt.Println("play music: " + music.Src)
}

//对外接口
func Play(player Player) {
	player.PlayMusic()
}

//在音乐播放基础上实现游戏播放
//定义游戏结构体
type VideoPlayer struct {
	Src string
}

//video的方法
func (video *VideoPlayer) PlaySound() {
	fmt.Println("play sound: " + video.Src)
}

//这里要实现调用play方法的时候，实现VideoPlayer的播放
//通过组合的方式，声明一个Video的适配器
type VideoPlayerAdapter struct {
	Video VideoPlayer
}

//继承Player interface, 调用VideoPlayer的方法
func (video *VideoPlayerAdapter) PlayMusic() {
	video.Video.PlaySound()
}

func TestPlay(t *testing.T) {
	music := &MusicPlayer{Src: "music.mp3"}
	Play(music)

	//使用Play来播放VideoPlayer
	video := &VideoPlayer{Src: "video.mp4"}
	adapter := &VideoPlayerAdapter{Video: *video}
	Play(adapter)
}
