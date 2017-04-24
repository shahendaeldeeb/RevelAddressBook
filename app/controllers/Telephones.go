package controllers

import (
	"github.com/revel/revel"
	"Revel-AddressBook/app/models"
	"strconv"
	"fmt"
)

type Telephone struct {
	*revel.Controller
}
func (t Telephone) ViewNumbers()revel.Result{
	usedDataBase := &models.HandlersVars{Db:Database}
	contact := models.ContactInfo{}
	id, err := strconv.ParseInt(t.Params.Route.Get("id"), 10, 64)
	models.CheckErr(err)
	contact.Id = int(id)
	Numbers := contact.GetNumbers(usedDataBase)
	return  t.RenderJSON(Numbers)
}
func (t Telephone) DeleteNumber()revel.Result{
	NumID, _ := strconv.ParseInt(t.Params.Route.Get("numId"), 10, 64)
	usedDataBase := &models.HandlersVars{Db:Database}
	number := models.Telephone{NumId:int(NumID)}
	found := number.DeleteValidation(t.Session["User"], usedDataBase)
	fmt.Println("found",found)
	if found{
		number.Delete(usedDataBase)
	}
	return t.RenderTemplate("Contacts/Home.html")
}
func (t Telephone) AddNumber()revel.Result{
	tele := models.Telephone{}
	t.Validation.Required(t.Params.Form.Get("NewNumber"))
	tele.Number = t.Params.Form.Get("NewNumber")
	id, err := strconv.ParseInt(t.Params.Route.Get("ContactId"), 10, 64)
	models.CheckErr(err)
	tele.ContactId = int(id)
	usedDataBase := &models.HandlersVars{Db:Database}
	LastId := tele.Add(usedDataBase)
	tele.NumId = int(LastId)
	return  t.RenderJSON(tele)
}