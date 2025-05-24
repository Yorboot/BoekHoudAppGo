package main

import (
	"github.com/signintech/gopdf"
	"log"
)

func generatePdf() {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	pdf.AddPage()

}
