package pgsql

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
	"github.com/z0ne-dev/mgx/v2"
	"hexagonal-todo/internal/adapter/config"
	"hexagonal-todo/internal/adapter/storage/pgsql/migrations"
)

func Connect(cfg *config.DBConfig) (*pgxpool.Pool, error) {
	dbPool, err := pgxpool.New(context.Background(), cfg.Connection)
	if err != nil {
		log.Error().Err(err).Msg("failed to connect to database")
		return nil, err
	}
	err = dbPool.Ping(context.Background())
	if err != nil {
		return nil, err
	}
	log.Info().Msg("successfully connected to database")

	if cfg.MigrationOnBoot {
		if err = execMigration(dbPool); err != nil {
			log.Error().Err(err).Msg("failed to apply migrations")
			return nil, err
		}
	}

	return dbPool, nil
}

func execMigration(dbPool *pgxpool.Pool) error {
	migrator, err := mgx.New(mgx.Migrations(
		migrations.Migration202406211108CreateTodoTable,
		migrations.Migration202406231150CreateUserTable,
	), mgx.Log(migrationLogger{}))
	if err != nil {
		return err
	}

	// output pending migration
	pendings, err := migrator.Pending(context.Background(), dbPool)
	if err != nil {
		return err
	}
	if len(pendings) == 0 {
		log.Info().Msg("no pending migrations")
		return nil
	}
	return migrator.Migrate(context.Background(), dbPool)
}
