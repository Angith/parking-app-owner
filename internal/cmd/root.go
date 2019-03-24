package cmd

import (
	"parking-app-owner/internal/api"
	"parking-app-owner/internal/data"
)

func execute() {
	db := prepareDB()
	app := api.App{
		ParkingRepo: data.NewRepo(db),
	}
	app.Serve()
}
