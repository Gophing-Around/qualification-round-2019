package main

import (
	"fmt"
)

func new_main() {
	files := []string{
		"a",
		"b",
		"c",
		"d",
		"e",
	}

	for _, fileName := range files {
		fmt.Printf("Processing input: %s\n", fileName)
		config := buildNewInput(fmt.Sprintf("./inputFiles/%s.in", fileName))
		fmt.Printf("CONFIG %+v\n", config)
		for _, photo := range config.photoList {
			fmt.Printf("PHOTO %+v\n", *photo)
		}
		break
	}
}

func buildNewInput(file string) *Config {
	config := Config{}
	ReadInput(file, 8000, 1000, func(lineReady <-chan bool, wordChannel <-chan string, readNextLine chan<- bool, doneChannel chan<- bool) {
		<-lineReady
		config.nPhotos = toint(<-wordChannel)
		readNextLine <- true

		config.photoList = make([]*Photo, 0)
		for id := 0; id < config.nPhotos; id++ {
			<-lineReady
			photo := Photo{ID: id}
			photo.Layout = <-wordChannel
			photo.NTags = toint(<-wordChannel)
			photo.Tags = make([]string, 0)
			for j := 0; j < photo.NTags; j++ {
				photo.Tags = append(photo.Tags, <-wordChannel)
			}
			config.photoList = append(config.photoList, &photo)
			readNextLine <- true
		}

		doneChannel <- true
	})
	return &config
}
