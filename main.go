package main

import (
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/r-vasquez/gen-invoice/dataModel/dianExtensions"
)

func handleError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

// Util: Print idented json
func prettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "  ")
	return string(s)
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

	byteValue, _ := ioutil.ReadAll(invoiceFile)

	// Parse Json Invoice to InvoiceControl Struct
	var invoiceControl dianExtensions.InvoiceControl
	e := json.Unmarshal(byteValue, &invoiceControl)
	handleError(e)

	fmt.Println("Input Invoice ->\n", prettyPrint(invoiceControl))

	fmt.Println("Output XML ->")
	outputXml, err := xml.MarshalIndent(invoiceControl, "", "    ")
	handleError(err)
	os.Stdout.Write(outputXml)
	_ = ioutil.WriteFile(*outputPath, outputXml, 0644)
}
