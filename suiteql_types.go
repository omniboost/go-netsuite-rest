package netsuite

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
