package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
	database "github.com/p-jirayusakul/go-clean-arch-template/database/sqlc"
	handlers "github.com/p-jirayusakul/go-clean-arch-template/internal/handlers/http"
	"github.com/p-jirayusakul/go-clean-arch-template/internal/handlers/http/request"
	"github.com/p-jirayusakul/go-clean-arch-template/pkg/common"
	"github.com/p-jirayusakul/go-clean-arch-template/pkg/config"
	"github.com/p-jirayusakul/go-clean-arch-template/pkg/middleware"
	"github.com/p-jirayusakul/go-clean-arch-template/pkg/utils"
	"github.com/p-jirayusakul/go-clean-arch-template/test/mockup"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

const uid = "a8b1e1f8-b757-4f17-a70a-7def0f2ffe9b"

const password = "123456"

func TestRegister(t *testing.T) {
	testCases := []struct {
		name          string
		body          string
		buildStubs    func(store *mockup.MockStore, body request.RegisterRequest)
		checkResponse func(t *testing.T, status int, err error)
	}{
		{
			name: "OK",
			body: `{"email":"test@email.com","password":"123456"}`,
			buildStubs: func(store *mockup.MockStore, body request.RegisterRequest) {
				store.EXPECT().IsEmailAlreadyExists(gomock.Any(), body.Email).Times(1).Return(false, nil)
				store.EXPECT().CreateAccount(gomock.Any(), gomock.Any()).Times(1).Return(uid, nil)
			},
			checkResponse: func(t *testing.T, status int, err error) {
				require.NoError(t, err)
				require.Equal(t, http.StatusCreated, status)
			},
		},
		{
			name: "Bad Request - this email is already used",
			body: `{"email":"test@email.com","password":"123456"}`,
			buildStubs: func(store *mockup.MockStore, body request.RegisterRequest) {
				store.EXPECT().IsEmailAlreadyExists(gomock.Any(), body.Email).Times(1).Return(true, nil)
			},
			checkResponse: func(t *testing.T, status int, err error) {
				require.Error(t, err)
				require.Equal(t, utils.ReplaceStringError(http.StatusBadRequest, err.Error()), common.ErrEmailIsAlreadyExists.Error())
			},
		},
		{
			name: "Bad Request - email invalid format",
			body: `{"email":"testemail.com","password":"123456"}`,
			buildStubs: func(store *mockup.MockStore, body request.RegisterRequest) {
			},
			checkResponse: func(t *testing.T, status int, err error) {
				require.Error(t, err)
				require.Equal(t, utils.ReplaceStringError(http.StatusBadRequest, err.Error()), "Key: 'RegisterRequest.Email' Error:Field validation for 'Email' failed on the 'email' tag")
			},
		},
		{
			name: "Internal Server Error",
			body: `{"email":"test@email.com","password":"123456"}`,
			buildStubs: func(store *mockup.MockStore, body request.RegisterRequest) {
				store.EXPECT().IsEmailAlreadyExists(gomock.Any(), body.Email).Times(1).Return(false, nil)
				store.EXPECT().CreateAccount(gomock.Any(), gomock.Any()).Times(1).Return("", pgx.ErrTxClosed)
			},
			checkResponse: func(t *testing.T, status int, err error) {
				require.Error(t, err)
				require.Equal(t, utils.ReplaceStringError(http.StatusInternalServerError, err.Error()), pgx.ErrTxClosed.Error())
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			cfg := config.InitConfigs(".env")
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			var dto request.RegisterRequest
			err := json.Unmarshal([]byte(tc.body), &dto)
			require.NoError(t, err)

			dbFactory := mockup.NewMockStore(ctrl)
			tc.buildStubs(dbFactory, dto)

			app := echo.New()
			app.Validator = middleware.NewCustomValidator()
			app.Use(middleware.ErrorHandler)

			req := httptest.NewRequest(http.MethodPost, "/api/v1/auth/register", strings.NewReader(tc.body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := app.NewContext(req, rec)
			handler := handlers.NewServerHttpHandler(app, &cfg, dbFactory)

			err = handler.Register(c)
			tc.checkResponse(t, c.Response().Status, err)
		})
	}
}

func TestLogin(t *testing.T) {
	hashedPassword, err := utils.HashPassword(password)
	require.NoError(t, err)
	testCases := []struct {
		name          string
		body          string
		buildStubs    func(store *mockup.MockStore, body request.LoginRequest)
		checkResponse func(t *testing.T, status int, err error)
	}{
		{
			name: "OK",
			body: `{"email":"test@email.com","password":"123456"}`,
			buildStubs: func(store *mockup.MockStore, body request.LoginRequest) {
				store.EXPECT().GetAccountByEmail(gomock.Any(), body.Email).Times(1).Return(&database.GetAccountByEmailRow{
					ID:       uid,
					Email:    body.Email,
					Password: hashedPassword,
				}, nil)
			},
			checkResponse: func(t *testing.T, status int, err error) {
				require.NoError(t, err)
				require.Equal(t, http.StatusOK, status)
			},
		},
		{
			name: "Unauthorized - username invalid",
			body: `{"email":"test9999@email.com","password":"123456"}`,
			buildStubs: func(store *mockup.MockStore, body request.LoginRequest) {
				store.EXPECT().GetAccountByEmail(gomock.Any(), body.Email).Times(1).Return(&database.GetAccountByEmailRow{}, pgx.ErrNoRows)
			},
			checkResponse: func(t *testing.T, status int, err error) {
				require.Error(t, err)
				require.Equal(t, utils.ReplaceStringError(http.StatusUnauthorized, err.Error()), common.ErrLoginFail.Error())
			},
		},
		{
			name: "Unauthorized - password invalid",
			body: `{"email":"test@email.com","password":"123456"}`,
			buildStubs: func(store *mockup.MockStore, body request.LoginRequest) {
				store.EXPECT().GetAccountByEmail(gomock.Any(), body.Email).Times(1).Return(&database.GetAccountByEmailRow{
					ID:       uid,
					Email:    body.Email,
					Password: "password invalid",
				}, nil)
			},
			checkResponse: func(t *testing.T, status int, err error) {
				require.Error(t, err)
				require.Equal(t, utils.ReplaceStringError(http.StatusUnauthorized, err.Error()), common.ErrLoginFail.Error())
			},
		},
		{
			name: "Bad Request - email invalid format",
			body: `{"email":"testemail.com","password":"123456"}`,
			buildStubs: func(store *mockup.MockStore, body request.LoginRequest) {
			},
			checkResponse: func(t *testing.T, status int, err error) {
				require.Error(t, err)
				require.Equal(t, utils.ReplaceStringError(http.StatusBadRequest, err.Error()), "Key: 'LoginRequest.Email' Error:Field validation for 'Email' failed on the 'email' tag")
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			cfg := config.InitConfigs(".env")
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			var dto request.LoginRequest
			err := json.Unmarshal([]byte(tc.body), &dto)
			require.NoError(t, err)

			dbFactory := mockup.NewMockStore(ctrl)
			tc.buildStubs(dbFactory, dto)

			app := echo.New()
			app.Validator = middleware.NewCustomValidator()
			app.Use(middleware.ErrorHandler)

			req := httptest.NewRequest(http.MethodPost, "/api/v1/auth/login", strings.NewReader(tc.body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := app.NewContext(req, rec)
			handler := handlers.NewServerHttpHandler(app, &cfg, dbFactory)

			err = handler.Login(c)
			tc.checkResponse(t, c.Response().Status, err)
		})
	}
}
