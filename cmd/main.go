package main

import (
	"encoding/json"
	"fmt"
	"go-architecture/config"
	"go-architecture/repository"
	"go-architecture/service"
	"go-architecture/usecase"
	"os"

	_ "github.com/go-sql-driver/mysql" // match with your database
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

func initRestAPI(cfg *config.ShareConfig) {
	// connect database
	dbSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		cfg.DB.User,
		cfg.DB.Password,
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.DBName)
	db, err := sqlx.Connect("mysql", dbSourceName) // match your DSN database
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	repo := repository.NewRepository(db, *cfg)
	usecase := usecase.NewUsecase(repo, *cfg)
	httpHandler := service.NewHttpHandler(usecase, *cfg)

	app := fiber.New()

	httpHandler.Router(app)

	app.Listen(":3000")

}

func InitConfig() *config.ShareConfig {
	cfg := config.Config{}
	path, err := os.Getwd()
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
	readConfig, err := os.ReadFile(path + "/env.json")
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	if err = json.Unmarshal(readConfig, &cfg); err != nil {
		log.Error(err)
		os.Exit(1)
	}

	var shareConfig config.ShareConfig
	if cfg.Environment == "development" {
		shareConfig = config.ShareConfig{
			DB: cfg.DB.Development,
		}
	} else {
		shareConfig = config.ShareConfig{
			DB: cfg.DB.Production,
		}
	}

	return &shareConfig

}

func main() {
	cfg := InitConfig()

	config.InitLog()
	initRestAPI(cfg)
}
