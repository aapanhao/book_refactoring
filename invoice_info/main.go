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

func htmlStatement(invoice CustomerInvoice, plays map[string]Play) string {
	return renderHtml(CreateStatementData(invoice, plays))
}

func renderHtml(data StatementData) string {
	result := fmt.Sprintf(`<h1>Statement for %s</h1>\n`, data.Customer)
	result += "<table>\n"
	result +=
		"<tr><th>play</th><th>seats</th><th>cost</th></tr>"
	for _, perf := range data.NewPerformances {
		result += fmt.Sprintf(` <tr><td>%s</td><td>%d</td>`, perf.Name, perf.Audience)
		result += fmt.Sprintf(`<td>%s</td></tr>\n`, Usd(perf.Amount))
	}
	result += "</table>\n"
	result += fmt.Sprintf(`<p>Amount owed is <em>%d</em></p>\n`, data.TotalAmount)
	result += fmt.Sprintf(`<p>You earned <em>%d</em> credits</p>\n`, data.TotalVolumeCredits)
	return result
}
