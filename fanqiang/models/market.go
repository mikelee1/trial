package models

type KLine struct {
	Id     int64
	Time   string
	Amount float32
	Count  int
	Open   float32
	Close  float32
	Low    float32
	High   float32
	Vol    float64
}

type MarketInfo struct {
	Data   []*KLine
	Status string
	Ch     string
	Ts     int64
}
