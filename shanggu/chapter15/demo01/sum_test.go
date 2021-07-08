package demo01

import "testing"

func TestSumNum(t *testing.T) {
	res := SumNum(3, 1)
	if res != 7 {
		t.Fatalf("Sum() 错误，希望为=%v，结果为=%v", 7, res)
	}

	t.Logf("测试成功")
}
