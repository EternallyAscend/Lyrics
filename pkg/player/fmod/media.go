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
				// TODO Media Length
				that.position = 0
				cc.SetMediaFMOD(that.path)
				return nil
			}
		}
		for index := range extensions.VideoSupportList {
			if extensions.VideoSupportList[index] == extension {
				that.music = false
				// TODO Media Length
				that.position = 0
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
		cc.PlayFMOD()
	}
}

func (that *Player) Pause() {

}

func (that *Player) Stop() {
	if that.music {

	} else {

	}
}

func (that *Player) Jump() {

}
