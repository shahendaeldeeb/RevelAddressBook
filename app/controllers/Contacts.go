package controllers
import (
	"github.com/revel/revel"
	"Revel-AddressBook/app/models"
	"strconv"
)
type Contacts struct {
	 *revel.Controller
 }
func (c Contacts) Home () revel.Result {
	var page string
	usedDataBase := &models.HandlersVars{Db:Database}
	page = c.Params.Route.Get("pageName")
	if page == ""{
		page = "Home"
	}
	if page == "About"{
		 return c.RenderTemplate("Contacts/About.html")
	}
	if page == "Contact"{
		return c.RenderTemplate("Contacts/Contact.html")
	}else{

		user := models.User{Username:c.Session["User"]}
		Contacts := user.GetAllContacts(usedDataBase)
                return c.Render(Contacts)
	}
}
func (c Contacts) DeleteContact() revel.Result{
	usedDataBase := &models.HandlersVars{Db:Database}
	contact := models.ContactInfo{}
	id, err := strconv.ParseInt(c.Params.Route.Get("id"), 10, 64)
 	models.CheckErr(err)
	contact.Id = int(id)
	found := contact.DeleteValidation(c.Session["User"], usedDataBase)
	if found {
		contact.Delete(usedDataBase)
	}
	return c.RenderTemplate("Contacts/Home.html")
}
func (c Contacts) AddContact() revel.Result{
	usedDataBase := &models.HandlersVars{Db:Database}
	var contactInfo models.ContactInfo
	c.Validation.Required(c.Params.Form.Get("contact.Name"))
	c.Validation.Required(c.Params.Form.Get("contact.Number"))
	c.Validation.Required(c.Params.Form.Get("contact.Email"))
	c.Validation.Required(c.Params.Form.Get("contact.Nationality"))
	c.Validation.Required(c.Params.Form.Get("contact.Address"))
        c.Validation.Email(c.Params.Form.Get("contact.Email"))

	c.Params.Bind(&contactInfo, "contact")
	contactInfo.Username = c.Session["User"]
	id := contactInfo.Add(usedDataBase)
	contactInfo.Id = int(id)

	tele := models.Telephone{Number:contactInfo.Number, ContactId:contactInfo.Id}
	tele.Add(usedDataBase)
	c.Session["Contact"] = string(id)
	return c.RenderTemplate("Contacts/Home.html")

}
