package model

import "time"

type Contact struct {
	ContactID    int        `db:"contact_id" json:"contact_id"`
	MembershipID *int       `db:"membership_id" json:"membership_id"`
	ContactType  *string    `db:"contact_type" json:"contact_type"`
	ContactValue *string    `db:"contact_value" json:"contact_value"`
	IsActive     *bool      `db:"is_active" json:"is_active"`
	CreatedDate  *time.Time `db:"created_date" json:"created_date"`
	CreatedBy    *string    `db:"created_by" json:"created_by"`
	UpdatedDate  *time.Time `db:"updated_date" json:"updated_date"`
	UpdatedBy    *string    `db:"updated_by" json:"updated_by"`
}
