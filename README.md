... Under development. Not MVP ready.

## Usage

```
go run main.go -help
  -invoice string
        Path to invoice JSON file
  -output string
        Path and name for your generated xml invoice file, example: ./invoice.xml
```

## Example using testData

`go run main.go -invoice testData/invoice.json -output ./invoice.xml ` 

This will take the invoice.json and creates an invoice.xml file according to Dian Requirements. 

```
 go run main.go -invoice testData/invoice.json -output ./invoice.xml 
```
