package model

import "time"

type Membership struct {
	MembershipID int        `db:"membership_id" json:"membership_id,omitempty"`
	Name         *string    `db:"name" json:"name"`
	Password     *string    `db:"password" json:"password"`
	Address      *string    `db:"address" json:"address"`
	IsActive     *bool      `db:"is_active" json:"is_active"`
	CreatedDate  *time.Time `db:"created_date" json:"created_date"`
	CreatedBy    *string    `db:"created_by" json:"created_by"`
	UpdatedDate  *time.Time `db:"updated_date" json:"updated_date"`
	UpdatedBy    *string    `db:"updated_by" json:"updated_by"`
}

type MembershipContact struct {
	MembershipID int     `db:"membership_id" json:"membership_id"`
	Name         *string `db:"name" json:"name"`
	Address      *string `db:"address" json:"address"`
	ContactType  *string `db:"contact_type" json:"contact_type"`
	ContactValue *string `db:"contact_value" json:"contact_value"`
}
