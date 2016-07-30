package wgx

import (
	"database/sql"
	"fmt"
	log "github.com/Sirupsen/logrus"
	_ "github.com/mattn/go-sqlite3"
	"gopkg.in/gorp.v1"
)

const (
	// FIXME
	DbFile = "wgx.sqlite3"
)

func InitDatabase() {
	db, err := sql.Open("sqlite3", DbFile)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}
	defer dbmap.Db.Close()

	dbmap.AddTableWithName(Evidence{}, "evidences").SetKeys(true, "Id")
	dbmap.AddTableWithName(Variant{}, "variants").SetKeys(true, "Id")
	dbmap.DropTables()

	dbmap.AddTableWithName(Genome{}, "genomes").SetKeys(true, "Id")
	err = dbmap.CreateTablesIfNotExists()
	if err != nil {
		log.Fatal(err)
	}

	err = InitEvidence(dbmap)
	if err != nil {
		log.Fatal(err)
	}
}

func GetDatabaseConnection() (*sql.DB, *gorp.DbMap, error) {
	db, err := sql.Open("sqlite3", DbFile)
	if err != nil {
		return nil, nil, &GenomeError{fmt.Sprintf("%s", err)}
	}
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}
	dbmap.AddTableWithName(Genome{}, "genomes").SetKeys(true, "Id")

	return db, dbmap, nil
}
