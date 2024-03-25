package addresses

import (
	"context"

	database "github.com/p-jirayusakul/go-clean-arch-template/database/sqlc"
	"github.com/p-jirayusakul/go-clean-arch-template/domain/entities"
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
