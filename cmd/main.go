package main

import (
	"flag"
	"log"

	"github.com/FOMARTEM/newssite-golang/internal/api"
	"github.com/FOMARTEM/newssite-golang/internal/config"
	"github.com/FOMARTEM/newssite-golang/internal/provider"
	"github.com/FOMARTEM/newssite-golang/internal/usecase"
)

func main() {
	// Считываем аргументы командной строки
	configPath := flag.String("config-path", "../config/config.yaml", "путь к файлу конфигурации")
	flag.Parse()

	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		log.Fatal(err)
	}

	prv := provider.NewProvider(cfg.DB.Host, cfg.DB.Port, cfg.DB.User, cfg.DB.Password, cfg.DB.DBname)
	use := usecase.NewUsecase(prv)
	srv := api.NewServer(cfg.IP, cfg.Port, use, cfg.API.SecretKey)

	srv.Run()
}
