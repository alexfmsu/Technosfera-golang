package user2

import (
    "database/sql"
)

var db sql.DB
type User2 struct{ 
    ID2 uint `myorm:"primary_key"`
    Login2 string `myorm:"column:username"`
    Info2 string `myorm:"null"`
    Balance2 int 
    Status2 int 
    SomeInnerFlag2 bool `myorm:"-"`
}

func SetDB(_db *sql.DB){
    db = *_db
}

func (u* User2) FindByPK(PK uint) error{
    stmt, err := db.Prepare("SELECT * FROM users2 where ID2 = ?")
    
    if err != nil{
        return err
    }

    row, err := stmt.Query(PK)

    if err != nil{
        return err
    }

    defer row.Close()

    for row.Next(){
        err = row.Scan(&u.ID2, &u.Login2, &u.Info2, &u.Balance2, &u.Status2)
        
        if err != nil {
            return err
        }

        break
    }
    
    return nil
}

func (u* User2) Create() error{
    stmt, err := db.Prepare("INSERT INTO users2(ID2, username, Info2, Balance2, Status2) VALUES(?, ?, ?, ?, ?)")
    
    if err != nil{
        return err
    }

    row, err := stmt.Exec(&u.ID2, &u.Login2, &u.Info2, &u.Balance2, &u.Status2)
    
    if err != nil{
        return err
    }

    last_id, err := row.LastInsertId()
    
    if err != nil{
        return err
    }

    u.ID = uint(last_id)

    return nil
}

func (u* User2) Update() error{
    stmt, err := db.Prepare("UPDATE users2 SET username=?, Info2=?, Balance2=?, Status2=? WHERE ID2 = ?")
    
    if err != nil{
        return err
    }

    _, err = stmt.Exec(&u.Login2, &u.Info2, &u.Balance2, &u.Status2, &u.ID2)

    if err != nil{
        return err
    }

    return nil
}
