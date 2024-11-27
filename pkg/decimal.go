package pkg

import (
	"database/sql"
	"fmt"

	"github.com/shopspring/decimal"
)

func ToDecimalFromNullString(ns sql.NullString) decimal.NullDecimal {
	var deci decimal.NullDecimal

	if err:=deci.Scan(ns); err != nil {
		fmt.Println(err)
	}

	return deci
}

func ToNullStringFromDecimal() {
	
}