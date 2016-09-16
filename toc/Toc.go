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
			listOfStrings = append(listOfStrings, fmt.Sprintf("%s%v. %s", parentIndex, index, segment.Name))
			listOfStrings = append(listOfStrings, Create(segment.Children, fmt.Sprintf("%s%v.", parentIndex, index), level+1))
			index++
		}
	}
	return strings.Join(listOfStrings, "\n")
}

// func leftPad(s string, padStr string, pLen int) string {
// 	return strings.Repeat(padStr, pLen) + s
// }
