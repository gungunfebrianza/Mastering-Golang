package main

import (
	"fmt"
	"runtime"
	"strconv"
	"unsafe"
)

func main() {
	fmt.Println(runtime.Compiler, runtime.GOARCH, runtime.GOOS)
	fmt.Println(strconv.IntSize)

	sizeInt := int(1999)
	sizeUint := uint(1993)
	sizeInt8 := int8(10)
	sizeUint8 := uint8(10)
	sizeInt32 := int32(20)
	sizeUInt32 := uint32(20)
	sizeInt64 := int64(123)
	sizeUint64 := uint64(123)
	c := "foo"
	d := struct {
		FieldA float32
		FieldB string
	}{0, "bar"}

	fmt.Printf("Size : %T, %d bytes \n", sizeInt, unsafe.Sizeof(sizeInt))       // 8 bytes
	fmt.Printf("Size : %T, %d bytes \n", sizeUint, unsafe.Sizeof(sizeUint))     // 8 bytes
	fmt.Printf("Size : %T, %d bytes \n", sizeInt8, unsafe.Sizeof(sizeInt8))     // 1 bytes
	fmt.Printf("Size : %T, %d bytes \n", sizeUint8, unsafe.Sizeof(sizeUint8))   // 1 bytes
	fmt.Printf("Size : %T, %d bytes \n", sizeInt32, unsafe.Sizeof(sizeInt32))   // 4 bytes
	fmt.Printf("Size : %T, %d bytes \n", sizeUInt32, unsafe.Sizeof(sizeUInt32)) // 4 bytes
	fmt.Printf("Size : %T, %d bytes \n", sizeInt64, unsafe.Sizeof(sizeInt64))   // 8 bytes
	fmt.Printf("Size : %T, %d bytes \n", sizeUint64, unsafe.Sizeof(sizeUint64)) // 8 bytes

	fmt.Printf("c: %T, %d\n", c, unsafe.Sizeof(c))
	fmt.Printf("d: %T, %d\n", d, unsafe.Sizeof(d))
}
