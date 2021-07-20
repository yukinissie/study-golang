package main

import "fmt"

type dog struct {
	name   string
	group  string
	height int
}

func bigger(d1 dog, d2 dog) string {
	diff := d1.height - d2.height
	switch {
	case diff > 5:
		return "より大きい"
	case diff < -5:
		return "より小さい"
	default:
		return "と同じ大きさ"
	}
}

func biggest(dogs []dog) dog {
	biggest := dogs[0]

	for i := 1; i < len(dogs); i++ {
		if dogs[i].height > biggest.height {
			biggest = dogs[i]
		}
	}
	return biggest
}

func main() {
	pome := dog{"ポメ", "ポメラニアン", 25}
	maru := dog{"マル", "マルチーズ", 22}
	shiba := dog{"シバ", "芝犬", 40}

	dogs := []dog{pome, maru, shiba}

	// for i := 0; i < len(dogs); i++ {
	// 	next := i + 1
	// 	if next >= len(dogs) {
	// 		next = 0
	// 	}
	// 	fmt.Printf("%sさんは%sさん%sです\n", dogs[i].name, dogs[next].name,
	// 		bigger(dogs[i], dogs[next]))
	// }
	fmt.Printf("%sさんが最大です\n", biggest(dogs).name)
}
