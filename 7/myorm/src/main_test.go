package main

import (
    "testing"
    "database/sql"
)

func F(){
    var err error
    db, err := sql.Open("mysql", "root:321678@tcp(localhost:3306)/msu-go-11?charset=utf8&interpolateParams=true")
    PanicOnErr(err)
    
    db.SetMaxOpenConns(10)

}

func TestF(t *testing.T) {

    // err = db.Ping()
    // PanicOnErr(err)
}