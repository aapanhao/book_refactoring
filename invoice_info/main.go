package invoice_info

import (
	"fmt"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func usd(currency int) string {
	printCurrency := message.NewPrinter(language.English)
	floatCurrency := float64(currency)
	return printCurrency.Sprintf("$%.2f", floatCurrency/100)
}

func playFor(performance Performance, plays map[string]Play) Play {
	return plays[performance.PlayID]
}

func amountFor(performance Performance, plays map[string]Play) int {
	result := 0

	switch playFor(performance, plays).PlayType {
	case "tragedy":
		result = 40000
		if performance.Audience > 30 {
			result += 1000 * (performance.Audience - 30)
		}
	case "comedy":
		result = 30000
		if performance.Audience > 20 {
			result += 10000 + 500*(performance.Audience-20)
		}
		result += 300 * performance.Audience
	default:
		panic(fmt.Sprintf("unknown type %s", playFor(performance, plays).PlayType))
	}
	return result
}

func volumeCreditsFor(performance Performance, plays map[string]Play) int {
	result := 0
	result += Max(performance.Audience-30, 0)
	if "comedy" == playFor(performance, plays).PlayType {
		result += performance.Audience / 5
	}
	return result
}

func totalVolumeCredits(data statementData, plays map[string]Play) int {
	result := 0
	for _, performance := range data.Performances {
		result += volumeCreditsFor(performance, plays)
	}
	return result
}

func totalAmount(data statementData, plays map[string]Play) int {
	result := 0
	for _, performance := range data.Performances {
		result += amountFor(performance, plays)
	}
	return result

}

type statementData struct {
	Customer     string
	Performances []Performance
}

func statement(invoice CustomerInvoice, plays map[string]Play) string {
	var statementData statementData
	statementData.Customer = invoice.Customer
	statementData.Performances = invoice.Performances
	return renderPlainText(statementData, plays)
}

func renderPlainText(data statementData, plays map[string]Play) string {

	result := fmt.Sprintf("Statement for %s\n", data.Customer)
	for _, performance := range data.Performances {
		result += fmt.Sprintf("  %s: %s (%d seats)\n", playFor(performance, plays).Name,
			usd(amountFor(performance, plays)), performance.Audience)
	}

	result += fmt.Sprintf("Amount owed is %s\n", usd(totalAmount(data, plays)))
	result += fmt.Sprintf("You earned %d credits\n", totalVolumeCredits(data, plays))

	return result
}
