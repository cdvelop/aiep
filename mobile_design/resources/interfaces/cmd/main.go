package main

import (
	"github.com/cdvelop/aiep/mobile_design/resources/interfaces/app"
	"github.com/cdvelop/aiep/mobile_design/resources/interfaces/database"
)

func main() {

	newDB := &database.DataBase{}

	app.AppRun(newDB)

}
