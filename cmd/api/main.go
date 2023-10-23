package main

import (
	"log"

	"github.com/nikolaev-roman/simbir-go/config"
	"github.com/nikolaev-roman/simbir-go/internal/server"
	"github.com/nikolaev-roman/simbir-go/pkg/db/postgres"
)

//	@title			Simbir.GO
//	@version		1
//	@description	service for transport rent

func main() {
	cfgFile, err := config.LoadConfig("config/config_local")
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}

	cfg, err := config.ParseConfig(cfgFile)
	if err != nil {
		log.Fatalf("ParseConfig: %v", err)
	}

	psqlDB, err := postgres.NewPsqlDB(cfg)
	if err != nil {
		log.Fatalf("Postgresql init: %s", err)
	} else {
		log.Println("Postgres connected")
	}

	postgres.MigrateDb(psqlDB)

	server := server.NewServer(cfg, psqlDB)

	if err = server.Run(); err != nil {
		log.Fatal(err)
	}

}
