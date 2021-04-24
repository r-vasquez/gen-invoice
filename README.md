... It's a proof of concept to check it's viability/ease of implementation using Go.
## Usage

```
go run main.go -help
  -invoice string
    	Path to invoice JSON file
  -output string
    	Path and name for your generated xml invoice file, example: ./invoice.xml
```

## Example using testData
```
 go run main.go -invoice test/testData/_invoice.json -output invoice.xml
```

This will take the _invoice.json file and creates an invoice.xml file that will look like:
```xml
<InvoiceControl>
    <sts:InvoiceAuthorization>18760000001</sts:InvoiceAuthorization>
    <sts:AuthorizationPeriod>
        <cbc:StartDate>2019-01-19</cbc:StartDate>
        <cbc:EndDate>2030-01-19</cbc:EndDate>
    </sts:AuthorizationPeriod>
    <sts:AuthorizedInvoices>
        <sts:Prefix>SETP</sts:Prefix>
        <sts:From>990000000</sts:From>
        <sts:To>995000000</sts:To>
    </sts:AuthorizedInvoices>
</InvoiceControl>
```
