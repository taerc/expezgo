package main

import (
	"encoding/binary"
	"fmt"
	"unsafe"
)

func main() {

	var data uint64 = 0x1122334455667788

	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(data))
	fmt.Println("bigendian:")
	for i, v := range buf {
		fmt.Println(fmt.Sprintf("%d : 0x%02x", i, v))
	}

	binary.LittleEndian.PutUint64(buf, uint64(data))
	fmt.Println("little:")
	for i, v := range buf {
		fmt.Println(fmt.Sprintf("%d : 0x%02x", i, v))
	}

	fmt.Println("testing ...")

	z := unsafe.Pointer(&data)
	b := *(*byte)(unsafe.Pointer(z))
	fmt.Printf("==%02x\n", b)
	b = *(*byte)(unsafe.Pointer(unsafe.Add(z, 7)))
	fmt.Printf("==%02x\n", b)

	unsafe.Sizeof(byte)

}
