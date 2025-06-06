package main

import (
	"log"
	"main/api"
	"main/api/fetch"
	"main/config/env"
)

func main() {
	adress := env.GetString("PG_ADRESS", ":8080")
	maxOpenConnections := env.GetInt("PG_MAX_OPEN_CONNS", 30)
	maxIdleConnections := env.GetInt("PG_MAX_IDLE_CONNS", 30)
	maxIdleTime := env.GetString("PG_IDLE_TIME", "15min")

	dbConfig := api.DbConfig{
		Address:            adress,
		MaxOpenConnections: maxOpenConnections,
		MaxIdleConnections: maxIdleConnections,
		MaxIdleTime:        maxIdleTime,
	}

	cfg := api.Config{
		Address: ":8080",
		Version: "0.0.1",
		Db:      dbConfig,
	}

	db := fetch.New(
		cfg.Db.Address,
		cfg.Db.MaxOpenConnections,
		cfg.Db.MaxIdleConnections,
		cfg.Db.MaxIdleTime,
	)

	defer db.Close()
	log.Printf("database connection pool established")

	store := fetch.NewPostGresStorage(db)

	app := &api.Application{
		Config: cfg,
		Store:  store,
	}

	mux := app.Mount()

	log.Fatal(app.Run(mux))
}
