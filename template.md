# テンプレート

## 実行コード

```go
package xxxx

import (
	"fmt"
)

func Execute(){
	fmt.Println("Problem xx 「xxxx」")
	fmt.Println("----")
	fmt.Println(calc(xxx))
}

func calc(num int) int{
  ans := xxxx
  return ans
}
```

## テストコード

```go
package euler1

import (
	"testing"
)

func TestCalc(t *testing.T) {
	result := calc(xxx)
	expext := xxxx

	if result != expext {
		t.Error("\n結果：", result, "\n期待値：", expext)
	}

	t.Log("xxxx テスト終了")
}

```
