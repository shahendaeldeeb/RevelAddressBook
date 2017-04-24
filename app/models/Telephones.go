package models

type Telephone struct {
	ContactId int   `db:"contactid" `
	Number string  	`db:"number" `
	NumId int	`db:"numid"`
}
func (tele Telephone) Delete(a *HandlersVars){
	stmt, err := a.Db.Prepare("delete from Telephones where Numid=?")
	CheckErr(err)
	_, err = stmt.Exec(tele.NumId)
	CheckErr(err)
}
func (tele *Telephone) Add(a *HandlersVars)int64{
	stmt2, err := a.Db.Prepare("INSERT Telephones SET number=? , contactid=?")
	CheckErr(err)
	stmt, err := stmt2.Exec(tele.Number, tele.ContactId)
	CheckErr(err)
	id, err := stmt.LastInsertId()
	CheckErr(err)
	return id
}
func (tele Telephone) DeleteValidation(LoggedUsername string, a *HandlersVars)bool{
	row, err := a.Db.Query("SELECT contactid FROM Telephones WHERE numid =?", tele.NumId)
	defer row.Close()
	if row.Next() {
		row.Scan(&tele.ContactId)
	}

	row2, err := a.Db.Query("SELECT username FROM Contacts WHERE id =?",tele.ContactId)
	defer row2.Close()
	CheckErr(err)
	result := false
	var UserName string
	if row2.Next() {
		row2.Scan(&UserName)
	}
	if UserName == LoggedUsername {
		result = true
	}
	return result
}
