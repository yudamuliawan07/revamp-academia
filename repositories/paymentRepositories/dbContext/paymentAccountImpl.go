package dbContext

import (
	"context"
	"net/http"

	"codeid.revampacademy/models"
)

type UserAccount struct {
	UserName      string  `json:"user_name"`
	AccountNumber string  `json:"account_number"`
	Description   string  `json:"description"`
	Saldo         float64 `json:"saldo"`
	Type          string  `json:"type"`
}

const listPaymentUsers_accountByUserName = `-- name: ListPaymentUsers_accountByUserName :many
SELECT 
		    u.user_name,
		    ua.usac_account_number, 
		    COALESCE(b.bank_code, f.fint_code) AS description,
		    ua.usac_saldo,
		    ua.usac_type
		FROM 
		    payment.users_account ua
		LEFT JOIN 
		    payment.bank b ON ua.usac_bank_entity_id = b.bank_entity_id
		LEFT JOIN 
		    payment.fintech f ON ua.usac_bank_entity_id = f.fint_entity_id
		LEFT JOIN
		    users.users u ON ua.usac_user_entity_id = u.user_entity_id
		WHERE
		    u.user_name = $1;
`

func (q *Queries) ListPaymentUsers_accountByUserName(ctx context.Context, userName string) ([]UserAccount, error) {
	rows, err := q.db.QueryContext(ctx, listPaymentUsers_accountByUserName, userName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []UserAccount
	for rows.Next() {
		var i UserAccount
		if err := rows.Scan(
			&i.UserName,
			&i.AccountNumber,
			&i.Description,
			&i.Saldo,
			&i.Type,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPaymentUsers_account = `-- name: GetPaymentUsers_account :one
SELECT 
		    u.user_name,
		    ua.usac_account_number, 
		    COALESCE(b.bank_code, f.fint_code) AS description,
		    ua.usac_saldo,
		    ua.usac_type
		FROM 
		    payment.users_account ua
		LEFT JOIN 
		    payment.bank b ON ua.usac_bank_entity_id = b.bank_entity_id
		LEFT JOIN 
		    payment.fintech f ON ua.usac_bank_entity_id = f.fint_entity_id
		LEFT JOIN
		    users.users u ON ua.usac_user_entity_id = u.user_entity_id
		WHERE
			ua.usac_account_number = $1;
`

// payment.users_account
func (q *Queries) GetPaymentUsers_account(ctx context.Context, usacAccountNumber string) (UserAccount, error) {
	row := q.db.QueryRowContext(ctx, getPaymentUsers_account, usacAccountNumber)
	var i UserAccount
	err := row.Scan(
		&i.UserName,
		&i.AccountNumber,
		&i.Description,
		&i.Saldo,
		&i.Type,
	)
	return i, err
}

const createPaymentUsers_account = `-- name: CreatePaymentUsers_account :one
INSERT INTO 
	payment.users_account (
	usac_bank_entity_id,
	usac_user_entity_id,
	usac_account_number,
	usac_saldo,
	usac_type
	)
VALUES ($1, $2, $3, $4, $5)
RETURNING 
	(SELECT user_name FROM users.users WHERE user_entity_id = $2) AS user_name,
	usac_account_number, 
	COALESCE((SELECT bank_code FROM payment.bank WHERE bank_entity_id = usac_bank_entity_id), 
			 (SELECT fint_code FROM payment.fintech WHERE fint_entity_id = usac_bank_entity_id)) AS description,
	usac_saldo,
	usac_type;
`

type CreatePaymentUsers_accountParams struct {
	UsacBankEntityID  int32   `db:"usac_bank_entity_id" json:"usacBankEntityID"`
	UsacUserEntityID  int32   `db:"usac_user_entity_id" json:"usacUserEntityID"`
	UsacAccountNumber string  `db:"usac_account_number" json:"usacAccountNumber"`
	UsacSaldo         float64 `db:"usac_saldo" json:"usacSaldo"`
	UsacType          string  `db:"usac_type" json:"usacType"`
}

func (q *Queries) CreatePaymentUsers_account(ctx context.Context, arg CreatePaymentUsers_accountParams) (*UserAccount, *models.ResponseError) {
	row := q.db.QueryRowContext(ctx, createPaymentUsers_account,
		arg.UsacBankEntityID,
		arg.UsacUserEntityID,
		arg.UsacAccountNumber,
		arg.UsacSaldo,
		arg.UsacType,
	)

	i := UserAccount{}
	err := row.Scan(
		&i.UserName,
		&i.AccountNumber,
		&i.Description,
		&i.Saldo,
		&i.Type,
	)

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return &UserAccount{
		UserName:      i.UserName,
		AccountNumber: i.AccountNumber,
		Description:   i.Description,
		Saldo:         i.Saldo,
		Type:          i.Type,
	}, nil
}

const updatePaymentUsers_accountPlus = `-- name: UpdatePaymentUsers_accountPlus :exec
UPDATE payment.users_account
SET usac_saldo = usac_saldo + $1
WHERE usac_account_number = $2
RETURNING
	(SELECT user_name FROM users.users WHERE user_entity_id = payment.users_account.usac_user_entity_id) AS user_name,
	usac_account_number, 
	COALESCE((SELECT bank_code FROM payment.bank WHERE bank_entity_id = usac_bank_entity_id), 
			 (SELECT fint_code FROM payment.fintech WHERE fint_entity_id = usac_bank_entity_id)) AS description,
	usac_saldo,
	usac_type;
`

type UpdatePaymentUsers_accountParams struct {
	Amount float64 `json:"amount"`
}

func (q *Queries) UpdatePaymentUsers_accountPlus(ctx context.Context, arg UpdatePaymentUsers_accountParams, usacAccountNumber string) (*UserAccount, error) {
	row := q.db.QueryRowContext(ctx, updatePaymentUsers_accountPlus, arg.Amount, usacAccountNumber)
	var i UserAccount
	err := row.Scan(
		&i.UserName,
		&i.AccountNumber,
		&i.Description,
		&i.Saldo,
		&i.Type,
	)
	if err != nil {
		return nil, err
	}
	return &i, nil
}

const updatePaymentUsers_accountMinus = `-- name: UpdatePaymentUsers_accountMinus :exec
UPDATE payment.users_account
SET usac_saldo = usac_saldo - $1
WHERE usac_account_number = $2
RETURNING
	(SELECT user_name FROM users.users WHERE user_entity_id = payment.users_account.usac_user_entity_id) AS user_name,
	usac_account_number, 
	COALESCE((SELECT bank_code FROM payment.bank WHERE bank_entity_id = usac_bank_entity_id), 
			 (SELECT fint_code FROM payment.fintech WHERE fint_entity_id = usac_bank_entity_id)) AS description,
	usac_saldo,
	usac_type;
`

func (q *Queries) UpdatePaymentUsers_accountMinus(ctx context.Context, arg UpdatePaymentUsers_accountParams, usacAccountNumber string) (*UserAccount, error) {
	row := q.db.QueryRowContext(ctx, updatePaymentUsers_accountMinus, arg.Amount, usacAccountNumber)
	var i UserAccount
	err := row.Scan(
		&i.UserName,
		&i.AccountNumber,
		&i.Description,
		&i.Saldo,
		&i.Type,
	)
	if err != nil {
		return nil, err
	}
	return &i, nil
}

const deletePaymentUsers_account = `-- name: DeletePaymentUsers_account :exec
DELETE FROM 
	payment.users_account
WHERE 
	usac_account_number = $1;

`

func (q *Queries) DeletePaymentUsers_account(ctx context.Context, usacAccountNumber string) error {
	_, err := q.db.ExecContext(ctx, deletePaymentUsers_account, usacAccountNumber)
	return err
}
