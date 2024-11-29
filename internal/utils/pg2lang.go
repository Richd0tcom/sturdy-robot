package utils

import (
	"fmt"
	"time"

	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/shopspring/decimal"
)

//TODO: add errors form Scan() in the validdation step

func DecimalToPGNumeric(d decimal.Decimal) pgtype.Numeric {

	nume:=  pgtype.Numeric{
		Int: d.BigInt(),
		Exp: d.Exponent(),
		Valid: true,
	}
	return  nume
}

func ParseUUID(src any) pgtype.UUID{

	var uuid pgtype.UUID

	err:= uuid.Scan(src)

	if err!= nil {
		panic(err)
	}

	return uuid
}

func PgUUIDToString(u pgtype.UUID) string {
	var uuid uuid.UUID

	err:= uuid.Scan(u.Bytes)
	if err!= nil {
		panic(err)
	}

	return uuid.String()
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

func StringToPGText(s string) pgtype.Text {

	var text pgtype.Text

	err:= text.Scan(s)

	if err!= nil {
		panic(err)
	}

	return text
}