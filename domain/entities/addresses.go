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

type AddressesQueryParams struct {
	PageNumber    int
	PageSize      int
	City          string
	StateProvince string
	PostalCode    string
	Country       string
	AccountsID    string
	OrderBy       string
	OrderType     string
}

type AddressesQueryResult struct {
	Data       []Addresses `json:"data"`
	TotalItems int         `json:"totalItems"`
	TotalPages int         `json:"totalPages"`
	PageNumber int         `json:"pageNumber"`
	PageSize   int         `json:"pageSize"`
}
