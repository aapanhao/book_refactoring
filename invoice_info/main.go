package invoice_info

import (
	"fmt"
)

func statement(invoice CustomerInvoice, plays map[string]Play) string {
	return renderPlainText(CreateStatementData(invoice, plays))
}

func renderPlainText(data StatementData) string {

	result := fmt.Sprintf("Statement for %s\n", data.Customer)
	for _, performance := range data.NewPerformances {
		result += fmt.Sprintf("  %s: %s (%d seats)\n", performance.Name,
			Usd(performance.Amount), performance.Audience)
	}

	result += fmt.Sprintf("Amount owed is %s\n", Usd(data.TotalAmount))
	result += fmt.Sprintf("You earned %d credits\n", data.TotalVolumeCredits)

	return result
}
