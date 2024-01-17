package invoice_info

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func Usd(currency int) string {
	printCurrency := message.NewPrinter(language.English)
	floatCurrency := float64(currency)
	return printCurrency.Sprintf("$%.2f", floatCurrency/100)
}
func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
