package repository

import (
	"context"

	"github.com/dh-atha/EmployeeAbsenceKNTest/internal/domain/model"
)

type ContactRepository interface {
	Create(ctx context.Context, emp *model.Contact) (*model.Contact, error)
	Update(ctx context.Context, emp *model.Contact) error
}
