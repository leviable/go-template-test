package main

import (
	"fmt"

	"github.com/leviable/go-template-test/src"
)

// title contains the name of the project
const title = "go-template-test"

/*
ProjectName returns the value of `title` string
*/
func ProjectName() string {
	return title
}

func main() {
	fmt.Printf("Running project: %s\n", src.ProjectName())
}
