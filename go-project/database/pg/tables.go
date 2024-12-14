package database

import (
	"database/sql"
	"fmt"
	"log"
)

type SyncTableStuct struct {
	success string
	err     error
}

func CreateTableSQL(db *sql.DB, s CreateTable, done chan SyncTableStuct) {
	sql := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS "%s" (`, s.TableName)

	for idx := range s.Columns {
		column := s.Columns[idx]
		sql += fmt.Sprintf(" %s %s ", column.Name, column.Type)
		for cix := range column.Constraints {
			constraints := column.Constraints[cix]
			sql += fmt.Sprintf(" %s ", constraints)
		}
		if idx+1 < len(s.Columns) {
			sql += ","
		}
	}
	sql += fmt.Sprintf(")")
	_, err := db.Exec(sql)
	if err != nil {
		// Send back an error if there's an issue creating the table
		done <- SyncTableStuct{
			success: "",
			err:     fmt.Errorf("failed to create table %s: %v", s.TableName, err),
		}
	}
	// Send back success message if the table is created successfully
	done <- SyncTableStuct{
		success: fmt.Sprintf("Table %s created successfully", s.TableName),
		err:     nil,
	}
}

var tableList = []CreateTable{
	UserTable,
}

func SyncTable(db *sql.DB) {

	syncDoneChan := make(chan SyncTableStuct, len(tableList))
	for ind := range tableList {
		//create table
		CreateTableSQL(db, tableList[ind], syncDoneChan)

	}
	for range tableList {
		singleChan := <-syncDoneChan // Wait for result from each goroutine

		// Check if there's an error or success
		if singleChan.err != nil {
			log.Fatal("Table is not synced:", singleChan.err.Error())
		} else {
			fmt.Println(singleChan.success)
		}
	}

	close(syncDoneChan)

}
