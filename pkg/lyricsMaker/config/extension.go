package config

import "lyrics/pkg/player/extensions"

var MusicExtension []string
var VideoExtension []string

func LoadExtension() {
	MusicExtension = extensions.MusicSupportList
	VideoExtension = extensions.VideoSupportList
}
