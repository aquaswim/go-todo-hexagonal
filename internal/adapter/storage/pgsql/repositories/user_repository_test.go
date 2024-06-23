package repositories

import (
	"context"
	"hexagonal-todo/internal/adapter/config"
	"hexagonal-todo/internal/adapter/storage/pgsql"
	"hexagonal-todo/internal/core/domain"
	"testing"
)

func newUserRepo(tb testing.TB) *userRepository {
	//todo: use https://github.com/pashagolub/pgxmock
	tb.Log("connecting to database")
	dbPool, err := pgsql.Connect(config.DBConfigFromENV())
	if err != nil {
		tb.Fatalf("fail to connect to db: %s", err)
	}
	return &userRepository{
		db: dbPool,
	}
}

func TestUserRepository_Positive(t *testing.T) {
	repo := newUserRepo(t)
	ctx := context.Background()
	var (
		email = "test@test.com"
		id    = int64(-1)
	)

	t.Cleanup(func() {
		if id != -1 {
			// remove created data
			t.Logf("removing created user with id: %d", id)
			_, err := repo.db.Exec(ctx, "DELETE FROM users WHERE id = $1", id)
			if err != nil {
				t.Fatalf("fail to remove created user: %s", err)
			}
		}
	})

	t.Run("create", func(t *testing.T) {
		user, err := repo.CreateUser(ctx, &domain.UserData{
			LoginCredential: domain.LoginCredential{
				Email:    email,
				Password: "password",
			},
			FullName: "dummy 1",
		})
		if err != nil {
			t.Fatalf("fail to create user: %s", err)
		}
		t.Logf("new user id: %d", user.Id)
		id = user.Id
	})

	t.Run("getById", func(t *testing.T) {
		_, err := repo.GetUserById(ctx, id)
		if err != nil {
			return
		}
		if err != nil {
			t.Fatalf("fail to getById: %s", err)
		}
	})

	t.Run("getByEmail", func(t *testing.T) {
		_, err := repo.GetUserByEmail(ctx, email)
		if err != nil {
			return
		}
		if err != nil {
			t.Fatalf("fail to getByEmail: %s", err)
		}
	})
}
