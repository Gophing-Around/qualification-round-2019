package main

import (
	"sort"
)

func algorithm(
	config *Config,
) *SlideShow {

	photoList := config.photoList

	sort.Slice(photoList, func(a, b int) bool {
		photoA := photoList[a]
		photoB := photoList[b]

		if photoA.Layout == "H" && photoB.Layout == "H" {
			return true
		}

		if photoA.Layout == "H" && photoB.Layout == "V" {
			return true
		}
		return false
	})

	slideShow := SlideShow{
		slides: make([][]int, 0),
	}

	slideIndex := 0

	for i := 0; i < len(photoList); i++ {
		// 	fmt.Printf(photoList[i].Layout)
		photo := photoList[i]
		if photo.Layout == "H" {
			slideShow.slides = append(slideShow.slides, []int{photo.ID})
			slideIndex++
		} else {
			var photoB *Photo
			if i < len(photoList) {
				photoB = photoList[i+1]
			}

			newslide := []int{photo.ID}
			if photoB != nil {
				newslide = append(newslide, photoB.ID)
			}
			slideShow.slides = append(slideShow.slides, newslide)
			i++
			// slide := slideShow.slides[slideIndex]
			// if len(slide) >= 2 {
			// 	slideIndex++
			// }

			// if len(slide) == 0 {
			// 	slide = make([]int,0)
			// }
			// slide = append(slide, photo.ID)
		}
	}

	return &slideShow
}
