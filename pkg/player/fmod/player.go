package fmod

import (
	"lyrics/pkg/player/cc"
	"lyrics/pkg/player/extensions"
	"time"
)

type Player struct {
	path     string
	music    bool
	length   uint32
	playing  bool
	speed    float64
	volume   float64
	position uint32
	tune     int64
}

func GeneratePlayerFMOD() *Player {
	// 加载启动环境
	cc.LaunchFMOD()
	// 加载支持文件列表
	extensions.LoadMusicSupportList()
	return &Player{
		path:     "",
		music:    false,
		length:   0,
		playing:  false,
		speed:    1,
		volume:   1,
		position: 0,
		tune:     0,
	}
}

func (that *Player) Listen() {
	go func() {
		for that.playing {
			time.Sleep(time.Millisecond * UpdatePositionGapMS)
			that.getPosition()
			that.getPlaying()
		}
	}()
}

func (that *Player) GetPath() string {
	return that.path
}

func (that *Player) Close() {
	that.Pause()
	cc.ExitFMOD()
}
