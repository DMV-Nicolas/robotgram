package db

import "context"

type Querier interface {
	CreateUser(ctx context.Context, arg CreateUserParams) (any, error)
	GetUser(ctx context.Context, arg GetUserParams) (User, error)
}

var _ Querier = (*Queries)(nil)
