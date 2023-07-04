package main

import (
  "fmt"
  "research/errorcode"
)

//go:generate echo GOGOGO!
//go:generate go run  exp_13_generate.go
//go:generate echo $GOARCH $GOOS $GOFILE $GOLINE $GOPACKAGE

func main() {

  fmt.Println("go run exp_13_generate.go")
  fmt.Println(errorcode.ErrorCode(errorcode.OK).String())
  }
