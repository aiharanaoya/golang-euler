package euler3

import (
	"fmt"
)

// Problem 3 「最大の素因数」
// 13195 の素因数は 5, 7, 13, 29 である.
// 600851475143 の素因数のうち最大のものを求めよ.

func Execute(){
	fmt.Println("Problem 3 「最大の素因数」")
	fmt.Println("----")
	fmt.Println(calc(600851475143))
}

func calc(num int) int{
	ans := 0

	for num > 1 {
		for ans < num {
			ans++

			if !isPrimeNumber(ans) {
				continue
			}
			for num % ans == 0 {
				num /= ans
			}
		}
	}
	return ans
}

func isPrimeNumber(num int) bool{
	if num == 2 {
		return true
	}
	for i := 2; i <= num; i++ {
		if num % i == 0 {
			return false
		}
		if i + 1 == num {
			return true
		}
	}
	return false
}
