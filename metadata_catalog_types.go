package netsuite

// RecordSchema is what the metadata catalog returns for one record type: the
// JSON schema (draft-06 hyper-schema) of the record as the REST API reads and
// writes it, with NetSuite's own extensions under x-ns- keys. It is the one
// place the REST API describes an account's custom fields - custbody_,
// custcol_, custentity_, custitem_ and custrecord_ - since SuiteQL exposes no
// table of custom field definitions.
type RecordSchema struct {
	Schema string `json:"$schema"`
	Type   string `json:"type"`

	// Properties are the fields of the record, keyed by the name the REST
	// API uses for them. A sublist such as journalEntry's line is a property
	// of type object whose items property holds the array of rows.
	Properties map[string]SchemaProperty `json:"properties"`

	// Filterable names the properties a collection GET on the record accepts
	// in its q parameter.
	Filterable []string `json:"x-ns-filterable"`

	// CustomRecord is set on a custom record type's schema, such as
	// customrecord_cseg1, the values record type of a custom segment.
	CustomRecord bool `json:"x-ns-custom-record,omitempty"`
}

// SchemaProperty is one property of a RecordSchema, or of an object nested
// in one. Which fields are set depends on what it describes: a scalar field
// has a Type and often a Format; a reference to another record is an object
// whose Properties are id, refName, externalId and links; a sublist is an
// object holding an items array whose Items describe a row; a link is an
// array whose Items carry a Ref to the nsLink schema.
type SchemaProperty struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Type        string `json:"type,omitempty"`
	Format      string `json:"format,omitempty"`
	Nullable    bool   `json:"nullable,omitempty"`
	ReadOnly    bool   `json:"readOnly,omitempty"`

	// Pattern is a regular expression a string property must match. Few
	// record types use it: message, task and phonecall do, a journal entry
	// does not.
	Pattern string `json:"pattern,omitempty"`

	// Enum lists the values a string property accepts. NetSuite uses it for
	// the id of a reference to a fixed list, such as a country code or a
	// schedule type.
	Enum []string `json:"enum,omitempty"`

	// Ref points at another schema in the catalog instead of describing the
	// property inline, as a path such as /services/rest/record/v1/metadata-catalog/nsLink.
	Ref string `json:"$ref,omitempty"`

	Items      *SchemaProperty           `json:"items,omitempty"`
	Properties map[string]SchemaProperty `json:"properties,omitempty"`

	// CustomField marks a field the account added itself rather than one
	// NetSuite ships with the record type. Its key in Properties is the
	// field's script ID in lower case.
	CustomField bool `json:"x-ns-custom-field,omitempty"`

	// Filterable and SublistKey are set on the object describing a sublist
	// row: which of its properties a sublist GET accepts in its q parameter,
	// and which identify a row when the parent record is written.
	Filterable []string          `json:"x-ns-filterable,omitempty"`
	SublistKey *SchemaSublistKey `json:"x-ns-sublistkey,omitempty"`
}

// SchemaSublistKey says how the rows of a sublist are identified when a
// record is written: Present is false for a sublist whose rows are replaced
// wholesale, and true for one whose rows are matched on the properties in
// Value.
type SchemaSublistKey struct {
	Present bool                   `json:"present"`
	Value   *SchemaSublistKeyValue `json:"value,omitempty"`
}

// SchemaSublistKeyValue names the properties that identify a sublist row, for
// a row being added and for one that already exists.
type SchemaSublistKeyValue struct {
	New      []string `json:"new"`
	Existing []string `json:"existing"`
}

// MetadataCatalogItem is one record type the metadata catalog lists. Its
// links point at the schema (application/schema+json), the swagger document
// (application/swagger+json) and the record endpoint the schema describes.
type MetadataCatalogItem struct {
	Name  string `json:"name"`
	Links Links  `json:"links"`
}
