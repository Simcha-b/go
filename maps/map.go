package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	// input := "go is fun and go is fast"
	// fmt.Println(WordCount(input))

	// mp := make(map[string]int)
	// mp["אלוקינו"] = 1
	// mp["לוחות הברית"]=2
	// mp ["אבות"] = 3
	// v := mp["אבות"]
	// fmt.Println(v)
	// delete(mp,"אבות")

	// nextInt := intSeq()
	// fmt.Println(nextInt())
	// fmt.Println(nextInt())
	// fmt.Println(nextInt())
	// newInts := intSeq()
	// fmt.Println(newInts())

		 mp := map[string]int{"a":1, "b":2}
	 fmt.Println(mp)
	 fmt.Println(len(mp))
	 clear(mp)
	 fmt.Println(mp)

	s := "שלום"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c ", s[i])
	}

	fmt.Println("Rune count:", utf8.RuneCountInString(s))

}

func WordCount(str string) map[string]int {
	res := make(map[string]int)
	strArr := strings.Fields(str)

	for _, word := range strArr {
		res[word]++
	}
	return res
}

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
