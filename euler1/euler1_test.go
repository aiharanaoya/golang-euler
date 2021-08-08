package euler1

import (
	"testing"
)

func TestCalc(t *testing.T) {
	result := calc(1000)
	expext := 233168

	if result != expext {
		t.Error("\n結果：", result, "\n期待値：", expext)
	}

	t.Log("euler1 テスト終了")
}
