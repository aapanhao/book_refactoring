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

func format(currency int) string {
	printCurrency := message.NewPrinter(language.English)
	floatCurrency := float64(currency)
	return printCurrency.Sprintf("$%.2f", floatCurrency)
}

func amountFor(performance Performance, play Play) int {
	result := 0

	switch play.PlayType {
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
		panic(fmt.Sprintf("unknown type %s", play.PlayType))
	}
	return result
}

func volumeCredit(performance Performance, play Play) int {
	result := 0
	result += Max(performance.Audience-30, 0)
	if "comedy" == play.PlayType {
		result += performance.Audience / 5
	}
	return result
}

func statement(invoice CustomerInvoice, plays map[string]Play) string {

	totalAmount := 0
	volumeCredits := 0
	result := fmt.Sprintf("Statement for %s\n", invoice.Customer)

	for _, performance := range invoice.Performances {
		play := plays[performance.PlayID]
		totalAmount += amountFor(performance, play)
		volumeCredits += volumeCredit(performance, play)
		result += fmt.Sprintf("  %s: %s (%d seats)\n", play.Name, format(amountFor(performance, play)/100), performance.Audience)
	}

	result += fmt.Sprintf("Amount owed is %s\n", format(totalAmount/100))
	result += fmt.Sprintf("You earned %d credits\n", volumeCredits)

	return result
}
