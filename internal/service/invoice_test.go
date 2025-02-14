package service

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/Richd0tcom/sturdy-robot/internal/config"
	db "github.com/Richd0tcom/sturdy-robot/internal/db/sqlc"
	_ "github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	// "golang.org/x/exp/rand"
	"math/rand/v2"

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


// SeedDatabase populates the database with test data
func SeedDatabase(ctx context.Context, q *db.Queries) error {
	// Seed random number generator
	// rand.New(time.Now().UnixNano())
	

	// Create 5 organizations
	orgs := make([]db.Organization, 0)
	for i := 1; i <= 5; i++ {
		org, err := q.CreateOrganization(ctx, db.CreateOrganizationParams{
			Name:  fmt.Sprintf("Organization %d", i),
			Email: fmt.Sprintf("org+%d@example.com", i),
		})
		if err != nil {
			return fmt.Errorf("error creating organization: %v", err)
		}
		orgs = append(orgs, org)
	}

	// Create 3-4 branches per organization
	branches := make([]db.Branch, 0)
	for _, org := range orgs {
		for j := 1; j <= rand.IntN(2)+3; j++ {
			branch, err := q.CreateBranch(ctx, db.CreateBranchParams{
				OrganizationID: org.ID,
				Name:          fmt.Sprintf("Branch %x-%d", org.ID.Bytes[:4], j),
				Address: pgtype.Text{
					String: fmt.Sprintf("Address %d-%d, City", org.ID.Bytes[:4], j),
					Valid:  true,
				},
			})
			if err != nil {
				return fmt.Errorf("error creating branch: %v", err)
			}
			branches = append(branches, branch)
		}
	}

	// Create 5 users per branch
	users := make([]db.User, 0)
	for _, branch := range branches {
		for k := 1; k <= 5; k++ {
			user, err := q.CreateUser(ctx, db.CreateUserParams{
				Name:     fmt.Sprintf("User %d-%d", branch.ID.Bytes[:4], k),
				Email:    fmt.Sprintf("user%d-%d@example.com", branch.ID.Bytes[:4], k),
				BranchID: branch.ID,
				Address: pgtype.Text{
					String: fmt.Sprintf("User Address %d-%d", branch.ID.Bytes[:4], k),
					Valid:  true,
				},
			})
			if err != nil {
				return fmt.Errorf("error creating user: %v", err)
			}
			users = append(users, user)

			// Create payment info for each user
			_, err = q.CreatePaymentInfo(ctx, db.CreatePaymentInfoParams{
				UserID:      user.ID,
				AccountNo:   fmt.Sprintf("ACC%d", rand.IntN(999999)),
				RoutingNo:   pgtype.Text{String: fmt.Sprintf("RT%d", rand.IntN(999999)), Valid: true},
				AccountName: user.Name,
				BankName:    "Test Bank",
			})
			if err != nil {
				return fmt.Errorf("error creating payment info: %v", err)
			}
		}
	}

	// Create categories (with some parent-child relationships)
	categories := make([]db.Category, 0)
	parentCategories := make([]db.Category, 0)
	
	// First create parent categories
	for _, branch := range branches {
		parentCat, err := q.CreateCategory(ctx, db.CreateCategoryParams{
			Name:        fmt.Sprintf("Parent Category %d", branch.ID.Bytes[:4]),
			BranchID:    branch.ID,
			Description: pgtype.Text{String: "Parent category description", Valid: true},
		})
		if err != nil {
			return fmt.Errorf("error creating parent category: %v", err)
		}
		parentCategories = append(parentCategories, parentCat)
		categories = append(categories, parentCat)

		// Create child categories
		for l := 1; l <= 3; l++ {
			childCat, err := q.CreateCategory(ctx, db.CreateCategoryParams{
				Name:        fmt.Sprintf("Child Category %d-%d", branch.ID.Bytes[:4], l),
				BranchID:    branch.ID,
				ParentID:    pgtype.UUID{Bytes: parentCat.ID.Bytes, Valid: true},
				Description: pgtype.Text{String: "Child category description", Valid: true},
			})
			if err != nil {
				return fmt.Errorf("error creating child category: %v", err)
			}
			categories = append(categories, childCat)
		}
	}

	// Create products
	products := make([]db.Product, 0)
	productTypes := []string{"physical", "service"}
	servicePricingModels := []string{"hourly", "per-project", "tiered"}

	productVIDs := make([]pgtype.UUID, 0)
	
	for _, category := range categories {
		for m := 1; m <= 3; m++ {
			productType := productTypes[rand.IntN(len(productTypes))]
			basePrice := decimal.NewFromFloat(float64(rand.IntN(1000)) + rand.Float64()).Round(2)
			
			product, err := q.CreateProduct(ctx, db.CreateProductParams{
				CategoryID:   category.ID,
				BranchID:    category.BranchID,
				Name:        fmt.Sprintf("Product %d-%d", category.ID.Bytes[:4], m),
				ProductType: productType,
				Sku:        fmt.Sprintf("SKU-%d-%d", category.ID.Bytes[:4], m),
				BasePrice:  pgtype.Numeric{Int: basePrice.BigInt(), Exp: basePrice.Exponent()},
				ServicePricingModel: pgtype.Text{
					String: servicePricingModels[rand.IntN(len(servicePricingModels))],
					Valid:  productType == "service",
				},
			})
			if err != nil {
				return fmt.Errorf("error creating product: %v", err)
			}
			products = append(products, product)

			// Create product versions
			for n := 1; n <= 2; n++ {
				priceAdj := decimal.NewFromFloat(float64(rand.IntN(100)) + rand.Float64()).Round(2)
				version, err := q.CreateProductVersion(ctx, db.CreateProductVersionParams{
					ProductID:        product.ID,
					BranchID:        product.BranchID,
					Name:            fmt.Sprintf("Version %d.%d", m, n),
					PriceAdjustment: pgtype.Numeric{Int: priceAdj.BigInt(), Exp: priceAdj.Exponent()},
					StockQuantity:   pgtype.Int4{ Int32: int32(rand.IntN(100) + 1), Valid: true,},
					ReorderPoint:    pgtype.Int4{ Int32: int32(rand.IntN(20) + 1), Valid: true,},
				})
				if err != nil {
					return fmt.Errorf("error creating product version: %v", err)
				}

				productVIDs = append(productVIDs, version.ID)
				// Create inventory records for physical products
				if product.ProductType == "physical" {
					cost := decimal.NewFromFloat(float64(rand.IntN(500)) + rand.Float64()).Round(2)
					_, err = q.CreateInventoryRecord(ctx, db.CreateInventoryRecordParams{
						VersionID: version.ID,
						BranchID:  version.BranchID,
						Quantity: pgtype.Int4{
							Int32: int32(rand.IntN(50) + 1),
							Valid: true,
						},
						UnitCost: pgtype.Numeric{
							Int:   cost.BigInt(),
							Exp:   cost.Exponent(),
							Valid: true,
						},
					})
					if err != nil {
						return fmt.Errorf("error creating inventory record: %v", err)
					}
				}
			}
		}
	}

	// Create currencies
	currencies := []struct {
		name   string
		code   string
		symbol string
	}{
		{"US Dollar", "USD", "$"},
		{"Euro", "EUR", "€"},
		{"British Pound", "GBP", "£"},
		{"Japanese Yen", "JPY", "¥"},
	}

	createdCurrencies := make([]db.Currency, 0)
	for _, c := range currencies {
		currency, err := q.CreateCurrency(ctx, db.CreateCurrencyParams{
			Name:   c.name,
			Code:   c.code,
			Symbol: pgtype.Text{String: c.symbol, Valid: true},
		})
		if err != nil {
			return fmt.Errorf("error creating currency: %v", err)
		}
		createdCurrencies = append(createdCurrencies, currency)
	}

	// Create customers
	customers := make([]db.Customer, 0)
	for _, branch := range branches {
		for o := 1; o <= 5; o++ {
			customer, err := q.CreateCustomer(ctx, db.CreateCustomerParams{
				Name:     fmt.Sprintf("Customer %d-%d", branch.ID.Bytes[:4], o),
				Email:    pgtype.Text{String: fmt.Sprintf("customer%d-%d@example.com", branch.ID.Bytes[:4], o), Valid: true},
				Phone:    pgtype.Text{String: fmt.Sprintf("+1555%06d", rand.IntN(1000000)), Valid: true},
				BranchID: branch.ID,
			})
			if err != nil {
				return fmt.Errorf("error creating customer: %v", err)
			}
			customers = append(customers, customer)
		}
	}

	// Create invoices and invoice items
	for _, customer := range customers {
		for p := 1; p <= 2; p++ {
			subtotal := decimal.Zero
			currency := createdCurrencies[rand.IntN(len(createdCurrencies))]
			
			invoice, err := q.CreateInvoice(ctx, db.CreateInvoiceParams{
				CustomerID:    pgtype.UUID{Bytes: customer.ID.Bytes, Valid: true},
				InvoiceNumber: fmt.Sprintf("INV-%d-%d", customer.ID.Bytes[:4], p),
				CurrencyID:    pgtype.UUID{Bytes: currency.ID.Bytes, Valid: true},
				Status:        "draft",
				Total:  		pgtype.Numeric{ Int: subtotal.BigInt(), Exp: subtotal.Exponent(), Valid: true},
				Subtotal:      pgtype.Numeric{ Int: subtotal.BigInt(), Exp: subtotal.Exponent(), Valid: true},
				DueDate:      pgtype.Timestamptz{Time: time.Now().AddDate(0, 0, 30), Valid: true},
			})
			if err != nil {
				return fmt.Errorf("error creating invoice: %v", err)
			}

			// Add 2-4 items to each invoice
			for x := 1; x <= rand.IntN(3)+2; x++ {
				// product := products[rand.IntN(len(products))]
				quantity := rand.IntN(5) + 1
				unitPrice := decimal.NewFromFloat(float64(rand.IntN(100)) + rand.Float64()).Round(2)
				
				itemSubtotal := unitPrice.Mul(decimal.NewFromInt(int64(quantity)))
				subtotal = subtotal.Add(itemSubtotal)
				fmt.Println("unitPriceSubtotal", subtotal)

				_, err = q.CreateInvoiceItem(ctx, db.CreateInvoiceItemParams{
					InvoiceID: invoice.ID,
					VersionID: productVIDs[rand.IntN(len(productVIDs))], // Note: In real implementation, you'd need to get a valid version ID
					Quantity:  int32(quantity),
					UnitPrice: pgtype.Numeric{Int: unitPrice.BigInt(), Exp: unitPrice.Exponent(), Valid: true},
					Subtotal:  pgtype.Numeric{Int: itemSubtotal.BigInt(), Exp: itemSubtotal.Exponent() ,Valid: true},
				})
				if err != nil {
					return fmt.Errorf("error creating invoice item: %v", err)
				}
			}

			// Update invoice total
			_, err = q.UpdateInvoice(ctx, db.UpdateInvoiceParams{
				ID:       invoice.ID,
				Subtotal: pgtype.Numeric{Int: subtotal.BigInt(), Exp: subtotal.Exponent(), Valid: true},
				Total:    pgtype.Numeric{Int: subtotal.BigInt(), Exp: subtotal.Exponent(), Valid: true},
			})
			if err != nil {
				return fmt.Errorf("error updating invoice total: %v", err)
			}
		}
	}

	return nil
}

func TestDatabaseSeed(t *testing.T) {
    ctx := context.Background()
    err := SeedDatabase(ctx, testQueries)
    require.NoError(t, err)
}