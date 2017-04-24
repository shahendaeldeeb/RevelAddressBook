package controllers
import(
	"github.com/revel/revel"
	"fmt"
	"Revel-AddressBook/app/models"
)
type Users struct{
	*revel.Controller
}

func (u Users) Login() revel.Result {
	return u.RenderTemplate("Users/TryLogin.html")
}
func (u Users) TryLogin() revel.Result {
	var Register bool
	var Error string
	u.Validation.Required(u.Params.Form.Get("user.Username"))
	u.Validation.Required(u.Params.Form.Get("user.Password"))

	usedDataBase := &models.GlobeVars
	usedDataBase.Db = Database

	user := models.User{Username:u.Params.Form.Get("user.Username"), Password:[]byte(u.Params.Form.Get("user.Password"))}
	if u.Params.Form.Get("signUp") != ""{
		Register, Error = user.SignUp(usedDataBase)
	}else if u.Params.Form.Get("login") != ""{
		Register, Error = user.Login(usedDataBase)
		fmt.Println("Register : " , Register)
	}
	if Register{
		//.set key-> User , value ->user.Username
		u.Session["User"] = u.Params.Form.Get("user.Username")
		return u.Redirect("/Contacts/%s","Home")
	}else{
		fmt.Println(Error)
		return u.Render(Error)
	}
}
func (u Users) Logout() revel.Result{
	u.Session["User"] = ""
	return u.Redirect("/users/login")
}