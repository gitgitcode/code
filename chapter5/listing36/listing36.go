// Sample program to show how to use an interface in Go.
package main

import (
	"fmt"
)

// notifier is an interface that defined notification
// type behavior.
type notifier interface {
	notify()//行为
}

// user defines a user in the program.
type user struct {
	name  string
	email string
}

// notify implements a method with a pointer receiver.
func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n",
		u.name,
		u.email)
}

// main is the entry point for the application.
func main() {
	// Create a value of type User and send a notification.
	u := user{"Bill", "bill@email.com"}

	sendNotification(u)//u没有实现 使用 sendNotification(&u)
	
	//不能将 u 类型是user 作为 sendNotification的参数类型
	//user 类型咩有实现notifier 
	// ./listing36.go:32: cannot use u (type user) as type
	//                     notifier in argument to sendNotification:
	//   user does not implement notifier
	//                          (notify method has pointer receiver)
	//notify方法使用指针接收者声明
}

// sendNotification accepts values that implement the notifier
// interface and sends notifications.
func sendNotification(n notifier) {
	n.notify()
}
//接收 notifier 接口类型的值 之后接口调用n.notify方法
