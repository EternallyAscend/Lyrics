package fmod

import "lyrics/pkg/player/cc"

func (that *Player) GetPosition() uint32 {
	that.position = cc.GetPositionFMOD()
	return that.position
}

func (that *Player) SetPosition(ms uint32) {
	if that.Playable() {
		that.position = ms
		cc.SetPositionFMOD(ms)
	}
}
