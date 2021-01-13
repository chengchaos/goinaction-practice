package helpers

import (
	"reflect"
	"unsafe"
)

/*
使用指针的方式获取 string 中的 bytes
看这里：https://zhuanlan.zhihu.com/p/270626496
 */
func String2Bytes(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len: sh.Len,
		Cap: sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}

func Bytes2StrPnt(b []byte) *string {
	return (*string)(unsafe.Pointer(&b))
}