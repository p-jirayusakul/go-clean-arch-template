package entities

type Addresses struct {
	ID            string  `json:"id"`
	StreetAddress *string `json:"street_address"`
	City          string  `json:"city"`
	StateProvince string  `json:"state_province"`
	PostalCode    string  `json:"postal_code"`
	Country       string  `json:"country"`
	AccountsID    *string `json:"accounts_id"`
}

type AddressesDto struct {
	ID            string  `json:"id"`
	StreetAddress *string `json:"street_address"`
	City          string  `json:"city"`
	StateProvince string  `json:"state_province"`
	PostalCode    string  `json:"postal_code"`
	Country       string  `json:"country"`
	AccountsID    string  `json:"accounts_id"`
}
