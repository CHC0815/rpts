package database

import (
	"database/sql"
	"log"

	"cophee.team/rpts/config"
	_ "github.com/go-sql-driver/mysql"
)

func ConnectToDatabase(config config.Config) *sql.DB {
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	return conn
}

func ShowAllTables(db *sql.DB) []string {
	results, err := db.Query("SHOW TABLES")
	if err != nil {
		panic(err.Error())
	}

	var res []string
	for results.Next() {
		var table string
		err = results.Scan(&table)
		if err != nil {
			panic(err.Error())
		}
		res = append(res, table)
	}

	return res
}

/*
	Wieviele Stunden hat jeder einzelne Mitarbeiter auf die einzelnen Projekte gebucht im Zeitraum x inkl. Gesamtzeit

	join Mitarbeiter with Projekt

*/

type HoursOfEmployeePerProjectResult struct {
	Id int
	Vorname string
	Nachname string
}

func HoursOfEmployeePerProject(db *sql.DB) []interface{} {
	query := "SELECT id, Vorname, Nachname FROM `dbo.Mitarbeiter`"
	results, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
	var ress []interface{} 
	for results.Next() {
		var res HoursOfEmployeePerProjectResult
		// expected result: id, Vorname, Nachname, Project, Hours
		err = results.Scan(&res.Id, &res.Vorname, &res.Nachname)
		if err != nil {
			panic(err.Error())
		}
		ress = append(ress, res)
	}
	return ress
}