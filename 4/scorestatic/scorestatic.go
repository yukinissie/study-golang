package main

import "fmt"

func scoreavg(scores []int) int {
	sum := 0
	for i := 0; i < len(scores); i++ {
		sum += scores[i]
	}
	return sum / len(scores)
}

func sbsqr(a int, b int) int {
	return (a - b) * (a - b)
}

// func avgvar(scores []int, avg int) int {
// 	sum := 0
// 	for i := 0; i < len(scores); i++ {

// 	}
// }

func sample(a int, b int) (hoge int, fuga int) {
	// hoge := 1
	// fuga := 2

	// if hoge {
	// 	hoge
	// }
	hoge = a
	fuga = b

	return
}

func main() {
	// mathscores := []int{40, 89, 77, 68, 72, 39, 58, 87, 93, 48, 65, 74, 60}
	// fmt.Printf("%d名による数学試験の結果:\n", len(mathscores))
	// fmt.Printf("平均点 %d点\n", scoreavg(mathscores))
	// fmt.Printf("分散   %d(点・点)\n")
	foo, bar := sample(1, 2)
	fmt.Printf("%d %d", foo, bar)
}
