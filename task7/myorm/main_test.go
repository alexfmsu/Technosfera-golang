package main

import (
	// "./User"
	// "database/sql"
	"testing"
)

// func TestRunMain(t *testing.T) {
// 	RunMain()
// }

func TestSettDB(t *testing.T) {
	err := SettDB("", "", "", "", "")

	if err != nil {
		// if err != nil {
		t.Errorf("expected \n%+v\n, got \n%+v\n", "123", "23")
	}
}

func TestFindByPK(t *testing.T) {
	// u := user.User{}

	// err := SetDB("localhost", "root", "321678", "tcp", "33062")

	// db, err := sql.Open("mysql", "root:321678@tcp(localhost:3306)/msu-go-11?charset=utf8&interpolateParams=true")

	// SetDB()
	// user.SetDB(db)

	// err = findByPK(u, 12)

	// if err != nil {

	// }
}
