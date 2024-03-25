package factories

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	database "github.com/p-jirayusakul/go-clean-arch-template/database/sqlc"
	"github.com/p-jirayusakul/go-clean-arch-template/pkg/config"
)

// Store defines all functions to execute db queries and transactions
type DBFactory interface {
	database.Querier
}

// SQLStore provides all functions to execute SQL queries and transactions
type SQLStore struct {
	connPool *pgxpool.Pool
	*database.Queries
}

// NewDBFactory creates a new store
func NewDBFactory(connPool *pgxpool.Pool) DBFactory {
	return &SQLStore{
		connPool: connPool,
		Queries:  database.New(connPool),
	}
}

func InitDatabase(cfg config.Config) *pgxpool.Pool {

	// connect to database
	source := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=disable TimeZone=Asia/Bangkok", cfg.DATABASE_USER, cfg.DATABASE_PASSWORD, cfg.DATABASE_HOST, cfg.DATABASE_PORT, cfg.DATABASE_NAME)
	conn, err := pgxpool.New(context.Background(), source)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	return conn
}
