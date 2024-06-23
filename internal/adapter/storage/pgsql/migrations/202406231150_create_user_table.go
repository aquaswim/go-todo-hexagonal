package migrations

import (
	"github.com/z0ne-dev/mgx/v2"
)

var Migration202406231150CreateUserTable = mgx.NewRawMigration("202406231150_create_user_table", `
CREATE TABLE IF NOT EXISTS users (
	id SERIAL PRIMARY KEY,
	email VARCHAR(255) UNIQUE NOT NULL,
	password VARCHAR(255) NOT NULL,
	full_name VARCHAR(255) NULL,
	created_at TIMESTAMP WITH TIME ZONE,
	updated_at TIMESTAMP WITH TIME ZONE
)
`)
