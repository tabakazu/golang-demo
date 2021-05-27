package datastore

import (
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/tabakazu/webapi-app/domain"
	"github.com/tabakazu/webapi-app/domain/model"
)

type userAccountRepository struct {
	db *sqlx.DB
}

func NewUserAccountRepository(db *sqlx.DB) domain.UserAccountRepository {
	return &userAccountRepository{db: db}
}

func (r *userAccountRepository) Create(e *model.UserAccountEntity) error {
	tx := r.db.MustBegin()
	result := tx.MustExec("INSERT users (family_name, given_name, created_at, updated_at) VALUES (?, ?, ?, ?)",
		e.FamilyName, e.GivenName, time.Now(), time.Now())

	if userID, err := result.LastInsertId(); err == nil {
		e.ID = uint(userID)
	} else {
		err = tx.Rollback()
		return err
	}

	tx.MustExec("INSERT accounts (user_id, email, password_digest, created_at, updated_at) VALUES (?, ?, ?, ?, ?)",
		e.ID, e.Email, e.PasswordDigest, time.Now(), time.Now())

	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

func (r *userAccountRepository) FindByID(e *model.UserAccountEntity, id uint) error {
	if err := r.db.Get(e, `SELECT u.id, u.family_name, u.given_name, a.email, a.password_digest
												 FROM users u
												 JOIN accounts a ON a.user_id = u.id
												 WHERE u.id = ?`, id); err != nil {
		return err
	}

	return nil
}

func (r *userAccountRepository) FindByEmail(e *model.UserAccountEntity, email string) error {
	if err := r.db.Get(e, `SELECT u.id, u.family_name, u.given_name, a.email, a.password_digest
												 FROM users u
												 JOIN accounts a ON a.user_id = u.id
												 WHERE a.email = ?`, email); err != nil {
		return err
	}

	return nil
}
