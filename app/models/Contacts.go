package models

import (
	"github.com/gocql/gocql"
)

type ContactInfo struct{
	Id 		gocql.UUID      `db:"contactid"`
	Name   		string   	`db:"name"`
	Number 		string  	`db:"number"`
	Email  		string   	`db:"email"`
	Nationality 	string   	`db:"nationality"`
	Address 	string   	`db:"address"`
	Username  	string   	`db:"username"`
}
type Page struct{
	Contacts []ContactInfo
	Numbers  []Telephone
}
func (contact ContactInfo) Add(a *gocql.Session)gocql.UUID{
	var err error
	contact.Id, err = gocql.RandomUUID()
	NumId, _ := gocql.RandomUUID()

	CheckErr(err)
	batch := gocql.NewBatch(gocql.LoggedBatch)
	batch.Query("INSERT INTO Mydatabase.contact_by_userid (name, email, nationality, address, username, contactid) VALUES (?, ?, ?, ?, ?, ?) ",contact.Name, contact.Email, contact.Nationality, contact.Address, contact.Username, contact.Id)
	batch.Query("INSERT INTO Mydatabase.numbers_by_contactid (contactid, numid, number) VALUES (?, ?, ?)",  contact.Id, NumId, contact.Number)
	err = a.ExecuteBatch(batch)
	return  contact.Id
}
func (contact ContactInfo) Delete(a *gocql.Session){
	 err := a.Query(`DELETE from Mydatabase.contact_by_userid where username=? AND contactid=?`,contact.Username, contact.Id).Exec()
	CheckErr(err)
}
func (contact ContactInfo) GetNumbers(a *gocql.Session) []Telephone {
	numbers := []Telephone{}
	iter := a.Query("select * from Mydatabase.numbers_by_contactid where contactid =?", contact.Id).Iter()
	var tele Telephone
	for iter.Scan(&tele.ContactId, &tele.NumId, &tele.Number) {
		numbers = append(numbers, Telephone{ContactId:tele.ContactId, Number:tele.Number, NumId:tele.NumId })
	}
	return numbers

}
