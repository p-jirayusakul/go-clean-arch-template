package factories

import (
	database "github.com/p-jirayusakul/go-clean-arch-template/database/sqlc"
	"github.com/p-jirayusakul/go-clean-arch-template/domain/repositories"
	db_repositories "github.com/p-jirayusakul/go-clean-arch-template/internal/repositories/db"
)

type DBFactory struct {
	AccountsRepo  repositories.AccountsRepository
	AddressesRepo repositories.AddressesRepository
}

func NewDBFactory(db *database.Queries) *DBFactory {
	var (
		AccountsRepo  = db_repositories.NewAccountsRepository(db)
		AddressesRepo = db_repositories.NewAddressesRepository(db)
	)

	return &DBFactory{
		AccountsRepo:  &AccountsRepo,
		AddressesRepo: &AddressesRepo,
	}
}
