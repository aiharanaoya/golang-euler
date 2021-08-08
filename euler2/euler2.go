package euler2

import (
	"fmt"
)

// Problem 2 「偶数のフィボナッチ数」
// フィボナッチ数列の項は前の2つの項の和である. 最初の2項を 1, 2 とすれば, 最初の10項は以下の通りである.
// 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...
// 数列の項の値が400万以下のとき, 値が偶数の項の総和を求めよ.

func Execute(){
	fmt.Println("Problem 2 「偶数のフィボナッチ数」")
	fmt.Println("----")
	fmt.Println(calc(4000000))
}

func calc(num int) int{
	ans := 2
	wa := 0
	slice := [] int{1, 2}

	for {
		wa = slice[(len(slice) - 1)] + slice[(len(slice) - 2)]

		if wa > num {
			break
		}

		slice = append(slice, wa)

		if(wa % 2 == 0){
			ans += wa
		}
	}

	return ans
}