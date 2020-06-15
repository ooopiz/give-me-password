package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

var (
	w    int
	n    int
	v, V bool
	c    string
)

func init() {

	flag.IntVar(&w, "w", 12, "Set `word` count")

	flag.BoolVar(&v, "v", false, "Show version then exit")
	flag.BoolVar(&V, "V", false, "Show version and configure options then exit")

	flag.Parse()

	if v {
		fmt.Println("Version 1.0.0")
		os.Exit(0)
	}
	if V {
		fmt.Println("Version 1.0.0 ...")
		os.Exit(0)
	}
}

func main() {
	//  number := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	//  capitalLetter := []string{
	//      "A", "B", "C", "D", "E", "F", "G", "H", "I", "J",
	//      "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T",
	//      "U", "V", "W", "X", "Y", "Z"}
	//  lowercaseLetter := []string{
	//      "a", "b", "c", "d", "e", "f", "g", "h", "i", "j",
	//      "k", "l", "m", "m", "o", "p", "q", "r", "s", "t",
	//      "u", "v", "w", "x", "y", "z"}
	//  symbol := []string{"!", "@", "#", "$", "%", "^", "&", "*"}

	number := []string{"2", "3", "4", "5", "6", "7", "8", "9"}
	//capitalLetter := []string{
	//	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J",
	//	"K", "L", "M", "N", "P", "Q", "R", "S", "T",
	//	"U", "V", "W", "X", "Y", "Z"}
	lowercaseLetter := []string{
		"a", "b", "c", "d", "e", "f", "g", "h", "j",
		"k", "l", "m", "m", "p", "q", "r", "s", "t",
		"u", "v", "w", "x", "y", "z"}
	//symbol := []string{"@", "#", "$", "%", "^", "&", "*"}

	combine := []string{}
	combine = append(combine, number...)
	combine = append(combine, lowercaseLetter...)
	//fmt.Println(combine)

	for i := 0; i < 5; i++ {
		combine = Shuffle(combine, w)
		//fmt.Println(combine)
		result := strings.Join(combine, "")
		fmt.Println(result)
	}
}

func Shuffle(slice []string, word int) []string {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	ret := make([]string, word)
	n := word
	for i := 0; i < n; i++ {
		randIndex := r.Intn(len(slice))
		ret[i] = slice[randIndex]
		slice = append(slice[:randIndex], slice[randIndex+1:]...)
	}
	return ret
}
