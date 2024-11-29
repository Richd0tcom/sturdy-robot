package requests

import db "github.com/Richd0tcom/sturdy-robot/internal/db/sqlc"

type CreateOrgWithUser struct {
	OrganaizationName string `json:"org_name"`
	OrganizationEmail string `json:"org_email"`
	OrganizationAddress string `json:"org_address"`

	BranchName string `json:"branch_name"`
	BranchAddress string `json:"branch_address"`

	Username string `json:"username"`
	UserEmail string `json:"user_email"`
}

type UserWithTokenRes struct {
	User db.User `json:"user"`

	AccessToken string `json:"access_token"`
}

type LoginReq struct {
	Email string `json:"email"`
}