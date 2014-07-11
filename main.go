package main

import (
	"code.google.com/p/test2doc/t2d"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please specify the code file or the root directory of your code.")
		return
	}

	path := os.Args[1]
	docCreator := t2d.NewDocCreator(
		t2d.NewFileWalker(),
		t2d.NewTestFileFilter(),
		t2d.NewTestNameExtractor(),
		t2d.NewCamelCaseChopper(),
		t2d.NewListDocFormatter(path))
	doc, err := docCreator.CreateDocFromFolder(path)
	fmt.Println(doc)
	if err != nil {
		fmt.Println(err)
	}
}
