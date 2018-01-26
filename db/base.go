package db

import (
	"log"
)

func Get(requestMethod string, requestURI string, vars map[string]string) interface{} {
	reply := make(map[string]interface{})
	reply["ver"] = 1.1

	myobj := make(map[string]interface{})
	myobj["id"] = 2
	myobj["name___"] = "PenguenUmut"

	reply["who"] = myobj

	myarr := []string{"a", "bb", "ccc"}
	reply["arr"] = myarr

	log.Println(reply)

	return reply
}
