package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	database "github.com/p-jirayusakul/go-clean-arch-template/database/sqlc"
	"github.com/p-jirayusakul/go-clean-arch-template/domain/entities"
)

type AddressesRepository struct {
	db *database.Queries
}

func NewAddressesRepository(db *database.Queries) AddressesRepository {
	return AddressesRepository{db: db}
}

func (x *AddressesRepository) CreateAddresses(ctx context.Context, addresses entities.AddressesDto) (result string, err error) {
	params := database.CreateAddressesParams{
		City:          addresses.City,
		StateProvince: addresses.StateProvince,
		PostalCode:    addresses.PostalCode,
		Country:       addresses.Country,
		AccountsID:    addresses.AccountsID,
	}

	if addresses.StreetAddress != nil {
		params.StreetAddress = pgtype.Text{String: *addresses.StreetAddress, Valid: true}
	}

	result, err = x.db.CreateAddresses(ctx, params)
	if err != nil {
		return
	}

	return
}

func (x *AddressesRepository) ListAddressesByAccountId(ctx context.Context, accountsID string) (result []entities.Addresses, err error) {

	r, err := x.db.ListAddressesByAccountId(ctx, accountsID)
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

func (x *AddressesRepository) GetAddressById(ctx context.Context, addressesID string) (result entities.Addresses, err error) {

	r, err := x.db.GetAddressById(ctx, addressesID)
	if err != nil {
		return
	}

	result = entities.Addresses{
		ID:            r.ID,
		City:          r.City,
		StateProvince: r.StateProvince,
		PostalCode:    r.PostalCode,
		Country:       r.Country,
	}

	if r.StreetAddress.Valid {
		result.StreetAddress = &r.StreetAddress.String
	}

	if r.AccountsID.Valid {
		result.AccountsID = &r.AccountsID.String
	}

	return
}

func (x *AddressesRepository) UpdateAddressById(ctx context.Context, addresses entities.AddressesDto) (err error) {

	params := database.UpdateAddressByIdParams{
		ID:            addresses.ID,
		City:          addresses.City,
		StateProvince: addresses.StateProvince,
		PostalCode:    addresses.PostalCode,
		Country:       addresses.Country,
		AccountsID:    addresses.AccountsID,
	}

	if addresses.StreetAddress != nil {
		params.StreetAddress = pgtype.Text{String: *addresses.StreetAddress, Valid: true}
	}

	err = x.db.UpdateAddressById(ctx, params)
	if err != nil {
		return
	}

	return
}

func (x *AddressesRepository) DeleteAddressesById(ctx context.Context, addressesID string) (err error) {

	err = x.db.DeleteAddressesById(ctx, addressesID)
	if err != nil {
		return
	}

	return
}

func (x *AddressesRepository) IsAddressesAlreadyExists(ctx context.Context, addresses entities.AddressesDto) (result bool, err error) {

	params := database.IsAddressesAlreadyExistsParams{
		ID:         addresses.ID,
		AccountsID: addresses.AccountsID,
	}

	result, err = x.db.IsAddressesAlreadyExists(ctx, params)
	if err != nil {
		return
	}

	return
}
