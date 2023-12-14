package db

import (
	"context"
)

type Querier interface {
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	GetUser(ctx context.Context, username string) (User, error)
	ListUsers(ctx context.Context, limit int) ([]User, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)
	DeleteUser(ctx context.Context, username string) error
}

var _ Querier = (*Queries)(nil)
