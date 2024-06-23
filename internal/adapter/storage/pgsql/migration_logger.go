package pgsql

import "github.com/rs/zerolog/log"

type migrationLogger struct{}

func (migrationLogger) Log(msg string, data map[string]any) {
	l := log.Info()
	for key, val := range data {
		l = l.Any(key, val)
	}
	l.Msg(msg)
}
