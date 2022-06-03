package extensions

var MusicSupportList []string

func LoadMusicSupportList() {
	MusicSupportList = []string{
		".wav", ".mp3",
	}
}
