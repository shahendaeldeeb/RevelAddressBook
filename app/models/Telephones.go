package models

import (
	"github.com/gocql/gocql"
)

type Telephone struct {
	ContactId gocql.UUID   	`db:"contactid" `
	Number string  		`db:"number" `
	NumId gocql.UUID	`db:"numid"`
}
func (tele Telephone) Delete(a *gocql.Session){
	err := a.Query("DELETE from Mydatabase.numbers_by_contactid where contactid=? AND Numid=?",tele.ContactId, tele.NumId).Exec()
	CheckErr(err)

}
func (tele *Telephone) Add(a *gocql.Session){
	var err error
	tele.NumId, err = gocql.RandomUUID()
	CheckErr(err)
	err = a.Query(`INSERT INTO Mydatabase.numbers_by_contactid (contactid, numid, number) VALUES (?, ?, ?) IF NOT EXISTS`,  tele.ContactId, tele.NumId, tele.Number).Exec()
	CheckErr(err)

}
