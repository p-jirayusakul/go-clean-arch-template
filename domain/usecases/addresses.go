package usecases

import "github.com/p-jirayusakul/go-clean-arch-template/domain/entities"

type AddressesUsecase interface {
	CreateAddresses(arg entities.AddressesDto) (id string, err error)
	ListAddressesAddresses(arg string) (result []entities.Addresses, err error)
	UpdateAddresses(arg entities.AddressesDto) (err error)
	DeleteAddresses(arg entities.AddressesDto) (err error)
}
