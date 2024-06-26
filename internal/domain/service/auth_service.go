package service

import (
	"context"
	"database/sql"
	"errors"

	"github.com/dh-atha/EmployeeAbsenceKNTest/internal/domain/constants"
	"github.com/dh-atha/EmployeeAbsenceKNTest/internal/domain/model"
	"github.com/dh-atha/EmployeeAbsenceKNTest/internal/domain/repository"
	"github.com/dh-atha/EmployeeAbsenceKNTest/internal/infrastructure/encrypter"
	"github.com/dh-atha/EmployeeAbsenceKNTest/internal/infrastructure/security"
)

type AuthServiceInterface interface {
	Login(ctx context.Context, employee *model.LoginRequest) (model.LoginResponse, error)
}

type AuthService struct {
	membershipRepo repository.MembershipRepository
	jwtService     security.JWTService
}

func NewAuthService(membershipRepo repository.MembershipRepository, jwtService security.JWTService) AuthServiceInterface {
	return &AuthService{
		membershipRepo: membershipRepo,
		jwtService:     jwtService,
	}
}

func (s *AuthService) Login(ctx context.Context, req *model.LoginRequest) (data model.LoginResponse, err error) {
	membership, err := s.membershipRepo.GetMembership(ctx, req.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			err = errors.Join(constants.ErrService, errors.New("no membership found"))
		}
		return data, err
	}

	if !encrypter.VerifyPassword(req.Password, *membership.Password) {
		return data, errors.Join(constants.ErrService, errors.New("invalid password"))
	}

	data.MembershipID, data.Name = membership.MembershipID, *membership.Name
	data.Token, data.ExpiredAt, err = s.jwtService.GenerateToken(*membership.Name, membership.MembershipID)
	return
}
