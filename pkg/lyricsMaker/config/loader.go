package config

import (
	"log"
	"os"
)

func LoadConfig() {
	log.Println("Loading font file.")
	//data := make([]byte, 1024)
	//file, err := os.Open(FontFile)
	//if nil != err {
	//	log.Println(err)
	//}
	//n, err := file.Read(data)
	//if nil != err {
	//	log.Println(err)
	//}
	//log.Println(n, data)
	//file.Close()
	err := os.Setenv(FontEnv, FontFile)
	if nil != err {
		log.Println(err)
	}
	log.Println("Loading extension config.")
	LoadExtension()
}

func Close() {
	err := os.Unsetenv(FontEnv)
	if err != nil {
		log.Println(err)
	}
}
