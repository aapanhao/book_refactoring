package invoice_info

import (
	"fmt"
)

type NewPerformance struct {
	Performance
	Play
	Amount int
}

type StatementData struct {
	Customer           string
	NewPerformances    []NewPerformance
	TotalAmount        int
	TotalVolumeCredits int
}

func CreateStatementData(invoice CustomerInvoice, plays map[string]Play) StatementData {
	var statementData StatementData
	statementData.Customer = invoice.Customer
	statementData.NewPerformances = addPerformancePlayInfo(invoice.Performances, plays)
	statementData.TotalAmount = totalAmount(statementData)
	statementData.TotalVolumeCredits = totalVolumeCredits(statementData)
	return statementData
}

func addPerformancePlayInfo(performances []Performance, plays map[string]Play) []NewPerformance {

	newPerformances := make([]NewPerformance, 0, cap(performances))
	for _, perf := range performances {
		newPerformance := NewPerformance{Performance: perf, Play: playFor(perf, plays)}
		newPerformance.Amount = amountFor(newPerformance)
		newPerformances = append(newPerformances, newPerformance)
	}
	return newPerformances
}

func playFor(performance Performance, plays map[string]Play) Play {
	return plays[performance.PlayId]
}

func amountFor(performance NewPerformance) int {
	result := 0

	switch performance.PlayType {
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
		panic(fmt.Sprintf("unknown type %s", performance.PlayType))
	}
	return result
}

func volumeCreditsFor(performance NewPerformance) int {
	result := 0
	result += Max(performance.Audience-30, 0)
	if "comedy" == performance.PlayType {
		result += performance.Audience / 5
	}
	return result
}

func totalAmount(data StatementData) int {
	result := 0
	for _, performance := range data.NewPerformances {
		result += amountFor(performance)
	}
	return result
}

func totalVolumeCredits(data StatementData) int {
	result := 0
	for _, performance := range data.NewPerformances {
		result += volumeCreditsFor(performance)
	}
	return result
}
