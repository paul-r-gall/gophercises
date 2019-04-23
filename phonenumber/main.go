package main

import (

	//import drivers??
	"fmt"
	"math/rand"
	"time"

	"strconv"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func intToStr(nums []int) string {
	s := ""
	for _, i := range nums {
		s += strconv.Itoa(i)
	}
	return s
}

func randPhn() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var num [10]int
	for i := range num {
		num[i] = r.Intn(10)
	}
	fmt.Println(num)

	parens := r.Intn(2) == 1
	dash := r.Intn(2) == 1
	space := r.Intn(2) == 1
	s := intToStr(num[0:3])
	if parens {
		s = "(" + s + ")"
	}
	if space {
		s += " "
	}
	s += intToStr(num[3:6])
	if dash {
		s += "-"
	}
	s += intToStr(num[6:])
	return s
}

func main() {
	fmt.Println(randPhn())
}
