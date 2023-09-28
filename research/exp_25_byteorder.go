package main

import (
	"encoding/binary"
	"fmt"
	"unsafe"
)

type GSFrameCodecConfig struct {
	StartDelimiterOffset int
	SendSequenceOffset   int
	RecvSequenceOffset   int
	FrameTypeOffset      int
	DataLengthOffset     int
	DataOffset           int
	EndDelimiterOffset   int
	FrameDelimiter       uint16
}

func GetGSFrameCodecConfig() *GSFrameCodecConfig {
	return &GSFrameCodecConfig{
		StartDelimiterOffset: 0,
		SendSequenceOffset:   2,
		RecvSequenceOffset:   10,
		FrameTypeOffset:      18,
		DataLengthOffset:     19,
		DataOffset:           15,
		FrameDelimiter:       0xEB90,
	}
}

type GSFrameCodec struct {
	StartTag uint16
	SendSeq  uint64 // inc
	RecvSeq  uint64 // inc
	Type     byte   // 0x00 request 0x01 response
	Length   uint32
	Data     []byte
	EndTag   uint16

	config *GSFrameCodecConfig
}

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
	var c byte = 0
	fmt.Println(unsafe.Sizeof(c))

	x := &GSFrameCodec{}

	fmt.Println("offset startTag :", unsafe.Offsetof(x.StartTag))
	fmt.Println("offset sendSeq:", unsafe.Offsetof(x.SendSeq))
	fmt.Println("offset recvSeq:", unsafe.Offsetof(x.RecvSeq))
	fmt.Println("offset type:", unsafe.Offsetof(x.Type))
	fmt.Println("offset recvSeq:", unsafe.Offsetof(x.Length))
	fmt.Println("offset recvSeq:", unsafe.Offsetof(x.EndTag))

}
