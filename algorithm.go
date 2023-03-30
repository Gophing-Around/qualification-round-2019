package main

import (
	"sort"

	"github.com/samber/lo"
)

func algorithm(
	config *Config,
) *SlideShow {

	photoList := config.photoList

	// sort.Slice(photoList, func(a, b int) bool {
	// 	photoA := photoList[a]
	// 	photoB := photoList[b]

	// 	if photoA.Layout == "H" && photoB.Layout == "H" {
	// 		return true
	// 	}

	// 	if photoA.Layout == "H" && photoB.Layout == "V" {
	// 		return true
	// 	}
	// 	return false
	// })

	sort.Slice(photoList, func(i, j int) bool {
		tagsPhotoA := photoList[i].Tags
		tagsPhotoB := photoList[j].Tags

		tagIntersection := lo.Intersect(tagsPhotoA, tagsPhotoB)
		tagDiffA, tagDiffB := lo.Difference(tagsPhotoA, tagsPhotoB)

		ntagIntersection := len(tagIntersection)
		ntagDiffA := len(tagDiffA)
		ntagDiffB := len(tagDiffB)
		return photoA < photoB
	})

	slideShow := SlideShow{
		slides: make([][]int, 0),
	}

	slideIndex := 0
	partialSlideIndex := -1

	for i := 0; i < len(photoList); i++ {
		// 	fmt.Printf(photoList[i].Layout)
		photo := photoList[i]
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

	return &slideShow
}
