package netsuite

import (
	"encoding/json"

	"github.com/cydev/zero"
	"github.com/omniboost/go-netsuite-rest/omitempty"
)

type JournalEntry struct {
	AccountingBook NSResource `json:"accountingBook,omitzero"`
	// AccountingBookDetail AccountingBookDetail `json:"accountingBookDetail,omitzero"`
	// AppliedRules AppliedRules `json:"appliedRules,omitzero"`
	ApprovalStatus         RecordRef        `json:"approvalStatus,omitzero"`
	Approved               bool             `json:"approved"`
	Class                  Classification   `json:"Class,omitzero"`
	CreatedDate            Date             `json:"createdDate,omitzero"`
	CreatedFrom            NSResource       `json:"createdFrom,omitzero"`
	Currency               Currency         `json:"currency,omitzero"`
	CustomForm             CustomForm       `json:"customForm,omitzero"`
	Department             Department       `json:"department,omitzero"`
	ExchangeRate           float64          `json:"exchangeRate,omitempty"`
	ExcludeFromGLNumbering bool             `json:"excludeFromGLNumbering,omitempty"`
	ExternalID             string           `json:"externalId,omitempty"`
	ID                     string           `json:"id,omitempty"`
	IsReversal             bool             `json:"isReversal,omitempty"`
	LastModifiedDate       Date             `json:"lastModifiedDate,omitzero"`
	Line                   JournalEntryLine `json:"line"`
	Location               Location         `json:"location,omitzero"`
	Memo                   string           `json:"memo"`
	// NextApprover           Employee         `json:"nextApprover,omitzero"`
	Nexus                Nexus            `json:"nexus,omitzero"`
	ParentExpenseAlloc   NSResource       `json:"parentExpenseAlloc,omitzero"`
	PostingPeriod        AccountingPeriod `json:"postingPeriod,omitzero"`
	RefName              string           `json:"refName,omitempty"`
	ReversalDate         Date             `json:"reversalDate,omitzero"`
	ReversalDefer        bool             `json:"reversalDefer,omitempty"`
	Subsidiary           Subsidiary       `json:"subsidiary,omitzero"`
	SubsidiaryTaxRegNum  string           `json:"subsidiaryTaxRegNum,omitempty"`
	TaxDetailsOverride   bool             `json:"taxDetailsOverride,omitempty"`
	TaxPointDate         Date             `json:"taxPointDate,omitzero"`
	TaxPointDateOverride bool             `json:"taxPointDateOverride,omitempty"`
	TaxRegOverride       bool             `json:"taxRegOverride,omitempty"`
	ToSubsidiary         Subsidiary       `json:"toSubsidiary,omitzero"`
	TranDate             Date             `json:"tranDate,omitzero"`
	TranID               string           `json:"tranId,omitempty"`
	Void                 bool             `json:"void,omitempty"`

	CustomFields map[string]NSResource `json:"-"`
}

// func (j JournalEntry) MarshalJSON() ([]byte, error) {
// 	// TODO: implement marshalling of cust bodies & cols
// 	return omitempty.MarshalJSON(j)
// }

// func (j *JournalEntry) UnmarshalJSON(data []byte) error {
// 	data, err := UnmarshalCustomFields(data, j.CustomFields)
// 	if err != nil {
// 		return err
// 	}
//
// 	type Alias JournalEntry
// 	return json.Unmarshal(data, (*Alias)(j))
// }

func (j JournalEntry) IsZero() bool {
	return zero.IsZero(j)
}

type JournalEntryLine struct {
	Links        Links                    `json:"links,omitempty"`
	Items        JournalEntryLineElements `json:"items"`
	TotalResults int                      `json:"totalResults,omitempty"`
}

type Currencies []Currency

// https://system.netsuite.com/help/helpcenter/en_US/APIs/REST_API_Browser/record/v1/2025.1/index.html#/definitions/currency
type Currency struct {
	// Displays the precision of the currency, which designates the number of
	// digits to the right of the decimal point used in currency transactions.
	// Precision can be zero or two. The level of decimal precision indicated is
	// used for inventory costing calculations to maintains consistency between
	// inventory costing and reporting. Values in report results are rounded to
	// the base currency precision. This rounding applies to currency values and
	// non-currency values, including formula column values. To change this
	// read-only field to a dropdown list through which you can change the
	// precision from zero or two, contact NetSuite Technical Support.
	CurrencyPrecision int `json:"currencyPrecision,omitempty"`
	// Enter the currency symbol and text to use for this currency. Include
	// spaces if you want to separate the symbol from the currency value. For
	// example, $ USD or $CAD.
	DisplaySymbol string `json:"displaySymbol,omitempty"`
	// Enter an exchange rate for this currency against the base currency of
	// this company, or if you use OneWorld, for this currency against the base
	// currency of the root parent subsidiary. The exchange rate is equal to the
	// base currency amount divided by the foreign currency amount. For example,
	// if your company is located in Canada (base currency) and you are defining
	// the U.S. dollar (foreign currency), and the current exchange rate is
	// 1.02 Canadian dollars to 1.00 U.S. dollar, the Default Exchange Rate for
	// the U.S. dollar is 1.02/1.00, or 1.02. This rate is the basis for rates
	// in the Currency Exchange Rates table that are used in foreign currency
	// transactions. If you use OneWorld, this rate also is the basis for rates
	// in the Consolidated Exchange Rates table that are used in consolidated
	// financials. For more information, see the help topic Currency Exchange
	// Rates.
	ExchangeRate float64 `json:"exchangeRate,omitempty"`
	ExternalID   string  `json:"externalId,omitempty"`
	// This field displays a sample of how currency amounts display for the
	// selected format. The decimal precision shown cannot be changed. Note:
	// The decimal precision shown is the precision used for both inventory
	// reporting and for costing calculations.
	FormatSample         string     `json:"formatSample,omitempty"`
	FxRateUpdateTimezone NSResource `json:"fxRateUpdateTimezone,omitzero"`
	ID                   string     `json:"id,omitempty"`
	// Check this box to update currency exchange rates daily.
	IncludeInFxRateUpdates bool `json:"includeInFxRateUpdates,omitzero"`
	// A check in this box indicates that the currency has been selected as an
	// anchor currency in the accounting preferences. To clear the box, change
	// the selection in the accounting preference under Use Triangulation
	// Calculation by NetSuite. If this currency is a designated anchor currency
	// and has been used in an exchange rate calculation, you cannot delete this
	// currency. For more information about triangulation and anchor currencies,
	// see the help topics Methods for Obtaining Exchange Rates andAnchor
	// Currencies.
	IsAnchorCurrency bool `json:"isAnchorCurrency,omitzero"`
	// Indicates that this currency is the company's base currency or in
	// OneWorld accounts, the base currency for a subsidiary. Note: After you
	// have entered transactions in foreign currencies, you cannot change a
	// base currency.
	IsBaseCurrency bool `json:"isBaseCurrency,omitzero"`
	// Check this box to make the currency record is inactive, or clear it to
	// make the record active. You cannot make a currency inactive if any open
	// transactions exist in that currency.
	IsInactive       bool       `json:"isInactive,omitzero"`
	LastModifiedDate Date       `json:"lastModifiedDate,omitzero"`
	Links            Links      `json:"links,omitempty"`
	Locale           NSResource `json:"locale,omitzero"`
	// Enter a unique name for the currency. Because many countries use
	// the same name for their currencies, you should use a combined name that
	// includes the country name or abbreviation as well as the name of the
	// currency. For example, pesos are the currency in the Philippines,
	// Uruguay, and Mexico. In the Name field, you might enter “Mexican peso.”
	Name string `json:"name,omitempty"`
	// Check this box to customize the currency format.
	OverrideCurrencyFormat bool   `json:"overrideCurrencyFormat,omitzero"`
	RefName                string `json:"refName,omitempty"`
	// Enter the three-letter International Standards Organization (ISO)
	Symbol          string     `json:"symbol,omitempty"`
	SymbolPlacement NSResource `json:"symbolPlacement,omitzero"`
}

func (c Currency) IsZero() bool {
	return zero.IsZero(c)
}

type AccountingPeriod struct {
	Links   Links  `json:"links"`
	ID      string `json:"id"`
	RefName string `json:"refName"`
}

func (p AccountingPeriod) IsZero() bool {
	return zero.IsZero(p)
}

type Subsidiaries []Subsidiary

type Subsidiary struct {
	Links              Links            `json:"links,omitzero"`
	ClassTranslation   ClassTranslation `json:"classTranslation,omitzero"`
	Country            NSResource       `json:"country,omitzero"`
	Currency           NSResource       `json:"currency,omitzero"`
	ID                 string           `json:"id,omitempty"`
	IsElimination      *bool            `json:"isElimination,omitempty"`
	IsInactive         *bool            `json:"isInactive,omitempty"`
	LastModifiedDate   Date             `json:"lastModifiedDate,omitzero"`
	LegalName          string           `json:"legalName,omitempty"`
	MainAddress        Address          `json:"mainAddress,omitzero"`
	Name               string           `json:"name,omitempty"`
	Parent             NSResource       `json:"parent,omitzero"`
	RefName            string           `json:"refName,omitempty"`
	ReturnAddress      ReturnAddress    `json:"returnAddress,omitzero"`
	ShippingAddress    ShippingAddress  `json:"shippingAddress,omitzero"`
	State              string           `json:"state,omitempty"`
	TranInternalPrefix string           `json:"tranInternalPrefix,omitempty"`
	URL                string           `json:"url,omitempty"`

	CustomFields map[string]any `json:"-"`
}

func (s *Subsidiary) UnmarshalJSON(data []byte) error {
	err := UnmarshalCustomFields(data, &s.CustomFields)
	if err != nil {
		return err
	}

	type Alias Subsidiary
	return json.Unmarshal(data, (*Alias)(s))
}

func (s Subsidiary) IsZero() bool {
	return zero.IsZero(s)
}

type ClassTranslationElement struct {
	Links    Links  `json:"links,omitempty"`
	Language string `json:"language,omitempty"`
	Locale   struct {
		ID      string `json:"id,omitempty"`
		RefName string `json:"refName,omitempty"`
	} `json:"locale,omitempty"`
}

func (s ClassTranslationElement) IsZero() bool {
	return zero.IsZero(s)
}

type ClassTranslation struct {
	Links        Links                     `json:"links,omitempty"`
	Items        []ClassTranslationElement `json:"items,omitempty"`
	TotalResults int                       `json:"totalResults,omitempty"`
}

func (s ClassTranslation) IsZero() bool {
	return zero.IsZero(s)
}

type CustomForm struct {
	ID      string `json:"id"`
	RefName string `json:"refName"`
}

func (c CustomForm) IsZero() bool {
	return zero.IsZero(c)
}

type JournalEntryLineElements []JournalEntryLineElement

type JournalEntryLineElement struct {
	Account                Account    `json:"account,omitzero"`
	Class                  RecordRef  `json:"class,omitzero"`
	Cleared                Bool       `json:"cleared,omitzero"`
	ClearedDate            Date       `json:"clearedDate,omitzero"`
	Credit                 float64    `json:"credit,omitempty"`
	CreditTax              float64    `json:"creditTax,omitempty"`
	Debit                  float64    `json:"debit,omitempty"`
	DebitTax               float64    `json:"debitTax,omitempty"`
	Department             RecordRef  `json:"Department,omitzero"`
	Eliminate              Bool       `json:"eliminate,omitzero"`
	EndDate                Date       `json:"endDate,omitzero"`
	Entity                 NSResource `json:"entity,omitzero"`
	GiftCertCode           NSResource `json:"giftCertCode,omitzero"`
	Item                   NSResource `json:"item,omitzero"`
	Line                   int        `json:"line,omitempty"`
	LineCreatedDate        Date       `json:"lineCreatedDate,omitzero"`
	LineLastModifiedDate   Date       `json:"lineLastModifiedDate,omitzero"`
	LineTaxCode            NSResource `json:"lineTaxCode,omitzero"`
	LineTaxRate            float64    `json:"lineTaxRate,omitempty"`
	Links                  Links      `json:"links,omitzero"`
	Location               Location   `json:"location,omitzero"`
	Memo                   string     `json:"memo"`
	RefName                string     `json:"refName,omitempty"`
	Residual               float64    `json:"residual,omitempty"`
	RevenueRecognitionRule NSResource `json:"revenueRecognitionRule,omitzero"`
	Schedule               NSResource `json:"schedule,omitzero"`
	ScheduleNum            NSResource `json:"scheduleNum,omitzero"`
	ScheduleType           NSResource `json:"scheduleType,omitzero"`
	StartDate              Date       `json:"startDate,omitzero"`
	TaxAccount             TaxAccount `json:"taxAccount,omitzero"`
	TaxBasis               float64    `json:"taxBasis,omitempty"`
	TotalAmount            float64    `json:"totalAmount,omitempty"`
}

func (j JournalEntryLineElement) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(j)
}

type Accounts []Account

// https://system.netsuite.com/help/helpcenter/en_US/APIs/REST_API_Browser/record/v1/2025.1/index.html#/definitions/account
type Account struct {
	AccountContextSearch         AccountContextSearchCollection `json:"accountContextSearch,omitzero"`
	AccountSearchDisplayName     string                         `json:"accountSearchDisplayName,omitempty"`
	AccountSearchDisplayNameCopy string                         `json:"accountSearchDisplayNameCopy,omitempty"`
	// Enter up to 31 characters for an account name that will appear on all
	// reports. If you want to use GL account numbers and do not see a Number
	// field above the Name field, go to Setup > Accounting > Preferences >
	// Accounting Preferences, and check the Use Account Numbers box. Save the
	// preference, and return to this page.
	AcctName string `json:"acctName,omitempty"`
	// Enter the number to identify this account. The number can be alphanumeric. The maximum number of characters is 60.
	AcctNumber       string     `json:"acctNumber,omitempty"`
	AcctType         NSResource `json:"acctType,omitzero"`
	AvailableBalance float64    `json:"availableBalance,omitempty"`
	// Check this box to signify this is an account used with ACH processing.
	// ACH processing is used for the Electronic Funds Transfer (EFT) and Direct
	// Deposit features.
	BACH                 bool           `json:"bAch,omitzero"`
	Balance              float64        `json:"balance,omitempty"`
	BillableExpensesAcct NSResource     `json:"billableExpensesAcct,omitzero"`
	CashFlowRate         NSResource     `json:"cashFlowRate,omitzero"`
	Category1099Misc     NSResource     `json:"category1099Misc,omitzero"`
	Class                Classification `json:"class,omitzero"`
	// Enter a number to reset the default check number that appears on
	// transactions such as checks and bill payments. By default, this field
	// shows the highest check number in the account plus one. The number you
	// enter in the Next Check Number field of the account record determines the
	// number that appears in the Check Number field on transactions linked to
	// that account.
	CurDocNum    int64      `json:"curDocNum,omitempty"`
	Currency     Currency   `json:"currency,omitzero"`
	DeferralAcct NSResource `json:"deferralAcct,omitzero"`
	Department   Department `json:"department,omitzero"`
	// Enter a description for this account.
	Description              string `json:"description,omitempty"`
	DisplayNameWithHierarchy string `json:"displayNameWithHierarchy,omitempty"`
	// Check this box to make this account an intercompany account. Intercompany
	// accounts are used to record transactions between subsidiaries. You can
	// post both intercompany transactions and non-intercompany transactions to
	// most intercompany accounts. Intercompany Accounts Receivable and
	// Intercompany Accounts Payable, however, can be used only for recording
	// amounts that are candidates for eliminations. For details, see the help
	// topic Intercompany Accounts.
	Eliminate  bool   `json:"eliminate,omitzero"`
	ExternalID string `json:"externalId,omitempty"`
	FullName   string `json:"fullName,omitempty"`
	// Check the Include Children box to share the account with all the
	// sub-subsidiaries associated with each subsidiary selected in the
	// Subsidiary field.
	GeneralRate     NSResource `json:"generalRate,omitzero"`
	ID              string     `json:"id,omitempty"`
	IncludeChildren bool       `json:"includeChildren,omitzero"`
	// If this will be an Other Current Asset account and you want the balance
	// of this account included in the total balance of the Inventory KPI,
	// select the Inventory box.
	Inventory bool `json:"inventory,omitzero"`
	// Check this box to inactivate this account. Inactive accounts do not show
	// in lists on transactions and records.
	IsInactive bool `json:"isInactive,omitzero"`
	// Check this box to make this account record solely for reporting
	// purposes. Summary accounts are useful when you want to create a
	// non-posting, inactive parent account that has active child accounts. If
	// you do not have a OneWorld account, new summary accounts cannot have an
	// opening balance, but you can convert an existing account with a
	// transaction balance into a summary account. In this case, you cannot
	// post additional transactions to the account. Summary accounts appear
	// with their children in the chart of accounts list. You cannot merge a
	// summary account into another account.
	IsSummary        bool `json:"isSummary,omitzero"`
	LastModifiedDate Date `json:"lastModifiedDate,omitzero"`
	// If the Use Legal Name in Account accounting preference is enabled at
	// Setup > Accounting > Preferences > Accounting Preferences, the Legal
	// Name field appears. You can enter up to 400 characters in this field
	// including special characters such as colon and semi colon. You can also
	// enter characters such as Éé,Çç, and 2 byte characters such as 会計、科目.
	// This field is useful in countries where the legal name of an entity is
	// required by law. The Legal Name field can be added to financial reports
	// where account is an available component. It is also available in
	// advanced searches, SuiteScript, and ODBC. System notes maintains an
	// audit trail specific to the activity on the Legal Name field.
	LegalName     string                         `json:"legalName,omitempty"`
	Links         Links                          `json:"links,omitempty"`
	Localizations AccountLocalizationsCollection `json:"localizations,omitzero"`
	Location      Location                       `json:"location,omitzero"`
	// Enter the dollar amount limit per transaction with this account.
	MMaxAmtPerTran float64 `json:"mMaxAmtPerTran,omitempty"`
	// Enter an opening balance for this account.
	OpeningBalance float64  `json:"openingBalance,omitempty"`
	Parent         *Account `json:"parent,omitzero"`
	// Check this box if you want to use the Confirm Transaction Matches and
	// Reconcile Account Statement pages for this account. The original
	// reconciliation pages are still available. To use the original pages,
	// clear the box.
	ReconcileWithMatching    bool                 `json:"reconcileWithMatching,omitzero"`
	RefName                  string               `json:"refName,omitempty"`
	RestrictToAccountingBook NSResourceCollection `json:"restrictToAccountingBook,omitzero"`
	// Check this box to select this account for open balance currency
	Revalue bool `json:"revalue,omitzero"`
	// Enter a default ACH message for this account.
	SAchMsg string `json:"sAchMsg,omitempty"`
	// Enter the bank account number, up to 20 digits.
	SBankCompanyID string `json:"sBankCompanyId,omitempty"`
	// Enter the name of your bank.
	SBankName string `json:"sBankName,omitempty"`
	// Enter the 9-digit routing number for this bank account.
	SBankRoutingNumber string               `json:"sBankRoutingNumber,omitempty"`
	SSpecAcct          NSResource           `json:"sSpecAcct,omitzero"`
	Subsidiary         SubsidiaryCollection `json:"subsidiary,omitzero"`
	// Enter the date of the opening balance of this account.
	TranDate Date `json:"tranDate,omitzero"`
	// This field displays the base unit assigned to the Unit Type. The
	// default unit cannot be changed.
	Unit      string    `json:"unit,omitempty"`
	UnitsType UnitsType `json:"unitsType,omitzero"`
}

func (a Account) IsZero() bool {
	return zero.IsZero(a)
}

type Invoice struct {
	Links Links `json:"links,omitzero"`
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
	AmountPaid              float64 `json:"amountPaid"`
	AmountRemaining         float64 `json:"amountRemaining"`
	AmountRemainingTotalBox float64 `json:"amountRemainingTotalBox"`
	BillAddress             string  `json:"billAddress"`
	BillingAddress          Address `json:"billingAddress"`
	BillingAddressText      string  `json:"billingAddress_text"`
	CreatedDate             Date    `json:"createdDate"`
	Currency                struct {
		Links   Links  `json:"links"`
		ID      string `json:"id"`
		RefName string `json:"refName"`
	} `json:"currency"`
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
	CustomForm CustomForm `json:"customForm,omitzero"`
	DueDate    Date       `json:"dueDate,omitzero"`
	Entity     struct {
		Links   Links  `json:"links,omitzero"`
		ID      string `json:"id"`
		RefName string `json:"refName,omitempty"`
	} `json:"entity"`
	// EstGrossProfit         float64     `json:"estGrossProfit"`
	// EstGrossProfitPercent  float64     `json:"estGrossProfitPercent"`
	// ExchangeRate           float64     `json:"exchangeRate"`
	// ExcludeFromGLNumbering Bool        `json:"excludeFromGLNumbering"`
	ID   string                `json:"id,omitempty"`
	Item InvoiceItemCollection `json:"item"`
	// LastModifiedDate       Date        `json:"lastModifiedDate"`
	Location Location `json:"location"`
	Memo     string   `json:"memo"`
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
	ToBeEmailed Bool    `json:"toBeEmailed"`
	ToBeFaxed   Bool    `json:"toBeFaxed"`
	ToBePrinted Bool    `json:"toBePrinted"`
	Total       float64 `json:"total"`
	// TotalAfterTaxes      float64 `json:"totalAfterTaxes"`
	// TotalCostEstimate    float64 `json:"totalCostEstimate"`
	TranDate   Date      `json:"tranDate"`
	TranID     string    `json:"tranId"`
	Department RecordRef `json:"Department,omitzero"`
	Class      RecordRef `json:"Class,omitzero"`

	CustomFields map[string]any `json:"-"`
}

func (i Invoice) MarshalJSON() ([]byte, error) {
	type Alias Invoice
	return MarshalCustomFields(Alias(i), i.CustomFields)
}

func (i *Invoice) UnmarshalJSON(data []byte) error {
	err := UnmarshalCustomFields(data, &i.CustomFields)
	if err != nil {
		return err
	}

	type Alias Invoice
	return json.Unmarshal(data, (*Alias)(i))
}

type CreditMemo struct {
	Account              Account                        `json:"account,omitzero"`
	AccountingBookDetail AccountingBookDetailCollection `json:"accountingBookDetail,omitzero"`
	// The handling cost automatically calculates depending on the shipping
	// method you select in the Ship Via field. To change the cost of handling,
	// go to Lists > Shipping Items and select the shipping method with the
	// handling cost you want to change.
	AltHandlingCost float64 `json:"altHandlingCost,omitempty"`
	// The shipping cost automatically calculates depending on the shipping
	// method you select in the Ship Via field above. To change the cost of a
	// shipping method, go to Lists > Shipping Items and select the shipping
	// method you want to change. If you use UPS Real-Time rates, shipments
	// over 150lbs are broken up into shipments less than or equal to 150lbs
	// for charging.
	AltShippingCost         float64 `json:"altShippingCost,omitempty"`
	AmountPaid              float64 `json:"amountPaid,omitempty"`
	AmountRemaining         float64 `json:"amountRemaining,omitempty"`
	AmountRemainingTotalBox float64 `json:"amountRemainingTotalBox,omitempty"`
	// NetSuite displays the amount of credit you applied below.
	Applied      float64                `json:"applied,omitempty"`
	AppliedRules AppliedRulesCollection `json:"appliedRules,omitzero"`
	Apply        ApplyCollection        `json:"apply,omitzero"`
	// Type an as of date for this customer credit.
	AsOfDate  Date   `json:"asOfDate,omitempty"`
	BillAddr1 string `json:"billAddr1,omitempty"`
	BillAddr2 string `json:"billAddr2,omitempty"`
	BillAddr3 string `json:"billAddr3,omitempty"`
	// The default billing address autofills this field from the customer's
	// record at Lists > Relationships > Customers. . To enter a different
	// address: * Select another address in the Bill To Select field. * Select
	// New in the Bill To Select field to enter a new billing address.
	BillAddress                   string         `json:"billAddress,omitempty"`
	BillAddressList               NSResource     `json:"billAddressList,omitzero"`
	BillAddressee                 string         `json:"billAddressee,omitempty"`
	BillAttention                 string         `json:"billAttention,omitempty"`
	BillCity                      string         `json:"billCity,omitempty"`
	BillCountry                   NSResource     `json:"billCountry,omitzero"`
	BillOverride                  string         `json:"billOverride,omitempty"`
	BillPhone                     string         `json:"billPhone,omitempty"`
	BillState                     string         `json:"billState,omitempty"`
	BillZip                       string         `json:"billZip,omitempty"`
	BillingAccount                BillingAccount `json:"billingAccount,omitzero"`
	BillingAddress                BillingAddress `json:"billingAddress,omitzero"`
	BillingAddressText            string         `json:"billingAddress_text,omitempty"`
	BulkProcSubmission            NSResource     `json:"bulkProcSubmission,omitzero"`
	CanHaveStackable              bool           `json:"canHaveStackable,omitzero"`
	Class                         Classification `json:"class,omitzero"`
	ContribPct                    string         `json:"contribPct,omitempty"`
	CouponCode                    CouponCode     `json:"couponCode,omitzero"`
	CreatedDate                   Date           `json:"createdDate,omitzero"`
	CreatedFrom                   NSResource     `json:"createdFrom,omitzero"`
	Currency                      Currency       `json:"currency,omitzero"`
	CustomForm                    NSResource     `json:"customForm"`
	DefaultILBAShipAddrVal        string         `json:"defaultILBAShipAddrVal,omitempty"`
	DefaultILBAShippingAddressKey string         `json:"defaultILBAShippingAddressKey,omitempty"`
	DefaultILShipMethKey          int64          `json:"defaultILShipMethKey,omitzero"`
	// Deferred Revenue: amount of revenue deferred on this transaction
	DeferredRevenue float64      `json:"deferredRevenue,omitempty"`
	Department      Department   `json:"department,omitzero"`
	DiscountAmount  float64      `json:"discountAmount,omitempty"`
	DiscountDate    Date         `json:"discountDate,omitzero"`
	DiscountItem    DiscountItem `json:"discountItem,omitzero"`
	// NetSuite enters the rate for the discount item you selected. You can
	// change the discount rate for this cash refund. Enter the discount as a
	// dollar amount like 10.00, or as a percentage like 10%.
	DiscountRate float64 `json:"discountRate,omitempty"`
	// NetSuite enters the amount discounted on this credit memo. If this
	// discount item is taxable, the discount is applied before taxes. If it is
	// not taxable, the discount is applied after taxes.
	DiscountTotal   float64    `json:"discountTotal,omitempty"`
	DueDate         Date       `json:"dueDate,omitzero"`
	Email           string     `json:"email,omitempty"`
	EndDate         Date       `json:"endDate,omitzero"`
	Entity          NSResource `json:"entity"`
	EntityTaxRegNum string     `json:"entityTaxRegNum,omitempty"`
	// Estimated Gross Profit: Read-only field that equals the revenue amount minus the Est. Cost. At the transaction level, it equals the gross profit of all lines, factoring transaction-level discounts and markups.
	EstGrossProfit float64 `json:"estGrossProfit,omitempty"`
	// Estimated Gross Profit Margin, as a percentage: Read-only field that equals the Est. Gross Profit divided by revenue, expressed as a percentage. At the transaction level, it equals the gross profit percent of all lines, factoring transaction-level discounts and markups.
	EstGrossProfitPercent float64 `json:"estGrossProfitPercent,omitempty"`
	// The currency's exchange rate is shown in this field. You can edit the
	// exchange rate for this transaction only, or you can update the currency
	// record with the exchange rate you enter here.
	ExchangeRate float64 `json:"exchangeRate,omitempty"`
	// Check this option to exclude this transaction and its subordinate
	// transactions from inclusion in all commission calculations.
	ExcludeCommission      bool       `json:"excludeCommission,omitzero"`
	ExcludeFromGLNumbering bool       `json:"excludeFromGLNumbering,omitzero"`
	ExternalId             string     `json:"externalId,omitempty"`
	HandlingCost           float64    `json:"handlingCost,omitempty"`
	HandlingTaxAmount      float64    `json:"handlingTaxAmount,omitempty"`
	ID                     string     `json:"id,omitempty"`
	IntercoStatus          NSResource `json:"intercoStatus,omitzero"`
	IntercoTransaction     NSResource `json:"intercoTransaction,omitzero"`
	// Check this box if you want to allow shipping to more than one address,
	// using alternative shipping methods, for this transaction. Checking
	// this box adds the fields Ship To, Carrier and Ship Via fields to the
	// Item line. Search "Multiple Shipping Routes" in the Help Center for
	// more information.
	IsMultiShipTo    bool                     `json:"isMultiShipTo,omitzero"`
	Item             CreditMemoItemCollection `json:"item,omitzero"`
	Job              NSResource               `json:"job,omitzero"`
	LastModifiedDate Date                     `json:"lastModifiedDate,omitzero"`
	LeadSource       NSResource               `json:"leadSource,omitzero"`
	Links            Links                    `json:"links,omitempty"`
	Location         Location                 `json:"location,omitzero"`
	// If you wish, enter a memo to describe this credit. It will appear on
	// reports such as the 2-line Accounts Receivable register that your
	// customers can see if you give them permission to log in and view their
	// transaction history.
	Memo string `json:"memo,omitempty"`
	// Select a customer message that will appear on the Credit Memo form. To
	// add additional choices to this list, go to Setup > Accounting >
	// Accounting Lists > New > Customer Message.
	Message       string     `json:"message,omitempty"`
	MessageSel    NSResource `json:"messageSel,omitzero"`
	MuccPromoCode string     `json:"muccPromoCode,omitempty"`
	Nexus         Nexus      `json:"nexus,omitzero"`
	Originator    string     `json:"originator,omitempty"`
	// For your customers' convenience, you can enter their purchase order
	// number here.
	OtherRefNum         string             `json:"otherRefNum,omitempty"`
	Partner             Partner            `json:"partner,omitzero"`
	Partners            PartnersCollection `json:"partners,omitzero"`
	PostingPeriod       AccountingPeriod   `json:"postingPeriod,omitzero"`
	PrevDate            Date               `json:"prevDate,omitzero"`
	PrevPartner         int64              `json:"prevPartner,omitempty"`
	PrevRep             int64              `json:"prevRep,omitempty"`
	PromoCode           PromotionCode      `json:"promoCode,omitzero"`
	PromoCodePluginImpl string             `json:"promoCodePluginImpl,omitempty"`
	// Recognized Revenue: cumulative amount of revenue recognized for this transaction
	RecognizedRevenue    float64 `json:"recognizedRevenue,omitempty"`
	RecurringBill        bool    `json:"recurringBill,omitzero"`
	RefName              string  `json:"refName,omitempty"`
	RevCommitStatus      string  `json:"revCommitStatus,omitempty"`
	RevCommitStatusDescr string  `json:"revCommitStatusDescr,omitempty"`
	// Check this check box to create a Revenue Commitment or Revenue
	// Commitment Reversal. * On a Sales Order, the Revenue Commitment
	// replaces an invoice * On a Return Authorization, the Revenue Commitment
	// Reversal replaces a credit memo. NetSuite creates all applicable
	// Revenue Recognition schedules the Revenue Commitment or Revenue
	// Commitment Reversal.
	RevRecOnRevCommitment bool `json:"revRecOnRevCommitment,omitzero"`
	// The possible values for this field are:
	// 1. Pending: indicates that no recognition has occurred. All revenue is
	// still deferred.
	// 2. In Progress: indicates that some recognition has occurred.
	// 3. Completed: indicates that all recognition has occurred. No deferred
	// revenue remains.
	RevenueStatus      string `json:"revenueStatus,omitempty"`
	RevenueStatusDescr string `json:"revenueStatusDescr,omitempty"`
	// You can change the sales effective date for this transaction. The sales
	// effective date determines which commission plan and historical sales team this
	// transaction applies to. If this credit memo is created from a return
	// authorization, the sales effective date from the return shown in the
	// Created From field is set by default for this credit memo.
	SalesEffectiveDate  Date                `json:"salesEffectiveDate,omitzero"`
	SalesRep            NSResource          `json:"salesRep,omitzero"`
	SalesReturn         string              `json:"salesReturn,omitempty"`
	SalesTeam           SalesTeamCollection `json:"salesTeam,omitzero"`
	ShipAddress         string              `json:"shipAddress,omitempty"`
	ShipAddressList     NSResource          `json:"shipAddressList,omitzero"`
	ShipDate            Date                `json:"shipDate,omitzero"`
	ShipGroup           ShipGroupCollection `json:"shipGroup,omitzero"`
	ShipIsResidential   bool                `json:"shipIsResidential,omitzero"`
	ShipMethod          ShipItem            `json:"shipMethod,omitzero"`
	ShipOverride        bool                `json:"shipOverride,omitzero"`
	ShippingAddress     ShippingAddress     `json:"shippingAddress,omitzero"`
	ShippingAddressText string              `json:"shippingAddress_text,omitempty"`
	// Enter the amount the customer was charged for shipping.
	ShippingCost           float64    `json:"shippingCost,omitempty"`
	ShippingCostOverridden bool       `json:"shippingCostOverridden,omitzero"`
	ShippingTaxAmount      float64    `json:"shippingTaxAmount,omitempty"`
	Source                 NSResource `json:"source,omitzero"`
	StartDate              Date       `json:"startDate,omitzero"`
	Status                 NSResource `json:"status,omitzero"`
	StoreOrder             string     `json:"storeOrder,omitempty"`
	Subsidiary             Subsidiary `json:"subsidiary,omitzero"`
	// This field shows the tax registration number of the transaction nexus.
	// NetSuite automatically populates this field based on the nexus lookup
	// logic. You can override the transaction nexus and tax registration
	// number that NetSuite automatically selects by checking the Nexus
	// Override box. When you select a different tax registration number in
	// the dropdown list, the corresponding nexus is automatically selected
	// in the Nexus field.
	SubsidiaryTaxRegNum string               `json:"subsidiaryTaxRegNum,omitempty"`
	Subtotal            float64              `json:"subtotal,omitempty"`
	TaxDetails          TaxDetailsCollection `json:"taxDetails,omitzero"`
	// Check this box to override the tax information on the Tax Details
	// subtab of the transaction. Only roles with at least the Edit level of
	// the Tax Details Tab permission can override the tax details.
	TaxDetailsOverride   bool `json:"taxDetailsOverride,omitzero"`
	TaxPointDate         Date `json:"taxPointDate,omitzero"`
	TaxPointDateOverride bool `json:"taxPointDateOverride,omitzero"`
	// Check this box to override the values in the Nexus and Subsidiary
	// Tax Reg. Number fields. Only roles with at least the Edit level of the
	// Tax Details Tab permission can override the values that NetSuite
	// automatically selects in these fields.
	TaxRegOverride bool `json:"taxRegOverride,omitzero"`
	// NetSuite multiplies the tax rate by the taxable total of line items
	// and enters it here.
	TaxTotal float64 `json:"taxTotal,omitempty"`
	Terms    Term    `json:"terms,omitzero"`
	// Check this box if you want to e-mail this credit memo. Then enter the
	// e-mail address in the space to the right of the check box. If you
	// don't check this box, the credit memo won't be e-mailed. You can
	// enter multiple e-mail addresses separated by a semicolon with no spaces.
	ToBeEmailed bool `json:"toBeEmailed,omitzero"`
	// Check this box if you want to fax this credit memo. Enter the fax
	// number in the space to the right of the check box if it doesn't
	// already appear. To fax NetSuite forms, an administrator must first set
	// up fax service at Setup > Company > Printing, Fax and Email
	// Preferences.
	ToBeFaxed bool `json:"toBeFaxed,omitzero"`
	// Check this box if you wish to save this in a queue of credit memos
	// to print. Otherwise, you can click Print below to submit and print
	// this credit memo at once.
	ToBePrinted bool `json:"toBePrinted,omitzero"`
	// NetSuite computes the total of line items and tax and then enters it here.
	Total           float64 `json:"total,omitempty"`
	TotalAfterTaxes float64 `json:"totalAfterTaxes,omitempty"`
	// Estimated Cost: Estimated cost of the specific number of items;
	// estimated rate x quantity = estimated cost.
	TotalCostEstimate float64 `json:"totalCostEstimate,omitempty"`
	// NetSuite inserts today's date as the date of this order. You can
	// type or pick another date.
	TranDate Date `json:"tranDate,omitzero"`
	// NetSuite increases the largest credit memo number by one. If you
	// wish, you can type another number. The next credit memo number will
	// revert to the standard pattern. You can enter a maximum of 45 characters
	// in this field.
	TranId           string  `json:"tranId,omitempty"`
	TranIsVsoeBundle bool    `json:"tranIsVsoeBundle,omitzero"`
	Unapplied        float64 `json:"unapplied,omitempty"`
	// The VSOE allocation amount can be calculated automatically by
	// checking the Auto Calculate VSOE Allocation box. Allocation is the
	// process to determine the VSOE price for items. The total VSOE amount
	// allocated is the revenue amount to be recognized for the bundle.
	VsoeAutoCalc      bool       `json:"vsoeAutoCalc,omitzero"`
	WebSite           string     `json:"webSite,omitempty"`
	WhichChargesToAdd NSResource `json:"whichChargesToAdd,omitzero"`

	CustomFields map[string]any `json:"-"`
}

func (c CreditMemo) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(c)
}

func (c *CreditMemo) UnmarshalJSON(data []byte) error {
	err := UnmarshalCustomFields(data, &c.CustomFields)
	if err != nil {
		return err
	}

	type Alias CreditMemo
	return json.Unmarshal(data, (*Alias)(c))
}

type Customers []Customer

type Customer struct {
	AddressBook AddressBookCollection `json:"addressBook,omitzero"`
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
	ExternalID string `json:"externalId,omitempty"`
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
	Subsidiary Subsidiary `json:"subsidiary,omitzero"`
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
	Email                  string `json:"email,omitempty"`
	Phone                  string `json:"phone,omitempty"`
	DefaultBillingAddress  string `json:"defaultbillingaddress,omitempty"`
	DefaultShippingAddress string `json:"defaultshippingaddress,omitempty"`
	Parent                 string `json:"parent,omitempty"`
	CustomerNumber         string `json:"custentity_nch_customer_number,omitempty"`
}

type InvoiceItemCollection Collection[InvoiceItems]

type InvoiceItems []InvoiceItem

type InvoiceItem struct {
	Links   Links   `json:"links,omitzero"`
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
	Item        InvoiceItemItemItem `json:"item,omitzero"`
	ItemSubType NSResource          `json:"itemSubType,omitzero"`
	ItemType    NSResource          `json:"itemType,omitzero"`
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

	CustomFields map[string]any `json:"-"`
}

func (i *InvoiceItem) UnmarshalJSON(data []byte) error {
	err := UnmarshalCustomFields(data, &i.CustomFields)
	if err != nil {
		return err
	}

	type Alias InvoiceItem
	return json.Unmarshal(data, (*Alias)(i))
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
	ID      string `json:"id"`
	RefName string `json:"refName"`
}

func (i InvoiceItemItemItem) IsZero() bool {
	return zero.IsZero(i)
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

func (r RecordRef) IsZero() bool {
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

func (c Classification) IsZero() bool {
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

func (d Department) IsZero() bool {
	return zero.IsZero(d)
}

type Locations []Location

type Location struct {
	Links                  []any  `json:"links,omitzero"`
	Fullname               string `json:"fullname,omitzero"`
	ID                     string `json:"id,omitzero"`
	Isinactive             string `json:"isinactive,omitzero"`
	Lastmodifieddate       string `json:"lastmodifieddate,omitzero"`
	Mainaddress            string `json:"mainaddress,omitzero"`
	Makeinventoryavailable string `json:"makeinventoryavailable,omitzero"`
	Name                   string `json:"name,omitzero"`
	Subsidiary             string `json:"subsidiary,omitzero"`
	Externalid             string `json:"externalid,omitempty"`
}

func (l Location) IsZero() bool {
	return zero.IsZero(l)
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

type SalesTaxItems []SalesTaxItem

type SalesTaxItem struct {
	CustomForm  CustomForm `json:"customForm,omitzero"`
	Description string     `json:"description"`
	ExternalID  string     `json:"externalId"`
	ID          string     `json:"id"`
	// Once the tax code record has been saved, you cannot change the value in the Tax Type field.
	ImportantNote string `json:"importantNote"`
	// Check this box to inactivate this record. Inactive records do not show on transactions and records for selection in lists.
	IsInactive bool   `json:"isInactive"`
	ItemType   string `json:"itemType"`
	Links      Links  `json:"links"`
	// Enter a name for this record.
	Name    string  `json:"name"`
	RefName string  `json:"refName"`
	TaxType TaxType `json:"taxType"`
}

func (s SalesTaxItem) IsZero() bool {
	return zero.IsZero(s)
}

// https://system.netsuite.com/help/helpcenter/en_US/APIs/REST_API_Browser/record/v1/2025.1/index.html#/definitions/taxType
type TaxType struct {
	Country           NSResource              `json:"country,omitzero"`
	CustomForm        CustomForm              `json:"customForm,omitzero"`
	Description       string                  `json:"description,omitempty"`
	DoesNotAddToTotal bool                    `json:"doesNotAddToTotal,omitempty"`
	ExternalID        string                  `json:"externalId,omitempty"`
	ID                string                  `json:"id,omitempty"`
	IsInactive        bool                    `json:"isInactive,omitempty"`
	Links             Links                   `json:"links,omitzero"`
	Name              string                  `json:"name,omitempty"`
	NexusAccounts     NexusAccountsCollection `json:"nexusAccounts,omitzero"`
	PostToItemCost    bool                    `json:"postToItemCost,omitempty"`
	RefName           string                  `json:"refName,omitempty"`
	ReverseCharge     bool                    `json:"reverseCharge,omitempty"`
	TaxInNetAmount    bool                    `json:"taxInNetAmount,omitempty"`
}

func (t TaxType) IsZero() bool {
	return zero.IsZero(t)
}

type NSResource struct {
	Links   Links  `json:"links,omitempty"`
	ID      string `json:"id"`
	RefName string `json:"refName,omitempty"`
}

func (n NSResource) IsZero() bool {
	return zero.IsZero(n)
}

type TaxAccount NSResource

type StatisticalJournalEntryCollection Collection[StatisticalJournalEntries]

type NSLink struct {
	Rel  string `json:"rel"`
	Href string `json:"href"`
}

type StatisticalJournalEntries []StatisticalJournalEntry

type StatisticalJournalEntry struct {
	// Check this box to update the balance of multiple statistical accounts at
	// the same time. The account balances you enter overwrite all previous
	// values. Absolute values are available to you through reporting and for
	// use in dynamic allocation schedules.
	AbsoluteUpdate            bool                     `json:"absoluteUpdate,omitempty"`
	AbsoluteUpdateJournalLink *StatisticalJournalEntry `json:"absoluteUpdateJournalLink,omitzero"`
	ApprovalStatus            RecordRef                `json:"approvalStatus,omitzero"`
	// Check this box to indicate that this journal entry is approved. If you
	// check this box, this journal entry will post immediately. If you do not
	// check this box, this journal entry must be approved before it posts.
	Approved               bool                                  `json:"approved,omitempty"`
	Class                  Classification                        `json:"Class,omitzero"`
	CreatedDate            Date                                  `json:"createdDate,omitzero"`
	CustomForm             CustomForm                            `json:"customForm,omitzero"`
	Department             Department                            `json:"Department,omitzero"`
	ExcludeFromGLNumbering bool                                  `json:"excludeFromGLNumbering,omitempty"`
	ExternalID             string                                `json:"externalId,omitempty"`
	ID                     string                                `json:"id,omitempty"`
	IsReversal             bool                                  `json:"isReversal,omitempty"`
	LastModifiedDate       Date                                  `json:"lastModifiedDate,omitzero"`
	Line                   StatisticalJournalEntryLineCollection `json:"line"`
	Links                  Links                                 `json:"links,omitzero"`
	Location               Location                              `json:"location,omitzero"`
	// If you wish, enter a memo to describe this journal entry. It will appear
	// on this transaction detail as well as reports such as a 2-line account
	// register.
	Memo              string                   `json:"memo"`
	NextApprover      Employee                 `json:"nextApprover,omitzero"`
	OffsetJournalLink *StatisticalJournalEntry `json:"offsetJournalLink,omitzero"`
	ParentStat        NSResource               `json:"parentStat,omitzero"`
	PostingPeriod     AccountingPeriod         `json:"postingPeriod,omitzero"`
	RefName           string                   `json:"refName,omitempty"`
	Reversal          *StatisticalJournalEntry `json:"reversal,omitzero"`
	// If this journal entry is intended to be reversed, enter the date for the
	// reversing entry to be posted. If Defer Entry is checked, this field is
	// mandatory.
	ReversalDate Date `json:"reversalDate,omitzero"`
	// Check this box to make the reversal a memorized transaction that
	// automatically occurs on the date entered in the required Reversal Date
	// field. Clear this box to make the reversal an immediately entered
	// transaction with the date in the Reversal Date field.
	ReversalDefer bool       `json:"reversalDefer,omitempty"`
	Subsidiary    Subsidiary `json:"subsidiary,omitzero"`
	// In this field, NetSuite keeps track of the difference between the debits
	// and credits you enter below. If this field does not equal 0.00, NetSuite
	// will not allow you to submit this transaction.
	Total float64 `json:"total,omitempty"`
	// NetSuite inserts today's date as the date of this journal entry. You
	// can enter or select another date.
	TranDate Date `json:"tranDate,omitzero"`
	// NetSuite increases the largest journal entry number by one. If you wish,
	// you can type another number. The next journal entry number will revert
	// to the standard pattern. You can enter a maximum of 45 characters in
	// this field.
	TranID    string    `json:"tranId,omitempty"`
	Unit      string    `json:"unit,omitempty"`
	UnitsType UnitsType `json:"unitsType,omitzero"`
	Void      bool      `json:"void,omitempty"`
}

// https://system.netsuite.com/help/helpcenter/en_US/APIs/REST_API_Browser/record/v1/2025.1/index.html#/definitions/statisticalJournalEntry-lineCollection

type StatisticalJournalEntryLineCollection Collection[StatisticalJournalEntryLineElements]

type StatisticalJournalEntryLineElements []StatisticalJournalEntryLineElement

// https://system.netsuite.com/help/helpcenter/en_US/APIs/REST_API_Browser/record/v1/2025.1/index.html#/definitions/statisticalJournalEntry-lineElement
type StatisticalJournalEntryLineElement struct {
	Account                Account    `json:"Account,omitzero"`
	Class                  RecordRef  `json:"Class,omitzero"`
	Debit                  float64    `json:"debitAmount,omitempty"`
	Department             Department `json:"Department,omitzero"`
	Eliminate              bool       `json:"eliminate,omitempty"`
	Entity                 NSResource `json:"entity,omitzero"`
	Line                   int        `json:"line,omitempty"`
	LineCreatedDate        Date       `json:"lineCreatedDate,omitzero"`
	LineLastModifiedDate   Date       `json:"lineLastModifiedDate,omitzero"`
	LineUnit               string     `json:"lineUnit,omitempty"`
	Links                  Links      `json:"links,omitzero"`
	Location               Location   `json:"location,omitzero"`
	Memo                   string     `json:"memo"`
	PreviewDebit           float64    `json:"previewDebitAmount,omitempty"`
	RefName                string     `json:"refName,omitempty"`
	RevenueRecognitionRule NSResource `json:"revenueRecognitionRule,omitzero"`
}

// https://system.netsuite.com/help/helpcenter/en_US/APIs/REST_API_Browser/record/v1/2025.1/index.html#/definitions/employee
type Employee struct {
}

// https://system.netsuite.com/help/helpcenter/en_US/APIs/REST_API_Browser/record/v1/2025.1/index.html#/definitions/unitsType
//
//	externalId External ID: string
//
// id Internal ID: string
// isInactive Inactive: boolean
//
//	Check this box to inactivate this unit record. When a record is inactive, it does not appear in lists and cannot be selected on records and transactions.
//
// lastModifiedDate Last Modified Date: string (date-time)
// links Links: nsLink
// nsLink
// name Type Name: string
//
//	Enter a name for the type of units you are creating. For example, enter "Length" to create units of Inch, Foot and Yard.
//
// refName Reference Name: string
// uom: unitsType-uomCollection
type UnitsType struct {
	ExternalID string `json:"externalId,omitempty"`
	ID         string `json:"id,omitempty"`
	// Check this box to inactivate this unit record. When a record is inactive,
	// it does not appear in lists and cannot be selected on records and
	// transactions.
	IsInactive       bool  `json:"isInactive,omitempty"`
	LastModifiedDate Date  `json:"lastModifiedDate,omitzero"`
	Links            Links `json:"links,omitempty"`
	// Enter a name for the type of units you are creating. For example, enter
	// "Length" to create units of Inch, Foot and Yard.
	Name    string                 `json:"name,omitempty"`
	RefName string                 `json:"refName,omitempty"`
	UOM     UnitsTypeUOMCollection `json:"uom,omitempty"`
}

func (u UnitsType) IsZero() bool {
	return zero.IsZero(u)
}

// https://system.netsuite.com/help/helpcenter/en_US/APIs/REST_API_Browser/record/v1/2025.1/index.html#/definitions/unitsType-uomCollection
// count Count: integer (int64)
// hasMore Has More Results: boolean
// items: unitsType-uomElement
// unitsType-uomElement
// links Links: nsLink
// nsLink
// offset Query Offset: integer (int64)
// totalResults Total Results: integer (int64)

type UnitsTypeUOMCollection Collection[UnitsTypeUOMElements]

type UnitsTypeUOMElements []UnitsTypeUOMElement

// https://system.netsuite.com/help/helpcenter/en_US/APIs/REST_API_Browser/record/v1/2025.1/index.html#/definitions/unitsType-uomElement
//
//	abbreviation Abbreviation: string
//
// baseUnit Base Unit: boolean
// conversionRate Conversion Rate {1}: number (float)
// inUse Used: boolean
// internalId Internal ID: integer (int64)
// links Links: nsLink
// nsLink
// pluralAbbreviation Plural Abbreviation: string
// pluralName Plural Name: string
// refName Reference Name: string
// unitName Name: string
type UnitsTypeUOMElement struct {
	Abbreviation       string  `json:"abbreviation,omitempty"`
	BaseUnit           bool    `json:"baseUnit,omitempty"`
	ConversionRate     float64 `json:"conversionRate,omitempty"`
	InUse              bool    `json:"inUse,omitempty"`
	InternalID         int     `json:"internalId,omitempty"`
	Links              Links   `json:"links,omitempty"`
	PluralAbbreviation string  `json:"pluralAbbreviation,omitempty"`
	PluralName         string  `json:"pluralName,omitempty"`
	RefName            string  `json:"refName,omitempty"`
	UnitName           string  `json:"unitName,omitempty"`
}

// count Count: integer (int64)
// hasMore Has More Results: boolean
// items: account-localizationsElement
// account-localizationsElement
// links Links: nsLink
// nsLink
// offset Query Offset: integer (int64)
// totalResults Total Results: integer (int64)

// https://system.netsuite.com/help/helpcenter/en_US/APIs/REST_API_Browser/record/v1/2025.1/index.html#/definitions/account-localizationsCollection
type AccountLocalizationsCollection Collection[AccountLocalizationsElements]

func (c AccountLocalizationsCollection) IsZero() bool {
	return zero.IsZero(c)
}

type AccountLocalizationsElements []AccountLocalizationsElement

// https://system.netsuite.com/help/helpcenter/en_US/APIs/REST_API_Browser/record/v1/2025.1/index.html#/definitions/account-localizationsElement
type AccountLocalizationsElement struct {
	AccountingContext NSResource `json:"accountingcontext,omitzero"`
	AcctName          string     `json:"acctname,omitempty"`
	AcctNumber        string     `json:"acctnumber,omitempty"`
	LegalName         string     `json:"legalName,omitempty"`
	Links             Links      `json:"links,omitempty"`
	Locale            NSResource `json:"locale,omitzero"`
	RefName           string     `json:"refName,omitempty"`
}

// https://system.netsuite.com/help/helpcenter/en_US/APIs/REST_API_Browser/record/v1/2025.1/index.html#/definitions/nsResourceCollection
type NSResourceCollection Collection[[]NSResource]

func (c NSResourceCollection) IsZero() bool {
	return zero.IsZero(c)
}

// https://system.netsuite.com/help/helpcenter/en_US/APIs/REST_API_Browser/record/v1/2025.1/index.html#/definitions/account-accountContextSearchCollection
type AccountContextSearchCollection Collection[AccountContextSearchElements]

func (c AccountContextSearchCollection) IsZero() bool {
	return zero.IsZero(c)
}

type AccountContextSearchElements []AccountContextSearchElement

// https://system.netsuite.com/help/helpcenter/en_US/APIs/REST_API_Browser/record/v1/2025.1/index.html#/definitions/account-accountContextSearchElement
type AccountContextSearchElement struct {
	AccountingContext NSResource `json:"accountingcontext,omitzero"`
	AcctName          string     `json:"acctname,omitempty"`
	AcctNumber        string     `json:"acctnumber,omitempty"`
	LegalName         string     `json:"legalName,omitempty"`
	Links             Links      `json:"links,omitempty"`
	Locale            NSResource `json:"locale,omitzero"`
	RefName           string     `json:"refName,omitempty"`
}

type Collection[T any] struct {
	Count        int      `json:"count"`
	HasMore      bool     `json:"hasMore"`
	Items        T        `json:"items"`
	Links        []NSLink `json:"links"`
	Offset       int      `json:"offset"`
	TotalResults int      `json:"totalResults"`
}

func (c Collection[T]) IsZero() bool {
	return zero.IsZero(c)
}

type SubsidiaryCollection Collection[Subsidiaries]

func (c SubsidiaryCollection) IsZero() bool {
	return zero.IsZero(c)
}

type NexusAccountsCollection Collection[NexusAccountsElements]

type NexusAccountsElements []NexusAccountElement

// https://system.netsuite.com/help/helpcenter/en_US/APIs/REST_API_Browser/record/v1/2025.1/index.html#/definitions/taxType-nexusAccountsElement
type NexusAccountElement struct {
	Links              Links   `json:"links,omitempty"`
	Nexus              Nexus   `json:"nexus,omitzero"`
	PayablesAccount    Account `json:"payablesAccount,omitzero"`
	ReceivablesAccount Account `json:"receivablesAccount,omitzero"`
	RefName            string  `json:"refName,omitempty"`
}

type Nexuses []Nexus

// https://system.netsuite.com/help/helpcenter/en_US/APIs/REST_API_Browser/record/v1/2025.1/index.html#/definitions/nexus
type Nexus struct {
	Country    NSResource `json:"country,omitzero"`
	CustomForm CustomForm `json:"customForm,omitzero"`
	// Enter a description for this record.
	Description string     `json:"description,omitempty"`
	ExternalID  string     `json:"externalId,omitempty"`
	ID          string     `json:"id,omitempty"`
	IsInactive  bool       `json:"isInactive,omitempty"`
	Links       Links      `json:"links,omitempty"`
	ParentNexus *Nexus     `json:"parentNexus,omitzero"`
	RefName     string     `json:"refName,omitempty"`
	State       NSResource `json:"state,omitzero"`
	TaxAgency   Vendor     `json:"taxAgency,omitzero"`
	// Check this box to use the fulfillment date as the tax point date for this
	// transaction. This overrides the default Nexus setting which uses the
	// transaction date to determine the tax point date. This setting is
	// typically used for jurisdictions which require the item fulfillment date
	// to be used to define the tax point date.
	TaxDateFromFulfillment bool `json:"taxDateFromFulfillment,omitempty"`
}

// https://system.netsuite.com/help/helpcenter/en_US/APIs/REST_API_Browser/record/v1/2025.1/index.html#/definitions/vendor
type Vendor struct {
	ID      string `json:"id,omitempty"`
	RefName string `json:"refName,omitempty"`
	Links   Links  `json:"links,omitempty"`
}

type NexusCollection Collection[[]Nexus]

// https://system.netsuite.com/help/helpcenter/en_US/APIs/REST_API_Browser/record/v1/2025.1/index.html#/definitions/creditMemo-accountingBookDetailCollection
type AccountingBookDetailCollection Collection[AccountingBookDetailElements]

type AccountingBookDetailElements []AccountingBookDetailElement

// https://system.netsuite.com/help/helpcenter/en_US/APIs/REST_API_Browser/record/v1/2025.1/index.html#/definitions/creditMemo-accountingBookDetailElement

type AccountingBookDetailElement struct {
	AccountingBook        NSResource `json:"accountingBook,omitzero"`
	ExchangeRate          float64    `json:"exchangeRate,omitempty"`
	Links                 Links      `json:"links,omitempty"`
	RefName               string     `json:"refName,omitempty"`
	RevRecOnRevCommitment bool       `json:"revRecOnRevCommitment,omitempty"`
	Subsidiary            Subsidiary `json:"subsidiary,omitzero"`
	TranIsVSOEBundle      bool       `json:"tranIsVsoeBundle,omitempty"`
}

type AppliedRulesCollection Collection[AppliedRulesElements]

type AppliedRulesElements []AppliedRulesElement

// https://system.netsuite.com/help/helpcenter/en_US/APIs/REST_API_Browser/record/v1/2025.1/index.html#/definitions/creditMemo-appliedRulesElement
type AppliedRulesElement struct {
	CreationDate        Date   `json:"creationDate,omitzero"`
	Details             string `json:"details,omitempty"`
	ExternalLogID       int    `json:"externalLogId,omitempty"`
	ID                  string `json:"id,omitempty"`
	Links               Links  `json:"links,omitempty"`
	RefName             string `json:"refName,omitempty"`
	RuleType            string `json:"ruleType,omitempty"`
	RuleTypeTranslation string `json:"ruleTypeTranslation,omitempty"`
	TransactionVersion  int    `json:"transactionVersion,omitempty"`
}

type ApplyCollection Collection[ApplyElements]

type ApplyElements []ApplyElement

// https://system.netsuite.com/help/helpcenter/en_US/APIs/REST_API_Browser/record/v1/2025.1/index.html#/definitions/creditMemo-applyElement
type ApplyElement struct {
	Amount      float64    `json:"amount,omitempty"`
	Apply       bool       `json:"applied,omitempty"`
	ApplyDate   Date       `json:"applyDate,omitzero"`
	CreatedFrom string     `json:"createdFrom,omitempty"`
	Currency    string     `json:"currency,omitempty"`
	Docs        NSResource `json:"docs,omitzero"`
	Due         float64    `json:"due,omitempty"`
	JobName     string     `json:"jobName,omitempty"`
	Line        int        `json:"line,omitempty"`
	Links       Links      `json:"links,omitempty"`
	RefName     string     `json:"refName,omitempty"`
	RefNumber   string     `json:"refNumber,omitempty"`
	Total       float64    `json:"total,omitempty"`
	Type        string     `json:"type,omitempty"`
}

// https://system.netsuite.com/help/helpcenter/en_US/APIs/REST_API_Browser/record/v1/2025.1/index.html#/definitions/billingAccount
type BillingAccount struct {
	// Clear this box to manually enter a name for this record. If you leave
	// this box marked, NetSuite assigns a name or number for this record based
	// on your settings at Setup > Set Up Auto-Generated Numbers.
	AutoName    bool   `json:"autoName,omitempty"`
	BillAddress string `json:"billAddress,omitempty"`
	// Select the billing address you want to use for this billing account. Options
	// are sourced from the customer record.
	BillAddressList string          `json:"billAddressList,omitempty"`
	BillingSchedule BillingSchedule `json:"billingSchedule,omitzero"`
	CashSaleForm    NSResource      `json:"cashSaleForm,omitzero"`
	Class           Classification  `json:"Class,omitzero"`
	CreatedBy       string          `json:"createdBy,omitzero"`
	CreatedDate     Date            `json:"createdDate,omitzero"`
	Currency        Currency        `json:"currency,omitzero"`
	CustomForm      CustomForm      `json:"customForm,omitzero"`
	Customer        Customer        `json:"customer,omitzero"`
	// If checked, this field indicates that this is the default billing
	CustomerDefault bool       `json:"customerDefault,omitempty"`
	Department      Department `json:"Department,omitzero"`
	DisplayName     string     `json:"displayName,omitempty"`
	ExternalID      string     `json:"externalId,omitempty"`
	Frequency       NSResource `json:"frequency,omitzero"`
	// The internal ID for this record is shown here. If you do not want to
	// show internal IDs, clear the Show Internal IDs box at Home > Set
	// Preferences.
	ID string `json:"id,omitempty"`
	// The number of the billing account. This number is automatically
	// generated upon save.
	IDNumber string `json:"idNumber,omitempty"`
	// Check this box if the billing account is no longer active.
	Inactive          bool       `json:"inactive,omitempty"`
	InvoiceForm       NSResource `json:"invoiceForm,omitzero"`
	LastBillCycleDate Date       `json:"lastBillCycleDate,omitzero"`
	LastBillDate      Date       `json:"lastBillDate,omitzero"`
	LastModifiedDate  Date       `json:"lastModifiedDate,omitzero"`
	Links             Links      `json:"links,omitempty"`
	Location          Location   `json:"location,omitzero"`
	// Enter a description for this billing account.
	Memo string `json:"memo,omitempty"`
	// The name of this billing account.
	Name              string     `json:"name,omitempty"`
	NextBillCycleDate NSResource `json:"nextBillCycleDate,omitzero"`
	RefName           string     `json:"refName,omitempty"`
	// Check this box if you want advance charges for this billing account
	// invoiced before the first scheduled bill date.
	RequestOffCycleInvoice bool `json:"requestOffCycleInvoice,omitempty"`
	// The shipping address for this billing account.
	ShipAddress string `json:"shipAddress,omitempty"`
	// Select the shipping address you want to use for this billing account. Options
	// are sourced from the customer record.
	ShipAddressList string `json:"shipAddressList,omitempty"`
	// Enter the date this person or company became a customer, lead or
	// prospect. If this person or company has a contract with you, enter the
	// start date of the contract. If you enter an estimate or an opportunity
	// for this customer, this field will be updated with the date of that
	// transaction.
	StartDate  Date       `json:"startDate,omitzero"`
	Subsidiary Subsidiary `json:"subsidiary,omitzero"`
}

type BillingSchedule struct {
}

// https://system.netsuite.com/help/helpcenter/en_US/APIs/REST_API_Browser/record/v1/2025.1/index.html#/definitions/creditMemo-billingAddress
type BillingAddress Address

type Addresses []Address

type Address struct {
	// Enter the address the way it should appear on forms. For employees, customers, partners, and vendors, what you enter here autofills on forms if this address is marked default for Shipping or Billing. Enter up to 50 characters. This field is required for the Online Bill Pay feature.
	Addr1 string `json:"addr1,omitempty"`
	// Enter an optional second address line the way it should appear on forms. For employees, customers, partners, and vendors, what you enter here autofills on forms if this address is marked default for Shipping or Billing. Enter up to 50 characters.
	Addr2 string `json:"addr2,omitempty"`
	Addr3 string `json:"addr3,omitempty"`
	// Enter the phone number.
	AddrPhone string `json:"addrPhone,omitempty"`
	// The values entered in the other address fields are displayed here.
	AddrText string `json:"addrText,omitempty"`
	// Enter the name of the entity that should appear on the shipping label here. This name appears on the shipping label below what you enter in the Attention field.
	Addressee string `json:"addressee,omitempty"`
	// Enter the name of the person to whom a shipment is addressed, as it should appear on shipping labels. This field is required for UPS Integration.
	Attention string `json:"attention,omitempty"`
	// Enter the city the way it should appear on all forms except checks.
	City             string     `json:"city,omitempty"`
	Country          NSResource `json:"country,omitzero"`
	ExternalID       string     `json:"externalId,omitempty"`
	LastModifiedDate Date       `json:"lastModifiedDate,omitzero"`
	Links            Links      `json:"links,omitempty"`
	// Check this box to disable the free-form address text field. When this field is disabled, text entered in the other address fields does not display in the Address text field. Clear this box to allow text entered in the address component fields to appear in the free-form address text field.
	Override bool   `json:"override,omitempty"`
	RefName  string `json:"refName,omitempty"`
	// Enter your company's state or province the way it should appear on all forms except checks.
	State string `json:"state,omitempty"`
	// Enter the postal code the way it should appear on all forms except checks.
	Zip string `json:"zip,omitempty"`
}

// code Code: string
//
//     The coupon code that is entered on a sales transaction to apply the promotion. Coupon codes can be multiple-use or single-use. Multiple-use codes can be used any number of times by any number of customers. Each single-use code can only be used in one transaction by a single customer.
// dateSent Date Sent: string (date)
//
//     This field shows the date the coupon code was sent to the recipient.
// externalId External ID: string
// id Internal ID: string
// links Links: nsLink
// nsLink
// promotion: promotionCode
// recipient: one of: [ customer, partner, vendor, nsResource, employee, contact ]
// refName Reference Name: string
// useCount Used Count: string
// used Used: boolean
//
//     The coupon has been used by the recipient when the Used box is checked.

// https://system.netsuite.com/help/helpcenter/en_US/APIs/REST_API_Browser/record/v1/2025.1/index.html#/definitions/couponCode
type CouponCode struct {
	// The coupon code that is entered on a sales transaction to apply the promotion. Coupon codes can be multiple-use or single-use. Multiple-use codes can be used any number of times by any number of customers. Each single-use code can only be used in one transaction by a single customer.
	Code string `json:"code,omitempty"`
	// This field shows the date the coupon code was sent to the recipient.
	DateSent   Date          `json:"dateSent,omitzero"`
	ExternalID string        `json:"externalId,omitempty"`
	ID         string        `json:"id,omitempty"`
	Links      Links         `json:"links,omitempty"`
	Promotion  PromotionCode `json:"promotion,omitzero"`
	Recipient  NSResource    `json:"recipient,omitzero"`
	RefName    string        `json:"refName,omitempty"`
	UseCount   string        `json:"useCount,omitempty"`
	// The coupon has been used by the recipient when the Used box is checked.
	Used bool `json:"used,omitempty"`
}

// https://system.netsuite.com/help/helpcenter/en_US/APIs/REST_API_Browser/record/v1/2025.1/index.html#/definitions/promotionCode
type PromotionCode struct {
}

// https://system.netsuite.com/help/helpcenter/en_US/APIs/REST_API_Browser/record/v1/2025.1/index.html#/definitions/discountItem
type DiscountItem struct {
}

// https://system.netsuite.com/help/helpcenter/en_US/APIs/REST_API_Browser/record/v1/2025.1/index.html#/definitions/creditMemo-itemCollection
type CreditMemoItemCollection Collection[CreditMemoItemElements]

type CreditMemoItemElements []CreditMemoItemElement

//  account: account
// amount Amount: number (double)
// catchUpPeriod: accountingPeriod
// chargeType: nsResource
// class: classification
// costEstimate Est. Cost: number (double)
// costEstimateRate Est. Rate: number (double)
// costEstimateType: object
//
// id Internal identifier: string PREFVENDORRATE , AVGCOST , PURCHORDERRATE , LASTPURCHPRICE , MEMBERDEFINED , CUSTOM , ITEMDEFINED , PURCHPRICE
// refName Reference Name: string
//
// deferRevRec Deferred Revenue: boolean
// department: department
// description Description: string
// doNotCreateRevenueElement Do Not Create Revenue Element: boolean
// estGrossProfit Est. Gross Profit: number (double)
// estGrossProfitPercent Est. Gross Profit Percent: number (double)
// giftCertFrom From: string
// giftCertMessage Gift Message: string
// giftCertNumber Code: string
// giftCertRecipientEmail Recipient Email: string
// giftCertRecipientName Recipient Name: string
// initOqpBucket Initial OQP Bucket: string
// inventoryDetail: creditMemo-item-inventoryDetail
// isClosed Closed: boolean
// isDropShipment Drop Shipment: boolean
// isOpen Is Opened: boolean
// item: one of: [ inventoryItem, assemblyItem, kitItem, nsResource, discountItem, markupItem, subtotalItem, descriptionItem, paymentItem, salesTaxItem, shipItem, downloadItem, giftCertificateItem, subscriptionPlan, nonInventorySaleItem, nonInventoryResaleItem, nonInventoryPurchaseItem, serviceSaleItem, serviceResaleItem, servicePurchaseItem, otherChargeSaleItem, otherChargeResaleItem, otherChargePurchaseItem ]
// itemSubtype: object
//
// id Internal identifier: string Sale , Purchase , Resale
// refName Reference Name: string
//
// itemType: object
//
// id Internal identifier: string Group , Description , Discount , EndGroup , GiftCert , Subtotal , Service , ShipItem , TaxItem , InvtPart ,
// refName Reference Name: string
//
// job: job
// licenseCode License Code: string
// line Transaction Line: integer (int64)
// linked Linked: boolean
// links Links: nsLink
// nsLink
// location: location
// marginal Marginal: boolean
// matrixType Matrix Type: string
// minQty Minimum Quantity: number (float)
// options Options: string
// oqpBucket Oqp Bucket: string
// orderDoc: nsResource
// orderLine Order Line: integer (int64)
// price: priceLevel
// printItems Print Items: boolean
// processedByRevCommit Processed by Rev Commit: boolean
// quantity Quantity: number (float)
// quantityAllocated Quantity Allocated: number (float)
// quantityAvailable Available: number (float)
// quantityCommitted Committed: number (float)
// quantityDemandAllocated Allocated Demand: number (float)
// quantityOnHand On Hand: number (float)
// rate Rate: number (double)
// rateIncludingTax Rate: number (double)
// rateSchedule Rate Schedule: string
// refName Reference Name: string
// revRecEndDate Rev. Rec. End Date: string (date)
// revRecSchedule: revRecSchedule
// revRecStartDate Rev. Rec. Start Date: string (date)
// revrec_recurrencetype Rev Rec Recurrence Type: string
// shipAddress: nsResource
// shipCarrier: object
//
// id Internal identifier: string ups , nonups
// refName Reference Name: string
//
// shipMethod: shipItem
// subscription: subscription
// subscriptionLine: subscriptionLine
// subsidiary: subsidiary
// taxAmount Tax Amount: number (double)
// taxDetailsReference Tax Details Reference: string
// units Units: string
// vsoeAllocation Allocation Amount: number (double)
// vsoeAmount Calculated Amount: number (double)
// vsoeDeferral: object
//
// id Internal identifier: string DEFERALLUNTIL , DEFERUNTIL
// refName Reference Name: string
//
// vsoeDelivered Delivered: boolean
// vsoeIsEstimate Estimate: boolean
// vsoePermitDiscount: object
//
// id Internal identifier: string IFDELIVERED , NEVER
// refName Reference Name: string
//
// vsoePrice Allocation Price: number (double)
// vsoeSOPGroup: object
//
// id Internal identifier: string EXCLUDE , NORMAL , SOFTWARE
// refName Reference Name: string

// https://system.netsuite.com/help/helpcenter/en_US/APIs/REST_API_Browser/record/v1/2025.1/index.html#/definitions/creditMemo-itemElement
type CreditMemoItemElement struct {
	Account                   Account          `json:"Account,omitzero"`
	Amount                    float64          `json:"amount,omitempty"`
	CatchUpPeriod             AccountingPeriod `json:"catchUpPeriod,omitzero"`
	ChargeType                NSResource       `json:"chargeType,omitzero"`
	Class                     Classification   `json:"Class,omitzero"`
	CostEstimate              float64          `json:"costEstimate,omitempty"`
	CostEstimateRate          float64          `json:"costEstimateRate,omitempty"`
	CostEstimateType          NSResource       `json:"costEstimateType,omitzero"`
	DeferRevRec               bool             `json:"deferRevRec,omitempty"`
	Department                Department       `json:"Department,omitzero"`
	Description               string           `json:"description,omitempty"`
	DoNotCreateRevenueElement bool             `json:"doNotCreateRevenueElement,omitempty"`
	EstGrossProfit            float64          `json:"estGrossProfit,omitempty"`
	EstGrossProfitPercent     float64          `json:"estGrossProfitPercent,omitempty"`
	GiftCertFrom              string           `json:"giftCertFrom,omitempty"`
	GiftCertMessage           string           `json:"giftCertMessage,omitempty"`
	GiftCertNumber            string           `json:"giftCertNumber,omitempty"`
	GiftCertRecipientEmail    string           `json:"giftCertRecipientEmail,omitempty"`
	GiftCertRecipientName     string           `json:"giftCertRecipientName,omitempty"`
	InitOQPBucket             string           `json:"initOqpBucket,omitempty"`
	InventoryDetail           InventoryDetail  `json:"inventoryDetail,omitzero"`
	IsClosed                  bool             `json:"isClosed,omitempty"`
	IsDropShipment            bool             `json:"isDropShipment,omitempty"`
	IsOpen                    bool             `json:"isOpen,omitempty"`
	Item                      NSResource       `json:"item,omitzero"`
	ItemSubtype               NSResource       `json:"itemSubtype,omitzero"`
	ItemType                  NSResource       `json:"itemType,omitzero"`
	Job                       Job              `json:"job,omitzero"`
	LicenseCode               string           `json:"licenseCode,omitempty"`
	Line                      int              `json:"line,omitempty"`
	Linked                    bool             `json:"linked,omitempty"`
	Links                     Links            `json:"links,omitempty"`
	Location                  Location         `json:"location,omitzero"`
	Marginal                  bool             `json:"marginal,omitempty"`
	MatrixType                string           `json:"matrixType,omitempty"`
	MinQty                    float64          `json:"minQty,omitempty"`
	Options                   string           `json:"options,omitempty"`
	OQPBucket                 string           `json:"oqpBucket,omitempty"`
	OrderDoc                  NSResource       `json:"orderDoc,omitzero"`
	OrderLine                 int              `json:"orderLine,omitempty"`
	Price                     PriceLevel       `json:"price,omitzero"`
	PrintItems                bool             `json:"printItems,omitempty"`
	ProcessedByRevCommit      bool             `json:"processedByRevCommit,omitempty"`
	Quantity                  float64          `json:"quantity,omitempty"`
	QuantityAllocated         float64          `json:"quantityAllocated,omitempty"`
	QuantityAvailable         float64          `json:"quantityAvailable,omitempty"`
	QuantityCommitted         float64          `json:"quantityCommitted,omitempty"`
	QuantityDemandAllocated   float64          `json:"quantityDemandAllocated,omitempty"`
	QuantityOnHand            float64          `json:"quantityOnHand,omitempty"`
	Rate                      float64          `json:"rate,omitempty"`
	RateIncludingTax          float64          `json:"rateIncludingTax,omitempty"`
	RateSchedule              string           `json:"rateSchedule,omitempty"`
	RefName                   string           `json:"refName,omitempty"`
	RevRecEndDate             Date             `json:"revRecEndDate,omitzero"`
	RevRecSchedule            RevRecSchedule   `json:"revRecSchedule,omitzero"`
	RevRecStartDate           Date             `json:"revRecStartDate,omitzero"`
	RevrecRecurrenceType      string           `json:"revrec_recurrencetype,omitempty"`
	ShipAddress               NSResource       `json:"shipAddress,omitzero"`
	ShipCarrier               NSResource       `json:"shipCarrier,omitzero"`
	ShipMethod                ShipItem         `json:"shipMethod,omitzero"`
	Subscription              Subscription     `json:"subscription,omitzero"`
	SubscriptionLine          SubscriptionLine `json:"subscriptionLine,omitzero"`
	Subsidiary                Subsidiary       `json:"subsidiary,omitzero"`
	TaxAmount                 float64          `json:"taxAmount,omitempty"`
	TaxDetailsReference       string           `json:"taxDetailsReference,omitempty"`
	Units                     string           `json:"units,omitempty"`
	VSOEAllocation            float64          `json:"vsoeAllocation,omitempty"`
	VSOEAmount                float64          `json:"vsoeAmount,omitempty"`
	VSOEDeferral              NSResource       `json:"vsoeDeferral,omitzero"`
	VSOEDelivered             bool             `json:"vsoeDelivered,omitempty"`
	VSOEIsEstimate            bool             `json:"vsoeIsEstimate,omitempty"`
	VSOEPermitDiscount        NSResource       `json:"vsoePermitDiscount,omitzero"`
	VSOEPrice                 float64          `json:"vsoePrice,omitempty"`
	VSOESOPGroup              NSResource       `json:"vsoeSOPGroup,omitzero"`

	CustomFields map[string]any `json:"-"`
}

func (c *CreditMemoItemElement) UnmarshalJSON(data []byte) error {
	err := UnmarshalCustomFields(data, &c.CustomFields)
	if err != nil {
		return err
	}

	type Alias CreditMemoItemElement
	return json.Unmarshal(data, (*Alias)(c))
}

type PartnersCollection Collection[PartnersElements]

type PartnersElements []PartnerElement

// https://system.netsuite.com/help/helpcenter/en_US/APIs/REST_API_Browser/record/v1/2025.1/index.html#/definitions/creditMemo-partnersElement
type PartnerElement struct {
	Contribution float64    `json:"contribution,omitempty"`
	IsPrimary    bool       `json:"isPrimary,omitempty"`
	Links        Links      `json:"links,omitempty"`
	Partner      Partner    `json:"partner,omitzero"`
	PartnerRole  NSResource `json:"partnerRole,omitzero"`
	RefName      string     `json:"refName,omitempty"`
}

type Partner struct {
}

type SalesTeamCollection Collection[SalesTeamElements]

type SalesTeamElements []SalesTeamElement

type SalesTeamElement struct {
}

type ShipGroupCollection Collection[ShipGroupElements]

type ShipGroupElements []ShipGroupElement

type ShipGroupElement struct {
}

type ShipItem struct {
}

type ReturnAddress Address

type ShippingAddress Address

type TaxDetailsCollection Collection[TaxDetailsElements]

type TaxDetailsElements []TaxDetailElement

type TaxDetailElement struct {
}

type Term struct {
}

type InventoryDetail struct {
}

type Job struct {
}

type PriceLevel struct {
}

type RevRecSchedule struct {
}

type Subscription struct {
}

type SubscriptionLine struct {
}

type AddressBookCollection Collection[AddressBookElements]

type AddressBookElements []AddressBookElement

// https://system.netsuite.com/help/helpcenter/en_US/APIs/REST_API_Browser/record/v1/2025.1/index.html#/definitions/customer-addressBookElement
type AddressBookElement struct {
	AddressBookAddress     Address `json:"addressBookAddress,omitzero"`
	AddressBookAddressText string  `json:"addressBookAddress_text,omitempty"`
	AddressID              string  `json:"addressId,omitempty"`
	DefaultBilling         bool    `json:"defaultBilling,omitempty"`
	DefaultShipping        bool    `json:"defaultShipping,omitempty"`
	ID                     int     `json:"id,omitempty"`
	InternalID             int     `json:"internalId,omitempty"`
	IsResidential          bool    `json:"isResidential,omitempty"`
	Label                  string  `json:"label,omitempty"`
	LastModifiedDate       Date    `json:"lastModifiedDate,omitzero"`
	Links                  Links   `json:"links,omitempty"`
	RefName                string  `json:"refName,omitempty"`
}
