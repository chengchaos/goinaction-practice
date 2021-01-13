package p8

import (
	"bytes"
	"fmt"
	"os"
)

func PracticeWriter() {

	// 创建一个 Buffer 值，病将一个字符串写入 Buffer 中国年
	// 使用实现了 io.Writer 的 Writer 方法
	var b bytes.Buffer
	b.Write([]byte("Hello "))

	// 使用 FPrintf 就爱那个一个字符串拼接到 Buffer 中
	// 将 bytes.Buffer 的地址作为 io.Writer 类型值传入
	fmt.Fprintf(&b, "world!")

	// 将 Buffer 中的内容输出到标准输出设备
	// 将 os.File 值的地址作为 io.Writer 类型值传入
	b.WriteTo(os.Stdout)
}
