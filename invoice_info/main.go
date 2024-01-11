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

func InitInvoice() {
	hamletPerformance := map[string]any{"playID": "hamlet", "audience": 55}
	asLikePerformance := map[string]any{"playID": "as-like", "audience": 35}
	othelloPerformance := map[string]any{"playID": "othello", "audience": 40}

	performances := []map[string]any{hamletPerformance, asLikePerformance, othelloPerformance}

	customerInvoice := map[string]any{"customer": "BigCo", "performances": performances}

	allInvoice := []map[string]any{customerInvoice}

	fmt.Println(allInvoice)
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

func InitInvoice2() {
	p1 := Performance{playID: "hamlet", audience: 55}
	p2 := Performance{playID: "as-like", audience: 35}
	p3 := Performance{playID: "othello", audience: 40}

	c1 := CustomerInvoice{customer: "BigCo", performance: []Performance{p1, p2, p3}}

	a1 := AllInvoice{customerInvoices: []CustomerInvoice{c1}}
	fmt.Println(a1)
}
