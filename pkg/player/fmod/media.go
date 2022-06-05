package fmod

import (
	"errors"
	"lyrics/pkg/player/cc"
	"lyrics/pkg/player/extensions"
	"lyrics/pkg/player/text"
	"strings"
)

func (that *Player) LoadMedia(path string) error {
	if that.Playable() {
		that.Stop()
	}
	if "" != path {
		dot := strings.LastIndex(path, ".")
		if -1 == dot {
			return errors.New(text.ErrorWrongPath)
		}
		extension := path[dot:]
		for index := range extensions.MusicSupportList {
			if extensions.MusicSupportList[index] == extension {
				that.path = path
				that.music = true
				that.position = 0
				cc.SetMediaFMOD(that.path)
				that.getLength()
				return nil
			}
		}
		return errors.New(text.ErrorNotSupport)
	} else {
		return errors.New(text.ErrorNoMedia)
	}
}

func (that *Player) Playable() bool {
	return "" != that.path
}

func (that *Player) Play() {
	if that.Playable() {
		if !that.playing {
			cc.PlayFMOD()
			go that.UpdateFMOD()
			that.playing = true
			that.Listen()
		} else {
			that.Pause()
		}
	}
}

func (that *Player) Pause() {
	cc.PauseFMOD()
	that.playing = false
}

func (that *Player) Stop() {
	that.SetPosition(0)
	that.Pause()
}

func (that *Player) getPlaying() {
	that.playing = cc.GetPlayingFMOD()
	//if !cc.GetPlayingFMOD() {
	//	if that.playing {
	//		that.playing = false
	//		go func() {
	//			time.Sleep(time.Millisecond * UpdateConfigFMOD)
	//			that.UpdateFMOD()
	//		}()
	//		return
	//	}
	//}
	//that.playing = true
}

func (that *Player) GetPlaying() bool {
	return that.playing
}

func (that *Player) getLength() {
	if that.Playable() {
		that.length = cc.GetLengthFMOD()
	}
}

func (that *Player) GetLength() uint32 {
	return that.length
}
