package utils

import (
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/shopspring/decimal"
)

func DecimalToPGNumeric(d decimal.Decimal) pgtype.Numeric {

	nume:=  pgtype.Numeric{
		Int: d.BigInt(),
		Exp: d.Exponent(),
		Valid: true,
	}
	fmt.Println("nume: ", nume)
	return  nume
	
}

func ParseUUID(src any) pgtype.UUID{

	var uuid pgtype.UUID

	err:= uuid.Scan(src)

	if err!= nil {
		fmt.Println("uuid err: ", err)
		panic(err)
	}

	return uuid
}

func ParseDate(src any) pgtype.Timestamptz {
	
	var date pgtype.Timestamptz
	t, err:=time.Parse(time.RFC3339, src.(string))
	if err!= nil {
		fmt.Println("time parse err: ", err)
		panic(err)

	}

	err= date.Scan(t)
	if err!= nil {
		fmt.Println("date scan err: ",err )
		panic(err)
	}

	return date
}