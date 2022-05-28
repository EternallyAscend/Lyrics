package config

var MusicExtension []string
var VideoExtension []string

func LoadExtension() {
	MusicExtension = []string{
		".mp3", ".wav",
	}
	VideoExtension = []string{}
}
