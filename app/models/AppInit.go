package models

import (
	"database/sql"
	"fmt"
)

type HandlersVars struct {
	Db *sql.DB
}
func CheckErr(err error) {
	if err != nil {
		defer func() {
			if msg:=recover() ; msg!=nil    {
				fmt.Printf("Recovered with message %s\n", msg)
			}
		} ()
		panic(err)
	}
}