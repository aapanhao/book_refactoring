package main

import "fmt"

func statement(invoice map[string]any, plays map[string]map[string]string) string {
	result := ""
	return result
}

func getStatement(invoiceFile string, playFile string) string {
	// 读文件

	// 调用statement，返回
	return ""
}

func readFile() {
	//fileName := "invoices.json"
	//file, err := os.ReadFile(fileName)
	//if err != nil {
	//	panic(err)
	//}
	//json.Unmarshal(file)
}

type AllInvoice struct {
	customerInvoices []CustomerInvoice
}

type CustomerInvoice struct {
	customer    string
	performance []Performance
}

type Performance struct {
	playID   string
	audience int
}

func InitInvoice() {
	p1 := Performance{playID: "hamlet", audience: 55}
	p2 := Performance{playID: "as-like", audience: 35}
	p3 := Performance{playID: "othello", audience: 40}

	c1 := CustomerInvoice{customer: "BigCo", performance: []Performance{p1, p2, p3}}

	a1 := AllInvoice{customerInvoices: []CustomerInvoice{c1}}
	fmt.Println(a1)
}
