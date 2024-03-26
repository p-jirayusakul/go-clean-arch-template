package factories

import (
	"context"
	"fmt"
)

type SearchAddressesRow struct {
	ID            string  `json:"id"`
	StreetAddress *string `json:"street_address"`
	City          string  `json:"city"`
	StateProvince string  `json:"state_province"`
	PostalCode    string  `json:"postal_code"`
	Country       string  `json:"country"`
	AccountsID    *string `json:"accounts_id"`
}

type SearchAddressesResult struct {
	Data       []SearchAddressesRow `json:"data"`
	TotalItems int64                `json:"totalItems"`
}

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

func (s *SQLStore) SearchAddresses(ctx context.Context, params SearchAddressesParams) (*SearchAddressesResult, error) {

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
	keys := make(map[string]interface{})

	if params.City != "" {
		keys["city"] = params.City
	}

	if params.StateProvince != "" {
		keys["state_province"] = params.StateProvince
	}

	if params.PostalCode != "" {
		keys["postal_code"] = params.PostalCode
	}

	if params.Country != "" {
		keys["country"] = params.Country
	}

	if params.AccountsID != "" {
		keys["accounts_id"] = params.AccountsID
	}

	where, args = s.AddCondition(keys, args)

	query := fmt.Sprintf("SELECT id, street_address, city, state_province, postal_code, country, accounts_id FROM public.addresses %s ORDER BY %s %s LIMIT $1 OFFSET $2;", where, order, orderType)
	rows, err := s.connPool.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []SearchAddressesRow{}
	for rows.Next() {
		var i SearchAddressesRow
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

	return &SearchAddressesResult{
		Data:       items,
		TotalItems: totalItems,
	}, nil
}
