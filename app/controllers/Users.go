package controllers
import(
	"github.com/revel/revel"
	"Revel-AddressBook/app/models"
)
type UsersController struct{
	*revel.Controller
}

func (u UsersController) Login() revel.Result {
	return u.RenderTemplate("Users/TryLogin.html")
}
func (u UsersController) TryLogin() revel.Result {
	var Register bool
	var Error string
	u.Validation.Required(u.Params.Form.Get("user.Username"))
	u.Validation.Required(u.Params.Form.Get("user.Password"))

	usedDataBase := models.SessionVariable

	user := models.User{Username:u.Params.Form.Get("user.Username"), Password:[]byte(u.Params.Form.Get("user.Password"))}
	if u.Params.Form.Get("signUp") != ""{
		Register, Error = user.SignUp(usedDataBase)
	}else if u.Params.Form.Get("login") != ""{
		Register, Error = user.Login(usedDataBase)
	}
	if Register{
		//.set key-> User , value ->user.Username
		u.Session["User"] = u.Params.Form.Get("user.Username")
		return u.Redirect("/Contacts/%s","Home")
	}else{
		return u.Render(Error)
	}
}
func (u UsersController) Logout() revel.Result{
	u.Session["User"] = ""
	return u.Redirect("/users/login")
}