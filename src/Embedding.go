package main

import "fmt"

type user struct {
	name  string
	email string
}
type admin struct {
	user  // 임베딩 타입
	level string
}
type notifier interface {
	notify()
}

func main() {

	ad := admin{
		user: user{
			name:  "Hoanh An",
			email: "hoanhan101@gmail.com",
		},
		level: "superuser",
	}

	// 내부 타입 메서드를 직접 사용 가능하다.
	ad.user.notify()
	ad.notify()

	sendNotification(&ad)
}

func (u *user) notify() {
	fmt.Printf("Sending user email To %s<%s>\n", u.name, u.email)
}
func sendNotification(n notifier) {
	n.notify()
}
