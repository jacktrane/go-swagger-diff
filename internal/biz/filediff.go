package biz

import (
	"encoding/json"
	"reflect"

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

func (fileDiff *FileDiff) Diff() (i, d, c map[string]any, err error) {
	// m1, err := getFileContent(fileDiff.filePath1)
	// if err != nil {
	// 	logger.Error(err)
	// 	return i, d, c, err
	// }

	// m2, err := getFileContent(fileDiff.filePath2)
	// if err != nil {
	// 	logger.Error(err)
	// 	return i, d, c, err
	// }
	// increase, decrease, change := diffContent(m1, m2)
	// i, d, c = make(map[string]any), make(map[string]any), make(map[string]any)
	// for k, v := range increase {
	// 	keys := strings.Split(k, placeHolder)
	// 	if len(keys) <= 1 {
	// 		i[keys[0]] = v
	// 		continue
	// 	}
	// 	temp := make(map[string]any)
	// 	for i := len(keys) - 1; i >= 0; i-- {
	// 		if i == len(keys)-1 {
	// 			temp[keys[i]] = v
	// 		} else {
	// 			t := make(map[string]any)
	// 			t[keys[i]] = temp
	// 			temp = t
	// 		}
	// 	}

	// 	// for {
	// 	// 	if v, ok := i[keys[0]]
	// 	// }
	// 	i[keys[0]] = temp[keys[0]]
	// 	logger.Info(temp[keys[0]])
	// }

	// logger.Info(i)
	// logger.Info(increase)
	// logger.Info(decrease)
	// logger.Info(change)

	m1, err := getFileMap(fileDiff.filePath1)
	if err != nil {
		logger.Error(err)
		return i, d, c, err
	}

	m2, err := getFileMap(fileDiff.filePath2)
	if err != nil {
		logger.Error(err)
		return i, d, c, err
	}
	// logger.Info(diffContent(m1, m2))
	i, d, c = diffContent(m1, m2)
	return i, d, c, nil
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

// var placeHolder = "##%%$$placeHolder$$%%##"

// func encodeMap(src map[string]any) map[string]string {
// 	res := make(map[string]string)
// 	for k, v := range src {
// 		switch vv := v.(type) {
// 		case map[string]any:
// 			tempRes := encodeMap(vv)
// 			for k1, v1 := range tempRes {
// 				res[k+placeHolder+k1] = v1
// 			}
// 		default:
// 			res[k] = fmt.Sprintf("%v", v)
// 		}
// 	}
// 	return res
// }

// func decodeMap(src map[string]string) map[string]any {
// 	dst := make(map[string]any)
// 	for k, v := range src {
// 		keys := strings.Split(k, placeHolder)
// 		if len(keys) <= 1 {
// 			dst[keys[0]] = v
// 			continue
// 		}
// 		temp := make(map[string]any)
// 		for i := len(keys) - 1; i >= 0; i-- {
// 			if i == len(keys)-1 {
// 				temp[keys[i]] = v
// 			} else {
// 				t := make(map[string]any)
// 				t[keys[i]] = temp
// 				temp = t
// 			}
// 		}

// 		// for {
// 		// 	if v, ok := i[keys[0]]
// 		// }
// 		dst[keys[0]] = temp[keys[0]]
// 		logger.Info(temp[keys[0]])
// 	}
// 	return dst
// }

// func getFileContent(filePath string) (map[string]string, error) {
// 	fileMap, err := getFileMap(filePath)
// 	if err != nil {
// 		logger.Error(err)
// 		return nil, err
// 	}

// 	return encodeMap(fileMap), nil
// }

type Modify struct {
	old any
	new any
}

// increase，以A的角度查看比B多的部分
// decrease，以A的角度查看比B少的部分
// change，以A的角度查看变化的部分
func diffContent(a, b map[string]any) (increase, decrease map[string]any, change map[string]any) {
	increase = make(map[string]any)
	decrease = make(map[string]any)
	change = make(map[string]any)
	for k, v := range a {
		switch vv := v.(type) {
		case map[string]any:
			bv, ok := b[k]
			if ok {
				bvv, ok := bv.(map[string]any)
				if ok {
					subI, subD, subC := diffContent(vv, bvv)
					if len(subI) > 0 {
						increase[k] = subI
					}
					if len(subD) > 0 {
						decrease[k] = subD
					}
					if len(subC) > 0 {
						change[k] = subC
					}
				} else { // 不同的处理
					change[k] = Modify{old: bv, new: vv}
				}
			} else { // 新增的处理
				increase[k] = vv
			}

		default:
			bv, ok := b[k]
			if !ok {
				increase[k] = v
			}

			if !reflect.DeepEqual(v, bv) {
				change[k] = Modify{old: bv, new: v}
			}
		}

	}

	for k, v := range b {
		switch vv := v.(type) {
		case map[string]any:
			av, ok := a[k]
			if ok {
				avv, ok := av.(map[string]any)
				if ok {
					subI, _, _ := diffContent(vv, avv)
					if len(subI) > 0 {
						decrease[k] = subI
					}
					// increase[k] = subD
					// change[k] = subC
				} else { // 不同的处理
				}
			} else { // 减少的处理
				decrease[k] = vv
			}

		default:
			_, ok := a[k]
			if !ok {
				decrease[k] = v
			}
		}
	}
	return
}
