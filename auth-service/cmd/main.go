package main

import (
	_ "auth-service/docs"
	"database/sql"
	"github.com/dkhvan-dev/alabs_project/auth-service/internal/api"
	"github.com/dkhvan-dev/alabs_project/auth-service/internal/db"
	"github.com/dkhvan-dev/alabs_project/auth-service/internal/users/service"
	"github.com/dkhvan-dev/alabs_project/auth-service/internal/users/store"
	"github.com/dkhvan-dev/alabs_project/common-libraries/logger"
	"github.com/dkhvan-dev/alabs_project/common-libraries/utils"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/robfig/cron/v3"
	"net/http"
)

const SERVER_PORT = "SERVER_PORT"

// @title User Service API
// @version 1.0
// @description User Service.
// @host localhost:8080
// @BasePath /api/v1
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /
func main() {
	start()
}

func start() {
	// init logger
	logger.InitLogger()
	commonLogger.Info("Starting server")

	// load environments
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	// init db
	db.InitDB()

	// init server
	port := ":" + utils.GetEnv(SERVER_PORT)
	userStore := store.New(db.DB)
	userService := service.New(userStore)
	srv := api.New(*userService, db.DB)
	r := mux.NewRouter()
	srv.AddRoutes(r)

	scheduler := cron.New()

	_, err := scheduler.AddFunc("30 * * * *", func() {
		pkg.Logger.Info("Starting schedule delete tokens")
		deleteTokens(db.DB)
		pkg.Logger.Info("Finishing schedule delete tokens")
	})

	if err != nil {
		panic(err)
	}

	scheduler.Start()

	pkg.Logger.Fatal(http.ListenAndServe(port, r).Error())
}

func deleteTokens(db *sql.DB) {
	if _, err := db.Exec("delete from t_tokens where expiration_deadline >= now()"); err != nil {
		panic(err)
	}
}
