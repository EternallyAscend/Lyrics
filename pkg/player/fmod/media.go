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
				that.GetLength()
				return nil
			}
		}
		for index := range extensions.VideoSupportList {
			if extensions.VideoSupportList[index] == extension {
				that.music = false
				that.position = 0
				that.GetLength()
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
		go cc.PlayFMOD()
		that.playing = true
		that.Listen()
	}
}

func (that *Player) Pause() {
	go cc.PauseFMOD()
	that.playing = false
}

func (that *Player) Stop() {
	that.SetPosition(0)
	that.Pause()
}

func (that *Player) GetPlaying() bool {
	that.playing = cc.GetPlayingFMOD()
	return that.playing
}

func (that *Player) GetLength() uint32 {
	if that.Playable() {
		that.length = cc.GetLengthFMOD()
		return that.length
	}
	return 0
}
