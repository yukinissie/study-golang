package main

import "fmt"

func main() {

	pattern := [3]string{"โ", "โ๏ธ", "๐"}

	fmt.Println("****ใใใใใๅๆใชในใ****")

	for me := 0; me < 3; me++ {

		fmt.Printf("็งใ%sใฎใจใ:\n", pattern[me])

		for you := 0; you < 3; you++ {
			score := (3 + me - you) % 3

			if score == 2 {
				fmt.Printf("ใใชใใ%sใชใ็งใฎๅใก\n", pattern[you])
			} else if score == 1 {
				fmt.Printf("ใใชใใ%sใชใ็งใฎ่ฒ ใ\n", pattern[you])
			} else {
				fmt.Printf("ใใชใใ%sใชใใใใ\n", pattern[you])
			}
		}
		fmt.Println()
	}
}
