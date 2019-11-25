// Sample program to show how a bytes.Buffer can also be used
// with the io.Copy function.
package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)
//拼接字符串
// main is the entry point for the application.
func main() {
	var b bytes.Buffer// bytes 包的buffer 类型b用于缓冲数据

	// Write a string to the buffer.
	b.Write([]byte("Hello"))//write 方法写入

	// Use Fprintf to concatenate a string to the Buffer.
	fmt.Fprintf(&b, "World!") //追加 接受io.write 类型
	//bytes.Buffer 类型指针实现了io.Writer接口
	//拼接字符串
	// Write the content of the Buffer to stdout.
	io.Copy(os.Stdout, &b)//写到终端
}

