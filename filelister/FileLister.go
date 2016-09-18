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

package filelister

import (
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"strings"
)

// Get returns all md files in a directory and its sub directories
func Get(basePath, subPath string) ([]MarkDownSegment, error) {
	combinedPath := path.Join(basePath, subPath)

	fileInfos, err := ioutil.ReadDir(combinedPath)
	if nil != err {
		return nil, err
	}

	if len(fileInfos) == 0 {
		return nil, nil
	}
	var result []MarkDownSegment
	for _, f := range fileInfos {
		if f.IsDir() {
			kids, err := Get(combinedPath, f.Name())
			if nil != err {
				return nil, err
			}
			if kids != nil {
				subsegment := createMarkDownSegment(f, combinedPath)
				subsegment.Children = kids
				result = append(result, subsegment)
			}
		} else {
			if strings.HasSuffix(f.Name(), ".md") {
				subsegment := createMarkDownSegment(f, combinedPath)
				result = append(result, subsegment)
			}
		}
	}
	return result, nil

}

func createMarkDownSegment(f os.FileInfo, combinedPath string) MarkDownSegment {
	return MarkDownSegment{
		Name:  calculateName(f.Name()),
		Path:  path.Join(combinedPath, f.Name()),
		IsToc: f.IsDir(),
	}
}

func calculateName(currentName string) string {
	if len(currentName) == 0 {
		return currentName
	}

	result := strings.TrimSuffix(currentName, ".md")

	re := regexp.MustCompile("(?P<prefix>[0-9]*_)(?P<name>.*)")
	result = re.ReplaceAllString(result, "${name}")

	return result
}
