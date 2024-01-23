package models

import (
	"sort"
	"strconv"
)

type Producer struct {
	Name       string `json:"name"`
	Cost       int    `json:"cost"`
	Production int    `json:"production"`
}

func (p *Producer) SetCost(arg string) {
	newArg, err := strconv.Atoi(arg)
	if err != nil {
		panic(err)
	}
	p.Cost = newArg
}

func (p *Producer) SetProduction(province *Province, arg string) {
	newArg, err := strconv.Atoi(arg)
	if err != nil {
		panic(err)
	}
	province.TotalProduction += newArg - p.Production
	p.Production = newArg
}

type ByCost []Producer

func (b ByCost) Len() int {
	return len(b)
}

func (b ByCost) Less(i, j int) bool {
	return b[i].Cost < b[j].Cost
}

func (b ByCost) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func SortProducers(producers []Producer) {
	sort.Sort(ByCost(producers))
}
