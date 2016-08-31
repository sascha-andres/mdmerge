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
