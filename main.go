package main

import (
	"exporting/db"
	"exporting/routes"
)

func main() {
	db.DataMigration()
	routes.MainRouting()
}
