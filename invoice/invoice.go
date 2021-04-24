package invoice

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func GenInvoice(invoiceFile *os.File) InvoiceControl {
	byteValue, _ := ioutil.ReadAll(invoiceFile)

	// Parse Json Invoice to InvoiceControl Struct
	var invoiceControl InvoiceControl
	e := json.Unmarshal(byteValue, &invoiceControl)
	handleError(e)

	// TODO : This will be deleted from here once we complete the invoice
	// Or leave it as -verbose
	fmt.Println("Input Invoice ->\n", prettyPrint(invoiceControl))
	fmt.Println("Output XML ->")
	outputXml, err := xml.MarshalIndent(invoiceControl, "", "    ")
	handleError(err)
	os.Stdout.Write(outputXml)
	// ------------------------------------------------------------------

	return invoiceControl
}

func WriteInvoiceToFile(invoiceControl InvoiceControl, outputPath *string) {
	outputXml, err := xml.MarshalIndent(invoiceControl, "", "    ")
	handleError(err)
	_ = ioutil.WriteFile(*outputPath, outputXml, 0644)
}

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
