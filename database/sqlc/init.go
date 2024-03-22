package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/p-jirayusakul/go-clean-arch-template/pkg/config"
)

func InitDatabase(cfg config.Config) *Queries {

	// connect to database
	source := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=disable TimeZone=Asia/Bangkok", cfg.DATABASE_USER, cfg.DATABASE_PASSWORD, cfg.DATABASE_HOST, cfg.DATABASE_PORT, cfg.DATABASE_NAME)
	conn, err := pgxpool.New(context.Background(), source)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	db := New(conn)

	return db
}
