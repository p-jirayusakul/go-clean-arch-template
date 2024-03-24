package addresses

import (
	"context"

	"github.com/p-jirayusakul/go-clean-arch-template/domain/entities"
	"github.com/p-jirayusakul/go-clean-arch-template/pkg/common"
)

func (x *addressesInteractor) CreateAddresses(arg entities.AddressesDto) (id string, err error) {
	ctx := context.Background()

	id, err = x.addressesRepo.CreateAddresses(ctx, arg)
	if err != nil {
		return
	}

	return
}

func (x *addressesInteractor) ListAddressesAddresses(arg string) (result []entities.Addresses, err error) {
	ctx := context.Background()

	result, err = x.addressesRepo.ListAddressesByAccountId(ctx, arg)
	if err != nil {
		return
	}

	return
}

func (x *addressesInteractor) UpdateAddresses(arg entities.AddressesDto) (err error) {
	ctx := context.Background()

	isAlreadyExists, err := x.addressesRepo.IsAddressesAlreadyExists(ctx, entities.AddressesDto{
		ID:         arg.ID,
		AccountsID: arg.AccountsID,
	})

	if err != nil {
		return
	}

	if !isAlreadyExists {
		return common.ErrDataNotFound
	}

	err = x.addressesRepo.UpdateAddressById(ctx, arg)
	if err != nil {
		return
	}

	return
}

func (x *addressesInteractor) DeleteAddresses(arg entities.AddressesDto) (err error) {
	ctx := context.Background()

	isAlreadyExists, err := x.addressesRepo.IsAddressesAlreadyExists(ctx, entities.AddressesDto{
		ID:         arg.ID,
		AccountsID: arg.AccountsID,
	})

	if err != nil {
		return
	}

	if !isAlreadyExists {
		return common.ErrDataNotFound
	}

	err = x.addressesRepo.DeleteAddressesById(ctx, arg.ID)
	if err != nil {
		return
	}

	return
}
