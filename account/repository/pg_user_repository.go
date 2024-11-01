package repository

import (
	"context"
	"log"
	"memrizr/account/model"
	"memrizr/account/model/apperrors"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type pGUserRepository struct {
	DB *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) model.UserRepository {
	return &pGUserRepository{
		DB: db,
	}
}

// reaches out to database SQLX api
func (r *pGUserRepository) Create(ctx context.Context, u *model.User) error {
	query := "INSERT INTO users (email, password) VALUES ($1, $2) RETURNING uid, email, password"

	row := r.DB.QueryRowContext(ctx, query, u.Email, u.Password)

	// if err := r.DB.GetContext(ctx, query, u.Email, u.Password); err != nil {
	if err := row.Scan(&u.UID, &u.Email, &u.Password); err != nil {
		log.Printf("Error inserting user: %v, query: %v, ctx: %v", err, query, ctx)

		// check unique constraint
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code.Name() == "unique_violation" {
			log.Printf("Unique constraint violation for email: %v. Reason: %v\n", u.Email, pqErr.Code.Name())
			return apperrors.NewConflict("email", u.Email)
		}

		log.Printf("Database error for email: %v, Reason: %v", u.Email, err)
		return apperrors.NewInternal()
	}
	return nil
}

// FindByID fetches user by id
func (r *pGUserRepository) FindByID(ctx context.Context, uid uuid.UUID) (*model.User, error) {
	user := &model.User{}

	query := "SELECT * FROM users WHERE uid = $1"

	if err := r.DB.GetContext(ctx, user, query, uid); err != nil {
		return user, apperrors.NewNotFound("uid", uid.String())
	}

	return user, nil
}

// FindByEmail retrieves user row by email address
func (r *pGUserRepository) FindByEmail(ctx context.Context, email string) (*model.User, error) {
    user := &model.User{}

    query := "SELECT * FROM users WHERE email=$1"

    if err := r.DB.GetContext(ctx, user, query, email); err != nil {
        log.Printf("Unable to get user with email address: %v. Err: %v\n", email, err)
        return user, apperrors.NewNotFound("email", email)
    }

    return user, nil
}

// Updates a user's properties
func (r *pGUserRepository) Update(ctx context.Context, u *model.User) error {
	query := `
		UPDATE users
		SET name=:name, email=:email, website=:website
		WHERE uid=:uid
		RETURNING *;
	`

	nameStatement, err := r.DB.PrepareNamedContext(ctx, query)

	if err != nil {
        log.Printf("Unable to prepare user update query: %v\n", err)
        return apperrors.NewInternal()
	}

	if err := nameStatement.GetContext(ctx, u, u); err != nil {
        log.Printf("Failed to update details for user: %v\n", u)
        return apperrors.NewInternal()
	}

	return nil
}