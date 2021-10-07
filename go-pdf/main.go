package main

import (
	"log"

	"github.com/signintech/gopdf"
)

func main()  {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{ PageSize: *gopdf.PageSizeA4 })  
	pdf.AddPage()
	var err error
	err = pdf.AddTTFFont("B Zar", "./fonts/bzar.ttf")
	if err != nil {
		log.Print(err.Error())
		return
	}
	pdf.SetX(30)
	pdf.SetY(40)
	pdf.Text("Link to example.com")
	pdf.AddExternalLink("http://example.com/", 27.5, 28, 125, 15)

	pdf.SetX(30)
	pdf.SetY(70)
	pdf.Text("Link to second page")
	pdf.AddInternalLink("anchor", 27.5, 58, 120, 15)

	pdf.AddPage()
	pdf.SetX(30)
	pdf.SetY(100)
	pdf.SetAnchor("anchor")
	pdf.Text("Anchor position")

	pdf.WritePdf("hello.tmp.pdf")

}