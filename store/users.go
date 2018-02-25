package store

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/blockloop/beer_ratings/models"

	sq "github.com/Masterminds/squirrel"
	"github.com/blockloop/scan"
	"github.com/pborman/uuid"
	"github.com/pkg/errors"
)

var (
	selectUserCols = scan.Columns(new(models.User))
	selectUsers    = sq.Select(selectUserCols...).From("users")

	insertUserCols = scan.Columns(new(models.User), "id")
	insertUsers    = sq.Insert("users").Columns(insertUserCols...)
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
	rows, err := selectUsers.
		Where(sq.Eq{"email": email}).
		RunWith(u.db).
		QueryContext(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to execute query")
	}

	var user models.User
	if err := scan.Row(&user, rows); err != nil {
		return nil, errors.Wrap(err, "failed to scan user")
	}

	return &user, nil
}

func (u *users) LookupByID(ctx context.Context, id int64) (*models.User, error) {
	rows, err := selectUsers.
		Where(sq.Eq{"id": id}).
		Limit(1).
		RunWith(u.db).
		QueryContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not select from db: %v", err)
	}
	var user models.User
	if err := scan.Row(&user, rows); err != nil {
		return nil, errors.Wrap(err, "failed to scan user")
	}

	return &user, nil
}

func (u *users) LookupByEmailAndPassword(ctx context.Context, email string, passwordHash string) (*models.User, error) {
	rows, err := selectUsers.
		Where(sq.Eq{
			"email":         email,
			"password_hash": passwordHash,
		}).
		Limit(1).
		RunWith(u.db).
		QueryContext(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to execute query")
	}

	var user models.User
	if err := scan.Row(&user, rows); err != nil {
		return nil, errors.Wrap(err, "failed to scan user")
	}

	return &user, nil
}

func (u *users) Create(ctx context.Context, email string, passwordHash string) (user *models.User, err error) {
	now := time.Now()

	user = &models.User{
		UUID:     uuid.New(),
		Username: email,
		Email:    email,
		Created:  now,
		Modified: now,
	}

	vals := scan.Values(insertUserCols, user)
	res, err := insertUsers.
		Values(vals...).
		RunWith(u.db).
		ExecContext(ctx)
	if err != nil {
		if ok, uerr := isUniqueConstraint(err); ok {
			return nil, uerr
		}
		return nil, fmt.Errorf("could not insert user: %v", err)
	}

	user.ID, err = res.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("could not determine last inserted id: %v", err)
	}

	return user, nil
}
