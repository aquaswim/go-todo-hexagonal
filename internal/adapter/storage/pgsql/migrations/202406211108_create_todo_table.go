package migrations

import (
	"github.com/z0ne-dev/mgx/v2"
)

var Migration202406211108CreateTodoTable = mgx.NewRawMigration("202406211108_create_todo_table", `
CREATE TABLE IF NOT EXISTS todos (
	id SERIAL PRIMARY KEY,
	title VARCHAR(255) NOT NULL,
	description TEXT,
	created_at TIMESTAMP WITH TIME ZONE,
	updated_at TIMESTAMP WITH TIME ZONE
)
`)
