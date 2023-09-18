package biz

import (
	"fmt"
	"testing"
)

func TestFileDiff(t *testing.T) {
	a, b := "../../static/swagger-v1.0.0.json", "../../static/swagger-v1.1.0.json"
	output := "runtime/swagger-diff.txt"
	f := NewFileDiff(a, b, output)
	fmt.Println(f.Diff())
}
