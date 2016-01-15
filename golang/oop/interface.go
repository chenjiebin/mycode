//非侵入式
//php,c++这些语言接口的实现都是强制性的，即使是在不同的包下也是不同的接口
package main

import (
	"fmt"
)

//播放器
type Player interface {
	Play(source string)
}

func play(source string, mtype string) {
	var p Player
	switch mtype {
	case "MP3":
		p = &MP3Player{}
	case "WAV":
		p = &WAVPlayer{}
	default:
		fmt.Println("Unsupported music type", mtype)
		return
	}
	p.Play(source)
}

//MP3播放器
type MP3Player struct {
	stat int
}

func (this *MP3Player) Play(source string) {
	fmt.Println("mp3 play", source)
}

//WAV播放器
type WAVPlayer struct {
	stat int
}

func (this *WAVPlayer) Play(source string) {
	fmt.Println("wav play", source)
}

func main() {
	source := "music file"
	play(source, "MP3")
	play(source, "WAV")
}
