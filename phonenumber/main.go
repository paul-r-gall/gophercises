package main

import (

	//import drivers??
	"fmt"
	"math/rand"
	"time"

	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

const DB_SIZE = 20
const DB_FILE = "pNums.db"

func intToStr(nums []int) string {
	s := ""
	for _, i := range nums {
		s += strconv.Itoa(i)
	}
	return s
}

func cleanPNum(s string) string {
	newS := ""
	mm := map[string]bool{" ": true, "(": true, ")": true, "-": true}
	for _, r := range s {
		c := string(r)
		if !mm[c] {
			newS += string(c)
		}
	}
	return newS
}

func randPNum() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var num [10]int
	for i := range num {
		num[i] = r.Intn(10)
	}

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

type pNum struct {
	gorm.Model
	Num string `gorm:"not null;unique"`
}

func main() {
	db, err := gorm.Open("sqlite3", DB_FILE)
	if err != nil {
		return
	}

	if !db.HasTable(&pNum{}) {
		fmt.Println("making tbl")
		db.AutoMigrate(&pNum{})
		fmt.Println("automigration completed")
		for i := 0; i < DB_SIZE; i++ {
			rNum := randPNum()
			//fmt.Println(rNum)
			db.Create(&pNum{Num: rNum})
		}
		fmt.Println("filling completed")
	}
	defer db.Close()

	pNums := []pNum{}
	db.Find(&pNums)

	//fmt.Println(pNums)

	for _, sNum := range pNums {
		fmt.Println(sNum.ID)
		fmt.Println(cleanPNum(sNum.Num))
		db.Model(&sNum).Update("num", cleanPNum(sNum.Num))
	}
}
