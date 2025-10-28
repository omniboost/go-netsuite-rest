package netsuite

import (
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

	// CustomFields map[string]NSResource `json:"customFields,omitempty"`
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

type Currency struct {
	Links   Links  `json:"links"`
	ID      string `json:"id"`
	RefName string `json:"refName"`
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
	Links   Links  `json:"links,omitempty"`
	ID      string `json:"id,omitempty"`
	RefName string `json:"refName,omitempty"`
}

func (s Subsidiary) IsZero() bool {
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
	DueDate    Date       `json:"dueDate,omitzero"`
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
	Department RecordRef `json:"Department,omitzero"`
	Class      RecordRef `json:"Class,omitzero"`
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
	ItemSubType SubType             `json:"itemSubType"`
	ItemType    Type                `json:"itemType"`
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

type SubType struct {
	ID      string `json:"id"`
	RefName string `json:"refName,omitempty"`
}

type Type struct {
	ID      string `json:"id"`
	RefName string `json:"refName,omitempty"`
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
	CustomForm  CustomForm `json:"customForm"`
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

// country: object
//
// id Internal identifier: string BR , AU , SG , JP , GB , IT , FR , CA , US
// refName Reference Name: string
//
// customForm: object
//
// id Internal identifier: string -22519
// refName Reference Name: string
//
// description Description: string
//
//     Enter a description for this record.
// doesNotAddToTotal Does Not Add to Transaction Total: boolean
//
//     Check this box for special tax types that do not impact the transaction total, such as withholding tax. Tax types with this property can be used on sales and purchase transactions.
// externalId External ID: string
// id Internal ID: string
// isInactive Inactive: boolean
//
//     Check this box to inactivate this record. Inactive records do not show on transactions and records for selection in lists.
// links Links: nsLink
// nsLink
// name Name: string
//
//     Enter a name for this record.
// nexusAccounts: taxType-nexusAccountsCollection
// postToItemCost Post to Item Cost: boolean
//
//     Check this box if the tax amounts for the tax type should be added to the related item cost. Tax types with this property can be used on purchase transactions only.
// refName Reference Name: string
// reverseCharge Reverse Charge: boolean
//
//     Check this box if reverse charge applies to transactions that have this tax type. Tax types with this property can be used on purchase transactions only.
// taxInNetAmount Tax Included in Net Amount: boolean
//
//     Check this box for special tax types where the tax amount is included in the item price. Tax types with this property can be used on sales and purchase transactions.

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
