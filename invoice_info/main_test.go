package invoice_info

import (
	"fmt"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"testing"
)

func TestInvoiceInfo(t *testing.T) {
	expectResult := `Statement for BigCo
  Hamlet: $650.00 (55 seats)
  As You Like It: $580.00 (35 seats)
  Othello: $500.00 (40 seats)
Amount owed is $1,730.00
You earned 47 credits
`

	play := ReadPlayFile()
	invoices := ReadInvoiceFile()
	for _, invoice := range invoices {
		actuallyResult := statement(invoice, play)
		if actuallyResult != expectResult {
			t.Fatalf("actuallyResult: (%s) not equal expectResult: (%s)", actuallyResult, expectResult)
		}
	}
}

func TestDivNumber(t *testing.T) {
	fmt.Println(10 / 3)
	printCurrency := message.NewPrinter(language.English)
	amount := 1234.5678

	// 使用货币格式化
	formattedAmount := printCurrency.Sprintf("%.2f", amount)

	fmt.Printf("Original amount: %.2f\n", amount)
	fmt.Printf("Formatted currency: %s\n", formattedAmount)
}
