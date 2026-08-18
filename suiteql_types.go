package netsuite

type SQLSubsidiaries []SQLSubsidiary

// SQLSubsidiary is a row of the SuiteQL subsidiary table. SuiteQL names the
// columns as they are spelled in the table, in lower case, rather than in the
// camel case the record endpoints use, so the tags here read "isinactive" where
// Subsidiary reads "isInactive". A column that is null for a row is left out of
// the response entirely, which is why every field is omitempty: the root
// subsidiary carries no parent.
type SQLSubsidiary struct {
	ID            string `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	FullName      string `json:"fullname,omitempty"`
	LegalName     string `json:"legalname,omitempty"`
	Parent        string `json:"parent,omitempty"`
	Country       string `json:"country,omitempty"`
	Currency      string `json:"currency,omitempty"`
	TranPrefix    string `json:"tranprefix,omitempty"`
	IsInactive    Bool   `json:"isinactive,omitempty"`
	IsElimination Bool   `json:"iselimination,omitempty"`
	Links         Links  `json:"links,omitempty"`
}

type SQLNexuses []SQLNexus

type SQLNexus struct {
	Country    string     `json:"country,omitzero"`
	CustomForm CustomForm `json:"customForm,omitzero"`
	// Enter a description for this record.
	Description string `json:"description,omitempty"`
	ExternalID  string `json:"externalId,omitempty"`
	ID          string `json:"id,omitempty"`
	IsInactive  Bool   `json:"isInactive,omitempty"`
	Links       Links  `json:"links,omitempty"`
	ParentNexus *Nexus `json:"parentNexus,omitzero"`
	RefName     string `json:"refName,omitempty"`
	State       string `json:"state,omitzero"`
	TaxAgency   string `json:"taxAgency,omitzero"`
	// Check this box to use the fulfillment date as the tax point date for this
	// transaction. This overrides the default Nexus setting which uses the
	// transaction date to determine the tax point date. This setting is
	// typically used for jurisdictions which require the item fulfillment date
	// to be used to define the tax point date.
	TaxDateFromFulfillment Bool `json:"taxDateFromFulfillment,omitempty"`
}
