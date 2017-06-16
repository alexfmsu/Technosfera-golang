package user

import (
    "database/sql"
)

var db sql.DB
type User struct{ 
    ID uint `myorm:"primary_key"`
    Login string `myorm:"column:username"`
    Info string `myorm:"null"`
    Balance int 
    Status int 
    SomeInnerFlag bool `myorm:"-"`}

func SetDB(_db *sql.DB){
    db = *_db
}
func (u* User) FindByPK(PK uint) error{
    var null_info sql.NullString
    
    err := db.QueryRow("SELECT username, Info, Balance, Status FROM users WHERE ID=?", PK).Scan(&u.Login, &null_info, &u.Balance, &u.Status)

    if err != nil{
        return err
    }
    
    if null_info.Valid{
        u.Info = null_info.String
    }

    return nil
}

func (u* User) Create() error{
    stmt, err := db.Prepare("INSERT INTO users(username, Info, Balance, Status) VALUES(?, ?, ?, ?)")
    
    if err != nil{
        return err
    }
    
    var null_info sql.NullString
    
    null_info.String = u.Info
    if u.Info != ""{
        null_info.Valid = true
    }else{
        null_info.Valid = false
    }
    
    
    row, err := stmt.Exec(&u.Login, &u.Info, &u.Balance, &u.Status)

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

func (u* User) Update() error{
    stmt, err := db.Prepare("UPDATE users SET username=?, Info=?, Balance=?, Status=? WHERE ID = ?")
    
    if err != nil{
        return err
    }
    
    var null_info sql.NullString
    
    null_info.String = u.Info
    if u.Info != ""{
        null_info.Valid = true
    }else{
        null_info.Valid = false
    }
    

    _, err = stmt.Exec(&u.Login, &null_info, &u.Balance, &u.Status, &u.ID)

    if err != nil{
        return err
    }

    return nil
}
