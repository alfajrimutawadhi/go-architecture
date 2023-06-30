package repository

import (
	"context"
	"go-architecture/common"
	"go-architecture/domain"

	"github.com/google/uuid"
)

func (repo *Repository) CreateUser(ctx context.Context, request domain.User) (id uuid.UUID,err error) {
	q := `
		INSERT INTO users (id, name, email, password, created_at) VALUES (
			UUID_TO_BIN(:id),
			:name,
			:email,
			TO_BASE64(AES_ENCRYPT(:password, :db_aes_key)),
			:created_at
		)
	`

	if request.Id == uuid.Nil {
		request.Id = uuid.New()
	}
	request.CreatedAt = common.CurrentTime()
	request.DBAesKey = repo.Config.DB.DBAesKey
	_, err = repo.DB.NamedExecContext(ctx, q, request)
	if err != nil {
		return uuid.Nil, common.WrapError(err, "sqlx", "ExecContext")
	}

	return request.Id, nil
}

func (repo *Repository) GetUserById(ctx context.Context, request domain.User) (result domain.User, err error) {
	q := `
		SELECT
			BIN_TO_UUID(id) id,
			name,
			email,
			AES_DECRYPT(FROM_BASE64(password), :db_aes_key) password,
			created_at
		FROM users
		WHERE id = UUID_TO_BIN(:id)
	`
	request.DBAesKey = repo.Config.DB.DBAesKey
	rows, err := repo.DB.NamedQueryContext(ctx, q, request)
	if err != nil {
		return domain.User{}, common.WrapError(err, "sqlx", "NamedQueryContext")
	}
	rows.Next()

	if err = rows.StructScan(&result); err != nil {
		return domain.User{}, common.WrapError(err, "sqlx", "Scan")
	}
	return
}
