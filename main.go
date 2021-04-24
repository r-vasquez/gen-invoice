package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/r-vasquez/gen-invoice/invoice"
)

func handleError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func setFlags() (*string, *string) {
	invoice := flag.String("invoice", "", "Path to invoice JSON file")

	output := flag.String("output", "", "Path and name for your generated xml invoice file, example: ./invoice.xml")
	flag.Parse()

	if *invoice == "" {
		log.Fatalln("Please provide a path to your invoice file")
	}

	if *output == "" {
		fmt.Println("Generated xml Invoice will be saved to default location: ./dian_invoice.xml")
		*output = "./dian_invoice.xml"
	}

	return invoice, output
}

func main() {
	invoicePath, outputPath := setFlags()

	// Open our json Invoice
	var invoiceFile *os.File
	var err error
	invoiceFile, err = os.Open(*invoicePath)
	handleError(err)
	defer invoiceFile.Close()

	invoiceControl := invoice.GenInvoice(invoiceFile)
	invoice.WriteInvoiceToFile(invoiceControl, outputPath)
}
