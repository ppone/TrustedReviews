package places

type placesData struct {
	factual_id       string "The Factual ID"
	name             string "Business/POI name"
	address          string "Address number and street name"
	address_extended string "Additional address, incl. suite numbers"
	locality         string "City, town or equivalent"
	region           string "State, province, territory, or equivalent"
	postcode         string "Postcode or equivalent (zipcode in US)"
	country          string
	neighborhood     string  "The neighborhood(s) or other informal geography in which this entity is found."
	tel              string  "Telephone number with local formatting"
	fax              string  "Fax number in local formatting"
	website          string  "Authority page (official website)"
	latitude         float64 "Latitude in decimal degrees (WGS84 datum). Value will not exceed 6 decimal places (0.111m)"
	longitude        float64 "Longitude in decimal degrees (WGS84 datum). Value will not exceed 6 decimal places (0.111m)"
	status           string  "Is the business a going concern: closed (0) or open (1). We are aware that this will prove confusing to electrical engineers. Deprecated, as we now expose only open businesses."
	hours_display    string  "Structured JSON representation of opening hours"
	chain_name       string  "Label indicating which chain (brand or franchise) this entity is a member of. See documentation for more information on Factual Chains."
	email            string  "Primary contact email address of organization"
	spost_town       string  "Town/place employed in postal addressing. May not reflect the formal geographic location of a place."
	category_ids     int     "Category IDs that classify this entity."
	admin_region     string  "Additional sub-division. Usually, but not always, a country sub-division"
	chain_id         string  "Indicates which chain (brand or franchise) this entity is a member of. See documentation for more information on Factual Chains."
	hours            string  "JSON representation of hours of operation"
	po_box           string  "PO Box. As they do not represent the physical location of a brick-and-mortar store, PO Boxes are often excluded from mobile use cases."
}

func func_name() {

}
