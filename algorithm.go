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

	minThreshold := 10

	maxIterations := len(photoList)
	currentIterations := 0
	for len(photoList) > 0 {
		currentIterations++
		photo := photoList[0]
		photoList = photoList[1:]

		if slideIndex != 0 {
			previousPhotoSlide := slideShow.slides[slideIndex-1]
			slidePhotoAIndex := previousPhotoSlide[0]
			slidePhotoA := config.photosMap[slidePhotoAIndex]
			slidePhotoBIndex := -1
			var slidePhotoB *Photo
			tagsB := make([]string, 0)
			if len(previousPhotoSlide) > 1 {
				slidePhotoBIndex = previousPhotoSlide[1]
				slidePhotoB = config.photosMap[slidePhotoBIndex]
				tagsB = slidePhotoB.Tags
			}

			previousSlideTags := lo.Uniq(lo.Union(slidePhotoA.Tags, tagsB))
			currentPhotoTags := photo.Tags

			tagsIntersect := lo.Intersect(currentPhotoTags, previousSlideTags)

			commonTags := len(tagsIntersect)
			differenceA := len(previousSlideTags) - commonTags
			differenceB := len(currentPhotoTags) - commonTags

			min := commonTags
			if differenceA < commonTags {
				min = differenceA
			}
			if differenceB < differenceA {
				min = differenceB
			}

			if min < minThreshold && currentIterations < maxIterations {
				// minThreshold--
				photoList = append(photoList, photo)
				continue
			}
		}

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
