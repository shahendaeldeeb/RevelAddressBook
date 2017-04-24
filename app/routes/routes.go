// GENERATED CODE - DO NOT EDIT
package routes

import "github.com/revel/revel"


type tUsers struct {}
var Users tUsers


func (_ tUsers) Login(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Users.Login", args).URL
}

func (_ tUsers) TryLogin(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Users.TryLogin", args).URL
}

func (_ tUsers) Logout(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Users.Logout", args).URL
}


type tTelephone struct {}
var Telephone tTelephone


func (_ tTelephone) ViewNumbers(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Telephone.ViewNumbers", args).URL
}

func (_ tTelephone) DeleteNumber(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Telephone.DeleteNumber", args).URL
}

func (_ tTelephone) AddNumber(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Telephone.AddNumber", args).URL
}


type tContacts struct {}
var Contacts tContacts


func (_ tContacts) Home(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Contacts.Home", args).URL
}

func (_ tContacts) DeleteContact(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Contacts.DeleteContact", args).URL
}

func (_ tContacts) AddContact(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Contacts.AddContact", args).URL
}


type tAppInit struct {}
var AppInit tAppInit


func (_ tAppInit) CloseDatabase(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("AppInit.CloseDatabase", args).URL
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


