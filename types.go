package netsuite

import (
	"github.com/cydev/zero"
	"github.com/omniboost/go-netsuite-rest/omitempty"
)

type JournalEntry struct {
	AccountingBook AccountingBook `json:"accountingBook,omitempty"`
	Approved       Bool           `json:"approved"`
	CreatedDate    Date           `json:"createdDate,omitempty"`
	Currency       Currency       `json:"currency,omitempty"`
	// CustbodyAdjustmentJournal      Bool `json:"custbody_adjustment_journal,omitempty"`
	// CustbodyCashRegister           Bool `json:"custbody_cash_register,omitempty"`
	// CustbodyFfBrExcludeTransaction Bool `json:"custbody_ff_br_exclude_transaction,omitempty"`
	// CustbodyNondeductibleRefTran   struct {
	// 	Links        Links         `json:"links"`
	// 	Count        int           `json:"count"`
	// 	HasMore      Bool          `json:"hasMore"`
	// 	Items        []interface{} `json:"items"`
	// 	Offset       int           `json:"offset"`
	// 	TotalResults int           `json:"totalResults"`
	// } `json:"custbody_nondeductible_ref_tran,omitempty"`
	// CustbodyNstsGawClkTrigApprove Bool `json:"custbody_nsts_gaw_clk_trig_approve"`
	// CustbodyNstsGawClkTrigReject  Bool `json:"custbody_nsts_gaw_clk_trig_reject"`
	// CustbodyNstsGawClkTrigSubmit  Bool `json:"custbody_nsts_gaw_clk_trig_submit"`
	// CustbodyNstsGawCreatedBy      struct {
	// 	Links   Links  `json:"links"`
	// 	ID      string `json:"id"`
	// 	RefName string `json:"refName"`
	// } `json:"custbody_nsts_gaw_created_by"`
	// CustbodyNstsGawIsDelegated   Bool `json:"custbody_nsts_gaw_is_delegated"`
	// CustbodyNstsGawNextApprovers struct {
	// 	Links        Links         `json:"links"`
	// 	Count        int           `json:"count"`
	// 	HasMore      Bool          `json:"hasMore"`
	// 	Items        []interface{} `json:"items"`
	// 	Offset       int           `json:"offset"`
	// 	TotalResults int           `json:"totalResults"`
	// } `json:"custbody_nsts_gaw_next_approvers"`
	// CustbodyNstsGawSuperappApproved Bool   `json:"custbody_nsts_gaw_superapp_approved"`
	// CustbodyNstsGawTriggerSuper     Bool   `json:"custbody_nsts_gaw_trigger_super"`
	// CustbodyReportTimestamp         string `json:"custbody_report_timestamp"`
	// CustbodySiiArticle61D           Bool   `json:"custbody_sii_article_61d"`
	// CustbodySiiArticle7273          Bool   `json:"custbody_sii_article_72_73"`
	// CustbodySiiNotReportedInTime    Bool   `json:"custbody_sii_not_reported_in_time"`
	// CustbodyVatrepTaxperiodTrn      string `json:"custbody_vatrep_taxperiod_trn"`
	// CustbodyVatrepTrnenabled        Bool   `json:"custbody_vatrep_trnenabled"`
	CustomForm             CustomForm `json:"customForm,omitempty"`
	ExchangeRate           float64    `json:"exchangeRate,omitempty"`
	ExcludeFromGLNumbering Bool       `json:"excludeFromGLNumbering,omitempty"`
	ID                     string     `json:"id,omitempty"`
	IsReversal             Bool       `json:"isReversal,omitempty"`
	// LastModifiedDate       Date             `json:"lastModifiedDate,omitempty"`
	Lines         JournalEntryLine `json:"line"`
	Memo          string           `json:"memo"`
	PostingPeriod PostingPeriod    `json:"postingPeriod,omitempty"`
	RefName       string           `json:"refName,omitempty"`
	ReversalDefer Bool             `json:"reversalDefer,omitempty"`
	Subsidiary    Subsidiary       `json:"subsidiary,omitempty"`
	TranDate      Date             `json:"tranDate,omitempty"`
	TranID        string           `json:"tranId,omitempty"`
	Void          Bool             `json:"void,omitempty"`
	// CustBody4              string           `json:"custbody4"`
}

func (j JournalEntry) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(j)
}

func (j JournalEntry) IsEmpty() bool {
	return zero.IsZero(j)
}

type JournalEntryLine struct {
	Links        Links                    `json:"links,omitempty"`
	Items        JournalEntryLineElements `json:"items"`
	TotalResults int                      `json:"totalResults,omitempty"`
}

type Currency struct {
	Links   Links  `json:"links"`
	ID      string `json:"id"`
	RefName string `json:"refName"`
}

func (c Currency) IsEmpty() bool {
	return zero.IsZero(c)
}

type PostingPeriod struct {
	Links   Links  `json:"links"`
	ID      string `json:"id"`
	RefName string `json:"refName"`
}

func (p PostingPeriod) IsEmpty() bool {
	return zero.IsZero(p)
}

type Subsidiary struct {
	Links   Links  `json:"links,omitempty"`
	ID      string `json:"id,omitempty"`
	RefName string `json:"refName,omitempty"`
}

func (s Subsidiary) IsEmpty() bool {
	return zero.IsZero(s)
}

type AccountingBook struct {
	Links   Links  `json:"links,omitempty"`
	ID      string `json:"id,omitempty"`
	RefName string `json:"refName,omitempty"`
}

func (a AccountingBook) IsEmpty() bool {
	return zero.IsZero(a)
}

type CustomForm struct {
	ID      string `json:"id"`
	RefName string `json:"refName"`
}

func (c CustomForm) IsEmpty() bool {
	return zero.IsZero(c)
}

type JournalEntryLineElements []JournalEntryLineElement

type JournalEntryLineElement struct {
	Links               Links     `json:"links,omitempty"`
	Account             Account   `json:"Account,omitempty"`
	Cleared             Bool      `json:"cleared,omitempty"`
	Credit              float64   `json:"credit,omitempty"`
	Custcol2663Isperson Bool      `json:"custcol_2663_isperson,omitempty"`
	Eliminate           Bool      `json:"eliminate,omitempty"`
	Line                int       `json:"line,omitempty"`
	Debit               float64   `json:"debit,omitempty"`
	Memo                string    `json:"memo"`
	Department          RecordRef `json:"Department,omitempty"`
	Class               RecordRef `json:"Class,omitempty"`
	CustCol1            string    `json:"custcol1,omitempty"`
	CustCol2            string    `json:"custcol2,omitempty"`
	CustCol3            string    `json:"custcol3,omitempty"`
	CustCol4            string    `json:"custcol4,omitempty"`
	CustCol5            string    `json:"custcol5,omitempty"`
}

func (j JournalEntryLineElement) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(j)
}

type Accounts []Account

type Account struct {
	Links      Links  `json:"links,omitempty"`
	ID         string `json:"id,omitempty"`
	RefName    string `json:"refName,omitempty"`
	AcctNumber string `json:"acctNumber,omitempty"`
	ExternalID string `json:"externalId,omitempty"`
}

func (a Account) IsEmpty() bool {
	return zero.IsZero(a)
}

type Invoice struct {
	Links Links `json:"links"`
	// Account struct {
	// 	Links   Links  `json:"links"`
	// 	ID      string `json:"id"`
	// 	RefName string `json:"refName"`
	// } `json:"account"`
	// AccountingBookDetail struct {
	// 	Links        Links         `json:"links"`
	// 	Items        []interface{} `json:"items"`
	// 	TotalResults int           `json:"totalResults"`
	// } `json:"accountingBookDetail"`
	// AmountPaid              float64 `json:"amountPaid"`
	// AmountRemaining         float64 `json:"amountRemaining"`
	// AmountRemainingTotalBox float64 `json:"amountRemainingTotalBox"`
	// BillingAddress          Address `json:"billingAddress"`
	// CreatedDate Date `json:"createdDate"`
	// Currency    struct {
	// 	Links   Links  `json:"links"`
	// 	ID      string `json:"id"`
	// 	RefName string `json:"refName"`
	// } `json:"currency"`
	// CustbodyAtlasExistCustHdn struct {
	// 	Links   Links  `json:"links"`
	// 	ID      string `json:"id"`
	// 	RefName string `json:"refName"`
	// } `json:"custbody_atlas_exist_cust_hdn"`
	// CustbodyAtlasNewCustHdn struct {
	// 	Links   Links  `json:"links"`
	// 	ID      string `json:"id"`
	// 	RefName string `json:"refName"`
	// } `json:"custbody_atlas_new_cust_hdn"`
	// CustbodyAtlasNoHdn struct {
	// 	Links   Links  `json:"links"`
	// 	ID      string `json:"id"`
	// 	RefName string `json:"refName"`
	// } `json:"custbody_atlas_no_hdn"`
	// CustbodyAtlasYesHdn struct {
	// 	Links   Links  `json:"links"`
	// 	ID      string `json:"id"`
	// 	RefName string `json:"refName"`
	// } `json:"custbody_atlas_yes_hdn"`
	// CustbodyEdocGenTransPdf       Bool   `json:"custbody_edoc_gen_trans_pdf"`
	// CustbodyEiDsTxnIdentifier     Bool   `json:"custbody_ei_ds_txn_identifier"`
	// CustbodyMxCfdiFolio           string `json:"custbody_mx_cfdi_folio"`
	// CustbodyMxTxnSatPaymentMethod struct {
	// 	Links   Links  `json:"links"`
	// 	ID      string `json:"id"`
	// 	RefName string `json:"refName"`
	// } `json:"custbody_mx_txn_sat_payment_method"`
	// CustbodyPsgEiEdocRecipient struct {
	// 	Links        Links         `json:"links"`
	// 	Count        int           `json:"count"`
	// 	HasMore      Bool          `json:"hasMore"`
	// 	Items        []interface{} `json:"items"`
	// 	Offset       int           `json:"offset"`
	// 	TotalResults int           `json:"totalResults"`
	// } `json:"custbody_psg_ei_edoc_recipient"`
	// CustbodySteRcsApplicable   Bool   `json:"custbody_ste_rcs_applicable"`
	// CustbodySteRcsInvoiceTexts string `json:"custbody_ste_rcs_invoice_texts"`
	// CustbodySteTransactionType struct {
	// 	Links   Links  `json:"links"`
	// 	ID      string `json:"id"`
	// 	RefName string `json:"refName"`
	// } `json:"custbody_ste_transaction_type"`
	CustomForm CustomForm `json:"customForm"`
	DueDate    Date       `json:"dueDate,omitempty"`
	Entity     struct {
		Links   Links  `json:"links,omitempty"`
		ID      string `json:"id"`
		RefName string `json:"refName,omitempty"`
	} `json:"entity"`
	// EstGrossProfit         float64     `json:"estGrossProfit"`
	// EstGrossProfitPercent  float64     `json:"estGrossProfitPercent"`
	// ExchangeRate           float64     `json:"exchangeRate"`
	// ExcludeFromGLNumbering Bool        `json:"excludeFromGLNumbering"`
	ID   string      `json:"id"`
	Item InvoiceItem `json:"item"`
	// LastModifiedDate       Date        `json:"lastModifiedDate"`
	// Location InvoiceLocation `json:"location"`
	Memo string `json:"memo"`
	// Nexus struct {
	// 	Links   Links  `json:"links"`
	// 	ID      string `json:"id"`
	// 	RefName string `json:"refName"`
	// } `json:"nexus"`
	// PostingPeriod struct {
	// 	Links   Links  `json:"links"`
	// 	ID      string `json:"id"`
	// 	RefName string `json:"refName"`
	// } `json:"postingPeriod"`
	// PrevDate           Date    `json:"prevDate"`
	// SalesEffectiveDate Date    `json:"salesEffectiveDate"`
	// ShipDate           Date    `json:"shipDate"`
	// ShipIsResidential  Bool    `json:"shipIsResidential"`
	// ShipOverride       Bool    `json:"shipOverride"`
	// ShippingAddress    Address `json:"shippingAddress"`
	// Status struct {
	// 	ID      string `json:"id"`
	// 	RefName string `json:"refName"`
	// } `json:"status"`
	Subsidiary          Subsidiary `json:"subsidiary"`
	SubsidiaryTaxRegNum string     `json:"subsidiaryTaxRegNum,omitempty"`
	Subtotal            float64    `json:"subtotal,omitempty"`
	// TaxDetails           InvoiceTaxDetails `json:"taxDetails,omitempty"`
	// TaxDetailsOverride   Bool    `json:"taxDetailsOverride"`
	// TaxPointDate         Date    `json:"taxPointDate"`
	// TaxPointDateOverride Bool    `json:"taxPointDateOverride"`
	// TaxRegOverride       Bool    `json:"taxRegOverride"`
	// TaxTotal             float64 `json:"taxTotal"`
	// ToBeEmailed          Bool    `json:"toBeEmailed"`
	// ToBeFaxed            Bool    `json:"toBeFaxed"`
	// ToBePrinted          Bool    `json:"toBePrinted"`
	// Total                float64 `json:"total"`
	// TotalAfterTaxes      float64 `json:"totalAfterTaxes"`
	// TotalCostEstimate    float64 `json:"totalCostEstimate"`
	TranDate   Date      `json:"tranDate"`
	TranID     string    `json:"tranId"`
	Department RecordRef `json:"Department,omitempty"`
	Class      RecordRef `json:"Class,omitempty"`
}

func (i Invoice) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(i)
}

type Customers []Customer

type Customer struct {
	AddressBook struct {
		Links        Links     `json:"links"`
		Items        []Address `json:"items"`
		TotalResults int       `json:"totalResults"`
	} `json:"addressBook"`
	// Aging  float64 `json:"aging"`
	// Aging1 float64 `json:"aging1"`
	// Aging2 float64 `json:"aging2"`
	// Aging3 float64 `json:"aging3"`
	// Aging4 float64 `json:"aging4"`
	// AlcoholRecipientType struct {
	// 	ID      string `json:"id"`
	// 	RefName string `json:"refName"`
	// } `json:"alcoholRecipientType"`
	// Balance float64 `json:"balance"`
	// Campaigns struct {
	// 	Links        Links         `json:"links"`
	// 	Items        []interface{} `json:"items"`
	// 	TotalResults int           `json:"totalResults"`
	// } `json:"campaigns"`
	CompanyName string `json:"companyName,omitempty"`
	// ContactList struct {
	// 	Links        Links         `json:"links"`
	// 	Count        int           `json:"count"`
	// 	HasMore      Bool          `json:"hasMore"`
	// 	Items        []interface{} `json:"items"`
	// 	Offset       int           `json:"offset"`
	// 	TotalResults int           `json:"totalResults"`
	// } `json:"contactList"`
	// CreditHoldOverride struct {
	// 	ID      string `json:"id"`
	// 	RefName string `json:"refName"`
	// } `json:"creditHoldOverride"`
	// Currency struct {
	// 	Links   Links  `json:"links"`
	// 	ID      string `json:"id"`
	// 	RefName string `json:"refName"`
	// } `json:"currency"`
	// CurrencyList struct {
	// 	Links Links `json:"links"`
	// 	Items []struct {
	// 		Links    Links   `json:"links"`
	// 		Balance  float64 `json:"balance"`
	// 		Currency struct {
	// 			Links   Links  `json:"links"`
	// 			ID      string `json:"id"`
	// 			RefName string `json:"refName"`
	// 		} `json:"currency"`
	// 		DepositBalance         float64 `json:"depositBalance"`
	// 		DisplaySymbol          string  `json:"displaySymbol"`
	// 		OverdueBalance         float64 `json:"overdueBalance"`
	// 		OverrideCurrencyFormat Bool    `json:"overrideCurrencyFormat"`
	// 		SymbolPlacement        struct {
	// 			ID      string `json:"id"`
	// 			RefName string `json:"refName"`
	// 		} `json:"symbolPlacement"`
	// 		UnbilledOrders float64 `json:"unbilledOrders"`
	// 	} `json:"items"`
	// 	TotalResults int `json:"totalResults"`
	// } `json:"currencyList"`
	// Custentity2663CustomerRefund      Bool   `json:"custentity_2663_customer_refund"`
	// Custentity2663DirectDebit         Bool   `json:"custentity_2663_direct_debit"`
	// CustentityEdocGenTransPdf         Bool   `json:"custentity_edoc_gen_trans_pdf"`
	// CustentityMxRfc                   string `json:"custentity_mx_rfc"`
	// CustentityPsgEiAutoSelectTempSm   Bool   `json:"custentity_psg_ei_auto_select_temp_sm"`
	// CustentityPsgEiEntityEdocStandard struct {
	// 	Links   Links  `json:"links"`
	// 	ID      string `json:"id"`
	// 	RefName string `json:"refName"`
	// } `json:"custentity_psg_ei_entity_edoc_standard"`
	// CustomForm struct {
	// 	ID      string `json:"id"`
	// 	RefName string `json:"refName"`
	// } `json:"customForm"`
	DateCreated *Date `json:"dateCreated,omitempty"`
	// DaysOverdue     int       `json:"daysOverdue"`
	// DefaultTaxReg   string    `json:"defaultTaxReg"`
	// DepositBalance  float64   `json:"depositBalance"`
	// EmailPreference struct {
	// 	ID      string `json:"id"`
	// 	RefName string `json:"refName"`
	// } `json:"emailPreference"`
	// EmailTransactions Bool   `json:"emailTransactions,omitempty"`
	// EntityID          string `json:"entityId,omitempty"`
	// EntityStatus      struct {
	// 	Links   Links  `json:"links"`
	// 	ID      string `json:"id"`
	// 	RefName string `json:"refName"`
	// } `json:"entityStatus"`
	// FaxTransactions          Bool `json:"faxTransactions,omitempty"`
	FirstName string `json:"firstName"`
	// GlobalSubscriptionStatus struct {
	// 	ID      string `json:"id"`
	// 	RefName string `json:"refName"`
	// } `json:"globalSubscriptionStatus"`
	// GroupPricing struct {
	// 	Links        Links         `json:"links"`
	// 	Items        []interface{} `json:"items"`
	// 	TotalResults int           `json:"totalResults"`
	// } `json:"groupPricing"`
	// ID string `json:"id"`
	// IsAutogeneratedRepresentingEntity Bool   `json:"isAutogeneratedRepresentingEntity"`
	// IsBudgetApproved                  Bool   `json:"isBudgetApproved"`
	IsInactive Bool `json:"isInactive"`
	IsPerson   Bool `json:"isPerson"`
	// ItemPricing                       struct {
	// 	Links        Links         `json:"links"`
	// 	Items        []interface{} `json:"items"`
	// 	TotalResults int           `json:"totalResults"`
	// } `json:"itemPricing"`
	// Language struct {
	// 	ID      string `json:"id"`
	// 	RefName string `json:"refName"`
	// } `json:"language"`
	LastModifiedDate *Date  `json:"lastModifiedDate,omitempty"`
	LastName         string `json:"lastName"`
	// OverdueBalance     float64    `json:"overdueBalance"`
	// PrintTransactions  Bool       `json:"printTransactions"`
	// ReceivablesAccount struct {
	// 	Links   Links  `json:"links"`
	// 	ID      string `json:"id"`
	// 	RefName string `json:"refName"`
	// } `json:"receivablesAccount"`
	// ShipComplete Bool `json:"shipComplete"`
	// ShippingCarrier struct {
	ID string `json:"id,omitempty"`
	// 	RefName string `json:"refName"`
	// } `json:"shippingCarrier"`
	Subsidiary Subsidiary `json:"subsidiary"`
	// TaxRegistration struct {
	// 	Links Links `json:"links"`
	// 	Items []struct {
	// 		Links Links `json:"links"`
	// 		ID    int   `json:"id"`
	// 		Nexus struct {
	// 			Links   Links  `json:"links"`
	// 			ID      string `json:"id"`
	// 			RefName string `json:"refName"`
	// 		} `json:"nexus"`
	// 		NexusCountry struct {
	// 			ID      string `json:"id"`
	// 			RefName string `json:"refName"`
	// 		} `json:"nexusCountry"`
	// 		TaxRegistrationNumber string `json:"taxRegistrationNumber"`
	// 	} `json:"items"`
	// 	TotalResults int `json:"totalResults"`
	// } `json:"taxRegistration"`
	// UnbilledOrders float64 `json:"unbilledOrders"`
	Email                  string `json:"email"`
	Phone                  string `json:"phone"`
	DefaultBillingAddress  string `json:"defaultbillingaddress"`
	DefaultShippingAddress string `json:"defaultshippingaddress"`
	Parent                 string `json:"parent"`
	CustomerNumber         string `json:"custentity_nch_customer_number"`
}

type InvoiceItem struct {
	Links        Links            `json:"links"`
	TotalResults int              `json:"totalResults"`
	Items        InvoiceItemItems `json:"items"`
}

type InvoiceItemItems []InvoiceItemItem

type InvoiceItemItem struct {
	Links   Links   `json:"links"`
	Account Account `json:"account"`
	Amount  float64 `json:"amount"`
	// CostEstimate     float64 `json:"costEstimate"`
	// CostEstimateRate float64 `json:"costEstimateRate"`
	// CostEstimateType struct {
	// 	ID      string `json:"id"`
	// 	RefName string `json:"refName"`
	// } `json:"costEstimateType"`
	// Custcol2663Isperson       Bool `json:"custcol_2663_isperson"`
	// CustcolSteItemTaxSchedule struct {
	// 	Links   Links  `json:"links"`
	// 	ID      string `json:"id"`
	// 	RefName string `json:"refName"`
	// } `json:"custcol_ste_item_tax_schedule"`
	// CustcolSteTaxItemType struct {
	// 	Links   Links  `json:"links"`
	// 	ID      string `json:"id"`
	// 	RefName string `json:"refName"`
	// } `json:"custcol_ste_tax_item_type"`
	// CustcolSteTaxItemTypeCode string `json:"custcol_ste_tax_item_type_code"`
	Description string `json:"description"`
	// ExcludeFromRateRequest Bool `json:"excludeFromRateRequest"`
	// InventoryDetail        struct {
	// 	Links               Links `json:"links"`
	// 	Inventoryassignment struct {
	// 		Links        Links         `json:"links"`
	// 		Items        []interface{} `json:"items"`
	// 		TotalResults int           `json:"totalResults"`
	// 	} `json:"inventoryassignment"`
	// } `json:"inventoryDetail"`
	Item        InvoiceItemItemItem `json:"item"`
	ItemSubType string              `json:"itemSubType"`
	ItemType    string              `json:"itemType"`
	// Line        int                 `json:"line"`
	// Marginal Bool `json:"marginal"`
	// Price struct {
	// 	Links   Links  `json:"links"`
	// 	ID      string `json:"id"`
	// 	RefName string `json:"refName"`
	// } `json:"price"`
	// PrintItems          Bool    `json:"printItems"`
	// Quantity            float64 `json:"quantity"`
	TaxAmount           float64 `json:"taxAmount"`
	TaxDetailsReference string  `json:"taxDetailsReference"`
	// Units               string  `json:"units"`
	// CustCol2 string `json:"custcol2"`
	// CustCol3 string `json:"custcol3"`
}

type InvoiceTaxDetails struct {
	Links Links `json:"links"`
	Items []struct {
		Links     Links   `json:"links"`
		LineName  string  `json:"lineName"`
		LineType  string  `json:"lineType"`
		NetAmount float64 `json:"netAmount"`
		TaxAmount float64 `json:"taxAmount"`
		TaxBasis  float64 `json:"taxBasis"`
		TaxCode   struct {
			Links   Links  `json:"links"`
			ID      string `json:"id"`
			RefName string `json:"refName"`
		} `json:"taxCode"`
		TaxDetailsReference struct {
			ID      string `json:"id"`
			RefName string `json:"refName"`
		} `json:"taxDetailsReference"`
		TaxRate float64 `json:"taxRate"`
		TaxType string  `json:"taxType"`
	} `json:"items"`
	TotalResults int `json:"totalResults"`
}

type InvoiceItemItemItem struct {
	// Links   Links  `json:"links"`
	ID      int    `json:"id"`
	RefName string `json:"refName"`
}

type InvoiceLocation struct {
	Links   Links  `json:"links"`
	ID      int    `json:"id"`
	RefName string `json:"refName"`
}

type RecordRef struct {
	ID string `json:"id"`
	// ExternalID string `json:"externalId"`
	// InternalID string `json:"id"`
}

func (r RecordRef) IsEmpty() bool {
	return zero.IsZero(r)
}

type Classifications []Classification

type Classification struct {
	Links            Links  `json:"links"`
	Fullname         string `json:"fullname"`
	ID               string `json:"id"`
	Includechildren  string `json:"includechildren"`
	Isinactive       string `json:"isinactive"`
	Lastmodifieddate string `json:"lastmodifieddate"`
	Name             string `json:"name"`
	Parent           string `json:"parent"`
	Subsidiary       string `json:"subsidiary"`
}

func (c Classification) IsEmpty() bool {
	return zero.IsZero(c)
}

type Departments []Department

type Department struct {
	Links            Links  `json:"links"`
	Fullname         string `json:"fullname"`
	ID               string `json:"id"`
	Includechildren  string `json:"includechildren"`
	Isinactive       string `json:"isinactive"`
	Lastmodifieddate string `json:"lastmodifieddate"`
	Name             string `json:"name"`
	Subsidiary       string `json:"subsidiary"`
	Custrecord1      string `json:"custrecord1,omitempty"`
	Parent           string `json:"parent,omitempty"`
}

func (d Department) IsEmpty() bool {
	return zero.IsZero(d)
}

type Addresses []Address

type Address struct {
	Links                    Links  `json:"links"`
	Addr1                    string `json:"addr1"`
	Addrtext                 string `json:"addrtext"`
	Attention                string `json:"attention"`
	City                     string `json:"city"`
	Country                  string `json:"country"`
	CustrecordNchCountrycode string `json:"custrecord_nch_countrycode"`
	CustrecordNchInvoiceFee  string `json:"custrecord_nch_invoice_fee"`
	Lastmodifieddate         string `json:"lastmodifieddate"`
	Nkey                     string `json:"nkey"`
	Override                 string `json:"override"`
	Recordowner              string `json:"recordowner"`
	Zip                      string `json:"zip"`
	Addressee                string `json:"addressee,omitempty"`
}

type Locations []Location

type Location struct {
	Links                  []interface{} `json:"links"`
	Fullname               string        `json:"fullname"`
	ID                     string        `json:"id"`
	Isinactive             string        `json:"isinactive"`
	Lastmodifieddate       string        `json:"lastmodifieddate"`
	Mainaddress            string        `json:"mainaddress"`
	Makeinventoryavailable string        `json:"makeinventoryavailable"`
	Name                   string        `json:"name"`
	Subsidiary             string        `json:"subsidiary"`
	Externalid             string        `json:"externalid,omitempty"`
}

// known Transaction record types:
// creditmemo
// vendorbill
// vendorcredit
// fxreval
// customerrefund
// deposit
// expensereport
// invoice
// journalentry
// advintercompanyjournalentry
// customerpayment

type CreditMemoTransactions []CreditMemoTransaction
type CreditMemoTransaction struct {
	Links                              Links  `json:"links"`
	AbbrevType                         string `json:"abbrevtype"`
	BalSegStatus                       string `json:"balsegstatus"`
	BillingAddress                     string `json:"billingaddress"`
	BillingStatus                      string `json:"billingstatus"`
	CloseDate                          string `json:"closedate"`
	CreatedBy                          string `json:"createdby"`
	CreatedDate                        string `json:"createddate"`
	CreatedFromMerge                   string `json:"createdfrommerge"`
	Currency                           string `json:"currency"`
	Custbody15699ExcludeFromEP         string `json:"custbody_15699_exclude_from_ep_process"`
	CustbodyBsDdrDoNotReprocess        string `json:"custbody_bs_ddr_donot_reprocess"`
	CustbodyEdocGenTransPdf            string `json:"custbody_edoc_gen_trans_pdf"`
	CustbodyEiDsTxnIdentifier          string `json:"custbody_ei_ds_txn_identifier"`
	CustbodyEinvoicingDefaultPay       string `json:"custbody_einvoicing_defaultpaymtxt"`
	CustbodyEinvoicingGlneanovt        string `json:"custbody_einvoicing_glneanovt"`
	CustbodyFfBrExcludeTrans           string `json:"custbody_ff_br_exclude_transaction"`
	CustbodyFinInvoiceRefNumber        string `json:"custbody_fin_invoicereferencenumber"`
	CustbodyInvoiceCompanyEmail        string `json:"custbody_invoice_company_email"`
	CustbodyInvoiceCompanyPhone        string `json:"custbody_invoice_company_phone"`
	CustbodyNchCollectionForeign       string `json:"custbody_nch_collection_foreigncust"`
	CustbodyNchCollectionHold          string `json:"custbody_nch_collection_hold"`
	CustbodyNchCreditReason            string `json:"custbody_nch_credit_reason"`
	CustbodyNchCustBillingAddrEmail    string `json:"custbody_nch_cust_billingaddressemail"`
	CustbodyNchCustomerNumber          string `json:"custbody_nch_customernumber"`
	CustbodyNchEinvoicingPayeeID       string `json:"custbody_nch_einvoicing_payee_busin_id"`
	CustbodyNchEinvoicingPayeeNm       string `json:"custbody_nch_einvoicing_payee_name"`
	CustbodyNchExternalResNbr          string `json:"custbody_nch_external_res_nbr"`
	CustbodyNchInDispute               string `json:"custbody_nch_indispute"`
	CustbodyNchPmsOriginalInvNo        string `json:"custbody_nch_pms_originalinvoiceno"`
	CustbodyNchSettlPayUpdate          string `json:"custbody_nch_settl__pay_update"`
	CustbodyNchSource                  string `json:"custbody_nch_source"`
	CustbodyNsNordInvTypeCode          string `json:"custbody_ns_nord_inv_type_code"`
	CustbodyPayeeBic                   string `json:"custbody_payee_bic"`
	CustbodyPsgEiTransEdocStd          string `json:"custbody_psg_ei_trans_edoc_standard"`
	CustbodyReportTimestamp            string `json:"custbody_report_timestamp"`
	CustbodyRmAppliedInvoice           string `json:"custbody_rm_applied_invoice"`
	CustbodyRmCustomStatusForPos       string `json:"custbody_rm_custom_status_for_pos"`
	CustbodyRmIsbCreateVendorBill      string `json:"custbody_rm_isb_create_vendorbill"`
	CustbodyRmIsbSendInvoice           string `json:"custbody_rm_isb_send_invoice"`
	CustbodyRmPrepayment               string `json:"custbody_rm_prepayment"`
	CustbodyRmProcessed                string `json:"custbody_rm_processed"`
	CustbodyRmVatDevStatus             string `json:"custbody_rm_vat_development_status"`
	CustbodySiiArticle61D              string `json:"custbody_sii_article_61d"`
	CustbodySiiArticle7273             string `json:"custbody_sii_article_72_73"`
	CustbodySiiIsThirdParty            string `json:"custbody_sii_is_third_party"`
	CustbodySiiNotReportedInTime       string `json:"custbody_sii_not_reported_in_time"`
	CustbodyStaCreditDueDate           string `json:"custbody_sta_credit_duedate"`
	CustbodyStaDeliveryDate            string `json:"custbody_sta_deliverydate"`
	CustbodyStaEinvoiceFactoring       string `json:"custbody_sta_einvoice_factoring"`
	CustbodyStaEinvoicingInvSend       string `json:"custbody_sta_einvoicing_inv_send_metho"`
	CustbodyStaEinvoicingPayeeAddr     string `json:"custbody_sta_einvoicing_payee_address"`
	CustbodyStaEinvoicingPayeeBban     string `json:"custbody_sta_einvoicing_payee_bban"`
	CustbodyStaEinvoicingPayeeBbanBic  string `json:"custbody_sta_einvoicing_payee_bban_bic"`
	CustbodyStaEinvoicingPayeeCity     string `json:"custbody_sta_einvoicing_payee_city"`
	CustbodyStaEinvoicingPayeeCountry  string `json:"custbody_sta_einvoicing_payee_country"`
	CustbodyStaEinvoicingPayeeEdi      string `json:"custbody_sta_einvoicing_payee_edi"`
	CustbodyStaEinvoicingPayeeIban     string `json:"custbody_sta_einvoicing_payee_iban"`
	CustbodyStaEinvoicingPayeeName     string `json:"custbody_sta_einvoicing_payee_name"`
	CustbodyStaEinvoicingPayeeVat      string `json:"custbody_sta_einvoicing_payee_vat"`
	CustbodyStaEinvoicingPayeeZip      string `json:"custbody_sta_einvoicing_payee_zip"`
	CustbodyStaEinvoicingPayerOvt      string `json:"custbody_sta_einvoicing_payer_ovt"`
	CustbodyStaEinvoicingReadyToSend   string `json:"custbody_sta_einvoicing_ready_tosend"`
	CustbodyStaEinvoicingSendConfirmed string `json:"custbody_sta_einvoicing_send_confirmed"`
	CustbodyStaFactoringInvNotToCust   string `json:"custbody_sta_factoring_inv_not_to_cust"`
	CustbodyStaFinSubtotal1            string `json:"custbody_sta_fin_subtotal1"`
	CustbodyStaFinSubtotal2            string `json:"custbody_sta_fin_subtotal2"`
	CustbodyStaFinTaxAmount1           string `json:"custbody_sta_fin_taxamount1"`
	CustbodyStaFinTaxAmount2           string `json:"custbody_sta_fin_taxamount2"`
	CustbodyStaFinTaxRate1             string `json:"custbody_sta_fin_taxrate1"`
	CustbodyStaFinTaxRate2             string `json:"custbody_sta_fin_taxrate2"`
	CustbodyStaFinTaxTotalAmountBase   string `json:"custbody_sta_fin_taxtotal_amount_base"`
	CustbodyStaFinTotal1               string `json:"custbody_sta_fin_total1"`
	CustbodyStaFinTotal2               string `json:"custbody_sta_fin_total2"`
	CustbodyStaFinVatBaseCurrency      string `json:"custbody_sta_fin_vat_base_currency"`
	CustbodyStaLocNoRefNo              string `json:"custbody_sta_loc_no_refno"`
	CustbodyStaNordRcText              string `json:"custbody_sta_nord_rctext"`
	CustbodyStaUseEN16931              string `json:"custbody_sta_use_en16931_2017_a1_2019"`
	CustbodyStaVatCategory1            string `json:"custbody_sta_vat_category1"`
	CustbodyStaVatCategory2            string `json:"custbody_sta_vat_category2"`
	CustomForm                         string `json:"customform"`
	CustomType                         string `json:"customtype"`
	DaysOpen                           string `json:"daysopen"`
	Email                              string `json:"email"`
	Entity                             string `json:"entity"`
	EstGrossProfit                     string `json:"estgrossprofit"`
	EstGrossProfitPercent              string `json:"estgrossprofitpercent"`
	ExchangeRate                       string `json:"exchangerate"`
	ExternalID                         string `json:"externalid"`
	ForeignPaymentAmountUnused         string `json:"foreignpaymentamountunused"`
	ForeignPaymentAmountUsed           string `json:"foreignpaymentamountused"`
	ForeignPaymentAmountUsedNoPost     string `json:"foreignpaymentamountusednopost"`
	ForeignTotal                       string `json:"foreigntotal"`
	ID                                 string `json:"id"`
	IsFinChrg                          string `json:"isfinchrg"`
	IsReversal                         string `json:"isreversal"`
	LastModifiedBy                     string `json:"lastmodifiedby"`
	LastModifiedDate                   string `json:"lastmodifieddate"`
	Memo                               string `json:"memo"`
	Memorized                          string `json:"memorized"`
	MergedIntoNewArrangements          string `json:"mergedintonewarrangements"`
	NeedsBill                          string `json:"needsbill"`
	Nexus                              string `json:"nexus"`
	OpeningBalanceTransaction          string `json:"openingbalancetransaction"`
	OrdPicked                          string `json:"ordpicked"`
	PaymentHold                        string `json:"paymenthold"`
	Posting                            string `json:"posting"`
	PostingPeriod                      string `json:"postingperiod"`
	PrevDate                           string `json:"prevdate"`
	PrintedPickingTicket               string `json:"printedpickingticket"`
	RecordType                         string `json:"recordtype"`
	ShipComplete                       string `json:"shipcomplete"`
	ShippingAddress                    string `json:"shippingaddress"`
	Source                             string `json:"source"`
	Status                             string `json:"status"`
	ToBeEmailed                        string `json:"tobeemailed"`
	ToBeFaxed                          string `json:"tobefaxed"`
	ToBePrinted                        string `json:"tobeprinted"`
	TotalCostEstimate                  string `json:"totalcostestimate"`
	TranDate                           string `json:"trandate"`
	TranDisplayName                    string `json:"trandisplayname"`
	TranID                             string `json:"tranid"`
	TransactionNumber                  string `json:"transactionnumber"`
	Type                               string `json:"type"`
	TypeBasedDocumentNumber            string `json:"typebaseddocumentnumber"`
	UseRevenueArrangement              string `json:"userevenuearrangement"`
	VisibleToCustomer                  string `json:"visibletocustomer"`
	Void                               string `json:"void"`
	Voided                             string `json:"voided"`
}

type InvoiceTransactions []InvoiceTransaction
type InvoiceTransaction struct {
	Links                              Links  `json:"links"`
	AbbrevType                         string `json:"abbrevtype"`
	BalSegStatus                       string `json:"balsegstatus"`
	BillingAddress                     string `json:"billingaddress"`
	BillingStatus                      string `json:"billingstatus"`
	CloseDate                          string `json:"closedate"`
	CreatedBy                          string `json:"createdby"`
	CreatedDate                        string `json:"createddate"`
	CreatedFromMerge                   string `json:"createdfrommerge"`
	Currency                           string `json:"currency"`
	Custbody15699ExcludeFromEP         string `json:"custbody_15699_exclude_from_ep_process"`
	CustbodyBsDdrDoNotReprocess        string `json:"custbody_bs_ddr_donot_reprocess"`
	CustbodyEdocGenTransPdf            string `json:"custbody_edoc_gen_trans_pdf"`
	CustbodyEiDsTxnIdentifier          string `json:"custbody_ei_ds_txn_identifier"`
	CustbodyEinvoicingDefaultPay       string `json:"custbody_einvoicing_defaultpaymtxt"`
	CustbodyFfBrExcludeTrans           string `json:"custbody_ff_br_exclude_transaction"`
	CustbodyFinInvoiceRefNumber        string `json:"custbody_fin_invoicereferencenumber"`
	CustbodyInvoiceCompanyEmail        string `json:"custbody_invoice_company_email"`
	CustbodyInvoiceCompanyPhone        string `json:"custbody_invoice_company_phone"`
	CustbodyNchCollectionForeign       string `json:"custbody_nch_collection_foreigncust"`
	CustbodyNchCollectionHold          string `json:"custbody_nch_collection_hold"`
	CustbodyNchCustBillingAddrEmail    string `json:"custbody_nch_cust_billingaddressemail"`
	CustbodyNchCustomerNumber          string `json:"custbody_nch_customernumber"`
	CustbodyNchEinvoicingPayeeID       string `json:"custbody_nch_einvoicing_payee_busin_id"`
	CustbodyNchEinvoicingPayeeNm       string `json:"custbody_nch_einvoicing_payee_name"`
	CustbodyNchInDispute               string `json:"custbody_nch_indispute"`
	CustbodyNchSettlPayUpdate          string `json:"custbody_nch_settl__pay_update"`
	CustbodyNchSource                  string `json:"custbody_nch_source"`
	CustbodyNsNordInvTypeCode          string `json:"custbody_ns_nord_inv_type_code"`
	CustbodyPayeeBic                   string `json:"custbody_payee_bic"`
	CustbodyPsgEiTransEdocStd          string `json:"custbody_psg_ei_trans_edoc_standard"`
	CustbodyReportTimestamp            string `json:"custbody_report_timestamp"`
	CustbodyRmAppliedInvoice           string `json:"custbody_rm_applied_invoice"`
	CustbodyRmCustomStatusForPos       string `json:"custbody_rm_custom_status_for_pos"`
	CustbodyRmIsbCreateVendorBill      string `json:"custbody_rm_isb_create_vendorbill"`
	CustbodyRmIsbSendInvoice           string `json:"custbody_rm_isb_send_invoice"`
	CustbodyRmPrepayment               string `json:"custbody_rm_prepayment"`
	CustbodyRmProcessed                string `json:"custbody_rm_processed"`
	CustbodyRmVatDevStatus             string `json:"custbody_rm_vat_development_status"`
	CustbodySiiArticle61D              string `json:"custbody_sii_article_61d"`
	CustbodySiiArticle7273             string `json:"custbody_sii_article_72_73"`
	CustbodySiiIsThirdParty            string `json:"custbody_sii_is_third_party"`
	CustbodySiiNotReportedInTime       string `json:"custbody_sii_not_reported_in_time"`
	CustbodyStaCreditDueDate           string `json:"custbody_sta_credit_duedate"`
	CustbodyStaDeliveryDate            string `json:"custbody_sta_deliverydate"`
	CustbodyStaEinvoiceFactoring       string `json:"custbody_sta_einvoice_factoring"`
	CustbodyStaEinvoicingInvSend       string `json:"custbody_sta_einvoicing_inv_send_metho"`
	CustbodyStaEinvoicingPayeeAddr     string `json:"custbody_sta_einvoicing_payee_address"`
	CustbodyStaEinvoicingPayeeBban     string `json:"custbody_sta_einvoicing_payee_bban"`
	CustbodyStaEinvoicingPayeeBbanBic  string `json:"custbody_sta_einvoicing_payee_bban_bic"`
	CustbodyStaEinvoicingPayeeCity     string `json:"custbody_sta_einvoicing_payee_city"`
	CustbodyStaEinvoicingPayeeCountry  string `json:"custbody_sta_einvoicing_payee_country"`
	CustbodyStaEinvoicingPayeeEdi      string `json:"custbody_sta_einvoicing_payee_edi"`
	CustbodyStaEinvoicingPayeeIban     string `json:"custbody_sta_einvoicing_payee_iban"`
	CustbodyStaEinvoicingPayeeName     string `json:"custbody_sta_einvoicing_payee_name"`
	CustbodyStaEinvoicingPayeeVat      string `json:"custbody_sta_einvoicing_payee_vat"`
	CustbodyStaEinvoicingPayeeZip      string `json:"custbody_sta_einvoicing_payee_zip"`
	CustbodyStaEinvoicingPayerOvt      string `json:"custbody_sta_einvoicing_payer_ovt"`
	CustbodyStaEinvoicingReadyToSend   string `json:"custbody_sta_einvoicing_ready_tosend"`
	CustbodyStaEinvoicingSendConfirmed string `json:"custbody_sta_einvoicing_send_confirmed"`
	CustbodyStaFactoringInvNotToCust   string `json:"custbody_sta_factoring_inv_not_to_cust"`
	CustbodyStaFinSubtotal1            string `json:"custbody_sta_fin_subtotal1"`
	CustbodyStaFinSubtotal2            string `json:"custbody_sta_fin_subtotal2"`
	CustbodyStaFinTaxAmount1           string `json:"custbody_sta_fin_taxamount1"`
	CustbodyStaFinTaxAmount2           string `json:"custbody_sta_fin_taxamount2"`
	CustbodyStaFinTaxRate1             string `json:"custbody_sta_fin_taxrate1"`
	CustbodyStaFinTaxRate2             string `json:"custbody_sta_fin_taxrate2"`
	CustbodyStaFinTaxTotalAmountBase   string `json:"custbody_sta_fin_taxtotal_amount_base"`
	CustbodyStaFinTotal1               string `json:"custbody_sta_fin_total1"`
	CustbodyStaFinTotal2               string `json:"custbody_sta_fin_total2"`
	CustbodyStaFinVatBaseCurrency      string `json:"custbody_sta_fin_vat_base_currency"`
	CustbodyStaLocNoRefNo              string `json:"custbody_sta_loc_no_refno"`
	CustbodyStaNordRcText              string `json:"custbody_sta_nord_rctext"`
	CustbodyStaUseEN16931              string `json:"custbody_sta_use_en16931_2017_a1_2019"`
	CustbodyStaVatCategory1            string `json:"custbody_sta_vat_category1"`
	CustbodyStaVatCategory2            string `json:"custbody_sta_vat_category2"`
	CustbodyStcAmountAfterDiscount     string `json:"custbody_stc_amount_after_discount"`
	CustbodyStcTaxAfterDiscount        string `json:"custbody_stc_tax_after_discount"`
	CustbodyStcTotalAfterDiscount      string `json:"custbody_stc_total_after_discount"`
	CustomForm                         string `json:"customform"`
	CustomType                         string `json:"customtype"`
	DaysOpen                           string `json:"daysopen"`
	Daysoverduesearch                  string `json:"daysoverduesearch"`
	DueDate                            string `json:"duedate"`
	Email                              string `json:"email"`
	Entity                             string `json:"entity"`
	EstGrossProfit                     string `json:"estgrossprofit"`
	EstGrossProfitPercent              string `json:"estgrossprofitpercent"`
	ExchangeRate                       string `json:"exchangerate"`
	ExternalID                         string `json:"externalid"`
	ForeignAmountPaid                  string `json:"foreignamountpaid"`
	ForeignAmountUnpaid                string `json:"foreignamountunpaid"`
	ForeignTotal                       string `json:"foreigntotal"`
	ID                                 string `json:"id"`
	IsFinChrg                          string `json:"isfinchrg"`
	IsReversal                         string `json:"isreversal"`
	LastModifiedBy                     string `json:"lastmodifiedby"`
	LastModifiedDate                   string `json:"lastmodifieddate"`
	Memo                               string `json:"memo"`
	Memorized                          string `json:"memorized"`
	MergedIntoNewArrangements          string `json:"mergedintonewarrangements"`
	NeedsBill                          string `json:"needsbill"`
	Nexus                              string `json:"nexus"`
	OpeningBalanceTransaction          string `json:"openingbalancetransaction"`
	OrdPicked                          string `json:"ordpicked"`
	PaymentHold                        string `json:"paymenthold"`
	Posting                            string `json:"posting"`
	PostingPeriod                      string `json:"postingperiod"`
	PrevDate                           string `json:"prevdate"`
	PrintedPickingTicket               string `json:"printedpickingticket"`
	RecordType                         string `json:"recordtype"`
	ShipComplete                       string `json:"shipcomplete"`
	ShipDate                           string `json:"shipdate"`
	ShippingAddress                    string `json:"shippingaddress"`
	Source                             string `json:"source"`
	Status                             string `json:"status"`
	Terms                              string `json:"terms"`
	ToBeEmailed                        string `json:"tobeemailed"`
	ToBeFaxed                          string `json:"tobefaxed"`
	ToBePrinted                        string `json:"tobeprinted"`
	TotalCostEstimate                  string `json:"totalcostestimate"`
	TranDate                           string `json:"trandate"`
	TranDisplayName                    string `json:"trandisplayname"`
	TranID                             string `json:"tranid"`
	TransactionNumber                  string `json:"transactionnumber"`
	Type                               string `json:"type"`
	TypeBasedDocumentNumber            string `json:"typebaseddocumentnumber"`
	UseRevenueArrangement              string `json:"userevenuearrangement"`
	VisibleToCustomer                  string `json:"visibletocustomer"`
	Void                               string `json:"void"`
	Voided                             string `json:"voided"`
}
