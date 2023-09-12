package biz

import (
	"encoding/json"

	"github.com/jacktrane/gocomponent/file_util"
	"github.com/jacktrane/gocomponent/logger"
)

type FileDiff struct {
	filePath1  string
	filePath2  string
	resultPath string
}

func NewFileDiff(filePath1 string, filePath2 string, resultPath string) *FileDiff {
	return &FileDiff{
		filePath1:  filePath1,
		filePath2:  filePath2,
		resultPath: resultPath,
	}
}

func (fileDiff *FileDiff) Diff() (string, error) {
	m1, err := getFileMap(fileDiff.filePath1)
	if err != nil {
		logger.Error(err)
		return "", err
	}

	m2, err := getFileMap(fileDiff.filePath2)
	if err != nil {
		logger.Error(err)
		return "", err
	}

	return "diff"
}

func (fileDiff *FileDiff) GenResult() {
	logger.Info("gen result")
}

func getFileMap(filePath string) (map[string]any, error) {
	m := make(map[string]any)
	content, err := file_util.GetFileContent(filePath)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	if err := json.Unmarshal(content, &m); err != nil {
		logger.Error(err)
		return nil, err
	}
	return m, nil
}

// increase，以A的角度查看比B多的部分
// decrease，以A的角度查看比B少的部分
func diffContent(a, b map[string]any) (increase, decrease map[string]any) {
	increase = make(map[string]any)
	decrease = make(map[string]any)
	for k, v := range a {
		if _, ok := b[k]; !ok {
			increase[k] = v
		}
	}
	for k, v := range b {
		if _, ok := a[k]; !ok {
			decrease[k] = v
		}
	}
	return
}
