package helpers

import (
	"log"
	"reflect"
	"testing"
)

func TestString2Bytes(t *testing.T) {

	str1 := "Hello world"
	bytes := String2Bytes(str1)

	str2 := string(bytes)
	log.Printf("str1: %s, str2: %s\n", str1, str2)

}

func TestBytes2String(t *testing.T) {
	str1 := "Hello world"
	bytes := []byte(str1)
	str2 := Bytes2StrPnt(bytes)

	log.Println("type:", reflect.TypeOf(str2))
	str3 := *str2
	log.Println("type 3 : ", reflect.TypeOf(str3))

	log.Println("str2 =>", str2)
	log.Println("len:", len(str3))

}