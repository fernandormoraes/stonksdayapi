package models

type Stock struct {
	Name              string `json:"name"`
	Price             string `json:"price"`
	ChangePoint       string `json:"changePoint"`
	ChangePercentage  string `json:"changePercentage"`
	TotalVol          string `json:"totalVol"`
	DayRangeLower     string `json:"dayRangeLower"`
	DayRangeHigher    string `json:"dayRangeHigher"`
	Week52RangeLower  string `json:"week52RangeLower"`
	Week52RangeHigher string `json:"week52RangeHigher"`
}
