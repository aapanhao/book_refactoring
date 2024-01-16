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

func statement(invoice CustomerInvoice, plays map[string]Play) string {

	totalAmount := 0
	volumeCredits := 0
	result := fmt.Sprintf("Statement for %s\n", invoice.Customer)
	for _, perf := range invoice.Performances {
		thisAmount := 0
		play := plays[perf.PlayID]

		switch play.PlayType {
		case "tragedy":
			thisAmount = 40000
			if perf.Audience > 30 {
				thisAmount += 1000 * (perf.Audience - 30)
			}
		case "comedy":
			thisAmount = 30000
			if perf.Audience > 20 {
				thisAmount += 10000 + 500*(perf.Audience-20)
			}
			thisAmount += 300 * perf.Audience
		default:
			panic(fmt.Sprintf("unknown type %s", play.PlayType))
		}

		volumeCredits += Max(perf.Audience-30, 0)
		if "comedy" == play.PlayType {
			volumeCredits += perf.Audience / 5
		}

		result += fmt.Sprintf("  %s: %s (%d seats)\n", play.Name, format(thisAmount/100), perf.Audience)
		totalAmount += thisAmount
	}

	result += fmt.Sprintf("Amount owed is %s\n", format(totalAmount/100))
	result += fmt.Sprintf("You earned %d credits\n", volumeCredits)

	return result
}
