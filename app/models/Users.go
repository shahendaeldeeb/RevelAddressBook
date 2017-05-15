package models

import (
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	"github.com/gocql/gocql"
)
type User struct {
	Username string `db:"username"`
	Password []byte `db:"userpassword"`
}
type LoginPage struct {
	Error string
}

func (info User) SignUp (a *gocql.Session)(bool, string){
	var result bool
	var Error string
	secret, _ := bcrypt.GenerateFromPassword(info.Password, bcrypt.DefaultCost)
	err := a.Query(`INSERT INTO Mydatabase.users (username, userpassword) VALUES (?, ?)`,
		info.Username, secret).Exec();

	if err != nil{
		Error = err.Error()
		result = false
	}else{
		result = true
	}
	return result, Error
}
func (info User) Login (a *gocql.Session)(bool, string){
	var result bool
	var Error string
	var dbPasswordHash []byte

	err := a.Query(`SELECT userpassword FROM Mydatabase.users WHERE username =?`, info.Username).Scan(&dbPasswordHash)
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
func (info User) GetAllContacts(a *gocql.Session)[]ContactInfo{
	var contact ContactInfo
	Contacts := []ContactInfo{}
	iter := a.Query(`select * from Mydatabase.contact_by_userid where username =? `, info.Username).Iter()

	for iter.Scan(&contact.Username, &contact.Id, &contact.Address, &contact.Email, &contact.Name, &contact.Nationality){
		Contacts = append(Contacts, contact)
	}
	if err := iter.Close(); err != nil {
		CheckErr(err)
	}
	return Contacts
}
