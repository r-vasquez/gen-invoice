package dianExtensions

import (
	"encoding/xml"
)

// Datos Resolución de Numeración de Facturas
// Invoice Authorization [int]: Número autorización: Número del código de la resolución otorgada para la numeración
// AuthorizationPeriod: Grupo de informaciones relativas a la fecha de autorización de la numeración
// AuthorizedInvoices: Grupo de informaciones del rango de numeración autorizado para este emisor
type InvoiceControl struct {
	XMLName              xml.Name            `json:"-" xml:"sts:InvoiceControl"`
	Text                 string              `json:"-" xml:",chardata"`
	InvoiceAuthorization int                 `json:"numResolucion" xml:"sts:InvoiceAuthorization"`
	AuthorizationPeriod  AuthorizationPeriod `json:"periodoAutorizacion" xml:"sts:AuthorizationPeriod"`
	AuthorizedInvoices   AuthorizedInvoices  `json:"facturasAutorizadas" xml:"sts:AuthorizedInvoices"`
}

// StartDate[string] (AAAA-MM-DD)
// EndDate[string] (AAAA-MM-DD)
type AuthorizationPeriod struct {
	Text      string `json:"-" xml:",chardata"`
	StartDate string `json:"fechaDesde" xml:"cbc:StartDate"`
	EndDate   string `json:"fechaHasta" xml:"cbc:EndDate"`
}

// Prefix [string]: Prefijo establecido para el establecimiento
// From [int]: Valor inicial del rango de numeración otorgado
// To [int]: Valor final del rango de numeración otorgado
type AuthorizedInvoices struct {
	Text   string `json:"-" xml:",chardata"`
	Prefix string `json:"prefijo" xml:"sts:Prefix"`
	From   int    `json:"desde" xml:"sts:From"`
	To     int    `json:"hasta" xml:"sts:To"`
}
