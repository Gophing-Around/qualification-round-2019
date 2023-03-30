package main

import "sort"

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
		oldOldSelectedTags := []*Tags{}
		for _, listTag := range config.listTags {
			if len(listTag.photos) != 0 {
				oldOldSelectedTags = append(config.listTags, listTag)
			}
		}
		config.listTags = oldOldSelectedTags
		if len(config.listTags) == 0 {
			break
		}
		for _, photo := range selectedTag.photos {
			if usedPhotos[photo.ID] {
				continue
			}
			usedPhotos[photo.ID] = true
			maxScoredTag := ""
			for _, photoTag := range photo.Tags {
				if photoTag == selectedTag.name {
					continue
				}
				if maxScoredTag == "" || config.tagsScores[photoTag] > config.tagsScores[maxScoredTag] {
					selectedTag = &Tags{
						name:   photoTag,
						photos: config.tagsMap[photoTag],
					}
				}
			}
			tags := []string{}
			for _, photoTag := range photo.Tags {
				if photoTag != oldSelectedTags.name {
					tags = append(tags, photoTag)
				}
			}
			photo.Tags = tags

			photos := []*Photo{}
			for _, selectedPhoto := range oldSelectedTags.photos {
				if photo.ID != selectedPhoto.ID {
					photos = append(photos, photo)
				}
			}
			oldSelectedTags.photos = photos

			oldSelectedTags = selectedTag

			if photo.Layout == "H" {
				slideShow.slides = append(slideShow.slides, []int{photo.ID})
				slideIndex++
			} else {
				if partialSlideIndex < 0 {
					newslide := []int{photo.ID}
					slideShow.slides = append(slideShow.slides, newslide)
					partialSlideIndex = slideIndex
					slideIndex++
					continue
				}

				// fmt.Printf("FOOOO %v %+v %d", partialSlideIndex, slideShow.slides, slideIndex)
				slideShow.slides[partialSlideIndex] = append(slideShow.slides[partialSlideIndex], photo.ID)
				partialSlideIndex = -1
			}
		}
	}

	return &slideShow
}
