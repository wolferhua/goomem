package goomem

import "testing"

func Test_isBaseType(t *testing.T) {
	if isBaseType(6) { //try a unit test on function
		t.Error("测试通过") // 如果不是如预期的那么就报错
	} else {
		t.Log("测试不过") //记录一些你期望记录的信息
	}
	if isBaseType(1.2) { //try a unit test on function
		t.Error("测试通过") // 如果不是如预期的那么就报错
	} else {
		t.Log("测试不过") //记录一些你期望记录的信息
	}
	if isBaseType("") { //try a unit test on function
		t.Error("测试通过") // 如果不是如预期的那么就报错
	} else {
		t.Log("测试不过") //记录一些你期望记录的信息
	}
	if isBaseType(mItem{}) { //try a unit test on function
		t.Error("测试通过") // 如果不是如预期的那么就报错
	} else {
		t.Log("测试不过") //记录一些你期望记录的信息
	}
}
