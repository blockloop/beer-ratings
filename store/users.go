package store

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/blockloop/boar-example/models"

	sq "github.com/Masterminds/squirrel"
	"github.com/pborman/uuid"
)

func NewUsers(db *sql.DB) Users {
	return &users{
		db: db,
	}
}

type users struct {
	db *sql.DB
}

func (u *users) LookupByEmail(ctx context.Context, email string) (*models.User, error) {
	panic("not implemented")
}

func (u *users) LookupByID(ctx context.Context, id int) (*models.User, error) {
	row := sq.Select("id,uuid,username,email,address_line_1,address_line_2,city,state,country,postal_code").
		From("users").
		Where("id = $1", id).
		RunWith(u.db).
		QueryRow()

	user := models.User{Address: &models.Address{}}
	if err := row.Scan(
		&user.ID,
		&user.UUID,
		&user.Username,
		&user.Email,
		&user.Address.Line1,
		&user.Address.Line2,
		&user.Address.City,
		&user.Address.State,
		&user.Address.Country,
		&user.Address.PostalCode,
	); err != nil {
		return nil, fmt.Errorf("could not scan row: %+v", err)
	}

	return &user, nil
}

func (u *users) LookupByEmailAndPassword(ctx context.Context, email string, passwordHash string) (*models.User, error) {
	panic("not implemented")
}

func (u *users) Create(ctx context.Context, email string, passwordHash string) (*models.User, error) {
	uid := uuid.New()
	now := time.Now()

	res, err := sq.Insert("users").
		Columns("email,username,password_hash,uuid,created,modified").
		Values(email, email, passwordHash, uid, now, now).
		RunWith(u.db).
		ExecContext(ctx)
	if err != nil {
		if ok, uerr := isUniqueConstraint(err); ok {
			return nil, uerr
		}
		return nil, fmt.Errorf("could not insert user: %v", err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("could not determine last inserted id: %v", err)
	}

	return &models.User{
		ID:       int(id),
		UUID:     uid,
		Username: email,
		Email:    email,
	}, nil
}
