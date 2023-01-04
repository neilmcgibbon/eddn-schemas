package v3

import (
	"time"
)

type Schema struct {
	SystemName  string    `json:"systemName"`
	StationName string    `json:"stationName"`
	MarketID    int64     `json:"marketId"`
	IsHorizons  bool      `json:"horizons"`
	IsOdyssey   bool      `json:"odyssey"`
	Timestamp   time.Time `json:"timestamp"`
	Commodities []Item    `json:"commodities"`
	Economies   []Economy `json:"economies"`
	Prohibited  []string  `json:"prohibited"`
}

type Item struct {
	BuyPrice         int32 `json:"buyPrice"`
	Demand           int32 `json:"demand"`
	DemandBracket    int
	MeanPrice        int32       `json:"meanPrice"`
	Name             string      `json:"name"`
	RawDemandBracket interface{} `json:"demandBracket"`
	RawStockbracket  interface{} `json:"stockBracket"`
	SellPrice        int32       `json:"sellPrice"`
	StatusFlags      []string    `json:"statusFlags"`
	Stock            int32       `json:"stock"`
	StockBracket     int
}
type Economy struct {
	Name       string  `json:"name"`
	Proportion float32 `json:"proportion"`
}

func (s *Schema) PostDecode() {
	// According to docs, StockBracket and DemandBracket can be of type int or string:
	// 		"A value of \"\" indicates that the commodity is not normally sold/purchased at this
	//		station, but is currently temporarily for sale/purchase"
	// So we unmarhall into an interface{} which will give us float64 or string, placing the value in
	// "RawStockBraket" and "RawDemandBracket". This PostDecode hook then places a true int value into
	// StockBracket and Demandbracket, changing the "" (empty string) identidier to -1

	for i, v := range s.Commodities {

		if _, ok := v.RawDemandBracket.(float64); ok {
			s.Commodities[i].DemandBracket = int(v.RawDemandBracket.(float64))
		} else {
			s.Commodities[i].DemandBracket = -1
		}

		if _, ok := v.RawStockbracket.(float64); ok {
			s.Commodities[i].StockBracket = int(v.RawStockbracket.(float64))
		} else {
			s.Commodities[i].StockBracket = -1
		}
	}
}

func (s *Schema) GetSchema() string {
	return "https://eddn.edcd.io/schemas/commodity/3"
}
