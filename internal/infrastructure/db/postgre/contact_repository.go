package postgre

import (
	"context"
	"fmt"

	"github.com/dh-atha/EmployeeAbsenceKNTest/internal/domain/model"
	"github.com/dh-atha/EmployeeAbsenceKNTest/internal/domain/repository"
	"github.com/jmoiron/sqlx"
)

type ContactRepositoryImpl struct {
	db *sqlx.DB
}

func NewContactRepository(db *sqlx.DB) repository.ContactRepository {
	return &ContactRepositoryImpl{db: db}
}

func (r *ContactRepositoryImpl) Update(ctx context.Context, emp *model.Contact) error {
	query := `
		UPDATE contacts SET
			membership_id = COALESCE(NULLIF(:membership_id, 0), membership_id),
			contact_type = COALESCE(:contact_type, contact_type),
			contact_value = COALESCE(:contact_value, contact_value),
			is_active = COALESCE(:is_active, is_active),
			updated_date = :updated_date,
			updated_by = :updated_by
		WHERE contact_id = :contact_id
	`

	res, err := r.db.NamedExecContext(ctx, query, emp)
	if err != nil {
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affected < 1 {
		return fmt.Errorf("no rows updated")
	}

	return nil
}

func (r *ContactRepositoryImpl) Create(ctx context.Context, emp *model.Contact) (*model.Contact, error) {
	query := `
		INSERT INTO contacts (membership_id, contact_type, contact_value, is_active, created_date, created_by) VALUES ($1, $2, $3, $4, $5, $6) RETURNING contact_id
	`

	var id int
	err := r.db.QueryRowContext(ctx, query,
		emp.MembershipID,
		emp.ContactType,
		emp.ContactValue,
		emp.IsActive,
		emp.CreatedDate,
		emp.CreatedBy,
	).Scan(&id)
	if err != nil {
		return nil, err
	}
	emp.ContactID = id

	return emp, nil
}
