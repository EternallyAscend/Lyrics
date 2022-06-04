package fmod

import (
	"lyrics/pkg/player/cc"
	"lyrics/pkg/player/extensions"
)

type Player struct {
	path     string
	music    bool
	length   int64
	playing  bool
	speed    float64
	volume   float64
	position int64
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

func (that *Player) Close() {
	cc.ExitFMOD()
}
