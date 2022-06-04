package config

import (
	"log"
	"os"
)

func LoadConfig() {
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
		log.Printf("Load config font file %s as %s failed: %s.\n", FontFile, FontEnv, err)
	}
	LoadExtension()
}

func Close() {
	err := os.Unsetenv(FontEnv)
	if err != nil {
		log.Println("Unmount config failed: ", err)
	}
}
