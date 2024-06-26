package repository

import (
	"context"

	"github.com/dh-atha/EmployeeAbsenceKNTest/internal/domain/model"
)

type MembershipRepository interface {
	Create(ctx context.Context, emp *model.Membership) (*model.Membership, error)
	GetAll(ctx context.Context) ([]*model.MembershipContact, error)
	GetMembership(ctx context.Context, username string) (*model.Membership, error)
}
