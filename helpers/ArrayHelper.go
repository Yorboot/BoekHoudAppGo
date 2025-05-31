package helpers

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
func SetTotals(PricesExcl []string, PricesIncl []string, btwAmounts []string) []string {
	totals := []string{
		"Totalen:",
		"Totaal btw: " + CaluclateTotalBtw(btwAmounts),
		"Totaal Excl. BTW: " + CalculateTotalExcl(len(PricesExcl), PricesExcl),
		"Totaal Incl. BTW: " + CalculateTotalIncl(len(PricesIncl), PricesIncl, btwAmounts),
	}
	return totals
}
