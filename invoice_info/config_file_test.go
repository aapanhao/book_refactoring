package main

import (
	"fmt"
	"testing"
)

func TestInitPlayInfo(t *testing.T) {
	res := InitPlay()
	fmt.Println(res)
}

func TestReadInvoiceFile(t *testing.T) {
	obj := ReadInvoiceFile()
	fmt.Println(obj)
}

func TestWriteInvoiceFile(t *testing.T) {
	if WriteInvoiceFile() != true {
		t.Fatal("err")
	}
}
