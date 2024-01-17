package invoice_info

import (
	"fmt"
)

type NewPerformance struct {
	Performance
	Play
	Amount        int
	VolumeCredits int
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

		calculator := createPerformanceCalculator(newPerformance.PlayType)
		newPerformance.Amount = calculator.Amount(newPerformance.Audience)
		newPerformance.VolumeCredits = calculator.VolumeCredit(newPerformance.Audience)
		newPerformances = append(newPerformances, newPerformance)
	}
	return newPerformances
}

func playFor(performance Performance, plays map[string]Play) Play {
	return plays[performance.PlayId]
}

func createPerformanceCalculator(playType string) Calculator {
	switch playType {
	case "tragedy":
		return TragedyCalculator{}
	case "comedy":
		return ComedyCalculator{}
	default:
		panic(fmt.Sprintf("unknown type %s", playType))
	}
}

type Calculator interface {
	Amount(int) int
	VolumeCredit(int) int
}

type TragedyCalculator struct {
}

func (TragedyCalculator) Amount(audience int) int {
	result := 40000
	if audience > 30 {
		result += 1000 * (audience - 30)
	}
	return result
}
func (TragedyCalculator) VolumeCredit(audience int) int {
	result := 0
	result += Max(audience-30, 0)
	return result
}

type ComedyCalculator struct {
}

func (ComedyCalculator) Amount(audience int) int {
	result := 30000
	if audience > 20 {
		result += 10000 + 500*(audience-20)
	}
	result += 300 * audience

	return result
}
func (ComedyCalculator) VolumeCredit(audience int) int {
	result := 0
	result += Max(audience-30, 0)
	result += audience / 5
	return result
}
func totalAmount(data StatementData) int {
	result := 0
	for _, performance := range data.NewPerformances {
		result += performance.Amount
	}
	return result
}

func totalVolumeCredits(data StatementData) int {
	result := 0
	for _, performance := range data.NewPerformances {
		result += performance.VolumeCredits
	}
	return result
}
