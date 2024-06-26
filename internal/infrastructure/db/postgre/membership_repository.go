package postgre

import (
	"context"

	"github.com/dh-atha/EmployeeAbsenceKNTest/internal/domain/model"
	"github.com/dh-atha/EmployeeAbsenceKNTest/internal/domain/repository"
	"github.com/jmoiron/sqlx"
)

type MembershipRepositoryImpl struct {
	db *sqlx.DB
}

func NewMembershipRepository(db *sqlx.DB) repository.MembershipRepository {
	return &MembershipRepositoryImpl{db: db}
}

func (r *MembershipRepositoryImpl) GetAll(ctx context.Context) ([]*model.MembershipContact, error) {
	var memberships []*model.MembershipContact
	query := `
		SELECT m.membership_id, name, address, contact_type, contact_value FROM memberships m 
		LEFT JOIN contacts c ON m.membership_id = c.membership_id WHERE m.is_active = true
	`

	err := r.db.SelectContext(ctx, &memberships, query)
	if err != nil {
		return nil, err
	}

	return memberships, nil
}

func (r *MembershipRepositoryImpl) Create(ctx context.Context, emp *model.Membership) (*model.Membership, error) {
	query := `
		INSERT INTO memberships (name, password, address, is_active, created_date, created_by) 
		VALUES ($1, $2, $3, $4, $5, $6) RETURNING membership_id
	`

	var id int
	err := r.db.QueryRowContext(ctx, query,
		emp.Name,
		emp.Password,
		emp.Address,
		emp.IsActive,
		emp.CreatedDate,
		emp.CreatedBy,
	).Scan(&id)
	if err != nil {
		return nil, err
	}
	emp.MembershipID = id

	return emp, nil
}

func (r *MembershipRepositoryImpl) GetMembership(ctx context.Context, username string) (*model.Membership, error) {
	query := `
		SELECT m.* FROM memberships m
		INNER JOIN contacts c ON m.membership_id = c.membership_id WHERE c.contact_value = $1
	`

	var membership model.Membership

	err := r.db.GetContext(ctx, &membership, query, username)
	if err != nil {
		return nil, err
	}

	return &membership, nil
}
