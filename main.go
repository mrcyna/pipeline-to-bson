package main

import (
	"fmt"
	"github.com/mrcyna/pipeline-to-bson/pipeline"
	"io/ioutil"
)

func main() {
	data, err := ioutil.ReadFile("example.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	fmt.Println(pipeline.Validate(string(data)))
}
