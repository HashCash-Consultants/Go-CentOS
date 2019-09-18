// Package db provides helpers to connect to test databases.  It has no
// internal dependencies on aurora and so should be able to be imported by
// any aurora package.
package db

import (
	"fmt"
	"log"
	"testing"

	"github.com/jmoiron/sqlx"
	// pq enables postgres support
	_ "github.com/lib/pq"
	db "github.com/hcnet/go/support/db/dbtest"
)

var (
	coreDB     *sqlx.DB
	coreUrl    *string
	auroraDB  *sqlx.DB
	auroraUrl *string
)

// Aurora returns a connection to the aurora test database
func Aurora(t *testing.T) *sqlx.DB {
	if auroraDB != nil {
		return auroraDB
	}
	postgres := db.Postgres(t)
	auroraUrl = &postgres.DSN
	auroraDB = postgres.Open()
	return auroraDB
}

// AuroraURL returns the database connection the url any test
// use when connecting to the history/aurora database
func AuroraURL() string {
	if auroraUrl == nil {
		log.Panic(fmt.Errorf("Aurora not initialized"))
	}
	return *auroraUrl
}

// HcnetCore returns a connection to the hcnet core test database
func HcnetCore(t *testing.T) *sqlx.DB {
	if coreDB != nil {
		return coreDB
	}
	postgres := db.Postgres(t)
	coreUrl = &postgres.DSN
	coreDB = postgres.Open()
	return coreDB
}

// HcnetCoreURL returns the database connection the url any test
// use when connecting to the hcnet-core database
func HcnetCoreURL() string {
	if coreUrl == nil {
		log.Panic(fmt.Errorf("HcnetCore not initialized"))
	}
	return *coreUrl
}
