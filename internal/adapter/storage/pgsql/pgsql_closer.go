package pgsql

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
	"hexagonal-todo/internal/core/port"
)

type closer struct {
	conn *pgxpool.Pool
}

func (c closer) Close() error {
	log.Info().Msg("closing pgsql connection")
	c.conn.Close()
	log.Info().Msg("pgsql connection closed")
	return nil
}

func NewCloser(pgconn *pgxpool.Pool) port.Closable {
	return &closer{
		conn: pgconn,
	}
}
