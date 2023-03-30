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

type SlideShow struct {
	nSlides int
	slides  [][]int
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
		id := i
		indexInFile := i + 1
		photoLine := lines[indexInFile]
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

func buildOutput(result *SlideShow) string {
	nSlides := len(result.slides)
	output := ""
	output += fmt.Sprintf("%d\n", nSlides)
	for _, slide := range result.slides {
		if len(slide) == 1 {
			output += fmt.Sprintf("%d\n", slide[0])
		} else {
			output += fmt.Sprintf("%d %d\n", slide[0], slide[1])
		}
	}

	return output
}
