package models

type ContactInfo struct{
	Id 		int      `db:"id"`
	Name   		string   `db:"name"`
	Number 		string   `db:"number"`
	Email  		string   `db:"email"`
	Nationality 	string   `db:"nationality"`
	Address 	string   `db:"address"`
	Username  	string   `db:"username"`

}
type Page struct{
	Contacts []ContactInfo
	Numbers  []Telephone
}

func (contact ContactInfo) Add(a *HandlersVars)int64{
	stmt, err := a.Db.Prepare("INSERT Contacts SET name=?, email=?, nationality=?, address=?, username=?")
	CheckErr(err)
	res, err := stmt.Exec(contact.Name, contact.Email, contact.Nationality, contact.Address, contact.Username)
	CheckErr(err)
	id, err := res.LastInsertId()
	CheckErr(err)
	return  id
}
func (contact ContactInfo) Delete(a *HandlersVars){
	stmt, err := a.Db.Prepare("delete from Contacts where id=?")
	CheckErr(err)
	_, err = stmt.Exec(contact.Id)
	CheckErr(err)
}
func (contact ContactInfo) DeleteValidation(LoggedUsername string, a *HandlersVars)bool{
	row, err := a.Db.Query("SELECT username FROM Contacts WHERE id =?", contact.Id)
	defer row.Close()
	CheckErr(err)
	result := false
	var UserName string
	if row.Next() {
		row.Scan(&UserName)
	}
	if UserName == LoggedUsername {
		result = true
	}
	return result
}
func (contact ContactInfo) GetNumbers(a *HandlersVars) []Telephone {
	numbers := []Telephone{}
	rows, err := a.Db.Query("select * from Telephones where contactid =?", contact.Id)
	CheckErr(err)
	var tele Telephone
	for rows.Next() {
		rows.Scan(&tele.ContactId, &tele.Number, &tele.NumId)
		numbers = append(numbers, Telephone{ContactId:tele.ContactId, Number:tele.Number, NumId:tele.NumId })
	}
	return numbers

}
