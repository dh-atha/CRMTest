package model

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token        string `json:"token"`
	ExpiredAt    string `json:"expired_at"`
	MembershipID int    `json:"membership_id"`
	Name         string `json:"name"`
}
