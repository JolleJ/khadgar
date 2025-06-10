package user

import "context"

type UserRepository interface {
	List(ctx *context.Context) []User
}
