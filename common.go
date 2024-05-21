package squareup

type AmountMoney struct {
	Amount   int    `json:"amount"`
	Currency string `json:"currency"`
}

type DeviceOptions struct {
	DeviceId          string      `json:"device_id"`
	CollectSignature  bool        `json:"collect_signature"`
	TipSettings       TipSettings `json:"tip_settings"`
	SkipReceiptScreen bool        `json:"skip_receipt_screen"`
}

type PaymentOptions struct {
	Autocomplete bool `json:"autocomplete"`
}

type BillingAddress struct {
	AddressLine1                 string `json:"address_line_1"`
	AddressLine2                 string `json:"address_line_2"`
	AddressLine3                 string `json:"address_line_3"`
	AdministrativeDistrictLevel1 string `json:"administrative_district_level_1"`
	AdministrativeDistrictLevel2 string `json:"administrative_district_level_2"`
	AdministrativeDistrictLevel3 string `json:"administrative_district_level_3"`
	Country                      string `json:"country"`
	FirstName                    string `json:"first_name"`
	LastName                     string `json:"last_name"`
	Locality                     string `json:"locality"`
	PostalCode                   string `json:"postal_code"`
	Sublocality                  string `json:"sublocality"`
	Sublocality2                 string `json:"sublocality_2"`
	Sublocality3                 string `json:"sublocality_3"`
}

type CustomerDetail struct {
	CustomerInitiated bool `json:"customer_initiated"`
	SellerKeyedIn     bool `json:"seller_keyed_in"`
}
