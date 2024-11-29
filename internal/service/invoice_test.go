package service

import (
	"context"
	"log"
	"testing"

	"github.com/Richd0tcom/sturdy-robot/internal/config"
	db "github.com/Richd0tcom/sturdy-robot/internal/db/sqlc"
	_ "github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)


var testQueries *db.Queries


func init(){
	config, err:= config.LoadConfig("../../cmd/.env")

	if err != nil {
		log.Fatal("could not read configs", err)
	}
	// testDB, err = sql.Open(config.DbDriver, config.DbUri)
	connPool, err:= pgxpool.New(context.Background(), config.DbUri)
	if err != nil {
		log.Fatal("cannot connect to database", err)
	}

	testQueries = db.New(connPool)
}

func TestOtherFunctions(t *testing.T) {

	var c context.Context = context.Background()

	//create organization
	org, err:= testQueries.CreateOrganization(c, db.CreateOrganizationParams{
		Name: "FC",
		Email: "bankersurate@gmail.com",
	})

	require.NoError(t,err)
	require.NotEmpty(t, org)

	branch, err := testQueries.CreateBranch(c, db.CreateBranchParams{
		OrganizationID: org.ID,
		Address: pgtype.Text{
			String: "somewhere",
		},
		Name: "Branch1",
	})

	require.NoError(t,err)
	require.NotEmpty(t, branch)

	user, err:= testQueries.CreateUser(c, db.CreateUserParams{
		Name: "Rich",
		BranchID: branch.ID,
		Email: "bankersurate@gmail.com",
	})

	require.NoError(t,err)
	require.NotEmpty(t, user)

	payInfo,err:= testQueries.CreatePaymentInfo(c, db.CreatePaymentInfoParams{
		UserID: user.ID,
		AccountNo: "9182736450",
		AccountName: user.Name,
		BankName: "YouBank",
	})

	require.NoError(t,err)
	require.NotEmpty(t, payInfo)

	cat, err:= testQueries.CreateCategory(c, db.CreateCategoryParams{
		BranchID: branch.ID,
		Name: "Phones",
	})

	require.NoError(t,err)
	require.NotEmpty(t, cat)

	product,err:= testQueries.CreateProduct(c, db.CreateProductParams{
		CategoryID: cat.ID,
		BranchID: branch.ID,
		Name: "nokia",
		ProductType: "phyisical",
		Sku: "NK-457",
	})

	require.NoError(t,err)
	require.NotEmpty(t, product)

	product_version,err:= testQueries.CreateProductVersion(c, db.CreateProductVersionParams{
		ProductID: product.ID,
		BranchID: product.BranchID,
		Name: "S22",
	})

	require.NoError(t,err)
	require.NotEmpty(t, product_version)


	cost, _:= decimal.NewFromString("50.00")
	
	inventory, err:= testQueries.CreateInventoryRecord(c, db.CreateInventoryRecordParams{
		VersionID: product_version.ID,
		BranchID: product_version.BranchID,
		Quantity: pgtype.Int4{
			Int32: 16,
		},
		UnitCost: pgtype.Numeric{
			Int: cost.BigInt(),
			Exp: cost.Exponent(),
		},

	})

	require.NoError(t,err)
	require.NotEmpty(t, inventory)

	currency, err:= testQueries.CreateCurrency(c, db.CreateCurrencyParams{
		Name: "Dollar",
		Symbol: pgtype.Text{
			String: "$",
		},
		Code: "USD",
	})

	require.NoError(t,err)
	require.NotEmpty(t, currency)

	customer, err:= testQueries.CreateCustomer(c, db.CreateCustomerParams{
		Name: "V",
		Email: pgtype.Text{
			String: "t@gmail.com",

		},
		Phone: pgtype.Text{
			String: "09065900578",
		},
		BranchID: branch.ID,
	})

	require.NoError(t,err)
	require.NotEmpty(t, customer)
}