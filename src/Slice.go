package main

import "fmt"

func main() {
	sliceTest()
}
func sliceTest() {
	x := make([]int, 7)
	for i := 0; i < 7; i++ {
		x[i] = i * 100
	}
	twohundred := &x[1]
	x = append(x, 800)
	x[1]++
	fmt.Printf("\n=> Slice and reference\n")
	fmt.Println("twohundred:", *twohundred, "x[1]:", x[1])

	fmt.Println(len(x), cap(x))
}

func copySlice() {
	member := make([]string, 5, 8)
	member[0] = "Gump"
	member[1] = "Jonny"
	member[2] = "Chris"
	member[3] = "Binary"
	member[4] = "Edgar"

	copyMember := make([]string, 3, 4)
	copy(copyMember, member)

	fmt.Printf("\n=> original a slice\n")
	inspectSlice(member)
	fmt.Printf("\n=> Copy a slice\n")
	inspectSlice(copyMember)
}

func sliceToSlice() {
	member := make([]string, 5, 8)
	member[0] = "Gump"
	member[1] = "Jonny"
	member[2] = "Chris"
	member[3] = "Binary"
	member[4] = "Edgar"

	sliceMember := member[2:4]
	sliceMember = append(sliceMember, "ADD")

	inspectSlice(member)
	inspectSlice(sliceMember)
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
	var member []string
	member = []string{"Gump", "Jonny", "Chris", "Binary", "Edgar"}
	member[2] = "gump"
	fmt.Println(member)
}

func makeDeclaration() {
	member := make([]string, 5, 10)
	member[0] = "Gump"
	member[1] = "Jonny"
	member[2] = "Chris"
	member[3] = "Binary"
	member[4] = "Edgar"
	fmt.Println(member)
	fmt.Printf("len = %v , cap = %v\n", len(member), cap(member))

	member = append(member, "GilDong")
	fmt.Println(member)
	fmt.Printf("len = %v , cap = %v\n", len(member), cap(member))

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

func inspectSlice(slice []string) {
	fmt.Printf("Length[%d] Capacity[%d]\n", len(slice), cap(slice))
	for i := range slice {
		fmt.Printf("[%d] %p %s\n", i, &slice[i], slice[i])
	}
}
