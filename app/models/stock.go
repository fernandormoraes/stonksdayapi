package models

type Stock struct {
	Name              string
	Price             string
	ChangePoint       string
	ChangePercentage  string
	TotalVol          string
	DayRangeLower     string
	DayRangeHigher    string
	Week52RangeLower  string
	Week52RangeHigher string
}
