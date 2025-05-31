package helpers

import (
	"fmt"
	"strconv"
	"time"
)

func CalculateExperationDate(factuurDatum string, verloopTermijn string) string {
	date, err := time.Parse("2006-01-02", factuurDatum)
	if err != nil {
		fmt.Println("error parsing date:" + err.Error())
	}
	int, er := strconv.Atoi(verloopTermijn)
	if er != nil {
		fmt.Println("error parsing date:" + er.Error())
	}
	expirationDate := date.AddDate(0, 0, int)
	return expirationDate.Format("2006-01-02")
}

func CalculateTotalExcl(count int, prices []string) string {
	countFloat := float64(count)
	totalPrijs := 0.0
	for _, price := range prices {
		priceFloat, err := strconv.ParseFloat(price, 64)
		if err != nil {
			continue // skip invalid prices
		}
		totalPrijs += priceFloat
	}
	total := countFloat * totalPrijs
	totalString := strconv.FormatFloat(total, 'f', 2, 64)
	return totalString
}
func CalculateTotalIncl(count int, prices []string, btwAmounts []string) string {
	countFloat := float64(count)
	totalPrijs := 0.0
	for i, price := range prices {
		priceFloat, err := strconv.ParseFloat(price, 64)
		if err != nil {
			continue
		}
		btwAmountFloat, err := strconv.ParseFloat(btwAmounts[i], 64)
		if err != nil {
			continue
		}
		totalPrijs += priceFloat + btwAmountFloat
	}
	total := countFloat * totalPrijs
	totalString := strconv.FormatFloat(total, 'f', 2, 64)
	return totalString
}
func CaluclateTotalBtw(btwPrices []string) string {
	totalBtw := 0.0
	for _, price := range btwPrices {
		priceFloat, err := strconv.ParseFloat(price, 64)
		if err != nil {
			continue // skip invalid prices
		}
		totalBtw += priceFloat
	}
	totalBtwString := strconv.FormatFloat(totalBtw, 'f', 2, 64)
	fmt.Println("Total BTW: " + totalBtwString)
	return totalBtwString
}
