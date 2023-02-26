package main

import (
	"flag"
	"fmt"
	"os"

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
	var gitRepoPath string
	flag.StringVar(&gitRepoPath, "repo", ".", "Git repository path")
	flag.Parse()

	fmt.Println(gitRepoPath)

	gitRepo, err := git.PlainOpen(gitRepoPath)

	fmt.Println(gitRepo)

	if err != nil {
		fmt.Println("Unable to open the git repo. Do you have permission to view it?")
		os.Exit(0)
	}

	fmt.Println("testing git repository")
}
