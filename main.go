package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	new_main()
	return

	files := []string{
		"a",
		"b",
		"c",
		"d",
		"e",
	}

	for _, fileName := range files {
		fmt.Printf("Processing input: %s\n", fileName)
		inputSet := readFile(fmt.Sprintf("./inputFiles/%s.in", fileName))

		config := buildInput(inputSet)
		// fmt.Printf("CONFIG %+v\n", config)
		printInputMetrics(config)

		result := algorithm(config)

		output := buildOutput(result)
		printResultMetrics(config)

		ioutil.WriteFile(fmt.Sprintf("./result/%s.out", fileName), []byte(output), 0644)
		fmt.Printf("Wrote output for: %s\n", fileName)
	}
}
