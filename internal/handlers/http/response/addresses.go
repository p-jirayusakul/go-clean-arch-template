package response

import "github.com/p-jirayusakul/go-clean-arch-template/domain/entities"

type SearchAddressesResponse struct {
	PageNumber int                  `json:"pageNumber"`
	PageSize   int                  `json:"pageSize"`
	TotalPages int                  `json:"totalPages"`
	TotalItems int                  `json:"totalItems"`
	Data       []entities.Addresses `json:"data"`
}
