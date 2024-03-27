package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	database "github.com/p-jirayusakul/go-clean-arch-template/database/sqlc"
	"github.com/p-jirayusakul/go-clean-arch-template/domain/entities"
	handlers "github.com/p-jirayusakul/go-clean-arch-template/internal/handlers/http"
	"github.com/p-jirayusakul/go-clean-arch-template/internal/handlers/http/request"
	"github.com/p-jirayusakul/go-clean-arch-template/internal/repositories/db"
	"github.com/p-jirayusakul/go-clean-arch-template/pkg/common"
	"github.com/p-jirayusakul/go-clean-arch-template/pkg/config"
	"github.com/p-jirayusakul/go-clean-arch-template/pkg/middleware"
	"github.com/p-jirayusakul/go-clean-arch-template/pkg/utils"
	"github.com/p-jirayusakul/go-clean-arch-template/test/mockup"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestCreateAddress(t *testing.T) {
	testCases := []struct {
		name          string
		body          string
		buildStubs    func(store *mockup.MockStore, body request.CreateAddressesRequest)
		checkResponse func(t *testing.T, status int, err error)
	}{
		{
			name: "OK",
			body: `{"address":"address","city":"city","province":"province","postalCode":"postalCode","country":"country"}`,
			buildStubs: func(store *mockup.MockStore, body request.CreateAddressesRequest) {
				store.EXPECT().IsAccountAlreadyExists(gomock.Any(), uid).Times(1).Return(true, nil)
				store.EXPECT().CreateAddresses(gomock.Any(), &database.CreateAddressesParams{
					StreetAddress: body.Address,
					City:          body.City,
					StateProvince: body.Province,
					PostalCode:    body.PostalCode,
					Country:       body.Country,
					AccountsID:    uid,
				}).Times(1).Return("uuid", nil)
			},
			checkResponse: func(t *testing.T, status int, err error) {
				require.NoError(t, err)
				require.Equal(t, http.StatusCreated, status)
			},
		},
		{
			name: "unauthorized - accounts id is invalid",
			body: `{"address":"address","city":"city","province":"province","postalCode":"postalCode","country":"country"}`,
			buildStubs: func(store *mockup.MockStore, body request.CreateAddressesRequest) {
				store.EXPECT().IsAccountAlreadyExists(gomock.Any(), uid).Times(1).Return(false, nil)
			},
			checkResponse: func(t *testing.T, status int, err error) {
				require.Error(t, err)
				require.Equal(t, utils.ReplaceStringError(http.StatusUnauthorized, err.Error()), common.ErrAccountIsInvalid.Error())
			},
		},
		{
			name: "Bad Request - city is required",
			body: `{"address":"address","province":"province","postalCode":"postalCode","country":"country"}`,
			buildStubs: func(store *mockup.MockStore, body request.CreateAddressesRequest) {
			},
			checkResponse: func(t *testing.T, status int, err error) {
				require.Error(t, err)
				require.Equal(t, utils.ReplaceStringError(http.StatusBadRequest, err.Error()), "Key: 'CreateAddressesRequest.City' Error:Field validation for 'City' failed on the 'required' tag")
			},
		},
		{
			name: "Bad Request - province is required",
			body: `{"address":"address","city":"city","postalCode":"postalCode","country":"country"}`,
			buildStubs: func(store *mockup.MockStore, body request.CreateAddressesRequest) {
			},
			checkResponse: func(t *testing.T, status int, err error) {
				require.Error(t, err)
				require.Equal(t, utils.ReplaceStringError(http.StatusBadRequest, err.Error()), "Key: 'CreateAddressesRequest.Province' Error:Field validation for 'Province' failed on the 'required' tag")
			},
		},
		{
			name: "Bad Request - postalCode is required",
			body: `{"address":"address","city":"city","province":"province","country":"country"}`,
			buildStubs: func(store *mockup.MockStore, body request.CreateAddressesRequest) {
			},
			checkResponse: func(t *testing.T, status int, err error) {
				require.Error(t, err)
				require.Equal(t, utils.ReplaceStringError(http.StatusBadRequest, err.Error()), "Key: 'CreateAddressesRequest.PostalCode' Error:Field validation for 'PostalCode' failed on the 'required' tag")
			},
		},
		{
			name: "Bad Request - country is required",
			body: `{"address":"address","city":"city","province":"province","postalCode":"postalCode"}`,
			buildStubs: func(store *mockup.MockStore, body request.CreateAddressesRequest) {
			},
			checkResponse: func(t *testing.T, status int, err error) {
				require.Error(t, err)
				require.Equal(t, utils.ReplaceStringError(http.StatusBadRequest, err.Error()), "Key: 'CreateAddressesRequest.Country' Error:Field validation for 'Country' failed on the 'required' tag")
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			cfg := config.InitConfigs(".env")
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			var dto request.CreateAddressesRequest
			err := json.Unmarshal([]byte(tc.body), &dto)
			require.NoError(t, err)

			dbFactory := mockup.NewMockStore(ctrl)
			distributor := mockup.NewMockTaskDistributor(ctrl)
			tc.buildStubs(dbFactory, dto)

			app := echo.New()
			app.Validator = middleware.NewCustomValidator()
			app.Use(middleware.ErrorHandler)

			req := httptest.NewRequest(http.MethodPost, "/api/v1/profile/addresses", strings.NewReader(tc.body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := app.NewContext(req, rec)
			c.Set("accountsID", uid)
			handler := handlers.NewServerHttpHandler(app, &cfg, distributor, dbFactory)

			err = handler.CreateAddresses(c)
			tc.checkResponse(t, c.Response().Status, err)
		})
	}
}

func TestListAddresses(t *testing.T) {
	testCases := []struct {
		name          string
		buildStubs    func(store *mockup.MockStore)
		checkResponse func(t *testing.T, status int, err error)
	}{
		{
			name: "OK",
			buildStubs: func(store *mockup.MockStore) {
				store.EXPECT().IsAccountAlreadyExists(gomock.Any(), uid).Times(1).Return(true, nil)

				var streetAddress *string
				tmp := "addresses"
				streetAddress = &tmp
				store.EXPECT().ListAddressesByAccountId(gomock.Any(), uid).Times(1).Return([]*database.ListAddressesByAccountIdRow{
					{
						ID:            "942524af-9df4-425a-8abc-77e940ef8fcb",
						StreetAddress: streetAddress,
						City:          "city",
						StateProvince: "provice",
						PostalCode:    "pastalCode",
						Country:       "country",
					},
				}, nil)
			},
			checkResponse: func(t *testing.T, status int, err error) {
				require.NoError(t, err)
				require.Equal(t, http.StatusOK, status)
			},
		},
		{
			name: "unauthorized - accounts id is invalid",
			buildStubs: func(store *mockup.MockStore) {
				store.EXPECT().IsAccountAlreadyExists(gomock.Any(), uid).Times(1).Return(false, nil)
			},
			checkResponse: func(t *testing.T, status int, err error) {
				require.Error(t, err)
				require.Equal(t, utils.ReplaceStringError(http.StatusUnauthorized, err.Error()), common.ErrAccountIsInvalid.Error())
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			cfg := config.InitConfigs(".env")
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			dbFactory := mockup.NewMockStore(ctrl)
			distributor := mockup.NewMockTaskDistributor(ctrl)
			tc.buildStubs(dbFactory)

			app := echo.New()
			app.Validator = middleware.NewCustomValidator()
			app.Use(middleware.ErrorHandler)

			req := httptest.NewRequest(http.MethodGet, "/api/v1/profile/addresses/me", nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := app.NewContext(req, rec)
			c.Set("accountsID", uid)
			handler := handlers.NewServerHttpHandler(app, &cfg, distributor, dbFactory)

			err := handler.ListAddresses(c)
			tc.checkResponse(t, c.Response().Status, err)
		})
	}
}

func TestSearchAddresses(t *testing.T) {
	testCases := []struct {
		name          string
		buildStubs    func(store *mockup.MockStore, body request.SearchAddressesRequest)
		checkResponse func(t *testing.T, status int, err error)
	}{
		{
			name: "OK",
			buildStubs: func(store *mockup.MockStore, body request.SearchAddressesRequest) {
				store.EXPECT().IsAccountAlreadyExists(gomock.Any(), uid).Times(1).Return(true, nil)

				var streetAddress *string
				tmp := "addresses"
				streetAddress = &tmp

				r := entities.AddressesQueryResult{
					Data: []entities.Addresses{
						{
							ID:            "942524af-9df4-425a-8abc-77e940ef8fcb",
							StreetAddress: streetAddress,
							City:          "city",
							StateProvince: "provice",
							PostalCode:    "pastalCode",
							Country:       "country",
						},
					},
					TotalItems: int(1),
				}

				store.EXPECT().SearchAddresses(gomock.Any(), db.SearchAddressesParams{
					City:       body.City,
					OrderBy:    body.OrderBy,
					OrderType:  body.OrderType,
					PageSize:   10,
					PageNumber: 0,
				}).Times(1).Return(&r, nil)
			},
			checkResponse: func(t *testing.T, status int, err error) {
				require.NoError(t, err)
				require.Equal(t, http.StatusOK, status)
			},
		},
		{
			name: "unauthorized - accounts id is invalid",
			buildStubs: func(store *mockup.MockStore, body request.SearchAddressesRequest) {
				store.EXPECT().IsAccountAlreadyExists(gomock.Any(), uid).Times(1).Return(false, nil)
			},
			checkResponse: func(t *testing.T, status int, err error) {
				require.Error(t, err)
				require.Equal(t, utils.ReplaceStringError(http.StatusUnauthorized, err.Error()), common.ErrAccountIsInvalid.Error())
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			cfg := config.InitConfigs(".env")
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			q := make(url.Values)
			q.Set("city", "city")
			q.Set("orderBy", "city")
			q.Set("orderType", "desc")

			var dto request.SearchAddressesRequest
			dto.City = "city"
			dto.OrderBy = "city"
			dto.OrderType = "desc"

			dbFactory := mockup.NewMockStore(ctrl)
			distributor := mockup.NewMockTaskDistributor(ctrl)
			tc.buildStubs(dbFactory, dto)

			app := echo.New()
			app.Validator = middleware.NewCustomValidator()
			app.Use(middleware.ErrorHandler)

			req := httptest.NewRequest(http.MethodGet, "/api/v1/profile/addresses?"+q.Encode(), nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := app.NewContext(req, rec)
			c.Set("accountsID", uid)
			handler := handlers.NewServerHttpHandler(app, &cfg, distributor, dbFactory)

			err := handler.SearchAddresses(c)
			tc.checkResponse(t, c.Response().Status, err)
		})
	}
}

func TestUpdateAddresses(t *testing.T) {
	addressesID := "e2109e75-1d9d-48fb-9e68-310d4720b015"
	testCases := []struct {
		name          string
		body          string
		buildStubs    func(store *mockup.MockStore, body request.UpdateAddressesRequest)
		checkResponse func(t *testing.T, status int, err error)
	}{
		{
			name: "OK",
			body: `{"address":"address","city":"city","province":"province","postalCode":"postalCode","country":"country"}`,
			buildStubs: func(store *mockup.MockStore, body request.UpdateAddressesRequest) {
				store.EXPECT().IsAccountAlreadyExists(gomock.Any(), uid).Times(1).Return(true, nil)

				store.EXPECT().IsAddressesAlreadyExists(gomock.Any(), &database.IsAddressesAlreadyExistsParams{
					ID:         addressesID,
					AccountsID: uid,
				}).Times(1).Return(true, nil)

				store.EXPECT().UpdateAddressById(gomock.Any(), &database.UpdateAddressByIdParams{
					ID:            addressesID,
					StreetAddress: body.Address,
					City:          body.City,
					StateProvince: body.Province,
					PostalCode:    body.PostalCode,
					Country:       body.Country,
					AccountsID:    uid,
				}).Times(1).Return(nil)
			},
			checkResponse: func(t *testing.T, status int, err error) {
				require.NoError(t, err)
				require.Equal(t, http.StatusOK, status)
			},
		},
		{
			name: "Not found - address id not found",
			body: `{"address":"address","city":"city","province":"province","postalCode":"postalCode","country":"country"}`,
			buildStubs: func(store *mockup.MockStore, body request.UpdateAddressesRequest) {
				store.EXPECT().IsAccountAlreadyExists(gomock.Any(), uid).Times(1).Return(true, nil)

				store.EXPECT().IsAddressesAlreadyExists(gomock.Any(), &database.IsAddressesAlreadyExistsParams{
					ID:         addressesID,
					AccountsID: uid,
				}).Times(1).Return(false, nil)
			},
			checkResponse: func(t *testing.T, status int, err error) {
				require.Error(t, err)
				require.Equal(t, utils.ReplaceStringError(http.StatusNotFound, err.Error()), common.ErrDataNotFound.Error())
			},
		},
		{
			name: "unauthorized - accounts id is invalid",
			body: `{"address":"address","city":"city","province":"province","postalCode":"postalCode","country":"country"}`,
			buildStubs: func(store *mockup.MockStore, body request.UpdateAddressesRequest) {
				store.EXPECT().IsAccountAlreadyExists(gomock.Any(), uid).Times(1).Return(false, nil)
			},
			checkResponse: func(t *testing.T, status int, err error) {
				require.Error(t, err)
				require.Equal(t, utils.ReplaceStringError(http.StatusUnauthorized, err.Error()), common.ErrAccountIsInvalid.Error())
			},
		},
		{
			name: "Bad Request - city is required",
			body: `{"address":"address","province":"province","postalCode":"postalCode","country":"country"}`,
			buildStubs: func(store *mockup.MockStore, body request.UpdateAddressesRequest) {
			},
			checkResponse: func(t *testing.T, status int, err error) {
				require.Error(t, err)
				require.Equal(t, utils.ReplaceStringError(http.StatusBadRequest, err.Error()), "Key: 'UpdateAddressesRequest.City' Error:Field validation for 'City' failed on the 'required' tag")
			},
		},
		{
			name: "Bad Request - province is required",
			body: `{"address":"address","city":"city","postalCode":"postalCode","country":"country"}`,
			buildStubs: func(store *mockup.MockStore, body request.UpdateAddressesRequest) {
			},
			checkResponse: func(t *testing.T, status int, err error) {
				require.Error(t, err)
				require.Equal(t, utils.ReplaceStringError(http.StatusBadRequest, err.Error()), "Key: 'UpdateAddressesRequest.Province' Error:Field validation for 'Province' failed on the 'required' tag")
			},
		},
		{
			name: "Bad Request - postalCode is required",
			body: `{"address":"address","city":"city","province":"province","country":"country"}`,
			buildStubs: func(store *mockup.MockStore, body request.UpdateAddressesRequest) {
			},
			checkResponse: func(t *testing.T, status int, err error) {
				require.Error(t, err)
				require.Equal(t, utils.ReplaceStringError(http.StatusBadRequest, err.Error()), "Key: 'UpdateAddressesRequest.PostalCode' Error:Field validation for 'PostalCode' failed on the 'required' tag")
			},
		},
		{
			name: "Bad Request - country is required",
			body: `{"address":"address","city":"city","province":"province","postalCode":"postalCode"}`,
			buildStubs: func(store *mockup.MockStore, body request.UpdateAddressesRequest) {
			},
			checkResponse: func(t *testing.T, status int, err error) {
				require.Error(t, err)
				require.Equal(t, utils.ReplaceStringError(http.StatusBadRequest, err.Error()), "Key: 'UpdateAddressesRequest.Country' Error:Field validation for 'Country' failed on the 'required' tag")
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			cfg := config.InitConfigs(".env")
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			var dto request.UpdateAddressesRequest
			err := json.Unmarshal([]byte(tc.body), &dto)
			require.NoError(t, err)

			dbFactory := mockup.NewMockStore(ctrl)
			distributor := mockup.NewMockTaskDistributor(ctrl)
			tc.buildStubs(dbFactory, dto)

			app := echo.New()
			app.Validator = middleware.NewCustomValidator()
			app.Use(middleware.ErrorHandler)

			req := httptest.NewRequest(http.MethodPut, "/api/v1/profile/addresses", strings.NewReader(tc.body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := app.NewContext(req, rec)
			c.Set("accountsID", uid)
			c.SetPath("/:id")
			c.SetParamNames("id")
			c.SetParamValues(addressesID)
			handler := handlers.NewServerHttpHandler(app, &cfg, distributor, dbFactory)

			err = handler.UpdateAddresses(c)
			tc.checkResponse(t, c.Response().Status, err)
		})
	}
}

func TestDeleteAddresses(t *testing.T) {
	addressesID := "e2109e75-1d9d-48fb-9e68-310d4720b015"
	testCases := []struct {
		name          string
		buildStubs    func(store *mockup.MockStore)
		checkResponse func(t *testing.T, status int, err error)
	}{
		{
			name: "OK",
			buildStubs: func(store *mockup.MockStore) {
				store.EXPECT().IsAccountAlreadyExists(gomock.Any(), uid).Times(1).Return(true, nil)

				store.EXPECT().IsAddressesAlreadyExists(gomock.Any(), &database.IsAddressesAlreadyExistsParams{
					ID:         addressesID,
					AccountsID: uid,
				}).Times(1).Return(true, nil)
				store.EXPECT().DeleteAddressesById(gomock.Any(), addressesID).Times(1).Return(nil)
			},
			checkResponse: func(t *testing.T, status int, err error) {
				require.NoError(t, err)
				require.Equal(t, http.StatusOK, status)
			},
		},
		{
			name: "Not found - address id not found",
			buildStubs: func(store *mockup.MockStore) {
				store.EXPECT().IsAccountAlreadyExists(gomock.Any(), uid).Times(1).Return(true, nil)

				store.EXPECT().IsAddressesAlreadyExists(gomock.Any(), &database.IsAddressesAlreadyExistsParams{
					ID:         addressesID,
					AccountsID: uid,
				}).Times(1).Return(false, nil)
			},
			checkResponse: func(t *testing.T, status int, err error) {
				require.Error(t, err)
				require.Equal(t, utils.ReplaceStringError(http.StatusNotFound, err.Error()), common.ErrDataNotFound.Error())
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			cfg := config.InitConfigs(".env")
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			dbFactory := mockup.NewMockStore(ctrl)
			distributor := mockup.NewMockTaskDistributor(ctrl)
			tc.buildStubs(dbFactory)

			app := echo.New()
			app.Validator = middleware.NewCustomValidator()
			app.Use(middleware.ErrorHandler)

			req := httptest.NewRequest(http.MethodDelete, "/api/v1/profile/addresses", nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := app.NewContext(req, rec)
			c.Set("accountsID", uid)
			c.SetPath("/:id")
			c.SetParamNames("id")
			c.SetParamValues(addressesID)
			handler := handlers.NewServerHttpHandler(app, &cfg, distributor, dbFactory)

			err := handler.DeleteAddresses(c)
			tc.checkResponse(t, c.Response().Status, err)
		})
	}
}
