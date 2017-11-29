package store

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/apex/log"
	"github.com/blockloop/boar-example/models"
	"github.com/blockloop/scan"

	sq "github.com/Masterminds/squirrel"
	"github.com/pborman/uuid"
)

func NewUsers(db *sql.DB, log log.Interface) Users {
	return &users{
		db:  db,
		log: log,
	}
}

type users struct {
	db  *sql.DB
	log log.Interface
}

func (u *users) LookupByEmail(ctx context.Context, email string) (*models.User, error) {
	rows, err := sq.Select("*").
		From("users").
		Where(sq.Eq{"email": email}).
		RunWith(u.db).
		QueryContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not execute query: %v", err)
	}

	var user models.User
	if err := scan.Row(&user, rows); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("could not scan row: %v", err)
	}

	return &user, nil
}

func (u *users) LookupByID(ctx context.Context, id int) (*models.User, error) {
	defer func(start time.Time) {
		log.WithFields(log.Fields{
			"query.name":     "find_user_by_id",
			"query.duration": time.Since(start).Nanoseconds(),
		}).Info("query duration")
	}(time.Now())

	rows, err := sq.Select("*").
		From("users").
		Where("id = $1", id).
		RunWith(u.db).
		QueryContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not select from db: %v", err)
	}

	var user models.User
	err = scan.Row(&user, rows)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("could not scan row: %+v", err)
	}

	return &user, nil
}

func (u *users) LookupByEmailAndPassword(ctx context.Context, email string, passwordHash string) (*models.User, error) {
	rows, err := sq.Select("*").
		From("users").
		Where(sq.Eq{"email": email, "password_hash": passwordHash}).
		RunWith(u.db).
		QueryContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not execute query: %v", err)
	}

	var user models.User
	if err := scan.Row(&user, rows); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("could not scan row: %v", err)
	}

	return &user, nil
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
