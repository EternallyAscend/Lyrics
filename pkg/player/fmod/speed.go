package fmod

import (
	"lyrics/pkg/player/cc"
	"math"
)

func (that *Player) AdjustPitchForFrequency() {
	pitch := 1.0
	if that.speed < 0.5 {
		pitch = 1 + math.E/that.speed
	} else if that.speed >= 0.5 && that.speed <= 2 {
		pitch = 1 / that.speed
	} else if that.speed > 2 && that.speed <= 4 {
		pitch = -that.speed * math.Exp(that.speed)
	}
	cc.SetPitchFMOD(float32(pitch))
}

func (that *Player) SetFrequency(frequency float64) {
	cc.SetFrequencyFMOD(float32(frequency))
	that.speed = frequency
	that.AdjustPitchForFrequency()
}

func (that *Player) GetFrequency() float64 {
	return that.speed
}
