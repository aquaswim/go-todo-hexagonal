package pgsql

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/z0ne-dev/mgx/v2"
	"hexagonal-todo/internal/adapter/config"
	"hexagonal-todo/internal/adapter/storage/pgsql/migrations"
	"log"
)

func Connect(cfg *config.Config) (*pgxpool.Pool, error) {
	dbPool, err := pgxpool.New(context.Background(), cfg.DBConnection)
	if err != nil {
		return nil, err
	}
	err = dbPool.Ping(context.Background())
	if err != nil {
		return nil, err
	}

	if cfg.DBMigrationOnBoot {
		if err = execMigration(dbPool); err != nil {
			return nil, err
		}
	}

	return dbPool, nil
}

func execMigration(dbPool *pgxpool.Pool) error {
	migrator, err := mgx.New(mgx.Migrations(
		migrations.Migration202406211108CreateTodoTable,
	))
	if err != nil {
		return err
	}

	// output pending migration
	pendings, err := migrator.Pending(context.Background(), dbPool)
	if err != nil {
		return err
	}
	if len(pendings) == 0 {
		log.Println("No pending migrations")
	}
	return migrator.Migrate(context.Background(), dbPool)
}
