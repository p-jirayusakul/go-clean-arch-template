package repositories

import (
	"context"

	"github.com/p-jirayusakul/go-clean-arch-template/domain/entities"
)

type AddressesRepository interface {
	CreateAddresses(ctx context.Context, addresses entities.AddressesDto) (result string, err error)
	ListAddressesByAccountId(ctx context.Context, accountsID string) (result []entities.Addresses, err error)
	GetAddressById(ctx context.Context, addressesID string) (result entities.Addresses, err error)
	UpdateAddressById(ctx context.Context, addresses entities.AddressesDto) (err error)
	DeleteAddressesById(ctx context.Context, addressesID string) (err error)
	IsAddressesAlreadyExists(ctx context.Context, addresses entities.AddressesDto) (result bool, err error)
}
