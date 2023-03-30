package main

import "fmt"

type Photo struct {
	ID     int
	Layout string
	Tags   []string
	NTags  int
}

type Config struct {
	nPhotos   int
	photoList []*Photo
	photosMap map[int]*Photo
}

func buildInput(inputSet string) *Config {
	lines := splitNewLines(inputSet)
	configLine := splitSpaces(lines[0])
	fmt.Printf("Config line: %v\n", configLine)

	nPhotos := toint(configLine[0])
	config := &Config{
		nPhotos:   nPhotos,
		photosMap: make(map[int]*Photo),
		photoList: make([]*Photo, nPhotos),
	}

	for i := 0; i < nPhotos; i++ {
		id := i + 1
		photoLine := lines[id]
		fmt.Printf(">>>> %s\n", photoLine)
		photoConfig := splitSpaces(photoLine)

		ntags := toint(photoConfig[1])
		tags := photoConfig[2:]

		photo := Photo{
			ID:     id,
			Layout: photoConfig[0],
			Tags:   tags,
			NTags:  ntags,
		}
		config.photoList[i] = &photo
		config.photosMap[id] = &photo
	}

	return config
}

func buildOutput(result int) string {
	return "42"
}
