package main

import (
	"flag"
	"fmt"
	"github.com/sascha-andres/mdmerge/filelister"
	"github.com/sascha-andres/mdmerge/mdmerger"
	"github.com/sascha-andres/mdmerge/toc"
	"log"
	"os"
)

func main() {

	var createtoc, printHeadlines bool

	flag.BoolVar(&createtoc, "createtoc", false, "Create a toc")
	flag.BoolVar(&printHeadlines, "headlines", true, "Print headlines")

	flag.Parse()

	/* get files for directory */
	dir, err := os.Getwd()
	if nil != err {
		log.Fatal("Could not get directory")
	}

	files, err := filelister.Get(dir, "")
	if nil != err {
		log.Fatal(fmt.Sprintf(" %v", err))
	}

	/* create toc */
	var tableOfContents string
	if createtoc {
		tableOfContents = toc.Create(files, "", 1)
	}

	/* build target md file */
	mdmerger.Print(files, tableOfContents, printHeadlines, "", 1)
}
