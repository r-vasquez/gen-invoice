// This package deals with the creation of the invoice as per Dian standard
// This package might be migrated to pkg/invoice if an external client will use it
package invoice

import (
	"encoding/xml"
)

// This element MUST be conveyed as the root element in any instance document based on this Schema expression
type Invoice struct {
	XMLName              xml.Name             `xml:"Invoice"`
	Text                 string               `xml:",chardata"`
	Xmlns                string               `xml:"xmlns,attr"`
	Cac                  string               `xml:"cac,attr"`
	Cbc                  string               `xml:"cbc,attr"`
	Ds                   string               `xml:"ds,attr"`
	Ext                  string               `xml:"ext,attr"`
	Sts                  string               `xml:"sts,attr"`
	Xades                string               `xml:"xades,attr"`
	Xades141             string               `xml:"xades141,attr"`
	Xsi                  string               `xml:"xsi,attr"`
	SchemaLocation       string               `xml:"schemaLocation,attr"`
	UBLExtensions        UBLExtensions        `xml:"UBLExtensions"`
	UBLVersionID         string               `xml:"UBLVersionID"`
	CustomizationID      string               `xml:"CustomizationID"`
	ProfileID            string               `xml:"ProfileID"`
	ProfileExecutionID   string               `xml:"ProfileExecutionID"`
	ID                   string               `xml:"ID"`
	UUID                 UUID                 `xml:"UUID"`
	IssueDate            string               `xml:"IssueDate"`
	IssueTime            string               `xml:"IssueTime"`
	InvoiceTypeCode      string               `xml:"InvoiceTypeCode"`
	Note                 string               `xml:"Note"`
	DocumentCurrencyCode DocumentCurrencyCode `xml:"DocumentCurrencyCode"`
	LineCountNumeric     string               `xml:"LineCountNumeric"`
	InvoicePeriod        InvoicePeriod        `xml:"InvoicePeriod"`
	BillingReference     []struct {
		Text                     string `xml:",chardata"`
		InvoiceDocumentReference struct {
			Text string `xml:",chardata"`
			ID   string `xml:"ID"`
			UUID struct {
				Text       string `xml:",chardata"`
				SchemeName string `xml:"schemeName,attr"`
			} `xml:"UUID"`
			IssueDate           string `xml:"IssueDate"`
			DocumentDescription string `xml:"DocumentDescription"`
		} `xml:"InvoiceDocumentReference"`
	} `xml:"BillingReference"`
	AccountingSupplierParty struct {
		Text                string `xml:",chardata"`
		AdditionalAccountID string `xml:"AdditionalAccountID"`
		Party               struct {
			Text      string `xml:",chardata"`
			PartyName []struct {
				Text string `xml:",chardata"`
				Name string `xml:"Name"`
			} `xml:"PartyName"`
			PhysicalLocation struct {
				Text    string `xml:",chardata"`
				Address struct {
					Text                 string `xml:",chardata"`
					ID                   string `xml:"ID"`
					CityName             string `xml:"CityName"`
					CountrySubentity     string `xml:"CountrySubentity"`
					CountrySubentityCode string `xml:"CountrySubentityCode"`
					AddressLine          struct {
						Text string `xml:",chardata"`
						Line string `xml:"Line"`
					} `xml:"AddressLine"`
					Country struct {
						Text               string `xml:",chardata"`
						IdentificationCode string `xml:"IdentificationCode"`
						Name               struct {
							Text       string `xml:",chardata"`
							LanguageID string `xml:"languageID,attr"`
						} `xml:"Name"`
					} `xml:"Country"`
				} `xml:"Address"`
			} `xml:"PhysicalLocation"`
			PartyTaxScheme struct {
				Text             string `xml:",chardata"`
				RegistrationName string `xml:"RegistrationName"`
				CompanyID        struct {
					Text             string `xml:",chardata"`
					SchemeAgencyID   string `xml:"schemeAgencyID,attr"`
					SchemeAgencyName string `xml:"schemeAgencyName,attr"`
					SchemeID         string `xml:"schemeID,attr"`
					SchemeName       string `xml:"schemeName,attr"`
				} `xml:"CompanyID"`
				TaxLevelCode struct {
					Text     string `xml:",chardata"`
					ListName string `xml:"listName,attr"`
				} `xml:"TaxLevelCode"`
				RegistrationAddress struct {
					Text                 string `xml:",chardata"`
					ID                   string `xml:"ID"`
					CityName             string `xml:"CityName"`
					CountrySubentity     string `xml:"CountrySubentity"`
					CountrySubentityCode string `xml:"CountrySubentityCode"`
					AddressLine          struct {
						Text string `xml:",chardata"`
						Line string `xml:"Line"`
					} `xml:"AddressLine"`
					Country struct {
						Text               string `xml:",chardata"`
						IdentificationCode string `xml:"IdentificationCode"`
						Name               struct {
							Text       string `xml:",chardata"`
							LanguageID string `xml:"languageID,attr"`
						} `xml:"Name"`
					} `xml:"Country"`
				} `xml:"RegistrationAddress"`
				TaxScheme struct {
					Text string `xml:",chardata"`
					ID   string `xml:"ID"`
					Name string `xml:"Name"`
				} `xml:"TaxScheme"`
			} `xml:"PartyTaxScheme"`
			PartyLegalEntity struct {
				Text             string `xml:",chardata"`
				RegistrationName string `xml:"RegistrationName"`
				CompanyID        struct {
					Text             string `xml:",chardata"`
					SchemeAgencyID   string `xml:"schemeAgencyID,attr"`
					SchemeAgencyName string `xml:"schemeAgencyName,attr"`
					SchemeID         string `xml:"schemeID,attr"`
					SchemeName       string `xml:"schemeName,attr"`
				} `xml:"CompanyID"`
				CorporateRegistrationScheme struct {
					Text string `xml:",chardata"`
					ID   string `xml:"ID"`
					Name string `xml:"Name"`
				} `xml:"CorporateRegistrationScheme"`
			} `xml:"PartyLegalEntity"`
			Contact struct {
				Text           string `xml:",chardata"`
				Name           string `xml:"Name"`
				Telephone      string `xml:"Telephone"`
				ElectronicMail string `xml:"ElectronicMail"`
				Note           string `xml:"Note"`
			} `xml:"Contact"`
		} `xml:"Party"`
	} `xml:"AccountingSupplierParty"`
	AccountingCustomerParty struct {
		Text                string `xml:",chardata"`
		AdditionalAccountID string `xml:"AdditionalAccountID"`
		Party               struct {
			Text      string `xml:",chardata"`
			PartyName struct {
				Text string `xml:",chardata"`
				Name string `xml:"Name"`
			} `xml:"PartyName"`
			PhysicalLocation struct {
				Text    string `xml:",chardata"`
				Address struct {
					Text                 string `xml:",chardata"`
					ID                   string `xml:"ID"`
					CityName             string `xml:"CityName"`
					CountrySubentity     string `xml:"CountrySubentity"`
					CountrySubentityCode string `xml:"CountrySubentityCode"`
					AddressLine          struct {
						Text string `xml:",chardata"`
						Line string `xml:"Line"`
					} `xml:"AddressLine"`
					Country struct {
						Text               string `xml:",chardata"`
						IdentificationCode string `xml:"IdentificationCode"`
						Name               struct {
							Text       string `xml:",chardata"`
							LanguageID string `xml:"languageID,attr"`
						} `xml:"Name"`
					} `xml:"Country"`
				} `xml:"Address"`
			} `xml:"PhysicalLocation"`
			PartyTaxScheme struct {
				Text             string `xml:",chardata"`
				RegistrationName string `xml:"RegistrationName"`
				CompanyID        struct {
					Text             string `xml:",chardata"`
					SchemeAgencyID   string `xml:"schemeAgencyID,attr"`
					SchemeAgencyName string `xml:"schemeAgencyName,attr"`
					SchemeID         string `xml:"schemeID,attr"`
					SchemeName       string `xml:"schemeName,attr"`
				} `xml:"CompanyID"`
				TaxLevelCode struct {
					Text     string `xml:",chardata"`
					ListName string `xml:"listName,attr"`
				} `xml:"TaxLevelCode"`
				RegistrationAddress struct {
					Text                 string `xml:",chardata"`
					ID                   string `xml:"ID"`
					CityName             string `xml:"CityName"`
					CountrySubentity     string `xml:"CountrySubentity"`
					CountrySubentityCode string `xml:"CountrySubentityCode"`
					AddressLine          struct {
						Text string `xml:",chardata"`
						Line string `xml:"Line"`
					} `xml:"AddressLine"`
					Country struct {
						Text               string `xml:",chardata"`
						IdentificationCode string `xml:"IdentificationCode"`
						Name               struct {
							Text       string `xml:",chardata"`
							LanguageID string `xml:"languageID,attr"`
						} `xml:"Name"`
					} `xml:"Country"`
				} `xml:"RegistrationAddress"`
				TaxScheme struct {
					Text string `xml:",chardata"`
					ID   string `xml:"ID"`
					Name string `xml:"Name"`
				} `xml:"TaxScheme"`
			} `xml:"PartyTaxScheme"`
			PartyLegalEntity struct {
				Text             string `xml:",chardata"`
				RegistrationName string `xml:"RegistrationName"`
				CompanyID        struct {
					Text             string `xml:",chardata"`
					SchemeAgencyID   string `xml:"schemeAgencyID,attr"`
					SchemeAgencyName string `xml:"schemeAgencyName,attr"`
					SchemeID         string `xml:"schemeID,attr"`
					SchemeName       string `xml:"schemeName,attr"`
				} `xml:"CompanyID"`
				CorporateRegistrationScheme struct {
					Text string `xml:",chardata"`
					Name string `xml:"Name"`
				} `xml:"CorporateRegistrationScheme"`
			} `xml:"PartyLegalEntity"`
			Contact struct {
				Text           string `xml:",chardata"`
				Name           string `xml:"Name"`
				Telephone      string `xml:"Telephone"`
				ElectronicMail string `xml:"ElectronicMail"`
			} `xml:"Contact"`
		} `xml:"Party"`
	} `xml:"AccountingCustomerParty"`
	TaxRepresentativeParty struct {
		Text                string `xml:",chardata"`
		PartyIdentification struct {
			Text string `xml:",chardata"`
			ID   struct {
				Text             string `xml:",chardata"`
				SchemeAgencyID   string `xml:"schemeAgencyID,attr"`
				SchemeAgencyName string `xml:"schemeAgencyName,attr"`
				SchemeID         string `xml:"schemeID,attr"`
				SchemeName       string `xml:"schemeName,attr"`
			} `xml:"ID"`
		} `xml:"PartyIdentification"`
	} `xml:"TaxRepresentativeParty"`
	Delivery struct {
		Text            string `xml:",chardata"`
		DeliveryAddress struct {
			Text                 string `xml:",chardata"`
			ID                   string `xml:"ID"`
			CityName             string `xml:"CityName"`
			CountrySubentity     string `xml:"CountrySubentity"`
			CountrySubentityCode string `xml:"CountrySubentityCode"`
			AddressLine          struct {
				Text string `xml:",chardata"`
				Line string `xml:"Line"`
			} `xml:"AddressLine"`
			Country struct {
				Text               string `xml:",chardata"`
				IdentificationCode string `xml:"IdentificationCode"`
				Name               struct {
					Text       string `xml:",chardata"`
					LanguageID string `xml:"languageID,attr"`
				} `xml:"Name"`
			} `xml:"Country"`
		} `xml:"DeliveryAddress"`
		DeliveryParty struct {
			Text      string `xml:",chardata"`
			PartyName struct {
				Text string `xml:",chardata"`
				Name string `xml:"Name"`
			} `xml:"PartyName"`
			PhysicalLocation struct {
				Text    string `xml:",chardata"`
				Address struct {
					Text                 string `xml:",chardata"`
					ID                   string `xml:"ID"`
					CityName             string `xml:"CityName"`
					CountrySubentity     string `xml:"CountrySubentity"`
					CountrySubentityCode string `xml:"CountrySubentityCode"`
					AddressLine          struct {
						Text string `xml:",chardata"`
						Line string `xml:"Line"`
					} `xml:"AddressLine"`
					Country struct {
						Text               string `xml:",chardata"`
						IdentificationCode string `xml:"IdentificationCode"`
						Name               struct {
							Text       string `xml:",chardata"`
							LanguageID string `xml:"languageID,attr"`
						} `xml:"Name"`
					} `xml:"Country"`
				} `xml:"Address"`
			} `xml:"PhysicalLocation"`
			PartyTaxScheme struct {
				Text             string `xml:",chardata"`
				RegistrationName string `xml:"RegistrationName"`
				CompanyID        struct {
					Text             string `xml:",chardata"`
					SchemeAgencyID   string `xml:"schemeAgencyID,attr"`
					SchemeAgencyName string `xml:"schemeAgencyName,attr"`
					SchemeID         string `xml:"schemeID,attr"`
					SchemeName       string `xml:"schemeName,attr"`
				} `xml:"CompanyID"`
				TaxLevelCode struct {
					Text     string `xml:",chardata"`
					ListName string `xml:"listName,attr"`
				} `xml:"TaxLevelCode"`
				TaxScheme struct {
					Text string `xml:",chardata"`
					ID   string `xml:"ID"`
					Name string `xml:"Name"`
				} `xml:"TaxScheme"`
			} `xml:"PartyTaxScheme"`
			PartyLegalEntity struct {
				Text             string `xml:",chardata"`
				RegistrationName string `xml:"RegistrationName"`
				CompanyID        struct {
					Text             string `xml:",chardata"`
					SchemeAgencyID   string `xml:"schemeAgencyID,attr"`
					SchemeAgencyName string `xml:"schemeAgencyName,attr"`
					SchemeID         string `xml:"schemeID,attr"`
					SchemeName       string `xml:"schemeName,attr"`
				} `xml:"CompanyID"`
				CorporateRegistrationScheme struct {
					Text string `xml:",chardata"`
					Name string `xml:"Name"`
				} `xml:"CorporateRegistrationScheme"`
			} `xml:"PartyLegalEntity"`
			Contact struct {
				Text           string `xml:",chardata"`
				Name           string `xml:"Name"`
				Telephone      string `xml:"Telephone"`
				Telefax        string `xml:"Telefax"`
				ElectronicMail string `xml:"ElectronicMail"`
				Note           string `xml:"Note"`
			} `xml:"Contact"`
		} `xml:"DeliveryParty"`
	} `xml:"Delivery"`
	DeliveryTerms struct {
		Text                       string `xml:",chardata"`
		SpecialTerms               string `xml:"SpecialTerms"`
		LossRiskResponsibilityCode string `xml:"LossRiskResponsibilityCode"`
		LossRisk                   string `xml:"LossRisk"`
	} `xml:"DeliveryTerms"`
	PaymentMeans struct {
		Text             string `xml:",chardata"`
		ID               string `xml:"ID"`
		PaymentMeansCode string `xml:"PaymentMeansCode"`
		PaymentDueDate   string `xml:"PaymentDueDate"`
		PaymentID        string `xml:"PaymentID"`
	} `xml:"PaymentMeans"`
	PrepaidPayment struct {
		Text       string `xml:",chardata"`
		ID         string `xml:"ID"`
		PaidAmount struct {
			Text       string `xml:",chardata"`
			CurrencyID string `xml:"currencyID,attr"`
		} `xml:"PaidAmount"`
		ReceivedDate  string `xml:"ReceivedDate"`
		PaidDate      string `xml:"PaidDate"`
		InstructionID string `xml:"InstructionID"`
	} `xml:"PrepaidPayment"`
	TaxTotal []struct {
		Text      string `xml:",chardata"`
		TaxAmount struct {
			Text       string `xml:",chardata"`
			CurrencyID string `xml:"currencyID,attr"`
		} `xml:"TaxAmount"`
		TaxSubtotal []struct {
			Text          string `xml:",chardata"`
			TaxableAmount struct {
				Text       string `xml:",chardata"`
				CurrencyID string `xml:"currencyID,attr"`
			} `xml:"TaxableAmount"`
			TaxAmount struct {
				Text       string `xml:",chardata"`
				CurrencyID string `xml:"currencyID,attr"`
			} `xml:"TaxAmount"`
			TaxCategory struct {
				Text      string `xml:",chardata"`
				Percent   string `xml:"Percent"`
				TaxScheme struct {
					Text string `xml:",chardata"`
					ID   string `xml:"ID"`
					Name string `xml:"Name"`
				} `xml:"TaxScheme"`
			} `xml:"TaxCategory"`
		} `xml:"TaxSubtotal"`
	} `xml:"TaxTotal"`
	LegalMonetaryTotal struct {
		Text                string `xml:",chardata"`
		LineExtensionAmount struct {
			Text       string `xml:",chardata"`
			CurrencyID string `xml:"currencyID,attr"`
		} `xml:"LineExtensionAmount"`
		TaxExclusiveAmount struct {
			Text       string `xml:",chardata"`
			CurrencyID string `xml:"currencyID,attr"`
		} `xml:"TaxExclusiveAmount"`
		TaxInclusiveAmount struct {
			Text       string `xml:",chardata"`
			CurrencyID string `xml:"currencyID,attr"`
		} `xml:"TaxInclusiveAmount"`
		PrepaidAmount struct {
			Text       string `xml:",chardata"`
			CurrencyID string `xml:"currencyID,attr"`
		} `xml:"PrepaidAmount"`
		PayableAmount struct {
			Text       string `xml:",chardata"`
			CurrencyID string `xml:"currencyID,attr"`
		} `xml:"PayableAmount"`
	} `xml:"LegalMonetaryTotal"`
	InvoiceLine []struct {
		Text             string `xml:",chardata"`
		ID               string `xml:"ID"`
		InvoicedQuantity struct {
			Text     string `xml:",chardata"`
			UnitCode string `xml:"unitCode,attr"`
		} `xml:"InvoicedQuantity"`
		LineExtensionAmount struct {
			Text       string `xml:",chardata"`
			CurrencyID string `xml:"currencyID,attr"`
		} `xml:"LineExtensionAmount"`
		FreeOfChargeIndicator string `xml:"FreeOfChargeIndicator"`
		Delivery              struct {
			Text             string `xml:",chardata"`
			DeliveryLocation struct {
				Text string `xml:",chardata"`
				ID   struct {
					Text       string `xml:",chardata"`
					SchemeID   string `xml:"schemeID,attr"`
					SchemeName string `xml:"schemeName,attr"`
				} `xml:"ID"`
			} `xml:"DeliveryLocation"`
		} `xml:"Delivery"`
		AllowanceCharge struct {
			Text                    string `xml:",chardata"`
			ID                      string `xml:"ID"`
			ChargeIndicator         string `xml:"ChargeIndicator"`
			AllowanceChargeReason   string `xml:"AllowanceChargeReason"`
			MultiplierFactorNumeric string `xml:"MultiplierFactorNumeric"`
			Amount                  struct {
				Text       string `xml:",chardata"`
				CurrencyID string `xml:"currencyID,attr"`
			} `xml:"Amount"`
			BaseAmount struct {
				Text       string `xml:",chardata"`
				CurrencyID string `xml:"currencyID,attr"`
			} `xml:"BaseAmount"`
		} `xml:"AllowanceCharge"`
		TaxTotal struct {
			Text      string `xml:",chardata"`
			TaxAmount struct {
				Text       string `xml:",chardata"`
				CurrencyID string `xml:"currencyID,attr"`
			} `xml:"TaxAmount"`
			TaxSubtotal struct {
				Text          string `xml:",chardata"`
				TaxableAmount struct {
					Text       string `xml:",chardata"`
					CurrencyID string `xml:"currencyID,attr"`
				} `xml:"TaxableAmount"`
				TaxAmount struct {
					Text       string `xml:",chardata"`
					CurrencyID string `xml:"currencyID,attr"`
				} `xml:"TaxAmount"`
				TaxCategory struct {
					Text      string `xml:",chardata"`
					Percent   string `xml:"Percent"`
					TaxScheme struct {
						Text string `xml:",chardata"`
						ID   string `xml:"ID"`
						Name string `xml:"Name"`
					} `xml:"TaxScheme"`
				} `xml:"TaxCategory"`
			} `xml:"TaxSubtotal"`
		} `xml:"TaxTotal"`
		Item struct {
			Text                      string `xml:",chardata"`
			Description               string `xml:"Description"`
			SellersItemIdentification struct {
				Text string `xml:",chardata"`
				ID   string `xml:"ID"`
			} `xml:"SellersItemIdentification"`
			AdditionalItemIdentification struct {
				Text string `xml:",chardata"`
				ID   struct {
					Text       string `xml:",chardata"`
					SchemeID   string `xml:"schemeID,attr"`
					SchemeName string `xml:"schemeName,attr"`
				} `xml:"ID"`
			} `xml:"AdditionalItemIdentification"`
			StandardItemIdentification struct {
				Text string `xml:",chardata"`
				ID   struct {
					Text           string `xml:",chardata"`
					SchemeAgencyID string `xml:"schemeAgencyID,attr"`
					SchemeID       string `xml:"schemeID,attr"`
					SchemeName     string `xml:"schemeName,attr"`
				} `xml:"ID"`
			} `xml:"StandardItemIdentification"`
		} `xml:"Item"`
		Price struct {
			Text        string `xml:",chardata"`
			PriceAmount struct {
				Text       string `xml:",chardata"`
				CurrencyID string `xml:"currencyID,attr"`
			} `xml:"PriceAmount"`
			BaseQuantity struct {
				Text     string `xml:",chardata"`
				UnitCode string `xml:"unitCode,attr"`
			} `xml:"BaseQuantity"`
		} `xml:"Price"`
		DocumentReference []struct {
			Text             string `xml:",chardata"`
			ID               string `xml:"ID"`
			IssueDate        string `xml:"IssueDate"`
			DocumentTypeCode string `xml:"DocumentTypeCode"`
			DocumentType     string `xml:"DocumentType"`
		} `xml:"DocumentReference"`
		PricingReference struct {
			Text                      string `xml:",chardata"`
			AlternativeConditionPrice struct {
				Text        string `xml:",chardata"`
				PriceAmount struct {
					Text       string `xml:",chardata"`
					CurrencyID string `xml:"currencyID,attr"`
				} `xml:"PriceAmount"`
				PriceTypeCode string `xml:"PriceTypeCode"`
				PriceType     string `xml:"PriceType"`
			} `xml:"AlternativeConditionPrice"`
		} `xml:"PricingReference"`
	} `xml:"InvoiceLine"`
}

// ------- Init: UBL Extensions ----------- //
type UBLExtensions struct {
	Text         string `xml:",chardata"`
	UBLExtension []struct {
		Text             string           `xml:",chardata"`
		ExtensionContent ExtensionContent `xml:"ExtensionContent"`
	} `xml:"UBLExtension"`
}
type ExtensionContent struct {
	Text           string         `xml:",chardata"`
	DianExtensions DianExtensions `xml:"DianExtensions"`
	Signature      Signature      `xml:"Signature"`
}

type DianExtensions struct {
	Text                  string                `xml:",chardata"`
	InvoiceControl        InvoiceControl        `xml:"sts:InvoiceControl"`
	InvoiceSource         InvoiceSource         `xml:"InvoiceSource"`
	SoftwareProvider      SoftwareProvider      `xml:"SoftwareProvider"`
	SoftwareSecurityCode  SoftwareSecurityCode  `xml:"SoftwareSecurityCode"`
	AuthorizationProvider AuthorizationProvider `xml:"AuthorizationProvider"`
	QRCode                string                `xml:"QRCode"`
}

// Invoice Authorization [int]: Authorization number given by the Dian
// AuthorizationPeriod: Information Group for authorized dates in which you can generate invoices
// AuthorizedInvoices: Information Group for authorized invoice numbers
type InvoiceControl struct {
	Text                 string              `json:"-" xml:",chardata"`
	InvoiceAuthorization int                 `json:"invoiceAuthorization" xml:"sts:InvoiceAuthorization"`
	AuthorizationPeriod  AuthorizationPeriod `json:"authorizationPeriod" xml:"sts:AuthorizationPeriod"`
	AuthorizedInvoices   AuthorizedInvoices  `json:"authorizedInvoices" xml:"sts:AuthorizedInvoices"`
}

// StartDate[string] (AAAA-MM-DD)
// EndDate[string] (AAAA-MM-DD)
type AuthorizationPeriod struct {
	Text      string `json:"-" xml:",chardata"`
	StartDate string `json:"startDate" xml:"cbc:StartDate"`
	EndDate   string `json:"endDate" xml:"cbc:EndDate"`
}

// Prefix [string]: Establishment Prefix
// From [int]: Initial value of the invoice numeration
// To [int]: End value of the invoice numeration
type AuthorizedInvoices struct {
	Text   string `json:"-" xml:",chardata"`
	Prefix string `json:"prefix" xml:"sts:Prefix"`
	From   int    `json:"from" xml:"sts:From"`
	To     int    `json:"to" xml:"sts:To"`
}

type InvoiceSource struct {
	Text               string             `xml:",chardata"`
	IdentificationCode IdentificationCode `xml:"IdentificationCode"`
}

type IdentificationCode struct {
	Text           string `xml:",chardata"`
	ListAgencyID   string `xml:"listAgencyID,attr"`
	ListAgencyName string `xml:"listAgencyName,attr"`
	ListSchemeURI  string `xml:"listSchemeURI,attr"`
}

type SoftwareProvider struct {
	Text       string     `xml:",chardata"`
	ProviderID ProviderID `xml:"ProviderID"`
	SoftwareID SoftwareID `xml:"SoftwareID"`
}

type ProviderID struct {
	Text             string `xml:",chardata"`
	SchemeAgencyID   string `xml:"schemeAgencyID,attr"`
	SchemeAgencyName string `xml:"schemeAgencyName,attr"`
	SchemeID         string `xml:"schemeID,attr"`
	SchemeName       string `xml:"schemeName,attr"`
}

type SoftwareID struct {
	Text             string `xml:",chardata"`
	SchemeAgencyID   string `xml:"schemeAgencyID,attr"`
	SchemeAgencyName string `xml:"schemeAgencyName,attr"`
}

type SoftwareSecurityCode struct {
	Text             string `xml:",chardata"`
	SchemeAgencyID   string `xml:"schemeAgencyID,attr"`
	SchemeAgencyName string `xml:"schemeAgencyName,attr"`
}

type AuthorizationProvider struct {
	Text                    string                  `xml:",chardata"`
	AuthorizationProviderID AuthorizationProviderID `xml:"AuthorizationProviderID"`
}

type AuthorizationProviderID struct {
	Text             string `xml:",chardata"`
	SchemeAgencyID   string `xml:"schemeAgencyID,attr"`
	SchemeAgencyName string `xml:"schemeAgencyName,attr"`
	SchemeID         string `xml:"schemeID,attr"`
	SchemeName       string `xml:"schemeName,attr"`
}

// TODO : Flat nested struct when used.
type Signature struct {
	Text       string `xml:",chardata"`
	ID         string `xml:"Id,attr"`
	SignedInfo struct {
		Text                   string `xml:",chardata"`
		CanonicalizationMethod struct {
			Text      string `xml:",chardata"`
			Algorithm string `xml:"Algorithm,attr"`
		} `xml:"CanonicalizationMethod"`
		SignatureMethod struct {
			Text      string `xml:",chardata"`
			Algorithm string `xml:"Algorithm,attr"`
		} `xml:"SignatureMethod"`
		Reference []struct {
			Text       string `xml:",chardata"`
			ID         string `xml:"Id,attr"`
			URI        string `xml:"URI,attr"`
			Type       string `xml:"Type,attr"`
			Transforms struct {
				Text      string `xml:",chardata"`
				Transform struct {
					Text      string `xml:",chardata"`
					Algorithm string `xml:"Algorithm,attr"`
				} `xml:"Transform"`
			} `xml:"Transforms"`
			DigestMethod struct {
				Text      string `xml:",chardata"`
				Algorithm string `xml:"Algorithm,attr"`
			} `xml:"DigestMethod"`
			DigestValue string `xml:"DigestValue"`
		} `xml:"Reference"`
	} `xml:"SignedInfo"`
	SignatureValue struct {
		Text string `xml:",chardata"`
		ID   string `xml:"Id,attr"`
	} `xml:"SignatureValue"`
	KeyInfo struct {
		Text     string `xml:",chardata"`
		ID       string `xml:"Id,attr"`
		X509Data struct {
			Text            string `xml:",chardata"`
			X509Certificate string `xml:"X509Certificate"`
		} `xml:"X509Data"`
	} `xml:"KeyInfo"`
	Object struct {
		Text                 string `xml:",chardata"`
		QualifyingProperties struct {
			Text             string `xml:",chardata"`
			Target           string `xml:"Target,attr"`
			SignedProperties struct {
				Text                      string `xml:",chardata"`
				ID                        string `xml:"Id,attr"`
				SignedSignatureProperties struct {
					Text               string `xml:",chardata"`
					SigningTime        string `xml:"SigningTime"`
					SigningCertificate struct {
						Text string `xml:",chardata"`
						Cert []struct {
							Text       string `xml:",chardata"`
							CertDigest struct {
								Text         string `xml:",chardata"`
								DigestMethod struct {
									Text      string `xml:",chardata"`
									Algorithm string `xml:"Algorithm,attr"`
								} `xml:"DigestMethod"`
								DigestValue string `xml:"DigestValue"`
							} `xml:"CertDigest"`
							IssuerSerial struct {
								Text             string `xml:",chardata"`
								X509IssuerName   string `xml:"X509IssuerName"`
								X509SerialNumber string `xml:"X509SerialNumber"`
							} `xml:"IssuerSerial"`
						} `xml:"Cert"`
					} `xml:"SigningCertificate"`
					SignaturePolicyIdentifier struct {
						Text              string `xml:",chardata"`
						SignaturePolicyId struct {
							Text        string `xml:",chardata"`
							SigPolicyId struct {
								Text       string `xml:",chardata"`
								Identifier string `xml:"Identifier"`
							} `xml:"SigPolicyId"`
							SigPolicyHash struct {
								Text         string `xml:",chardata"`
								DigestMethod struct {
									Text      string `xml:",chardata"`
									Algorithm string `xml:"Algorithm,attr"`
								} `xml:"DigestMethod"`
								DigestValue string `xml:"DigestValue"`
							} `xml:"SigPolicyHash"`
						} `xml:"SignaturePolicyId"`
					} `xml:"SignaturePolicyIdentifier"`
					SignerRole struct {
						Text         string `xml:",chardata"`
						ClaimedRoles struct {
							Text        string `xml:",chardata"`
							ClaimedRole string `xml:"ClaimedRole"`
						} `xml:"ClaimedRoles"`
					} `xml:"SignerRole"`
				} `xml:"SignedSignatureProperties"`
			} `xml:"SignedProperties"`
		} `xml:"QualifyingProperties"`
	} `xml:"Object"`
}

// -------- End: UBL Extensions --------- //
type UUID struct {
	Text       string `xml:",chardata"`
	SchemeID   string `xml:"schemeID,attr"`
	SchemeName string `xml:"schemeName,attr"`
}

type DocumentCurrencyCode struct {
	Text           string `xml:",chardata"`
	ListAgencyID   string `xml:"listAgencyID,attr"`
	ListAgencyName string `xml:"listAgencyName,attr"`
	ListID         string `xml:"listID,attr"`
}

type InvoicePeriod struct {
	Text      string `xml:",chardata"`
	StartDate string `xml:"StartDate"`
	EndDate   string `xml:"EndDate"`
}
