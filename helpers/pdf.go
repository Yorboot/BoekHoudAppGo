package helpers

import (
	"fmt"
	"github.com/signintech/gopdf"
	"os"
)

func GeneratePdf(companyInfo []string, invoiceInfo []string, items [][]string, totals []string) error {
	pdf := &gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	lineHeight := 20.0
	fontPath := "./fonts/OpenSans-Regular.ttf"
	fontName := "OpenSans"
	if _, err := os.Stat(fontPath); os.IsNotExist(err) {
		return fmt.Errorf("font file not found at: %s", fontPath)
	}
	err := pdf.AddTTFFont(fontName, fontPath)
	if err != nil {
		return fmt.Errorf("failed to add font: %v", err)
	}

	err = pdf.SetFont("OpenSans", "", 8)
	if err != nil {
		return fmt.Errorf("failed to set font: %v", err)
	}

	if _, err := os.Stat("output.pdf"); err == nil {
		fmt.Printf("File exists\n")
		os.Remove("output.pdf")
		fmt.Print("Removed existing file: output.pdf\n")
	} else {
		fmt.Printf("File does not exist\n")
	}
	pdf.AddPage()
	pdf.SetY(5)
	pdf.SetX(5)
	pdf.Br(lineHeight)
	addListInfo(pdf, companyInfo, lineHeight)
	pdf.SetY(pdf.GetY() + lineHeight)
	addListInfo(pdf, invoiceInfo, lineHeight)
	pdf.SetY(pdf.GetY() + lineHeight)
	addTableInfo(pdf, items, lineHeight)
	pdf.SetY(pdf.GetY() + lineHeight)
	addListInfo(pdf, totals, lineHeight)

	return pdf.WritePdf("output.pdf")
}
func addListInfo(pdf *gopdf.GoPdf, info []string, lineHeight float64) {
	if len(info) == 0 {
		fmt.Println("no information provided")
	}
	for i, val := range info {
		fmt.Print(i)
		fmt.Println("Val: " + val)
		pdf.SetX(10.0)
		err := pdf.Cell(nil, val)
		pdf.Br(lineHeight)
		fmt.Println("Y position: ", pdf.GetY())
		if err != nil {
			fmt.Println(err)
		}
	}
}
func addTableInfo(pdf *gopdf.GoPdf, tableRows [][]string, lineHeight float64) {
	if len(tableRows) == 0 {
		fmt.Println("no table data provided")
		return
	}

	table := pdf.NewTableLayout(5, pdf.GetY()+lineHeight, lineHeight, len(tableRows))
	tableCols := []string{"itemNaam", "aantal", "btwTarief", "prijsPerStukExcl", "prijsPerStukIncl", "totaalExcl", "totaalIncl"}
	colWidth := 80.0
	for _, col := range tableCols {
		table.AddColumn(col, colWidth, "left")
		pdf.Br(lineHeight)
	}
	for _, row := range tableRows {
		paddedRow := make([]string, len(tableCols))
		copy(paddedRow, row)
		table.AddRow(paddedRow)
	}
	table.DrawTable()
}
