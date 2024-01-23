package models

import (
	"encoding/json"
	"os"
	"province_work"
	"strconv"
)

type Province struct {
	Name            string     `json:"name"`
	Producers       []Producer `json:"producers"`
	Demand          int        `json:"demand"`
	Price           int        `json:"price"`
	TotalProduction int
}

func (p *Province) SetDemand(demand string) {
	newDemand, err := strconv.Atoi(demand)
	if err != nil {
		panic(err)
	}
	p.Demand = newDemand
}

func (p *Province) SetPrice(price string) {
	newPrice, err := strconv.Atoi(price)
	if err != nil {
		panic(err)
	}
	p.Price = newPrice
}

func (p *Province) SetTotalProduction(arg int) {
	p.TotalProduction = arg
}

func (p *Province) Shortfall() int {
	return p.Demand - p.TotalProduction
}

func (p *Province) Profit() int {

	return p.DemandPrice() - p.DemandCost()
}

func (p *Province) DemandPrice() int {
	return p.SupportProduction() * p.Price
}

func (p *Province) SupportProduction() int {
	return province_work.Min(p.Demand, p.TotalProduction)
}

func (p *Province) DemandCost() int {
	SortProducers(p.Producers)
	remindDemand := p.Demand

	result := 0
	for _, p := range p.Producers {
		if remindDemand > p.Production {
			result += p.Production * p.Cost
			remindDemand -= p.Production
		} else {
			result += remindDemand * p.Cost
			break
		}
	}
	return result
}

func InitProvinceData() Province {
	dataPath := "data.json"
	province := ReadProvinceProducerData(dataPath)

	for _, producers := range province.Producers {
		province.TotalProduction += producers.Production
	}
	return province
}

func ReadProvinceProducerData(filePath string) Province {
	file, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	var province Province
	err = json.Unmarshal(file, &province)
	if err != nil {
		panic(err)
	}
	return province
}
