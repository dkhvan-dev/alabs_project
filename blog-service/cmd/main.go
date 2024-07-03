package cmd

import (
	"blogs-api/internal/db"
	"blogs-api/internal/utils"
	"blogs-api/pkg"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"net/http"
)

const SERVER_PORT = "SERVER_PORT"

func main() {
	start()
}

func start() {
	// init logger
	pkg.InitLogger()
	pkg.Logger.Info("Starting server")

	// load environments
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	// init db
	db.InitDB()

	port := ":" + utils.GetEnv(SERVER_PORT)
	r := mux.NewRouter()

	pkg.Logger.Fatal(http.ListenAndServe(port, r).Error())
}
