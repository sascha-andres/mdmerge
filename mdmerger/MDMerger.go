// Copyright 2016 Sascha Andres

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
