package db

import (
	"context"
	"fmt"

	"github.com/p-jirayusakul/go-clean-arch-template/domain/entities"
)

type SearchAddressesParams struct {
	PageNumber    int
	PageSize      int
	City          string
	StateProvince string
	PostalCode    string
	Country       string
	AccountsID    string
	OrderBy       string
	OrderType     string
}

func (s *SQLStore) SearchAddresses(ctx context.Context, params SearchAddressesParams) (*entities.AddressesQueryResult, error) {

	var where string
	order := params.OrderBy
	orderType := params.OrderType

	if params.OrderBy == "" {
		order = "updated_at"
	}

	if params.OrderType == "" {
		orderType = "DESC"
	}

	args := []interface{}{params.PageSize, params.PageNumber}

	// key คือ column ส่วน value คือค่าที่ได้จาก params
	keys := map[string]string{
		"city":           params.City,
		"state_province": params.StateProvince,
		"postal_code":    params.PostalCode,
		"country":        params.Country,
		"accounts_id":    params.AccountsID,
	}

	where, args = s.AddCondition(keys, args)

	query := fmt.Sprintf("SELECT id, street_address, city, state_province, postal_code, country, accounts_id FROM public.addresses %s ORDER BY %s %s LIMIT $1 OFFSET $2;", where, order, orderType)
	rows, err := s.connPool.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []entities.Addresses{}
	for rows.Next() {
		var i entities.Addresses
		if err := rows.Scan(
			&i.ID,
			&i.StreetAddress,
			&i.City,
			&i.StateProvince,
			&i.PostalCode,
			&i.Country,
			&i.AccountsID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	var totalItems int64
	if len(items) > 0 {
		where, args = s.AddCondition(keys, []interface{}{})
		queryTotal := fmt.Sprintf("SELECT count(id) as total FROM public.addresses %s", where)
		rowTotalItems := s.connPool.QueryRow(ctx, queryTotal, args...)
		err = rowTotalItems.Scan(&totalItems)
		if err != nil {
			return nil, err
		}
	}

	return &entities.AddressesQueryResult{
		Data:       items,
		TotalItems: int(totalItems),
	}, nil
}
