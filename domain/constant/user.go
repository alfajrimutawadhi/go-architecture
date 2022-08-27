package constant

const (
	INSERT_USER = `
		INSERT INTO users (id, name, email, password, created_at) VALUES (
			UUID_TO_BIN(:id),
			:name,
			:email,
			TO_BASE64(AES_ENCRYPT(:password, :db_aes_key)),
			:created_at
		)
	`

	SELECT_USER_BY_ID = `
		SELECT
			BIN_TO_UUID(id) id,
			name,
			email,
			AES_DECRYPT(FROM_BASE64(password), :db_aes_key) password,
			created_at
		FROM users
		WHERE id = UUID_TO_BIN(:id)
	`
)