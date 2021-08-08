package euler1

import (
	"fmt"
)

// Problem 1 「3と5の倍数」
// 10未満の自然数のうち, 3 もしくは 5 の倍数になっているものは 3, 5, 6, 9 の4つがあり, これらの合計は 23 になる.
// 同じようにして, 1000 未満の 3 か 5 の倍数になっている数字の合計を求めよ.

func Execute(){
	fmt.Println("Problem 1 「3と5の倍数」")
	fmt.Println("----")
	fmt.Println(calc(1000))
}

func calc(num int) int{
	ans := 0

	for i := 1; i < num; i++ {
		if i % 3 == 0 || i % 5 == 0 {
			ans += i
		}
	}

	return ans
}