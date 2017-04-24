package controllers

import (
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	"fmt"
	"github.com/revel/revel"
)
type AppInit struct {
	*revel.Controller
}
var Database *sql.DB
func InitDB() {
	var err error
	Database, err = sql.Open("mysql", "root:shahenda_hassan@/mydatabase")
	if err != nil {
		fmt.Println(err.Error())
	}


}
func (a AppInit)CloseDatabase() revel.Result{
	Database.Close()
	return nil
}


