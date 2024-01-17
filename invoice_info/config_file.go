package invoice_info

import (
	"encoding/json"
	"fmt"
	"os"
)

type CustomerInvoice struct {
	Customer     string        `json:"customer"`
	Performances []Performance `json:"performances"`
}

type Performance struct {
	PlayId   string `json:"playID"`
	Audience int    `json:"audience"`
}

type Play struct {
	Name     string `json:"name"`
	PlayType string `json:"type"`
}

func ReadFile(filePath string, obj any) {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(obj)
	if err != nil {
		panic(err)
	}
}

func ReadFileV2(filePath string, obj any) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(file, obj)
	if err != nil {
		panic(err)
	}
}

func ReadInvoiceFile() []CustomerInvoice {
	filePath := "invoices.json"
	var obj []CustomerInvoice
	ReadFile(filePath, &obj)
	return obj
}

func ReadInvoiceFileV2() []CustomerInvoice {
	filePath := "invoices.json"
	var obj []CustomerInvoice
	ReadFileV2(filePath, &obj)
	return obj
}

func ReadPlayFile() map[string]Play {
	filePath := "play.json"
	var obj map[string]Play
	ReadFile(filePath, &obj)
	return obj
}

func InitPlay() map[string]Play {
	playInfo := make(map[string]Play)
	playInfo["hamlet"] = Play{Name: "Hamlet", PlayType: "tragedy"}
	playInfo["as-like"] = Play{Name: "As You Like It", PlayType: "comedy"}
	playInfo["othello"] = Play{Name: "Othello", PlayType: "tragedy"}

	return playInfo
}

func InitInvoice() CustomerInvoice {
	p1 := Performance{PlayId: "hamlet", Audience: 55}
	p2 := Performance{PlayId: "as-like", Audience: 35}
	p3 := Performance{PlayId: "othello", Audience: 40}

	return CustomerInvoice{Customer: "BigCo", Performances: []Performance{p1, p2, p3}}
}

func WriteInvoiceFile() bool {
	obj := InitInvoice()
	objList := []CustomerInvoice{obj}
	fmt.Println(objList)
	file, err := os.Create("invoiceNew.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	err = encoder.Encode(objList)
	if err != nil {
		panic(err)
	}
	return true
}
