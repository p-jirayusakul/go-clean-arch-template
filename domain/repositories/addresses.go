package repositories

import (
	"context"

	"github.com/p-jirayusakul/go-clean-arch-template/domain/entities"
)

type AddressesRepository interface {
	CreateAddresses(ctx context.Context, p entities.AddressesDto) (string, error)
	ListAddressesByAccountId(ctx context.Context, p string) ([]entities.Addresses, error)
	GetAddressById(ctx context.Context, p string) (entities.Addresses, error)
	UpdateAddressById(ctx context.Context, p entities.AddressesDto) error
	DeleteAddressesById(ctx context.Context, p string) error
	IsAddressesAlreadyExists(ctx context.Context, p entities.AddressesDto) (bool, error)
}
