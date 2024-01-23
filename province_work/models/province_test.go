package models

import (
	"fmt"
	"testing"
)

func TestReadProvinceProducerData(t *testing.T) {
	provinceProducerData := InitProvinceData()

	fmt.Println(provinceProducerData)

	name := provinceProducerData.Name
	if name != "Asia" {
		t.Errorf("province name expect %s, got %s\n", "Asia", name)
	}

	totalProduction := provinceProducerData.TotalProduction
	if totalProduction != 25 {
		t.Errorf("province totalProduction expect %d, got %d\n", 25, totalProduction)
	}

	shortfall := provinceProducerData.Shortfall()
	if shortfall != 5 {
		t.Errorf("province shortfall expect %d, got %d\n", 5, shortfall)
	}

	profit := provinceProducerData.Profit()
	if profit != 230 {
		t.Errorf("Profit expect %d, get %d", 230, profit)
	}
}
