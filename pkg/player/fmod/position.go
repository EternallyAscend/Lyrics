package fmod

import "lyrics/pkg/player/cc"

// GetPosition 获取音频位置，单位毫秒，外部调用，更新外部值使用
func (that *Player) GetPosition() uint32 {
	return that.position
}

// getPosition 获取音频位置，单位毫秒，内部使用，更新内部值使用
func (that *Player) getPosition() {
	that.position = cc.GetPositionFMOD()
}

// SetPosition 设置位置，单位毫秒
func (that *Player) SetPosition(ms uint32) {
	if that.Playable() {
		that.position = ms
		cc.SetPositionFMOD(ms)
	}
}
