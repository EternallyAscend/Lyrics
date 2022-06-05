package fmod

import "lyrics/pkg/player/cc"

func (that *Player) SetVolume(volume float32) {
	cc.SetVolumeFMOD(volume)
	that.volume = float64(volume)
}

func (that *Player) GetVolume() float32 {
	return float32(that.volume)
}
