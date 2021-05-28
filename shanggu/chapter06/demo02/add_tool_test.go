package main

import "testing"

// 单元测试
func TestAddTool(t *testing.T) {
	tool := addTool()

	if ans := tool(10); ans != 10 {
		t.Error("addTool() should be equal 10")
	}
}
