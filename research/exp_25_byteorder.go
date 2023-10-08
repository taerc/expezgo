package main

import (
	"bytes"
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
	// Data     *byte
	// EndTag   uint16

	// config *GSFrameCodecConfig
}

func ednian() {

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
	// fmt.Println("offset recvSeq:", unsafe.Offsetof(x.EndTag))

}

func packCase() {
	frame := GSFrameCodec{}
	frame.StartTag = 0xEB90
	frame.SendSeq = 1
	frame.RecvSeq = 1
	frame.Type = 0
	frame.Length = 8
	// frame.Data = make([]byte, 8)
	// frame.EndTag = 0xEB90
	data := make([]byte, 16)
	for i := 0; i < 16; i++ {
		data[i] = byte(int('A') + i)
	}

	buf := &bytes.Buffer{}
	err := binary.Write(buf, binary.LittleEndian, &frame)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = binary.Write(buf, binary.LittleEndian, data)
	if err != nil {
		fmt.Println("000")
		fmt.Println(err.Error())
	}
	err = binary.Write(buf, binary.LittleEndian, frame.StartTag)
	if err != nil {
		fmt.Println("1111")
		fmt.Println(err.Error())
	}

	fmt.Println("len ", buf.Len())

}

func main() {
	packCase()

}
