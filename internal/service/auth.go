package service

import (
	"context"

	db "github.com/Richd0tcom/sturdy-robot/internal/db/sqlc"
	"github.com/Richd0tcom/sturdy-robot/internal/handlers/requests"
	"github.com/Richd0tcom/sturdy-robot/internal/utils"
	"github.com/jackc/pgx/v5/pgtype"
)

func RegisterUser(ctx context.Context, args requests.CreateOrgWithUser, st db.Store) (requests.UserWithTokenRes, error) {
	
	org, err:= st.CreateOrganization(ctx, db.CreateOrganizationParams{
		Name: args.OrganaizationName,
		Email: args.OrganizationEmail,
	})

	if err != nil {
		return requests.UserWithTokenRes{}, err
	}

	branch, err:= st.CreateBranch(ctx, db.CreateBranchParams{
		OrganizationID: org.ID,
		Name: args.BranchName,
		Address: utils.StringToPGText(args.BranchAddress),
	})

	if err != nil {
		return requests.UserWithTokenRes{}, err
	}

	user, err:= st.CreateUser(ctx, db.CreateUserParams{
		BranchID: branch.ID,
		Name: args.Username,
		Email: args.UserEmail,
	})

	if err != nil {
		return requests.UserWithTokenRes{}, err
	}

	u:= make(map[string]pgtype.UUID)
	u["user_id"] = user.ID
	u["branch_id"] = user.BranchID

	token, err:= utils.GenerateToken(u)

	if err != nil {
		return requests.UserWithTokenRes{}, err
	}

	return requests.UserWithTokenRes{
		User: user,
		AccessToken: token,
	}, nil

}

func GetUserSession(ctx context.Context, args requests.LoginReq, st db.Store) (requests.UserWithTokenRes, error) {
	user,err := st.GetUserByEmail(ctx, args.Email)

	if err != nil {
		return requests.UserWithTokenRes{}, err
	}

	u:= make(map[string]pgtype.UUID)
	u["user_id"] = user.ID
	u["branch_id"] = user.BranchID

	token, err:= utils.GenerateToken(u)

	if err != nil {
		return requests.UserWithTokenRes{}, err
	}

	return requests.UserWithTokenRes{
		User: user,
		AccessToken: token,
	}, nil
}