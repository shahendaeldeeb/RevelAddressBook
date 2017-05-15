// GENERATED CODE - DO NOT EDIT
package routes

import "github.com/revel/revel"


type tTelephoneController struct {}
var TelephoneController tTelephoneController


func (_ tTelephoneController) ViewNumbers(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TelephoneController.ViewNumbers", args).URL
}

func (_ tTelephoneController) DeleteNumber(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TelephoneController.DeleteNumber", args).URL
}

func (_ tTelephoneController) AddNumber(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TelephoneController.AddNumber", args).URL
}


type tContactsController struct {}
var ContactsController tContactsController


func (_ tContactsController) Home(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("ContactsController.Home", args).URL
}

func (_ tContactsController) DeleteContact(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("ContactsController.DeleteContact", args).URL
}

func (_ tContactsController) AddContact(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("ContactsController.AddContact", args).URL
}


type tAppInit struct {}
var AppInit tAppInit


func (_ tAppInit) CloseDatabase(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("AppInit.CloseDatabase", args).URL
}


type tUsersController struct {}
var UsersController tUsersController


func (_ tUsersController) Login(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("UsersController.Login", args).URL
}

func (_ tUsersController) TryLogin(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("UsersController.TryLogin", args).URL
}

func (_ tUsersController) Logout(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("UsersController.Logout", args).URL
}


type tTestRunner struct {}
var TestRunner tTestRunner


func (_ tTestRunner) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.Index", args).URL
}

func (_ tTestRunner) Suite(
		suite string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	return revel.MainRouter.Reverse("TestRunner.Suite", args).URL
}

func (_ tTestRunner) Run(
		suite string,
		test string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	revel.Unbind(args, "test", test)
	return revel.MainRouter.Reverse("TestRunner.Run", args).URL
}

func (_ tTestRunner) List(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.List", args).URL
}


type tStatic struct {}
var Static tStatic


func (_ tStatic) Serve(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.Serve", args).URL
}

func (_ tStatic) ServeModule(
		moduleName string,
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModule", args).URL
}


