package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadInput(t *testing.T) {
	t.Run("Read input and take only the first word of each line", func(t *testing.T) {
		result := make([]string, 0)

		handleWordsOnLine := func(lineReady <-chan bool, wordChannel <-chan string, readNextLine chan<- bool, doneChannel chan<- bool) {
			for range lineReady {
				readLine := true
				isFirstWord := true
				for readLine {
					select {
					case word := <-wordChannel:
						if isFirstWord {
							result = append(result, word)
							isFirstWord = false
						}
					default:
						readLine = false
					}
				}
				readNextLine <- true
			}
			doneChannel <- true
		}

		ReadInput("mock/file_mock.txt", 8, 10, handleWordsOnLine)

		require.Len(t, result, 6)
		require.Equal(t, `0-0`, result[0])
		require.Equal(t, `1-0`, result[1])
		require.Equal(t, `2-0`, result[2])
		require.Equal(t, `3-0`, result[3])
		require.Equal(t, `4-0`, result[4])
		require.Equal(t, `5-0`, result[5])
	})
}
