package main

import (
	"fmt"
	"sort"
)

func algorithm(
	config *Config,
) *SlideShow {

	slideShow := SlideShow{
		slides: make([][]int, 0),
	}

	sort.Slice(config.listTags, func(i, j int) bool {
		tagA := config.listTags[i]
		tagB := config.listTags[j]

		return len(tagA.photos) > len(tagB.photos)
	})

	slideIndex := 0
	partialSlideIndex := -1
	usedPhotos := map[int]bool{}

	selectedTag := config.listTags[0]
	oldSelectedTags := selectedTag
	for {
		if len(config.listTags) == 0 {
			break
		}

		selectedPhoto := selectedTag.photos[0]
		fmt.Printf("SELECTED PHOTO %s - %d\n", selectedTag.name, selectedPhoto.ID)
		if len(selectedTag.photos) > 1 {
			selectedTag.photos = selectedTag.photos[1:]
		} else if len(selectedTag.photos) == 1 {
			selectedTag.photos = []*Photo{selectedTag.photos[0]}
		} else {
			newConfigListLtags := []*Tags{}
			for _, listTag := range config.listTags {
				if len(listTag.photos) != 0 {
					newConfigListLtags = append(config.listTags, listTag)
				}
			}
			config.listTags = newConfigListLtags
			selectedTag = config.listTags[0]
		}

		// for _, photo := range selectedTag.photos {
		if usedPhotos[selectedPhoto.ID] {
			continue
		}
		usedPhotos[selectedPhoto.ID] = true

		maxScoredTag := ""
		for _, photoTag := range selectedPhoto.Tags {
			if photoTag == selectedTag.name {
				continue
			}
			fmt.Printf("SHOULD SELECTED NEW ONE", photoTag)

			if maxScoredTag == "" || config.tagsScores[photoTag] > config.tagsScores[maxScoredTag] {
				fmt.Printf("NEW TAG SELECTED", photoTag)
				selectedTag = &Tags{
					name:   photoTag,
					photos: config.tagsMap[photoTag],
				}
			}
		}
		tags := []string{}
		for _, photoTag := range selectedPhoto.Tags {
			if photoTag != oldSelectedTags.name {
				tags = append(tags, photoTag)
			}
		}
		selectedPhoto.Tags = tags

		photos := []*Photo{}
		for _, selectedPhoto := range oldSelectedTags.photos {
			if selectedPhoto.ID != selectedPhoto.ID {
				photos = append(photos, selectedPhoto)
			}
		}
		oldSelectedTags.photos = photos

		oldSelectedTags = selectedTag

		if selectedPhoto.Layout == "H" {
			slideShow.slides = append(slideShow.slides, []int{selectedPhoto.ID})
			slideIndex++
		} else {
			if partialSlideIndex < 0 {
				newslide := []int{selectedPhoto.ID}
				slideShow.slides = append(slideShow.slides, newslide)
				partialSlideIndex = slideIndex
				slideIndex++
				continue
			}

			// fmt.Printf("FOOOO %v %+v %d", partialSlideIndex, slideShow.slides, slideIndex)
			slideShow.slides[partialSlideIndex] = append(slideShow.slides[partialSlideIndex], selectedPhoto.ID)
			partialSlideIndex = -1
		}
		// }
	}

	return &slideShow
}
