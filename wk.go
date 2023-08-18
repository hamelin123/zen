package main

import (
	"fmt"
	"strings"
)

func mainp() {
	// Workshop Println
	name := "AV"
	year := 1999
	rt := 99.022
	t := "xx - re"
	hero := true
	emoji := '🤣'

	fmt.Println("เรื่่อง : ", name)
	fmt.Println("ปี : ", year)
	fmt.Println("เรต : ", rt)
	fmt.Println("ประเภท : ", t)
	fmt.Println("ฮีโร่ : ", hero)
	fmt.Println("================================")
	// Workshop Printf
	fmt.Printf("ชื่อเรื่อง : %#v\n ", name)
	fmt.Printf("ปี : %#v\n", year)
	fmt.Printf("เรต : %#v\n", rt)
	fmt.Printf("ประเภท : %#v\n", t)
	fmt.Printf("ฮีโร่ : %#v\n", hero)
	fmt.Printf("Emoji : %c\n", emoji)
	fmt.Println("================================")
	// Workshop if else
	num := 40

	if num <= 50 {
		fmt.Println("Bad")
	} else if num <= 60 {
		fmt.Println(" Not Bad")
	} else if num <= 70 {
		fmt.Println("Good")
	} else {
		fmt.Println("Very Good")
	}
	fmt.Println("================================")
	rat := 9.9

	if rat < 5.0 {
		fmt.Println("Dis 😀")
	} else if rat >= 5.0 && rat <= 7.0 {
		fmt.Println("Nor 😁")
	} else if rat >= 7.0 && rat <= 10.0 {
		fmt.Println("GOOD 😎")
	} else {
		fmt.Println("KUY 😑")
	}
	fmt.Println("================================")

	// Workshop Switch case
	switch rad := 8.4; {
	case rad == 5.0:
		fmt.Println("Dis 😀")
	case rad >= 5.1 && rad <= 7.0:
		fmt.Println("Nor 😁")
	case rad >= 7.1 && rad <= 10.0:
		fmt.Println("GOOD 😎")
	default:
		fmt.Println("KUY 😑")
	}
	fmt.Println("================================")

	ss := emot(8.3)
	fmt.Println(ss)
	fmt.Println("================================")
	// Workshop Array

	mo := [3]string{"ad", "mv", "ss"}

	moo(mo)
	mo[1] = "sss"
	fmt.Printf("Show := %#v\n", mo[1])
	fmt.Println("================================")
	// Workshop Array 1

	moos := [3]string{"ssss", "eeee", "wwww"}

	fmt.Printf("Show := %#v\n", moos[0])
	fmt.Printf("Show := %#v\n", moos[1])
	fmt.Printf("Show := %#v\n", moos[2])
	moos[1] = "ww"
	fmt.Printf("Show := %#v\n", moos[1])
	fmt.Println("================================")
	// Workshop Loop

	mt := [3]string{"AC", "AD", "FA"}
	for m := 0; m < len(mt); m++ {

		fmt.Println(mt)

	}
	for _, mt := range mt {

		fmt.Println("Movie = ", mt)

	}
	// Workshop Slices

	xs := []float64{1, 2, 3, 4, 5, 6, 7}
	ys := []float64{8, 9, 10, 11, 12}

	fmt.Print("vote = ", xs, ys)
	// vote คือการรวม  xs,ys(...)คือกการเเพิ่มทั้งหมด
	vote := append(xs, ys...)
	fmt.Printf("vote = %v\n", vote)
	// กำหนดการแสดงผล
	fmt.Printf("Vote 5-8 = %v\n", vote[5:9])
	fmt.Println("================================")
	// Workshop Structures
	// ตั้งตัวแปล ชื่อหนัง 2 เรื่องและเกก็บคนล่ะค่า
	moivess := moviee{
		mov: "avv",
		y:   1990,
		r:   6.9,
		pp:  []string{"Ac", "sss"},
		bbb: true,
	}
	moivesss := moviee{
		mov: "scs",
		y:   1991,
		r:   9.5,
		pp:  []string{"ac", "dssd"},
		bbb: true,
	}
	// จากกนั้นใหห้ทำกการรวทั้ง 2 เรื่องไว้ในตัวแปลเดียว
	var msv []moviee = []moviee{moivess, moivesss}
	// สุดท้ายใ้แสดงนังออกมาแบบ loop
	for _, ms := range msv {
		fmt.Printf("%#v\n", ms)
	}
	fmt.Println("================================")
	// Workshop Method
	me := moviee{
		mov: "avv",
		y:   1990,
		r:   6.9,
		pp:  []string{"Ac", "sss"},
		bbb: true,
	}
	me.ifs()
	// Workhop Pointer
	mes := &mooa{
		movs: "avv",
		ys:   1990,
		rs:   6.9,
		vs:   []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		pps:  []string{"Ac", "sss"},
		bbbs: true,
	}
	mes.addvs(9)
	fmt.Println("vose = ", mes.vs)
	fmt.Println("================================")
	// Workshop Interface
	var sv interface {
		ddvs(rs float64)
	}
	fmt.Println(sv)
	fmt.Println("================================")
	// Workshop map

	s := "If andd dptl oif'e ddldfks If and eee rrr ss gnjr;a jfg"
	w := WordCount(s)
	fmt.Printf("%#v\n", w)
	fmt.Println("================================")
}

// Workshop function

func emot(reat float64) string {

	switch {
	case reat == 5.0:
		return "Dis 😀"
	case reat >= 5.1 && reat <= 7.0:
		return "Nor 😁"
	case reat >= 7.1 && reat <= 10.0:
		return "GOOD 😎"
	default:
		return "KUY 😑"
	}

	// Workshop Method

}

// Workshop function Array
func moo(mos [3]string) {

	fmt.Printf("Show := %#v\n", mos[0])
	fmt.Printf("Show := %#v\n", mos[1])
	fmt.Printf("Show := %#v\n", mos[2])

}

// Workshop Type Structures

type moviee struct {
	mov string
	y   int
	r   float64
	pp  []string
	bbb bool
}

// Workshop Method

func (me moviee) ifs() {
	fmt.Printf("%s (%d) - %.2f\n", me.mov, me.y, me.r)
	fmt.Println("PPP = ")
	for _, g := range me.pp {
		fmt.Printf("\t%s\n", g)
	}
}

// Workhop Pointer
type mooa struct {
	movs string
	ys   int
	rs   float32
	vs   []float64
	pps  []string
	bbbs bool
}

// ตรงนี้ไว้เพิ่มค่าา addvs ด้านบน
func (ms *mooa) addvs(rs float64) {
	ms.vs = append(ms.vs, rs)

}

// Workshop map นับคำซ้ำ
func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	r := map[string]int{}

	for _, w := range words {
		r[w] = r[w] + 1
	}
	return r
}
