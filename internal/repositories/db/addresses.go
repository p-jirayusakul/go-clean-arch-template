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

func (x *AddressesRepository) CreateAddresses(ctx context.Context, p entities.AddressesDto) (string, error) {
	params := database.CreateAddressesParams{
		City:          p.City,
		StateProvince: p.StateProvince,
		PostalCode:    p.PostalCode,
		Country:       p.Country,
		AccountsID:    p.AccountsID,
	}

	if p.StreetAddress != nil {
		params.StreetAddress = pgtype.Text{String: *p.StreetAddress, Valid: true}
	}

	r, err := x.db.CreateAddresses(ctx, params)
	if err != nil {
		return "", err
	}

	return r, nil
}

func (x *AddressesRepository) ListAddressesByAccountId(ctx context.Context, p string) ([]entities.Addresses, error) {

	r, err := x.db.ListAddressesByAccountId(ctx, p)
	if err != nil {
		return []entities.Addresses{}, err
	}

	result := []entities.Addresses{}
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

	return result, nil
}

func (x *AddressesRepository) GetAddressById(ctx context.Context, p string) (entities.Addresses, error) {

	r, err := x.db.GetAddressById(ctx, p)
	if err != nil {
		return entities.Addresses{}, err
	}

	result := entities.Addresses{
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

	return result, nil
}

func (x *AddressesRepository) UpdateAddressById(ctx context.Context, p entities.AddressesDto) error {

	params := database.UpdateAddressByIdParams{
		ID:            p.ID,
		City:          p.City,
		StateProvince: p.StateProvince,
		PostalCode:    p.PostalCode,
		Country:       p.Country,
		AccountsID:    p.AccountsID,
	}

	if p.StreetAddress != nil {
		params.StreetAddress = pgtype.Text{String: *p.StreetAddress, Valid: true}
	}

	err := x.db.UpdateAddressById(ctx, params)
	if err != nil {
		return err
	}

	return nil
}

func (x *AddressesRepository) DeleteAddressesById(ctx context.Context, p string) error {

	err := x.db.DeleteAddressesById(ctx, p)
	if err != nil {
		return err
	}

	return nil
}

func (x *AddressesRepository) IsAddressesAlreadyExists(ctx context.Context, p entities.AddressesDto) (bool, error) {

	params := database.IsAddressesAlreadyExistsParams{
		ID:         p.ID,
		AccountsID: p.AccountsID,
	}

	r, err := x.db.IsAddressesAlreadyExists(ctx, params)
	if err != nil {
		return false, err
	}

	return r, nil
}
