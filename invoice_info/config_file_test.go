package invoice_info

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
	fmt.Printf("%T, %v", obj, obj)
}

func TestReadInvoiceFileV2(t *testing.T) {
	obj := ReadInvoiceFileV2()
	fmt.Printf("%T, %v", obj, obj)
}

func TestReadPlayFile(t *testing.T) {
	obj := ReadPlayFile()
	fmt.Printf("%T, %v", obj, obj)
}

func TestWriteInvoiceFile(t *testing.T) {
	if WriteInvoiceFile() != true {
		t.Fatal("err")
	}
}
