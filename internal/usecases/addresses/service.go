package addresses

import (
	"context"

	database "github.com/p-jirayusakul/go-clean-arch-template/database/sqlc"
	"github.com/p-jirayusakul/go-clean-arch-template/domain/entities"
	"github.com/p-jirayusakul/go-clean-arch-template/pkg/common"
)

func (x *addressesInteractor) CreateAddresses(addresses entities.AddressesDto) (id string, err error) {
	ctx := context.Background()

	params := database.CreateAddressesParams{
		City:          addresses.City,
		StateProvince: addresses.StateProvince,
		PostalCode:    addresses.PostalCode,
		Country:       addresses.Country,
		AccountsID:    addresses.AccountsID,
	}

	if addresses.StreetAddress != nil {
		params.StreetAddress.String = *addresses.StreetAddress
		params.StreetAddress.Valid = true
	}

	id, err = x.dbFactory.CreateAddresses(ctx, params)
	if err != nil {
		return
	}

	return
}

func (x *addressesInteractor) ListAddressesAddresses(addressesID string) (result []entities.Addresses, err error) {
	ctx := context.Background()

	r, err := x.dbFactory.ListAddressesByAccountId(ctx, addressesID)
	if err != nil {
		return
	}

	for _, data := range r {
		arg := entities.Addresses{
			ID:            data.ID,
			City:          data.City,
			StateProvince: data.StateProvince,
			PostalCode:    data.PostalCode,
			Country:       data.Country,
		}

		if data.StreetAddress.Valid {
			arg.StreetAddress = &data.StreetAddress.String
		}

		if data.AccountsID.Valid {
			arg.AccountsID = &data.AccountsID.String
		}

		result = append(result, arg)
	}

	return
}

func (x *addressesInteractor) UpdateAddresses(addresses entities.AddressesDto) (err error) {
	ctx := context.Background()

	isAlreadyExists, err := x.dbFactory.IsAddressesAlreadyExists(ctx, database.IsAddressesAlreadyExistsParams{
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
		City:          addresses.City,
		StateProvince: addresses.StateProvince,
		PostalCode:    addresses.PostalCode,
		Country:       addresses.Country,
		AccountsID:    addresses.AccountsID,
	}

	if addresses.StreetAddress != nil {
		params.StreetAddress.String = *addresses.StreetAddress
		params.StreetAddress.Valid = true
	}

	err = x.dbFactory.UpdateAddressById(ctx, params)
	if err != nil {
		return
	}

	return
}

func (x *addressesInteractor) DeleteAddresses(addresses entities.AddressesDto) (err error) {
	ctx := context.Background()

	isAlreadyExists, err := x.dbFactory.IsAddressesAlreadyExists(ctx, database.IsAddressesAlreadyExistsParams{
		ID:         addresses.ID,
		AccountsID: addresses.AccountsID,
	})

	if err != nil {
		return
	}

	if !isAlreadyExists {
		return common.ErrDataNotFound
	}

	err = x.dbFactory.DeleteAddressesById(ctx, addresses.ID)
	if err != nil {
		return
	}

	return
}
