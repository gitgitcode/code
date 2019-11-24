// Sample program to show how to declare methods and how the Go
// compiler supports them.
package main

import (
	"fmt"
)

// user defines a user in the program.
//定义一个用户类型
type user struct {
	name  string
	email string
}

// notify implements a method with a value receiver.
//使用值接受者实现一个方法
func (u user) notify() {
	fmt.Printf("Sending User Email To %s<%s>\n",
		u.name,
		u.email)
}
//值接收者 
//指针接受者

// changeEmail implements a method with a pointer receiver.
func (u *user) changeEmail(email string) {
	u.email = email
}

// main is the entry point for the application.
func main() {
	// Values of type user can be used to call methods
	// declared with a value receiver.
	bill := user{"Bill", "bill@email.com"}
	bill.notify()

	// Pointers of type user can also be used to call methods
	// declared with a value receiver.
	lisa := &user{"Lisa", "lisa@email.com"}
	lisa.notify()

	// Values of type user can be used to call methods
	// declared with a pointer receiver.
	bill.changeEmail("bill@newdomain.com")
	bill.notify()

	// Pointers of type user can be used to call methods
	// declared with a pointer receiver.
	lisa.changeEmail("lisa@newdomain.com")
	lisa.notify()
}
