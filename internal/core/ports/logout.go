package ports

import "context"

type LogoutService interface {
	Logout(ctx context.Context, phone string) error
}
