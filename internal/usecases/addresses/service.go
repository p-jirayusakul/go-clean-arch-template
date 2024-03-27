package addresses

import (
	"context"
	"math"

	database "github.com/p-jirayusakul/go-clean-arch-template/database/sqlc"
	"github.com/p-jirayusakul/go-clean-arch-template/domain/entities"
	"github.com/p-jirayusakul/go-clean-arch-template/internal/repositories/db"
	"github.com/p-jirayusakul/go-clean-arch-template/pkg/common"
)

func (s *addressesInteractor) CreateAddresses(addresses entities.AddressesDto) (id string, err error) {
	ctx := context.Background()

	params := database.CreateAddressesParams{
		City:          addresses.City,
		StreetAddress: addresses.StreetAddress,
		StateProvince: addresses.StateProvince,
		PostalCode:    addresses.PostalCode,
		Country:       addresses.Country,
		AccountsID:    addresses.AccountsID,
	}

	id, err = s.store.CreateAddresses(ctx, &params)
	if err != nil {
		return
	}

	return
}

func (s *addressesInteractor) ListAddressesAddresses(addressesID string) (result []entities.Addresses, err error) {
	ctx := context.Background()

	r, err := s.store.ListAddressesByAccountId(ctx, addressesID)
	if err != nil {
		return
	}

	for _, data := range r {
		arg := entities.Addresses{
			ID:            data.ID,
			StreetAddress: data.StreetAddress,
			City:          data.City,
			StateProvince: data.StateProvince,
			PostalCode:    data.PostalCode,
			Country:       data.Country,
			AccountsID:    data.AccountsID,
		}

		result = append(result, arg)
	}

	return
}

func (s *addressesInteractor) SearchAddresses(addressesQuery entities.AddressesQueryParams) (result *entities.AddressesQueryResult, err error) {
	ctx := context.Background()

	pageNumber := addressesQuery.PageNumber
	pageSize := addressesQuery.PageSize

	if addressesQuery.PageSize == 0 {
		pageSize = common.PAGE_SIZE
	}

	if pageSize > common.MAX_PAGE_SIZE {
		pageSize = common.MAX_PAGE_SIZE
	}

	if pageNumber > 0 {
		pageNumber = (pageNumber - 1) * pageSize
	}

	arg := db.SearchAddressesParams{
		PageNumber:    pageNumber,
		PageSize:      pageSize,
		City:          addressesQuery.City,
		PostalCode:    addressesQuery.PostalCode,
		StateProvince: addressesQuery.StateProvince,
		Country:       addressesQuery.Country,
		AccountsID:    addressesQuery.AccountsID,
		OrderBy:       addressesQuery.OrderBy,
		OrderType:     addressesQuery.OrderType,
	}

	result, err = s.store.SearchAddresses(ctx, arg)
	if err != nil {
		return
	}

	// convert some data before response
	var totalPages int
	if len(result.Data) > 0 {
		result.TotalItems = int(result.TotalItems)
		totalPages = int(math.Ceil(float64(result.TotalItems) / float64(pageSize)))
	} else {
		result.TotalItems = 0
		totalPages = 0
	}

	if addressesQuery.PageNumber == 0 {
		pageNumber = 1
	} else {
		pageNumber = addressesQuery.PageNumber
	}

	result.PageNumber = pageNumber
	result.PageSize = pageSize
	result.TotalPages = totalPages

	return
}

func (s *addressesInteractor) UpdateAddresses(addresses entities.AddressesDto) (err error) {
	ctx := context.Background()

	isAlreadyExists, err := s.store.IsAddressesAlreadyExists(ctx, &database.IsAddressesAlreadyExistsParams{
		ID:         addresses.ID,
		AccountsID: addresses.AccountsID,
	})

	if err != nil {
		return
	}

	if !isAlreadyExists {
		return common.ErrDataNotFound
	}

	params := database.UpdateAddressByIdParams{
		ID:            addresses.ID,
		StreetAddress: addresses.StreetAddress,
		City:          addresses.City,
		StateProvince: addresses.StateProvince,
		PostalCode:    addresses.PostalCode,
		Country:       addresses.Country,
		AccountsID:    addresses.AccountsID,
	}

	err = s.store.UpdateAddressById(ctx, &params)
	if err != nil {
		return
	}

	return
}

func (s *addressesInteractor) DeleteAddresses(addresses entities.AddressesDto) (err error) {
	ctx := context.Background()

	params := &database.IsAddressesAlreadyExistsParams{
		ID:         addresses.ID,
		AccountsID: addresses.AccountsID,
	}
	isAlreadyExists, err := s.store.IsAddressesAlreadyExists(ctx, params)

	if err != nil {
		return
	}

	if !isAlreadyExists {
		return common.ErrDataNotFound
	}

	err = s.store.DeleteAddressesById(ctx, addresses.ID)
	if err != nil {
		return
	}

	return
}
