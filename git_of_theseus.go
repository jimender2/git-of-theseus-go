package main

import (
	"fmt"

	"github.com/go-git/go-git/v5"
)

var ignoreFileTypes [14]string = [14]string{
	"*.json",
	"*.md",
	"*.ps",
	"*.eps",
	"*.txt",
	"*.xml",
	"*.xsl",
	"*.rss",
	"*.xslt",
	"*.xsd",
	"*.wsdl",
	"*.wsf",
	"*.yaml",
	"*.yml",
}

func main() {
	fmt.Println("Hello, world.")

	// loop through everything in ignoreFileTypes
	for _, ignoreFileType := range ignoreFileTypes {
		fmt.Println(ignoreFileType)
	}

	_, err := git.PlainOpen("")

	if err != nil {
		fmt.Println("test")
	} else {
		fmt.Println("2")
	}

}
