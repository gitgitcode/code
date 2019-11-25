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
//使用 值 接受者实现一个方法
func (u user) notify() {
	//func 和函数之间的参数 成为接收者
	fmt.Printf("Sending User Email To %s<%s>\n",
		u.name,
		u.email)
}
//值接收者 
//指针接受者

// changeEmail implements a method with a pointer receiver.
//指针 接收者
func (u *user) changeEmail(email string) {
	//指向user类型的指针 来调用notify方法
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
	//(&bill),changeEmail(bill@newdomain.com)
	bill.notify()//根据方法的接收者来使用返回操作及定义返回值类型

	// Pointers of type user can be used to call methods
	// declared with a pointer receiver.
	lisa.changeEmail("lisa@newdomain.com")
	lisa.notify()
}
