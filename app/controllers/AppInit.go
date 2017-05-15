package controllers

import (
	"github.com/gocql/gocql"
	_"github.com/go-sql-driver/mysql"
	"github.com/revel/revel"
	"Revel-AddressBook/app/models"
	"fmt"
)
type AppInit struct {
	*revel.Controller
}
func InitDB(){
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "Mydatabase"
	var err error
	models.SessionVariable, err = cluster.CreateSession()
	if err != nil{
		fmt.Println(err.Error())
	}
}
func (a AppInit)CloseDatabase() revel.Result{
	 defer models.SessionVariable.Close()
	return nil
}


