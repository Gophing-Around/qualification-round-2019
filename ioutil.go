package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

type ReadInputHandler func(lineReady <-chan bool, wordChannel <-chan string, readNextLine chan<- bool, doneChannel chan<- bool)

func ReadInput(fileName string, bufferSize int, maxWordsOnLine int, handler ReadInputHandler) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	wordChannel := make(chan string, maxWordsOnLine)
	lineReady := make(chan bool, 1)
	readNextLine := make(chan bool, 1)
	doneChannel := make(chan bool, 1)

	go handler(lineReady, wordChannel, readNextLine, doneChannel)

	reader := bufio.NewReader(file)
	buffer := make([]byte, bufferSize)

	var currentWordBuilder strings.Builder
	var seenSpace bool

	for {
		bytesRead, err := reader.Read(buffer)

		if err != nil {
			if err == io.EOF {
				break
			}

			log.Fatalln(err)
		}

		chunk := buffer[:bytesRead]

		for _, b := range chunk {
			switch b {
			case '\r':
				continue
			case ' ':
				seenSpace = true
			case '\n':
				seenSpace = false
				wordChannel <- currentWordBuilder.String()
				currentWordBuilder.Reset()
				lineReady <- true
				<-readNextLine
			default:
				if seenSpace {
					seenSpace = false
					newString := currentWordBuilder.String()
					if newString != "" {
						wordChannel <- newString
						currentWordBuilder.Reset()
					}
				}

				currentWordBuilder.WriteByte(b)
			}
		}
	}

	if currentWordBuilder.Len() > 0 {
		wordChannel <- currentWordBuilder.String()
		lineReady <- true
		<-readNextLine
	}

	close(wordChannel)
	close(readNextLine)
	close(lineReady)

	<-doneChannel

	close(doneChannel)
}
