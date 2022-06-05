package fmod

func (that *Player) UpdateFMOD() {
	that.SetVolume(that.GetVolume())
	that.SetFrequency(that.GetFrequency())
}
