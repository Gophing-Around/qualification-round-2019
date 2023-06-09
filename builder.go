package main

import (
	"fmt"
	"math"
)

type Photo struct {
	ID            int
	Layout        string
	Tags          []string
	NTags         int
	Score         int
	ScoreFloating float64
}

type Config struct {
	nPhotos    int
	photoList  []*Photo
	photosMap  map[int]*Photo
	tagsMap    map[string][]*Photo
	listTags   []*Tags
	tagsScores map[string]int
}

type Tags struct {
	name   string
	photos []*Photo
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

	tagsScoresFloating := make(map[string]float64)
	tagsScores := make(map[string]int)
	tagsMap := make(map[string][]*Photo)
	listTags := make([]*Tags, 0)
	for _, photo := range config.photoList {
		for _, tag := range photo.Tags {
			tagScore, _ := tagsScores[tag]
			tagScore++
			tagsScores[tag] = tagScore

			tagScoreFloating, _ := tagsScoresFloating[tag]
			tagScoreFloating += (float64(1) / float64(len(config.photoList)))
			tagsScoresFloating[tag] = tagScoreFloating

			if tagsMap[tag] == nil {
				tagsMap[tag] = []*Photo{}
			}
			tagsMap[tag] = append(tagsMap[tag], photo)
		}
	}

	for tag, photo := range tagsMap {
		listTags = append(listTags, &Tags{
			name:   tag,
			photos: photo,
		})
	}

	config.listTags = listTags
	config.tagsMap = tagsMap
	config.tagsScores = tagsScores

	for _, photo := range config.photoList {
		photoThreshold := 0.0
		photoScore := 0
		for _, tag := range photo.Tags {
			tagScore := tagsScores[tag]
			photoScore += tagScore

			tagScoreFloating := tagsScoresFloating[tag]
			if tagScoreFloating > 0.5 {
				photoThreshold++
			} else {
				photoThreshold--
			}

			photoScore += tagScore
		}

		photo.ScoreFloating = 100 - math.Abs(float64(photoThreshold))
		photo.Score = (photoScore / photo.NTags)
		// fmt.Printf("PHOTO#%d - score %d tags: %d\n", id, photo.Score, photo.NTags)
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
