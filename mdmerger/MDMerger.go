package mdmerger

import (
	"fmt"
	"github.com/sascha-andres/mdmerge/filelister"
	"io/ioutil"
	"log"
	"strings"
)

func Print(items []filelister.MarkDownSegment, toc string, printHeadlines bool, parentIndex string, level int) {
	for _, segment := range items {
		if !segment.IsToc {
			content, err := ioutil.ReadFile(segment.Path)
			if err != nil {
				log.Fatal(err.Error())
			}
			fmt.Println(string(content))
			fmt.Println("")
		}
	}
	if 0 < len(toc) {
		fmt.Println(toc)
	}
	for _, segment := range items {
		index := 1
		if segment.IsToc {
			if printHeadlines {
				fmt.Println(leftPad(fmt.Sprintf(" %s%v. %s", parentIndex, index, segment.Name), "#", level))
			}
			Print(segment.Children, "", printHeadlines, fmt.Sprintf("%s%v.", parentIndex, index), level+1)
			index++
		}
	}
}

func leftPad(s string, padStr string, pLen int) string {
	return strings.Repeat(padStr, pLen) + s
}
