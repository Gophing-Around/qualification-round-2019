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
		photoA := photoList[i]
		photoB := photoList[j]

		return photoA.Score < photoB.Score
	})

	slideShow := SlideShow{
		slides: make([][]int, 0),
	}

	slideIndex := 0
	partialSlideIndex := -1

	for _, photo := range photoList {
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

	sort.Slice(slideShow.slides, func(i, j int) bool {
		slideItemA := slideShow.slides[i]
		slideItemB := slideShow.slides[j]

		tagA0 := config.photosMap[slideItemA[0]].Tags
		tagB0 := config.photosMap[slideItemB[0]].Tags

		tagA1 := make([]string, 0)
		if len(slideItemA) > 1 {
			tagA1 = config.photosMap[slideItemA[1]].Tags
		}
		tagB1 := make([]string, 0)
		if len(slideItemB) > 1 {
			tagB1 = config.photosMap[slideItemB[1]].Tags
		}

		tagsSlideA := lo.Uniq(lo.Union(tagA0, tagA1))
		tagsSlideB := lo.Uniq(lo.Union(tagB0, tagB1))

		return len(tagsSlideA) < len(tagsSlideB)
	})

	// for i := 0; i < len(photoList); i++ {
	// 	// 	fmt.Printf(photoList[i].Layout)
	// 	photo := photoList[i]
	// 	if photo.Layout == "H" {
	// 		slideShow.slides = append(slideShow.slides, []int{photo.ID})
	// 		slideIndex++
	// 	} else {
	// 		if partialSlideIndex < 0 {
	// 			newslide := []int{photo.ID}
	// 			slideShow.slides = append(slideShow.slides, newslide)
	// 			partialSlideIndex = slideIndex
	// 			slideIndex++
	// 			continue
	// 		}

	// 		// fmt.Printf("FOOOO %v %+v %d", partialSlideIndex, slideShow.slides, slideIndex)
	// 		slideShow.slides[partialSlideIndex] = append(slideShow.slides[partialSlideIndex], photo.ID)
	// 		partialSlideIndex = -1
	// 	}
	// }

	return &slideShow
}
