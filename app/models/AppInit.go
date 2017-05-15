package models

import (
	"fmt"
	"github.com/gocql/gocql"
)
var (
	SessionVariable *gocql.Session
)

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