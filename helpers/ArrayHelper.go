package helpers

import (
	"fmt"
	"strconv"
	"time"
)

func SetBusinessInfo(bedrijfNaam string, straatHuisnummer string, postcodeWoonplaats string) []string {
	bedrijfInfo := []string{
		"Bedrijfgegevens:",
		"Bedrijfsnaam: " + bedrijfNaam,
		"Straat + Huisnummer: " + straatHuisnummer,
		"Postcode + Woonplaats: " + postcodeWoonplaats,
	}
	return bedrijfInfo
}
func SetInvoiceInfo(factuurNummer string, factuurDatum string, verloopTermijn string, kvkNummer string) []string {
	invoiceInfo := []string{
		"Factuurgegevens:",
		"Factuurnummer: " + factuurNummer,
		"Factuurdatum: " + factuurDatum,
		"Verlooptermijn: " + verloopTermijn,
		"Vervaldatum: " + CalculateExperationDate(factuurDatum, verloopTermijn),
		"KvK-nummer: " + kvkNummer,
	}
	return invoiceInfo
}
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
