package main

import (
	"testing"
)

func TestInvoiceInfo(t *testing.T) {
	actuallyResult := getStatement("invoices.json", "play.json")
	expectResult := `Statement for BigCo
  Hamlet: $650.00 (55 seats)
  As You Like It: $580.00 (35 seats)
  Othello: $500.00 (40 seats)
Amount owed is $1,730.00
You earned 47 credits`
	if actuallyResult != expectResult {
		t.Fatalf("actuallyResult: (%s) not equal expectResult: (%s)", actuallyResult, expectResult)
	}
}

func TestCreateInvoiceInfo(t *testing.T) {
	InitInvoice()
}
