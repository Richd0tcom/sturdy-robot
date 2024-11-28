package utils

import (
	"math/rand"
	"strings"
	"time"

	"github.com/gofrs/uuid"
)

// This file generates random test data for testing


const alphabet = "abcdefghijklmnopqrstuvwxyz"
const numphabet = "abcdefghijklmnopqrstuvwxyz1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ"

//initialises the seed of the random packeage to make sure that value are indeed random
func init(){
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

//Generates a random interger between min and max values
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generates a random string of length n
func RandomString(n int ) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomOwner generates a random owner name
func RandomOwner() string {
	return RandomString(6)
}

// RandomMoney generates a random amount of money
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// RandomCurrency generates a random currency code
func RandomCurrency() string {
	currencies := []string{"USD", "EUR", "CAD"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}

func RandomInvoiceNumber() string {
	var sb strings.Builder
	k := len(numphabet)

	for i := 0; i < 6; i++ {
		c := numphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func NewRandomUUID() uuid.UUID {
	uid, _:= uuid.NewV4()
	return uid
}