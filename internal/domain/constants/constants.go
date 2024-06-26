package constants

import "errors"

var (
	MembershipIDJWTKey   = "membership_id"
	MembershipNameJWTKey = "membership_name"
)

var (
	ErrService = errors.New("error service: ")
)
