// /internal/core/ports/user.go

package ports

import (
	"context"

	"github.com/RiteshDevOpsEngineer/ecom/internal/core/domain"
)

// UserService defines the interface for user-related operations
type UserService interface {
	GetUserByID(id int) (domain.User, error)
	FindByID(id uint) (*domain.User, error)
	Logout(ctx context.Context, phone string) error
}
