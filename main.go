package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/mikkyang/id3-go"
)

func main() {
	for _, arg := range os.Args[1:] {
		mp3, err := id3.Open(arg)
		if err != nil {
			log.Println(err)
			continue
		}
		filename := filepath.Base(arg)
		name := filename[:len(filename)-len(filepath.Ext(filename))]
		println(name)
		mp3.SetArtist(name)
		mp3.SetTitle(name)
		mp3.SetAlbum(name)
		mp3.Close()
	}
}
