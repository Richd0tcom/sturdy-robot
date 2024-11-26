package service

import (
	"context"
	"database/sql"
	"log"
	"testing"

	"github.com/Richd0tcom/sturdy-robot/internal/config"
	db "github.com/Richd0tcom/sturdy-robot/internal/db/sqlc"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
)


var testQueries *db.Queries

var testDB *sql.DB

func init(){
	config, err:= config.LoadConfig("../..")
	// /Users/richdotcom/go/src/github.com/Richd0tcom/SafeX-Pay/.env
	// /Users/richdotcom/go/src/github.com/Richd0tcom/SafeX-Pay/db/sqlc/main_test.go
	if err != nil {
		log.Fatal("could not read configs", err)
	}
	testDB, err = sql.Open(config.DbDriver, config.DbUri)
	if err != nil {
		log.Fatal("cannot connect to database", err)
	}

	testQueries = db.New(testDB)

	
}

func TestCreateInvoice(t *testing.T) {

	//create organization
	org, err:= testQueries.CreateOrganization(context.Background(), db.CreateOrganizationParams{
		Name: "FC",
		Email: "bankersurate@gmail.com",
	})

	require.NoError(t,err)
	require.NotEmpty(t, org)

	branch, err := testQueries.CreateBranch(context.Background(), db.CreateBranchParams{
		OrganizationID: org.ID,
		Address: sql.NullString{
			String: "somewhere",
		},
		Name: "Branch1",
	})

	require.NoError(t,err)
	require.NotEmpty(t, branch)

	user, err:= testQueries.CreateUser(context.Background(), db.CreateUserParams{
		Name: "Rich",
		BranchID: branch.ID,
		Email: "bankersurate@gmail.com",
	})

	require.NoError(t,err)
	require.NotEmpty(t, user)

	
}