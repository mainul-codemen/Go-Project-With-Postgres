package postgres

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pressly/goose"
)

type Storage struct {
	db *sqlx.DB
}

func NewStorageDB(db *sqlx.DB) *Storage {
	return &Storage{db: db}
}

// NewTestStorage returns a Storage that uses an isolated database for testing purposes
// and a teardown function
func NewTestStorage(dbstring, migrationDir string) (*Storage, func()) {
	db, teardown := MustNewDevelopmentDB(dbstring, migrationDir)
	db.SetMaxOpenConns(5)
	db.SetConnMaxLifetime(time.Hour)

	return NewStorageDB(db), teardown
}

// MustNewDevelopmentDB creates a new isolated database for the use of a package test
// The checking of dbconn is expected to be done in the package test using this
func MustNewDevelopmentDB(ddlConnStr, migrationDir string) (*sqlx.DB, func()) {
	const driver = "postgres"

	dbName := generateRandomString()
	fmt.Println("dbName: ", dbName)
	ddlDB := sqlx.MustConnect(driver, ddlConnStr)
	ddlDB.MustExec(fmt.Sprintf(`CREATE DATABASE %s`, dbName))
	if err := ddlDB.Close(); err != nil {
		panic(err)
	}

	connStr := addDBName(ddlConnStr, dbName)
	db := sqlx.MustConnect(driver, connStr)

	if err := goose.SetDialect(driver); err != nil {
		panic(err)
	}

	if err := goose.Run("up", db.DB, migrationDir); err != nil {
		panic(err)
	}

	tearDownFn := func() {
		if err := db.Close(); err != nil {
			log.Fatalf("failed to close database connection: %s", err.Error())
		}
		ddlDB, err := sqlx.Connect(driver, ddlConnStr)
		if err != nil {
			log.Fatalf("failed to connect database: %s", err.Error())
		}

		if _, err = ddlDB.Exec(fmt.Sprintf(`DROP DATABASE %s`, dbName)); err != nil {
			log.Fatalf("failed to drop database: %s", err.Error())
		}

		if err = ddlDB.Close(); err != nil {
			log.Fatalf("failed to close DDL database connection: %s", err.Error())
		}
	}

	return db, tearDownFn
}

// replaceDBName replaces the dbname option in connection string with given db name in parameter.
func addDBName(connStr, dbName string) string {
	return fmt.Sprintf("%s%s", strings.Trim(connStr, " "), dbName)
}

func generateRandomString() string {
	rand.Seed(time.Now().Unix())

	//Lowercase and Uppercase Both
	charSet := "abcdedfghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	length := 20
	var output strings.Builder
	for i := 0; i < length; i++ {
		random := rand.Intn(len(charSet))
		randomChar := charSet[random]
		output.WriteString(string(randomChar))
	}
	defer output.Reset()
	return string(output.String())
}
