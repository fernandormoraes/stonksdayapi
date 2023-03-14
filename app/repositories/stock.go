package repositories

import (
	"math/big"
	"regexp"
	"strings"

	"github.com/anaskhan96/soup"
	"github.com/fernandormoraes/stonksdayapi/app/models"
	"github.com/sirupsen/logrus"
)

type IStockRepository interface {
	GetStock(stockName string) (models.Stock, error)
}

type StockRepository struct {
}

func NewStockRepository() *StockRepository {
	return &StockRepository{}
}

const url string = "https://www.marketwatch.com/investing/stock/"

func (r StockRepository) GetStock(stockName string) (models.Stock, error) {
	page, err := soup.Get(url + stockName + "?mod=quote_search")

	if err != nil {
		logrus.Errorf(err.Error())
	}

	doc := soup.HTMLParse(page)

	price := getActualPrice(doc)

	rat := big.NewRat(1, 1)

	if _, ok := rat.SetString(price); !ok {
		logrus.Println("Failed to parse the string!")
	}

	changePoint := getChangePoint(doc)

	changePercentage := getChangePercentage(doc)

	totalVolDoc := getVolDoc(doc)

	rangeLower, rangeHigher := getRangeLowerAndHigher(doc.FindAll("div", "class", "range__header"))

	range52WeekLower, range52WeekHigher := get52WeekRangeLowerAndHigher(doc.FindAll("div", "class", "range__header"))

	return models.Stock{
		Name:              getName(doc),
		Price:             rat.FloatString(2),
		ChangePoint:       changePoint,
		ChangePercentage:  changePercentage,
		TotalVol:          totalVolDoc,
		DayRangeLower:     rangeLower,
		DayRangeHigher:    rangeHigher,
		Week52RangeLower:  range52WeekLower,
		Week52RangeHigher: range52WeekHigher,
	}, nil
}

func getVolDoc(root soup.Root) string {
	totalVolDoc := root.Find("div", "class", "range__header").Find("span", "class", "primary").Text()

	matched, err := regexp.MatchString("\\d+[.]*\\d*[a-zA-Z]*", totalVolDoc)

	if err != nil {
		return err.Error()
	}

	if matched {
		r, _ := regexp.Compile(`\d+[.]*\d*[a-zA-Z]*`)

		totalVolDoc = r.FindString(totalVolDoc)
	}

	return totalVolDoc
}

func getChangePercentage(root soup.Root) string {
	changePercentageDoc := root.Find("span", "class", "change--percent--q")

	changePercentageStr := strings.Join(strings.Split(changePercentageDoc.Find("bg-quote", "field", "percentchange").Text(), ","), "")

	changePercentage := changePercentageStr[:len(changePercentageStr)-1]

	return changePercentage
}

func getChangePoint(root soup.Root) string {
	changePointDoc := root.Find("span", "class", "change--point--q")

	changePoint := strings.Join(strings.Split(changePointDoc.Find("bg-quote", "field", "change").Text(), ","), "")

	return changePoint
}

func getActualPrice(root soup.Root) string {
	intraDayPriceClass := root.Find("h2", "class", "intraday__price")

	value := intraDayPriceClass.Find("bg-quote", "class", "value")

	price := strings.Join(strings.Split(value.Text(), ","), "")

	return price
}

func getName(root soup.Root) string {
	return root.Find("h1", "class", "company__name").Text()
}

func getRangeLowerAndHigher(listRoot []soup.Root) (string, string) {
	lower := listRoot[1].FindAll("span", "class", "primary")[0].Text()
	higher := listRoot[1].FindAll("span", "class", "primary")[1].Text()

	return lower, higher
}

func get52WeekRangeLowerAndHigher(listRoot []soup.Root) (string, string) {
	lower := listRoot[2].FindAll("span", "class", "primary")[0].Text()
	higher := listRoot[2].FindAll("span", "class", "primary")[1].Text()

	return lower, higher
}
