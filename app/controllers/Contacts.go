package controllers
import (
	"github.com/revel/revel"
	"Revel-AddressBook/app/models"
	"github.com/gocql/gocql"
)
type ContactsController struct {
	 *revel.Controller
 }
func (c ContactsController) Home () revel.Result {
	var page string
	usedDataBase := models.SessionVariable
	page = c.Params.Route.Get("pageName")
	if page == "" {
		page = "Home"
	}
	switch page {
	case "About" :
		return c.RenderTemplate("ContactsController/About.html")
	case "Contact" :
		return c.RenderTemplate("ContactsController/Contact.html")
	case "Home" :
		user := models.User{Username:c.Session["User"]}
		Contacts := user.GetAllContacts(usedDataBase)
		return c.Render(Contacts)
	default:
		return c.RenderTemplate("errors/404.html")

	}
}

func (c ContactsController) DeleteContact() revel.Result{
	var err error
	usedDataBase := models.SessionVariable
	contact := models.ContactInfo{}
	contact.Id, err = gocql.ParseUUID(c.Params.Route.Get("contactid"))
	contact.Username = c.Session["User"]
 	models.CheckErr(err)
	contact.Delete(usedDataBase)
	return c.RenderTemplate("Contacts/Home.html")
}
func (c ContactsController) AddContact() revel.Result{
	usedDataBase := models.SessionVariable
	var contactInfo models.ContactInfo
	c.Validation.Required(c.Params.Form.Get("contact.Name"))
	c.Validation.Required(c.Params.Form.Get("contact.Number"))
	c.Validation.Required(c.Params.Form.Get("contact.Email"))
	c.Validation.Required(c.Params.Form.Get("contact.Nationality"))
	c.Validation.Required(c.Params.Form.Get("contact.Address"))
        c.Validation.Email(c.Params.Form.Get("contact.Email"))

	c.Params.Bind(&contactInfo, "contact")
	contactInfo.Username = c.Session["User"]
	contactInfo.Id = contactInfo.Add(usedDataBase)
	c.Session["Contact"] = contactInfo.Id.String()

	return c.RenderTemplate("Contacts/Home.html")

}
