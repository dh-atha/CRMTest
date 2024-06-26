package service

import (
	"context"
	"time"

	"github.com/dh-atha/EmployeeAbsenceKNTest/internal/domain/constants"
	"github.com/dh-atha/EmployeeAbsenceKNTest/internal/domain/model"
	"github.com/dh-atha/EmployeeAbsenceKNTest/internal/domain/repository"
)

type ContactServiceInterface interface {
	Create(ctx context.Context, emp *model.Contact) (*model.Contact, error)
	Update(ctx context.Context, emp *model.Contact) error
}

type ContactServiceImpl struct {
	repo repository.ContactRepository
}

func NewContactService(repo repository.ContactRepository) ContactServiceInterface {
	return &ContactServiceImpl{repo: repo}
}

func (s *ContactServiceImpl) Create(ctx context.Context, emp *model.Contact) (*model.Contact, error) {
	var (
		name, _ = ctx.Value(constants.MembershipNameJWTKey).(string)
	)

	now := time.Now().UTC()
	emp.CreatedDate = &now
	emp.CreatedBy = &name

	return s.repo.Create(ctx, emp)
}

func (s *ContactServiceImpl) Update(ctx context.Context, emp *model.Contact) error {
	var (
		name, _ = ctx.Value(constants.MembershipNameJWTKey).(string)
	)
	now := time.Now().UTC()
	emp.UpdatedBy = &name
	emp.UpdatedDate = &now

	return s.repo.Update(ctx, emp)
}
