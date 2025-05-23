package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"strconv"
	"strings"
)

func main() {
	a := app.New()
	w := a.NewWindow("Factuurformulier")
	var items [][]string
	// Bedrijfsgegevens
	bedrijfNaam := widget.NewEntry()
	straatHuisnummer := widget.NewEntry()
	postcodeWoonplaats := widget.NewEntry()

	// Factuurgegevens
	factuurNummer := widget.NewEntry()
	factuurDatum := widget.NewEntry()
	verloopTermijn := widget.NewEntry()
	kvkNummer := widget.NewEntry()
	vervalDatum := widget.NewLabel("Automatisch berekend")

	// Itemgegevens
	itemNaam := widget.NewEntry()
	aantal := widget.NewEntry()
	btwTarief := widget.NewSelect([]string{"0%", "9%", "21%"}, nil)
	prijsPerStukExcl := widget.NewEntry()
	prijsPerStukIncl := widget.NewEntry()
	totaalExcl := widget.NewEntry()
	totaalIncl := widget.NewEntry()
	bedrijfInfoText := canvas.NewText("Bedrijfsgegevens:", color.White)
	bedrijfInfoText.TextStyle = fyne.TextStyle{Bold: true}
	generaalText := canvas.NewText("Factuurgegevens:", color.White)
	generaalText.TextStyle = fyne.TextStyle{Bold: true}
	totaalText := canvas.NewText("Totalen:", color.White)
	totaalText.TextStyle = fyne.TextStyle{Bold: true}
	itemText := canvas.NewText("Items:", color.White)
	itemText.TextStyle = fyne.TextStyle{Bold: true}
	// Totale bedragen
	totaalExclBtw := widget.NewEntry()
	totaalBtw := widget.NewEntry()
	totaalInclBtw := widget.NewEntry()

	bedrijfInfoForm := widget.NewForm(
		widget.NewFormItem("Bedrijfsnaam", bedrijfNaam),
		widget.NewFormItem("Straat + Huisnummer", straatHuisnummer),
		widget.NewFormItem("Postcode + Woonplaats", postcodeWoonplaats),
	)
	factuurForm := widget.NewForm(
		widget.NewFormItem("Factuurnummer", factuurNummer),
		widget.NewFormItem("Factuurdatum", factuurDatum),
		widget.NewFormItem("Verlooptermijn (dagen)", verloopTermijn),
		widget.NewFormItem("Vervaldatum", vervalDatum),
		widget.NewFormItem("KvK-nummer", kvkNummer),
	)
	itemForm := widget.NewForm(
		widget.NewFormItem("Itemnaam", itemNaam),
		widget.NewFormItem("Aantal", aantal),
		widget.NewFormItem("BTW-tarief", btwTarief),
		widget.NewFormItem("Prijs/stuk (excl. btw)", prijsPerStukExcl),
		widget.NewFormItem("Prijs/stuk (incl. btw)", prijsPerStukIncl),

		widget.NewFormItem("Prijs totaal (excl. btw)", totaalExcl),
		widget.NewFormItem("Prijs totaal (incl. btw)", totaalIncl),
	)
	addItemButton := widget.NewButton("Voeg item toe", func() {
		btwStr := strings.TrimSuffix(btwTarief.Selected, "%")
		btwPercentage, err := strconv.ParseFloat(btwStr, 64)
		if err != nil {
			fmt.Println("Invalid input for BTW percentage:", btwTarief.Selected)
			return
		}

		// Parse price per unit (excl. btw)
		prijsFloat, err := strconv.ParseFloat(prijsPerStukExcl.Text, 64)
		if err != nil {
			fmt.Println("Invalid input for price per unit:", prijsPerStukExcl.Text)
			return
		}

		prijsIncl := prijsFloat * (1 + btwPercentage/100)
		prijsInclStr := fmt.Sprintf("%.2f", prijsIncl)
		prijsPerStukIncl.SetText(prijsInclStr)
		fmt.Println("Price incl. btw:", prijsInclStr)
		// Optional: calculate total prices (if aantal is numeric)
		aantalFloat, err := strconv.ParseFloat(aantal.Text, 64)
		if err != nil {
			fmt.Println("Invalid input for aantal:", aantal.Text)
			return
		}
		totaalExclVal := prijsFloat * aantalFloat
		totaalInclVal := prijsIncl * aantalFloat
		totaalExcl.SetText(fmt.Sprintf("%.2f", totaalExclVal))
		totaalIncl.SetText(fmt.Sprintf("%.2f", totaalInclVal))

		item := []string{
			itemNaam.Text,
			aantal.Text,
			btwTarief.Selected,
			prijsPerStukExcl.Text,
			prijsInclStr,
			totaalExcl.Text,
			totaalIncl.Text,
		}
		items = append(items, item)

		// Clear the input fields
		itemNaam.SetText("")
		aantal.SetText("")
		btwTarief.SetSelected("0%")
		prijsPerStukExcl.SetText("")
		prijsPerStukIncl.SetText("")
		totaalExcl.SetText("")
		totaalIncl.SetText("")
		fmt.Print("Succesfully added item: ", item)
	})
	// Create the totals section as another form (after bold label)
	totalenForm := widget.NewForm(
		widget.NewFormItem("Totaal excl. btw", totaalExclBtw),
		widget.NewFormItem("Totaal btw-bedrag", totaalBtw),
		widget.NewFormItem("Totaal incl. btw", totaalInclBtw),
	)

	// Combine everything into a vertical container
	form := container.NewVBox(
		bedrijfInfoText,
		bedrijfInfoForm,
		generaalText,
		factuurForm,
		itemText,
		itemForm,
		addItemButton,
		totaalText, // bold label
		totalenForm,
	)

	// Optional: update vervaldatum when verlooptermijn changes
	verloopTermijn.OnChanged = func(s string) {
		if days, err := strconv.Atoi(s); err == nil {
			vervalDatum.SetText("+" + strconv.Itoa(days) + " dagen na factuurdatum")
		} else {
			vervalDatum.SetText("Ongeldige invoer")
		}
	}

	w.SetContent(form)
	w.Resize(fyne.NewSize(600, 700))
	w.ShowAndRun()
}
