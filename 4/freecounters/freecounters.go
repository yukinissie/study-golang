package main

import "fmt"

func counter() func() {
	ctr := 0
	fmt.Println("カウンタを初期化しました。")
	return func() {
		ctr++
		fmt.Println(ctr)
	}
}

func freecounter(start int) func(int) int {
	ctr := start
	fmt.Printf("フリーカウンターを%dからはじめます\n", ctr)
	return func(add int) int {
		fmt.Printf("%dを足して", add)
		ctr += add
		return ctr
	}
}

func main() {
	freecnt := freecounter(10)
	fmt.Println(freecnt(2))
	fmt.Println(freecnt(5))
	fmt.Println(freecnt(7))
}
