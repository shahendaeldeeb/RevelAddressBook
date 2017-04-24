package models

import (
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)
var GlobeVars HandlersVars
type User struct {
	Username string `db:"username"`
	Password []byte `db:"userpassword"`
}
type LoginPage struct {
	Error string
}

func (info User) SignUp (a *HandlersVars)(bool, string){
	var result bool
	var Error string
	secret, _ := bcrypt.GenerateFromPassword(info.Password, bcrypt.DefaultCost)
	stmt, err := a.Db.Prepare("INSERT Users SET username=? , userpassword=?")
	_, err = stmt.Exec(info.Username, secret)
	if err != nil{
		Error = err.Error()
		result = false
	}else{
		result = true
	}
	return result, Error
}
func (info User) Login (a *HandlersVars)(bool, string){
	var result bool
	var Error string
	var dbPasswordHash []byte

	row, err := a.Db.Query("SELECT userpassword FROM Users WHERE username =?", info.Username)
	defer row.Close()
	if row.Next(){
		row.Scan(&dbPasswordHash)
	}
	if err != nil {
		Error = err.Error()
		result = false
	}else if dbPasswordHash == nil{
		Error = " No such user found with the username : " + info.Username
	}else {
		err := bcrypt.CompareHashAndPassword(dbPasswordHash, info.Password)
		if err != nil{
			Error = " Wrong Password "
		}else {
			result = true
		}
	}
	return result, Error
}
func (info User) GetAllContacts(a *HandlersVars)[]ContactInfo{
	var contact ContactInfo
	Contacts := []ContactInfo{}
	rows, err := a.Db.Query("select * from Contacts where username =? ", info.Username)
	CheckErr(err)
	for rows.Next() {
		rows.Scan(&contact.Name, &contact.Email, &contact.Nationality, &contact.Address, &contact.Username, &contact.Id)
		Contacts = append(Contacts, contact)
	}
	return Contacts
}
