package main

import "fmt"

func main() {
	//defaultDeclaration()
	//makeDeclaration()
	//nilSlice()
	sliceTest()
}
func sliceTest() {

}

func sliceAppend() {
	oldMember := []string{"Gump", "Jonny", "Chris"}
	newMember := []string{"Binary", "Edgar"}

	member := append(oldMember, newMember...)
	fmt.Printf("oldMember = %s\n", oldMember)
	fmt.Printf("newMember = %s\n", newMember)
	fmt.Printf("Member    = %s\n", member)

	//member := make([]string, 3, 5) // len: 3 cap: 5인 slice 생성
	//member[0] = "Gump"
	//member[1] = "Jonny"
	//member[2] = "Chris"
	//fmt.Printf("member = %s, len = %d, cap = %d\n", member, len(member), cap(member))
	//
	//newMember := append(member,"Binary")
	//fmt.Printf("member = %s, len = %d, cap = %d\n", newMember, len(newMember), cap(newMember))

	a := make([]int, 3, 4) // len: 3 cap: 4인 slice 생성
	a[0] = 10
	a[1] = 20
	a[2] = 30

	b := append(a, 40) // a에 여분의 용량이 남으므로 내부배열 공유
	c := append(a, 50) // a에 여분의 용량이 남으므로 내부배열 공유
	d := append(c, 60) // c에 여분의 용량이 남지 않으므로 새로운 내부배열 할당

	fmt.Println(a, len(a), cap(a)) // [10 20 30] 3 4
	fmt.Println(b, len(b), cap(b)) // [10 20 30 50] 4 4
	fmt.Println(c, len(c), cap(c)) // [10 20 30 50] 4 4
	fmt.Println(d, len(d), cap(d)) // [10 20 30 50 60] 5 8

	fmt.Println(&a[0], &a[1], &a[2])
	fmt.Println(&b[0], &b[1], &b[2], &b[3])
	fmt.Println(&c[0], &c[1], &c[2], &c[3])
	fmt.Println(&d[0], &d[1], &d[2], &d[3], &d[4])
}

func defaultDeclaration() {
	var name []string
	name = []string{"Gump", "Jonny", "Chris", "Binary", "Edgar"}
	name[2] = "gump"
	fmt.Println(name)
}

func makeDeclaration() {
	name := make([]string, 5, 10)
	name[0] = "Gump"
	name[1] = "Jonny"
	name[2] = "Chris"
	name[3] = "Binary"
	name[4] = "Edgar"
	fmt.Println(name)
	fmt.Printf("len = %v , cap = %v\n", len(name), cap(name))

	name = append(name, "GilDong")
	fmt.Println(name)
	fmt.Printf("len = %v , cap = %v\n", len(name), cap(name))

}

func nilSlice() {
	var a []int
	var b []int = make([]int, 0)
	var c []int = []int{}

	fmt.Println(a, b, c)  // [] [] []
	fmt.Println(a == nil) // true
	fmt.Println(b == nil) // false
	fmt.Println(c == nil) // false
}
