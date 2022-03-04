package test

//测试类需要 满足 1） 测试源文件的命名格式为 **_test.go 2) 测试的函数为 Test**( t testing.T)

import (
	"strings"
	"testing"
	"text/template"

	"github.com/genkimaru/golang-learning/reflect/example3/tools"
)

func TestRepeat(t *testing.T) {

	t.Log("Rolling all gutter balls... (expected score: 0)")

	funcs := template.FuncMap{
		"trim":    strings.Trim,
		"lower":   strings.ToLower,
		"repeat":  strings.Repeat,
		"replace": strings.Replace,
	}
	result, _ := tools.Eval(`{{ "hello" 4 | repeat }}`, funcs)

	if result != "hellohellohellohello" {
		t.Errorf("Expected hello of 4 times, but it was %s instead.", result)
	}

}
