package main

import (
	"fmt"
	"math"
	"strconv"
)

// fmt คือ package ที่ต้องใช้อันดับแรก

// menu
// 1 การแทนค่าตัวแปล
// 2 การใช้ตัวแปลมากข้ึน
// 3 การเพิ่ม emoji ละเข้าใจประเภทการใช้งาน
// 4 การใช้ Printf
// 5 การใช้ if else
// 6 เครื่องหมายทางตรรกศาสตร์
// 7 ลำดับความสำคัญของตัวดำเนินการทางคณิตศาสตร์
// 8 การใข้ Switch case
// 9 วิธีการใช้ funtion
// 10 เรื่อง Array
// 11 การใช้ Loop
// 12 เรียน Slices
// 13 เรียน Structures
// 14 เรียน Method
// 15 เรียน Pointer
// 16 เรียน interface
// 17 เรียน error
// 18 เรียน map
// 19 เรียน ตัวแปลค่า
func main() {
	// 1 การแทนค่าตัวแปล
	// GO  นั้นไม่่จำเเป็ฯต้องกำหนด int หรืออื่นๆ เพราะมันสามารถอ่านค่าและกำหนดตัวแปรให้เราได้เอง เช่นด้านล่่าง
	zen, sen := 1, 2
	sum := zen + sen
	fmt.Println(sum)
	// GO  นั้น ไม่สามารถกำหนดตัวแปรขึ้นมาเฉยๆล้วไม่ใช้ไม่ได้
	// ต่่อมา ตรงนี้คือการ ต่อ string
	an := "ANS" + " MAAS"
	fmt.Println(an)
	// ต่อมาจะเป็นการเเติมคำให้ดูสวย
	fmt.Println("SUM :", sum)
	// หรือ
	ans := "ANS ="
	fmt.Println(ans, sum)
	fmt.Println("================================")
	// 2 การใช้ตัวแปลมากข้ึน
	name := "zqq"
	age := 13
	fmt.Println("I'm", name, "i", age, "year old")

	ap := "Apple"
	price := 22.13
	chack := true
	fmt.Println("Is", ap, "Price =", price, chack)
	fmt.Println("================================")
	// เป็นต้น
	// 3 การเพิ่ม emoji ละเข้าใจประเภทการใช้งาน
	// ต่อมาเป็นกการเรียนเรื่อง rune
	// วิธีการใส่ emoji ใน code แบะให้แสดงออกมา
	// ด้านบนให้แสดงเป็นตัวเลข ที่ใช้ '' char
	// ด้านล่างจะให้แสดง emoji ที่ใช้ "" ให้แสดงเป็น string
	em := '😀'
	emo := "🤣"
	fmt.Println(em)
	fmt.Println(emo)
	fmt.Println("================================")
	// 4 การใช้ Printf
	// ต่อมาเป็นการใช้ Printf

	a, b, c, d, e, f := "tttt", 1234, 22.33, '😀', "😀", true

	fmt.Printf("a = %s\n", a)
	fmt.Printf("b = %d\n", b)
	fmt.Printf("c = %f\n", c)
	fmt.Printf("c* = %.2f\n", c)
	fmt.Printf("d = %c\n", d)
	fmt.Printf("e = %s\n", e)
	fmt.Printf("f = %#T\n", f)
	// ต่อมา หากไม่รู้ว่าต้องใส่ % แทนค่าการแสดงตัวแปลยังไงให้ใส่ %#V เป็นต้น
	fmt.Println("================================")
	fmt.Printf("a = %#v\n", a)
	fmt.Printf("b = %#v\n", b)
	fmt.Printf("c = %#v\n", c)
	fmt.Printf("c* = %.2f\n", c)
	fmt.Printf("d = %#v\n", d)
	fmt.Printf("e = %s\n", e)
	fmt.Printf("f = %#v\n", f)
	fmt.Println("================================")
	// ส่วนต่อมาหากไม่รู้ว่า แต่ล่ะตัวแปลนั้น type อะไรให้เราทำำการ %#T เพื่อแสดงค่า Type นั้นๆออกมา
	fmt.Printf("a = %#T\n", a)
	fmt.Printf("b = %#T\n", b)
	fmt.Printf("c = %#T\n", c)
	fmt.Printf("c* = %.2f\n", c)
	fmt.Printf("d = %#T\n", d)
	fmt.Printf("e = %#T\n", e)
	fmt.Printf("f = %#T\n", f)
	fmt.Println("================================")
	// หรือเพื่อความระเอียด
	fmt.Printf("Type %#T --> a = %#v\n", a, a)
	fmt.Printf("Type %#T --> b = %#v\n", b, b)
	fmt.Printf("Type %#T --> c = %#v\n", c, c)
	fmt.Printf("c* = %.2f\n", c)
	fmt.Printf("Type %#T --> d = %#v\n", d, d)
	fmt.Printf("Type %#T --> e = %s\n", e, e)
	fmt.Printf("Type %#T --> f = %#v\n", f, f)
	fmt.Println("================================")
	// 5 การใช้ if else
	// ต่อมาจะมาเรียน if else

	num := 39
	// == 2 ตัวเเปลว่า มีค่า เท่ากันหรือไม่
	if num == 12 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
	fmt.Println("================================")
	// ต่อมาหาเลขคู่

	// ตัวแปล%2 คือการหาเลขคู่ หารลงตัว
	if num%2 == 0 {
		fmt.Println("คู่")
	} else {
		fmt.Println("เดี่ยว")
	}
	fmt.Println("================================")

	// ต่อมา เรียน else if
	if num%2 == 0 {
		fmt.Println("คู่")
	} else if num == 39 {
		fmt.Println("อะไรหว่าา")
	} else {
		fmt.Println("เดี่ยว")
	}
	fmt.Println("================================")

	// ต่อมาเรียนการใช้ math
	lim := 120.52
	p := math.Pow(2, 2)

	if p < lim {
		fmt.Println("🤣")
	} else if p == 80 {
		fmt.Println("OK")
	} else {
		fmt.Println("😀")
	}
	fmt.Println("================================")
	if p < lim {
		fmt.Println("🤣 speed you is: ", p)
	} else if p == 80 {
		fmt.Println("OK speed you is: ", p)
	} else {
		fmt.Println("😀 speed you is: ", p)
	}
	fmt.Println("================================")

	today := "mon"

	if today == "say" {
		fmt.Println(" OK ")
	} else if today == "mon" {
		fmt.Println(" Have a Good Day ")

	} else {

		fmt.Println(" Pim Kuy Rai")

	}
	fmt.Println("================================")
	// 6 เครื่องหมายทางตรรกศาสตร์
	// ! คือการกลับค่า หรือ ไม่เท่ากับ
	// + บวก ใช้เพื่อบวก เช่น a + b
	// - ลบ ใช้เพื่อลบ เช่น a – b
	// * คูณ ใช้เพื่อคูณ เช่น a * b
	// / หาร ใช้เพื่อหาร เช่น a / b
	// % หารแบบเอาเศษ (Modulo) ใช้เพื่อหารแบบเอาเศษที่เหลือ เช่น a % b
	// ++ เพิ่มค่าทีละ 1 ใช้เพื่อเพิ่มค่าอีก 1 ให้กับตัวแปร เช่น a = b++ หากตัว แปร b = 5 ค่าของตัวแปร a จะเท่ากับ 6
	// i += 5 i = i +5 10
	// -- ลดค่าทีละ 1 ใช้เพื่อลดค่าอีก 1 ให้กับตัวแปร เช่น a = b-- หากตัว แปร b = 5 ค่าของตัวแปร a จะเท่ากับ 4
	// f -= gf = f - g 8.75
	// += บวกค่าด้วยตัวแปร เช่น a += b คือ การบวกค่าตัวแปร a ด้วยตัวแปร b
	// -= ลดค่าด้วยตัวแปร เช่น a -= b คือ การลดค่าตัวแปร a ด้วยตัวแปร b
	// j *= (i - 3) j = j * (i - 3) 14
	// *= คูณค่าด้วยตัวแปร เช่น a *= b คือ การคูณค่าตัวแปร a ด้วยตัวแปร b
	// /= หารค่าด้วยตัวแปร เช่น a /= b คือ การหารค่าตัวแปร a ด้วยตัวแปร b
	// f /=3 f = f / 3 1.833333
	// %= หารแบบเอาเศษ ด้วยตัวแปร เช่น a %= b คือ การหารแบบเอาเศษค่าตัวแปร a ด้วยตัวแปร b
	// i % = (j - 2) i = i % (j - 2) 0
	// == เท่ากับ ตัวอย่าง a == b เป็นการเปรียบเทียบการเท่ากับระหว่างตัวแปร a และ b ซึ่งจะให้ค่าผลลัพธ์เป็นจริงหรือเท็จ
	// != ไม่เท่ากับ ตัวอย่าง a != b เป็นการเปรียบเทียบการไม่เท่ากับระหว่างตัวแปร a และ b ซึ่งจะให้ค่าผลลัพธ์เป็นจริงหรือเท็จ
	// < น้อยกว่า ตัวอย่าง a < b เป็นการเปรียบเทียบการน้อยกว่าระหว่างตัวแปร a และ b ซึ่งจะให้ค่าผลลัพธ์เป็นจริงหรือเท็จ
	// i < j True 1
	// <= น้อยกว่าหรือเท่ากับ ตัวอย่าง a <= b เป็นการเปรียบเทียบการน้อยกว่าหรือเท่ากับระหว่างตัวแปร a และ b ซึ่งจะให้ค่าผลลัพธ์เป็นจริงหรือเท็จ
	// (i + j) >=k True 1
	// > มากกว่า ตัวอย่าง a > b เป็นการเปรียบเทียบการมากกว่าระหว่างตัวแปร a และ b ซึ่งจะให้ค่าผลลัพธ์เป็นจริงหรือเท็จ
	// (j + k) > (i + 5) False 0
	// >= มากกว่าหรือเท่ากับ ตัวอย่าง a >= b เป็นการเปรียบเทียบการมากกว่าหรือเท่ากับระหว่างตัวแปร a และ b ซึ่งจะให้ค่าผลลัพธ์เป็นจริงหรือเท็จ
	// j == 2 False 0
	// && และ (and) ตัวอย่าง (a<2) && (b>3) เปรียบเทียบระหว่าง ค่า 2 ค่าโดยถูกเชื่อมด้วย "และ" แล้วให้ผลลัพธ์ เป็นจริงหรือเท็จ
	// || หรือ (or) ตัวอย่าง (a<2) || (b>3) เปรียบเทียบระหว่าง ค่า 2 ค่าโดยถูกเชื่อมด้วย "หรือ" แล้วให้ผลลัพธ์ เป็นจริงหรือเท็จ
	// ! นิเสธ (not) ตัวอย่าง !a จะได้ผลลัพธ์ตรงข้ามของค่า a เช่น หากตัวแปร a มีค่าเป็นจริง !a ก็จะให้ค่าเป็นเท็จ
	// ตารางที่ 5 หลักการดำเนินการด้วยเครื่องหมาย &&
	// ค่า A	ค่า B	ผลลัพธ์ของ A && B
	// T		T			T
	// T		F			F
	// F		T			F
	// F		F			F
	// ตารางที่ 6 หลักการดำเนินการด้วยเครื่องหมาย ||
	// ค่า A	ค่า B	ผลลัพธ์ของ A || B
	// T		T			T
	// T		F			T
	// F		T			T
	// F		F			F
	// ตารางที่ 7 หลักการดำเนินการด้วยเครื่องหมาย !
	// ค่า A	ผลลัพธ์ของ !A
	// T			F
	// F			T
	// นิพจน์								ผลลัพธ์
	// (i >=6) && (c == 'w')				1
	// (i >= 6) || (c == 119)				1
	// (f < 11) && (i > 100)				0
	// (c != 'p') || ((i + f) <= 10)		1
	// f > 5								1

	// 7 ลำดับความสำคัญของตัวดำเนินการทางคณิตศาสตร์

	// ลำดับที่				เครื่องหมาย			ลำดับการทำงาน
	// 1						( )					-
	// 2					 ++ ,  --			ซ้ายไปขวา
	// 3					* ,  / ,  %			ซ้ายไปขวา
	// 4					   + ,  -			ซ้ายไปขวา
	// 5				 < ,  <= ,  > ,  >=		ซ้ายไปขวา
	// 6					 == ,  !=			ซ้ายไปขวา
	// 7						&&				ซ้ายไปขวา
	// 8						||				ซ้ายไปขวา
	// 9	   = ,  += ,  -= ,  *= ,  /= ,  %=  ซ้ายไปขวา

	// ลำดับ							นิพจน์								นิพจน์ที่ถูกประมวลผลก่อน
	// 1			ans = 2 * ((i / 5) + (4 * (j - 2)) % (j - 2)	(i / 5), (j - 2) และ (j - 2)
	// 2			ans = 2 * (4 + (4 * (6)) % 6 )							  4 * 6
	// 3			ans = 2 * (4 + 24 % 6)								     24 % 6
	// 4			ans 2 * (4 + 0)									          4 + 0
	// 5			ans = 2 * 4												  2 * 4
	// 6			ans = 8

	// 8 การใข้ Switch case

	switch today {
	case "tue":
		fmt.Println(" No ")
	case "mon":
		fmt.Println("Have a Good Day")
	default:
		fmt.Printf("NOOOOOOO I Like is : %#v OK \n", today)
	}
	fmt.Println("================================")
	// หรือสามารถรวมกกันได้แบบนี้
	switch today {
	case "tue", "wet":
		fmt.Println(" No ")
	case "mon":
		fmt.Println("Have a Good Day")
	default:
		fmt.Printf("NOOOOOOO I Like is : %#v OK ", today)
	}
	fmt.Println("================================\n")

	//หรือเช็คตัวเลขแบบด้านล่าง
	switch ss := 1; {
	case ss == 1:
		fmt.Println("OK")
	case ss <= 1:
		fmt.Println("Low")
	case ss >= 1:
		fmt.Println("Hine")
	default:
		fmt.Println("NO")
	}
	fmt.Println("================================")

	// ดึงค่ามาจาก func add ที่อยู่ด้านล่าง
	q, bb := add(42, 50)
	fmt.Println("Ans = ", q, bb)
	// zz, xx ใส่ตัวอักษรที่ต้องการแสดงได้เลย เนื่องจากกำหนดจากหหนังสือ ttt เรียบร้อยแล้ว
	zz, xx := ttt("😏", "😥")
	fmt.Println(zz, xx)
	// yy, vv (ใส่เลขที่จะ +  ในตัวแปล spen ที่อยู่หนังสือเล่มที่จดสูตรได้เลย เพราะเราทำการกำหนดค่าไว้แล้ว)
	yy, vv := spen(10 + 40)
	fmt.Println(yy, vv)
	// uu (ใส่เลขที่จะ +  ในตัวแปล adds ที่อยู่หนังสือเล่มที่จดสูตรได้เลย เพราะเราทำการกำหนดค่าไว้แล้ว)
	uu := adds(39, 50)
	fmt.Println(uu)
	// แสดง type ของ adds ออกมา
	fmt.Printf("Type = %#T\n", adds)
	// กำหนดค่าจาก sddr และนำ สูตรจาก adds มาใส่และทำการแสดงค่าออกมา
	re := sddr(adds)
	fmt.Println(" re = ", re)
	// รับค่าทั้ง 2 ตัวมาจาก zxc และนำมาแสดงผล
	i, in := zxc()

	fmt.Println(i())
	fmt.Println(in())
	// สร้างตัวแปลรับค่าของ emote ขึ้นมาและทำการแสดงค่า
	ss := emote(8.3)
	fmt.Println(ss)
	fmt.Println("================================")
	// 10 เรื่อง Array
	// การสร้าง Array นั้น สามารถสร้างและพร้อมใช้งานได้ในทันที
	// Array นั้นจะนับแบบ 0 1 2 3 4 5 เเป็นต้น
	var skill [5]string = [5]string{"A", "S", "D", "F", "G"}
	fmt.Printf("Type : %#v\n ", skill)
	// ตัวอย่างการนับ Array
	cv := skill[0]
	fmt.Printf("Ans = %#v\n", cv)
	var skills [5]int
	fmt.Printf("Type : %#v\n ", skills)
	// วิธ๊หาขนาดของ Array

	le := len(skill)

	fmt.Println(le)
	// เราสามารถดึง Array  ที่ type ตรงกันมาแสดงได้เลย ถ้าไม่ตรงกกันไม่สามารุแสดงออกกกมาได้
	ssm(skill)
	// ต่อมาเป็นวิธีสร้าง Array ง่ายๆ
	sks := [3]string{"qq", "ww", "vv"}
	fmt.Printf("SShow = %#v", sks)
	fmt.Println("================================")
	// 11 การใช้ Loop
	// i ++ = i+1 และ หยุดกก่อนจะถึง 5
	for i := 0; i < 5; i++ {
		fmt.Println("i = ", i)
	}
	// ต้องประกาศ ตัวแปรด้านนอกก่อนจะนำไปใส่ loop
	jj := 0
	for j := 0; j < 5; j++ {

		// jj = j + jj
		jj += j
		fmt.Println("j = ", j, "jj = ", jj)

	}
	fmt.Println("sum = ", jj)
	// ต่อมาทำรูปดาว 5 แถว
	// sss คือการกำหนด ว่าดาวมีกี่แถว
	sss := 5
	// y = 0 คือกำนดเลข 0 คือค่าเริ่มต้น จากกนั้น y < sss คือถ้า y มีค่าน้อยกกว่า  sss ให้ y + จำนวนแถวเพิ่ม
	for y := 0; y < sss; y++ {
		// ต่อมากการซ้อน for กำหนดให้ x นั้นมีค่าช่องว่าง = 0 ต่อมา หาก x มีค่าน้อยกว่า  sss  ให้ทำการ - y และ - 1 ช่อง สุดท้ายให้ x+ เพิ่ม
		for x := 0; x < sss-y-1; x++ {
			// Print ค่าช่องว่าง
			fmt.Print(" ")
		}
		// ต่อมา z  มีค่า = 0 หาก z มีค่าน้อยกกว่า 2 ให้ทำการ * yและ +1 เพื่อเพิ่่มดาวตามจำนวนบรรทัด และ z++ เเพิ่มจำนวน
		for z := 0; z < 2*y+1; z++ {
			// Print รูปดาวเข้าไป
			fmt.Print("*")
		}
		// สุดท้ายให้ดาวนั้นเว้นบรรทัดเพิ่่มความสวยงาม
		fmt.Println()
	}

	// ต่อมาเป็นกการ Print ข้อมูลของ Array เป็นทีล่ะแถว
	for h := 0; h < len(skill); h++ {
		fmt.Println(skill[h])
	}
	// เขียนแบบย่อรูปง่ายกว่า
	for hh := range skill {
		fmt.Println(skill[hh])
	}
	// เขียนรูปแบบย่อและแบ่งประเภท
	for ii, hh := range skill {
		fmt.Println("num = ", ii, "TY = ", hh)
	}
	fmt.Println("================================")

	// 12 เรียน Slices
	// วิธ๊การใช้ เหมือน Array แต่สามารถเพิ่มหรือลดขนาดได้ เช่น
	sksw := []string{"qq", "ww", "vv"}
	fmt.Printf("SShow = %#v\n", sksw)
	// หรือกการเพิ่มจำนวนข้อมูล
	fmt.Printf("All = %d -- val = %#v \n", len(sksw), sksw)
	// append จะเเป็นคำำสั่่งเเพิ่มของfmt
	sksw = append(sksw, "tttt", "aassdw", "sadsas")
	fmt.Printf("All = %d -- val = %#v \n", len(sksw), sksw)
	// เป็นการระบุว่าจำให้โชว์ตรงไหนบ้าง
	s1 := sksw[0:2]
	fmt.Printf("Show = %#v\n", len(s1), s1)
	sl := len(sksw)
	fmt.Printf(" ss = %#v\n", len(sksw[:sl]))
	fmt.Println("================================")
	// กำหนดช่องโดย make
	sli := make([]int, 3)
	// ใส่ข้อมูล ไว้ที่ช่อง 1
	sli[1] = 42
	fmt.Printf("sli = %#v\n", sli)
	fmt.Println("================================")

	// 13 เรียน Structures
	// ดึง มาจาก type ที่อยู่ด้านล่าง
	cf := coures{
		n: "zen",
		i: "IDK",
		p: 99990,
	}
	// เปลี่ยนค่าตัวแปล i
	cf.i = "Is"
	// หลังจากกนั้นดึงมาใช้
	fmt.Println("name = ", cf.n)
	fmt.Println("Is = ", cf.i)
	fmt.Println("Price = ", cf.p)
	// ดึงค่าทั้งหมดออกกกมา
	fmt.Printf("Show  = %#v\n", cf)
	// 14 เรียน Method
	// การดึง func ด้านล้างออกมาแสดงโดยใช้ cf เป็ฯตัวเชื่อม
	// เก็บ คำสั่งไว้ใน dd
	dd := cf.discount()
	fmt.Println(dd)
	// แสดงสิ่งที่อยู่ด้านล่าง
	cf.info()
	// นำตัวเลขเข้ามาและแสดงผล
	es := cf.dis(5000)
	fmt.Println(es)
	fmt.Println("================================")
	// 15 เรียน Pointer
	pint := 9999
	// การเพิ่ม & ไว้ด้านหน้า pint คือกการระบุตำแหน่งของท่ีอยู่ที่จัดเก็บ
	fmt.Println("Pint = ", pint, &pint)
	// หรือเราสามารถเเก็ฐค้าไว้ได้
	// * คือการเก็บค่า int และถือที่อยู๋ของ pint เเลยทำให้ค่า pint เปลี่ยน
	var adddd *int = &pint
	*adddd = 5000
	fmt.Println(pint, adddd)
	// ต่อให้ตัวแปลเหหมือนกันแต่มันกกก็คือคนล่ะค่า ดูจากตัวเก็ฐข้อมูลด้านลัง
	cp(pint)

	// ต่อมาจะรับค่า type zaq จากด้านล้างขึ้นมาโชว์
	// new = การรับค่า return มาใช้
	z := new(zaq)
	fmt.Printf("Show = #%v", z)

	z.price = 590
	zzz := zaqw(z)
	fmt.Println(zzz)
	fmt.Println(z.price)
	fmt.Println("================================")
	// 16 เรียน interface
	// interdface สามารถประการตัวแปลที่ต่างกันได้เลย หรือ เก็บตัวแปลอะไรก็ได้ หรือใช้ any
	var zx any
	// แต่เวลาประการนั้น ต้องนำมาใช้ 2 ตัวไม่งั้นไม่แสดงผล
	zx = "asdsads"
	fmt.Printf("%#T %#v ", zx, zx)

	zx = 123
	fmt.Printf("%#T %#v \n", zx, zx)
	// รับค่า kk มาจากด้านล่าง
	var kk any
	// แปลงค่า kk เป็น any
	kk = 20
	fmt.Println(kk.(int))
	// จากนั้นทำการแสดงผลลัพธ์ออกมาแต่ต้องกำหนด kk.int เพื่อให้แสดงออกมาถูกต้อง หรือแปลง func ด้านล่างเป็น any

	var kks any
	kks = 15.9
	zaw(kks)
	fmt.Println("================================")

	// 17 เรียน error

	// ry, err := er(1, 0)

	// if err != nil {
	// 	fmt.Println("Error is ", err)
	// 	return
	// }
	// fmt.Println(ry, err)
	// fmt.Println("================================")
	// 18 เรียน map
	var mp map[string]int = map[string]int{}
	fmt.Println(mp)
	// ห้ามกกำหนด map ค่าซ้ำ
	mp["Ans"] = 42
	fmt.Printf(" = %#v \n", mp)
	vmp := mp["Ans"]
	fmt.Printf("aaa = %#v\n", vmp)
	mp["aaa"] = 50
	fmt.Println(mp)

	// วธ๊เช็คค่าใน map
	nk, ko := mp["Ans"]
	if ko {
		fmt.Println("Yeess")
		fmt.Println(" IS : ", nk)
	} else {
		fmt.Println("NO")
	}

	// เช็คง่ายกว่านั้นคือ

	if mp == nil {
		fmt.Println("Not")
	} else {
		fmt.Println(mp)
	}
	fmt.Println("================================")
	// 19 เรียน ตัวแปลค่า
	var is int = 256
	fmt.Printf("Type : %#T val: %v\n", is, is)
	var fs float64 = float64(is)
	fmt.Printf("Type : %#T val %v\n", fs, fs)
	var us uint8 = uint8(fs)
	fmt.Printf("Type : %#T val %v\n", us, us)
	//
	vfs := "499"
	st, err := strconv.Atoi(vfs)

	if err != nil {
		fmt.Println("err : ", err)
		return
	} else {
		fmt.Println(st)
	}
	// หรือแปลง int เป็น string ตรงๆ
	isu := 30
	iis := strconv.Itoa(isu)

	fmt.Println(iis)
	fmt.Println("================================")
}

// เปรียบเทียบ main คือกระดานดำที่คอยดึงข้อมูลจาก หนังสือหรือ funtion ที่เราเขียนแยกไว้ เอามาแสดง
// ส่วน func add คือหนังสือที่เราคอยจดบันทึกสูตรเอาไว้ และ ดึงเเอาไปแสดงหรือเขียนบนกระดานดำ
// ถ้าเพิ่มการรับค่า 2 ค่าขึ้นไปให้ทำการใส่ () ด้วย
// return คือการคืนค่าวนกลับไป
func add(x float64, y float64) (float64, string) {
	fmt.Println("Test", x, y)
	v := x + y
	return v, "Test2"
}

// ถ้าเกิด type เหมือนกัน เราสามารถละการใส่ได้ การเรียกใช้งานก็เหมือนกัน แค่สลับ
func ttt(x, y string) (string, string) {
	return y, x
}

// ต่อมาเป็นการ return แบบเปลือยคือ return แบบเปล่าๆ
func spen(set int) (x int, y int) {
	x = set + int(math.Pow(20, 3))
	y = set - int(math.Pow(30, 2))
	return x, y
}

// var จะมองเป็นค่าๆ นึง เช่น adds คือ funtion ที่เรากำหนดมา
var adds = func(x, y float64) float64 {
	return x + y

}

// func sddr ตัวนี้เราสามารถกำหนดค่าด้านในได้และ return ค่าขึ้นไปแสดงบนกกระดานดำได้ func(float64, float64) float64) เป็นผลการแสดง type จาก adds
func sddr(fn func(float64, float64) float64) float64 {

	ll := fn(23, 50)

	return ll
}

// ต่อมาจะเป็นกการ ruturn 2 ค่า
// func นั้นเเวลาขึ้นต้อง () ทุกครั้ง ส่่วน func() int  นั้นเป็นค่าเปล่าๆที่ให้แปลเป็ฯ int
// ต่อมาจะใ้มันreturn 2 ค่าแบบเดียวกัน
// การรับค่านั้น ทั้ง 2 ตัวจะต้องมี type ที่เหมือนกัน

// สามาารถทำบบด้านล้างก็ได้
// func ing() int {
// 	return 1
// }

// func ings() int {
// 	return 2
// }

func zxc() (func() int, func() int) {
	sum := 0
	return func() int {
			sum = sum + 1
			return sum
		}, func() int {
			return sum
		}
}

// วิธีการรับค่า จาก 64  เป็น str
// ขั้นตอนแรกให้สร้างตัวแปรขึึ้นมา ชื่อ emote จากนั้นให้สร้าง ตัวรับค่าที่ชื่อ reat ขึ้นมา ให้รับค่า float64 และ คืนค่าเป็น string
// จากนั้นสร้างสูตรที่ประมวลผลลัพธ์ และใหห้ทำการ return ค่าขึ้นไปแสดงบนกกระดานดำ
func emote(reat float64) string {

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

}

// วิธีการสร้าง func เพื่อใช้ Array และดึงไปใช้
func ssm(sm [5]string) {

	fmt.Printf("Show = %#v\ns", sm)

}

// ประการ type ของ structures เป็นการห่อตัวแปล

type coures struct {
	n string
	i string
	p float64
}

// รับค่าจากด้านบน cf ที่แปลงค่ามาเเป็น coures
func (cf coures) discount() float64 {
	// นำค่าจาก cf.p มาลบค่าตัวเลข และ ทำการเก็บไว้ในตัวแปล pp
	pp := cf.p - 10000
	fmt.Println("discoust = ", pp)
	// นำค่า return ไปด้านบน
	return pp
}

// สามารถสร้าง method ได้เท่าที่ต้องการ
func (cf coures) info() {
	fmt.Println("name = ", cf.n)
	fmt.Println("Is = ", cf.i)
	fmt.Println("Price = ", cf.p)
}

// สามารถเเพิ่มค่าเข้าไปได้

func (cf coures) dis(es float64) float64 {
	// นำค่าจากที่รับเข้ามา เข้าสูตร
	eess := cf.p - es + 50

	fmt.Println("DIS = ", eess)

	return eess
}

// เรื่อง Pointer
func cp(pint int) {
	// ต้องใช้ตัวปลเเดียวกันดงลงมาใช้สูตร
	pint = pint - 9999
	fmt.Println("sss = ", pint, &pint)

}

// type ที่ใช้กับ pointer
type zaq struct {
	names string
	foi   float64
	price int
}

// การเเปลี่ยนแปลงจะขึ้นอยู่ภายในเท่านั้น pointer
func zaqw(z *zaq) int {

	z.price = z.price - 6000
	fmt.Println("zaq ", z.price)
	return z.price
}

// 16 เรียน interface
func zvb(kk int) {
	fmt.Println(kk)
}

// อีกกกรณี
func zaw(kks any) {
	o, ok := kks.(int)
	// o มีค่า = kks.(int) ให้ kks แปลงค่าจาก any เป็น int
	if ok {
		o = o + 40
		fmt.Println(o)
	} else {
		fmt.Println("Not Int")
	}
	// วิธ๊ง่ายกว่าคือ swicth

	switch pk := kks.(type) {
	case int:
		o := pk + 20
		fmt.Println("sss = ", o)
	case string:
		o := pk + "HI"
		fmt.Println("test = ", o)
	default:
		fmt.Println("Not type")
	}
}

// วิิธีจัดการ error
// error มี type เป็น inteface
func er(ay, by int) (int, error) {
	if by == 0 {
		erro := merror{}
		return 0, erro
	}
	ry := ay / by
	return ry, nil
}

type merror struct{}

func (ero merror) Error() string {

	return "MyError"
}

// 18 เรียน map
