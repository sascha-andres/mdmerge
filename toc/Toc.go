package toc

import (
	"fmt"
	"github.com/sascha-andres/mdmerge/filelister"
	"strings"
)

func Create(items []filelister.MarkDownSegment, parentIndex string, level int) string {
	index := 1
	listOfStrings := make([]string, 0)
	for _, segment := range items {
		if segment.IsToc {
			listOfStrings = append(listOfStrings, leftPad(fmt.Sprintf("%s%v. %s", parentIndex, index, segment.Name), "#", level))
			listOfStrings = append(listOfStrings, Create(segment.Children, fmt.Sprintf("%s%v.", parentIndex, index), level+1))
			index++
		}
	}
	return strings.Join(listOfStrings, "\n")
}

func leftPad(s string, padStr string, pLen int) string {
	return strings.Repeat(padStr, pLen) + s
}
