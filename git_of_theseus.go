package main

import (
	"flag"
	"fmt"
	"os"

	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
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

type tuple struct {
	name, value interface{}
}

type void struct{}

var member void

func main() {
	var gitRepoPath string
	var gitRepoBranch string
	flag.StringVar(&gitRepoPath, "repo", ".", "Git repository path")
	flag.StringVar(&gitRepoBranch, "branch", "master", "Git branch name")
	flag.Parse()

	fmt.Println(gitRepoPath)

	gitRepo, err := git.PlainOpen(gitRepoPath)

	fmt.Println(gitRepo)

	if err != nil {
		fmt.Println("Unable to open the git repo. Do you have permission to view it?")
		os.Exit(0)
	}

	fmt.Println("Opening git repository branch")
	fmt.Println("Implementing later")
	// fmt.Println(gitRepoBranch)
	// gitRepo.Branch(gitRepoBranch)
	// if err != nil {
	// 	fmt.Println("Unable to find the branch specified. Falling back to the default branch")
	// }
	// vh, err := gitRepo.Head()
	// fmt.Println(vh)
	// fmt.Println(err)

	fmt.Println("Looping through all commits.")
	ref, err := gitRepo.Head()

	cIter, err := gitRepo.Log(&git.LogOptions{From: ref.Hash()})

	var count = 0

	set := make(map[tuple]void)
	// ... just iterates over the commits, printing it
	err = cIter.ForEach(func(c *object.Commit) error {
		// fmt.Println(c)
		// fmt.Println(c.Author)
		count = count + 1
		commitdate := c.Committer.When
		cohorttup := tuple{"cohort", commitdate}
		set[cohorttup] = struct{}{}

		authorName := c.Author.Name
		authortup := tuple{"author", authorName}
		set[authortup] = struct{}{}

		emailSplit := strings.Split(c.Author.Email, "@")
		domainName := emailSplit[len(emailSplit)-1]
		domaintup := tuple{"domain", domainName}
		set[domaintup] = struct{}{}

		return nil
	})

	fmt.Println(count)

	fmt.Println("end")
}
