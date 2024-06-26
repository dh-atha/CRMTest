package service

import (
	"context"
	"time"

	"github.com/dh-atha/EmployeeAbsenceKNTest/internal/domain/constants"
	"github.com/dh-atha/EmployeeAbsenceKNTest/internal/domain/model"
	"github.com/dh-atha/EmployeeAbsenceKNTest/internal/domain/repository"
	"github.com/dh-atha/EmployeeAbsenceKNTest/internal/infrastructure/encrypter"
)

type MembershipServiceInterface interface {
	Create(ctx context.Context, emp *model.Membership) (*model.Membership, error)
	GetAll(ctx context.Context) ([]*model.MembershipContact, error)
}

type MembershipServiceImpl struct {
	repo repository.MembershipRepository
}

func NewMembershipService(repo repository.MembershipRepository) MembershipServiceInterface {
	return &MembershipServiceImpl{repo: repo}
}

func (s *MembershipServiceImpl) Create(ctx context.Context, emp *model.Membership) (*model.Membership, error) {
	var (
		name, _ = ctx.Value(constants.MembershipNameJWTKey).(string)
	)

	*emp.Password = encrypter.GetMD5Hash(*emp.Password)
	now := time.Now().UTC()
	emp.CreatedDate = &now
	emp.CreatedBy = &name

	return s.repo.Create(ctx, emp)
}

func (s *MembershipServiceImpl) GetAll(ctx context.Context) ([]*model.MembershipContact, error) {
	return s.repo.GetAll(ctx)
}
