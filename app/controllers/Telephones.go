package controllers

import (
	"github.com/revel/revel"
	"Revel-AddressBook/app/models"
	"github.com/gocql/gocql"
)

type TelephoneController struct {
	*revel.Controller
}
func (t TelephoneController) ViewNumbers()revel.Result{
	var err error
	usedDataBase := models.SessionVariable
	contact := models.ContactInfo{}
	contact.Id, err = gocql.ParseUUID(t.Params.Route.Get("id"))
	models.CheckErr(err)
	Numbers := contact.GetNumbers(usedDataBase)
	return  t.RenderJSON(Numbers)
}
func (t TelephoneController) DeleteNumber()revel.Result{
	NumID, err :=gocql.ParseUUID(t.Params.Route.Get("numId"))
	models.CheckErr(err)
	usedDataBase := models.SessionVariable
	number := models.Telephone{NumId:NumID}
	number.ContactId, _ = gocql.ParseUUID(t.Params.Route.Get("contactId"))
	number.Delete(usedDataBase)
	return t.RenderTemplate("Contacts/Home.html")
}
func (t TelephoneController) AddNumber()revel.Result{
	tele := models.Telephone{}
	t.Validation.Required(t.Params.Form.Get("NewNumber"))
	tele.Number = t.Params.Form.Get("NewNumber")
	id, err := gocql.ParseUUID(t.Params.Route.Get("ContactId"))
	models.CheckErr(err)
	tele.ContactId = id
	usedDataBase := models.SessionVariable
	 tele.Add(usedDataBase)
	return  t.RenderJSON(tele)
}